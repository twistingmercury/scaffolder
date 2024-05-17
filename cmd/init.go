package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const ( // template variables
	ModuleName  = `{{module_name}}`
	BinName     = `{{bin_name}}`
	Description = `{{description}}`
	VendorName  = `{{vendor_name}}`
)

const ( // flag names
	moduleNameFlag  = "module-name"
	binNameFlag     = "bin-name"
	descriptionFlag = "description"
	vendorNameFlag  = "vendor-name"
	helpFlag        = "help"
	tmplFlag        = "template"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new project using a project from github.com as a template.",
	Run:   CreateProject,
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().SortFlags = false

	initCmd.Flags().StringP(moduleNameFlag, "m", "", "[Required] Name of the go module (e.g. github.com/username/project-name) as set in go.mod file")
	initCmd.Flags().StringP(binNameFlag, "b", "", "[Required] The name of the binary file to be compiled. Sets the service name in the conf.go file.")
	initCmd.Flags().StringP(tmplFlag, "t", "https://github.com/twistingmercury/gobasetmpl.git", "[Required] The project to clone from github.com.")
	initCmd.Flags().StringP(descriptionFlag, "d", "", "A brief description of the project, set in the dockerfile as a lable.")
	initCmd.Flags().StringP(vendorNameFlag, "v", "", "The name of the vendor, set in the dockerfile as a lable.")
	initCmd.Flags().BoolP(helpFlag, "h", false, "Help for init command.")
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

	if moduleName == "" {
		_ = cmd.Help()
		return
	}
	if binName == "" {
		_ = cmd.Help()
		return
	}
	if gitPath == "" {
		_ = cmd.Help()
		return
	}
	if vendorName == "" {
		vendorName = "TODO: provide a vendor name"
	}
	if description == "" {
		description = fmt.Sprintf("TODO: provide a description for %s", binName)
	}

	fmt.Printf("module name: %s\nbin name: %s\nvendor name: %s\ndescription: %s\ntemplate: %s\n", moduleName, binName, vendorName, description, gitPath)
	err := Clone(gitPath, binName)
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

	err = ReplaceInFiles(binName, moduleName, binName, description, vendorName)
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
func Clone(gitPath, binName string) error {
	cloneCmd := exec.Command("git", "clone", gitPath, binName)
	err := cloneCmd.Start()
	if err != nil {
		return err
	}

	err = cloneCmd.Wait()
	if err != nil {
		return err
	}

	err = os.RemoveAll("./" + binName + "/.git")
	if err != nil {
		return err
	}

	return os.Remove("./" + binName + "/go.sum")
}

// ReplaceInFiles inspects all files in a directory and replaces tokens with the provided values.
func ReplaceInFiles(path, moduleName, binName, descr, vendorName string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			newText, err := ReplaceTokens(path, moduleName, binName, descr, vendorName)
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
func ReplaceTokens(path, moduleName, binName, descr, vendorName string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	text := string(bytes)
	updateText := strings.ReplaceAll(text, ModuleName, moduleName)
	updateText = strings.ReplaceAll(updateText, BinName, binName)
	updateText = strings.ReplaceAll(updateText, Description, descr)
	updateText = strings.ReplaceAll(updateText, VendorName, vendorName)

	return updateText, nil
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
