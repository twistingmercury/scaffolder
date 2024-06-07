package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const ( // flag names
	moduleNameFlag  = "module-name"
	binNameFlag     = "bin-name"
	descriptionFlag = "description"
	vendorNameFlag  = "vendor-name"
	helpFlag        = "help"
	tmplFlag        = "template"
)

type Token []string

type TokenInfo struct {
	Keys  Token
	Value string
}

func (t TokenInfo) Regex() *regexp.Regexp {
	return regexp.MustCompile(strings.Join(t.Keys, "|"))
}

type TemplateInfo struct {
	ModuleName  TokenInfo
	BinName     TokenInfo
	Description TokenInfo
	VendorName  TokenInfo
	RootDir     string
	GitPath     string
}

var ( // template variables
	ModuleNameTokens  = Token{`MODULE_NAME`, `{{module_name}}`}
	BinNameTokens     = Token{`BIN_NAME`, `{{bin_name}}`}
	DescriptionTokens = Token{`IMG_DESCRIPTION`, `{{description}}`}
	VendorNameTokens  = Token{`IMG_VENDOR_NAME`, `{{vendor_name}}`}
)

func NewTemplateInfo(gitPath, rootDir, moduleName, binName, vendorName, description string) (te TemplateInfo, err error) {
	if moduleName == "" {
		return te, errors.New("no module name provided")
	}
	if binName == "" {
		return te, errors.New("no bin name provided")
	}
	if gitPath == "" {
		return te, errors.New("no git project URL provided")
	}
	if rootDir == "" {
		return te, errors.New("no git project root provided")
	}
	if vendorName == "" {
		vendorName = "TODO: provide a vendor name"
	}
	if description == "" {
		description = fmt.Sprintf("TODO: provide a description for %s", binName)
	}

	te = TemplateInfo{
		ModuleName:  TokenInfo{ModuleNameTokens, moduleName},
		BinName:     TokenInfo{BinNameTokens, binName},
		Description: TokenInfo{DescriptionTokens, description},
		VendorName:  TokenInfo{VendorNameTokens, vendorName},
		RootDir:     rootDir,
		GitPath:     gitPath,
	}
	return
}

func NewInitCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "init",
		Short: "Initializes a new project using a Go template project from github.com.",
		Run:   CreateProject,
	}
	cmd.Flags().SortFlags = false

	cmd.Flags().StringP(moduleNameFlag, "m", "", "Name of the go module (e.g. github.com/username/project-name) as set in go.mod file")
	cmd.Flags().StringP(binNameFlag, "b", "", "The name of the binary file to be compiled. Sets the service name in the conf.go file.")
	cmd.Flags().StringP(tmplFlag, "t", "https://github.com/twistingmercury/go-basic-tmpl.git", "The project to clone from github.com.")
	cmd.Flags().StringP(descriptionFlag, "d", "", "A brief description of the project, set in the Dockerfile as a label.")
	cmd.Flags().StringP(vendorNameFlag, "v", "", "The name of the vendor, set in the Dockerfile as a label.")
	cmd.Flags().BoolP(helpFlag, "h", false, "Help for init command.")

	return &cmd
}

// CreateProject creates a new project using a project from GitHub as a template.
func CreateProject(cmd *cobra.Command, _ []string) {
	if cmd.Flag(helpFlag).Changed {
		fmt.Println("Help for init command")
		_ = cmd.Help()
		return
	}

	moduleName := strings.ToLower(cmd.Flag(moduleNameFlag).Value.String())
	binName := strings.ToLower(cmd.Flag(binNameFlag).Value.String())
	description := cmd.Flag(descriptionFlag).Value.String()
	vendorName := cmd.Flag(vendorNameFlag).Value.String()
	gitPath := cmd.Flag(tmplFlag).Value.String()
	rootDir := binName

	templateInfo, err := NewTemplateInfo(gitPath, rootDir, moduleName, binName, vendorName, description)
	if err != nil {
		fmt.Println("error creating template info:", err)
		return
	}

	fmt.Printf("module name: %s\nbin name: %s\nvendor name: %s\ndescription: %s\ntemplate: %s\n", moduleName, binName, vendorName, description, gitPath)
	err = Clone(templateInfo)
	if err != nil {
		fmt.Println("error cloning project:", err)
		return
	}
	fmt.Println("project cloned successfully.")

	err = MakeExecutable(binName + "/_build/build.sh")
	if err != nil {
		fmt.Println("error making build.sh executable:", err)
		return
	}

	err = MakeExecutable(binName + "/_build/common.sh")
	if err != nil {
		fmt.Println("error making common.sh executable:", err)
		return
	}

	err = ReplaceInFiles(templateInfo)

	if err != nil {
		fmt.Println("error replacing tokens:", err)
		return
	}

	fmt.Println("tokens replaced successfully.")

	err = GoModTidy(binName)
	if err != nil {
		fmt.Println("error running go mod tidy:", err)
		return
	}
}

// Clone clones a project from GitHub.com using the provided git path.
func Clone(ti TemplateInfo) error {
	cloneCmd := exec.Command("git", "clone", ti.GitPath, ti.RootDir)
	err := cloneCmd.Start()
	if err != nil {
		return err
	}

	err = cloneCmd.Wait()
	if err != nil {
		return err
	}

	err = os.RemoveAll("./" + ti.RootDir + "/.git")
	if err != nil {
		return err
	}

	return os.Remove("./" + ti.RootDir + "/go.sum")
}

// ReplaceInFiles inspects all files in a directory and replaces tokens with the provided values.
func ReplaceInFiles(ti TemplateInfo) error {
	err := filepath.Walk(ti.RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			newText, err := ReplaceTokens(path, ti)
			if err != nil {
				return err
			}
			err = os.WriteFile(path, []byte(newText), info.Mode())
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// ReplaceTokens replaces tokens in a file with the provided values.
func ReplaceTokens(filePath string, ti TemplateInfo) (string, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	var originalTxt = string(bytes)

	if len(originalTxt) == 0 {
		log.Printf("file %s is empty", filePath)
		return "", nil
	}

	updatedText := ti.ModuleName.Regex().ReplaceAllString(originalTxt, ti.ModuleName.Value)
	updatedText = ti.BinName.Regex().ReplaceAllString(updatedText, ti.BinName.Value)
	updatedText = ti.VendorName.Regex().ReplaceAllString(updatedText, ti.VendorName.Value)
	updatedText = ti.Description.Regex().ReplaceAllString(updatedText, ti.Description.Value)

	return updatedText, nil
}

// GoModTidy runs `go mod tidy` in the provided path.
func GoModTidy(path string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = path
	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Wait()
}

// MakeExecutable makes a file executable.
func MakeExecutable(path string) error {
	return os.Chmod(path, 0755)
}
