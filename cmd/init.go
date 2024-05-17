package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new project using a project from github.com as a template.",
	Run:   createProject,
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().String(moduleNameFlag, "", "[Required] Name of the go module (e.g. github.com/username/project-name) as set in go.mod file")
	initCmd.Flags().String(binNameFlag, "", "[Required] The name of the binary file to be compiled. Sets the service name in the conf.go file.")
	initCmd.Flags().String(descriptionFlag, "", "A brief description of the project, set in the dockerfile as a lable.")
	initCmd.Flags().String(vendorNameFlag, "", "The name of the vendor, set in the dockerfile as a lable.")
	initCmd.Flags().Bool(helpFlag, false, "Help for init command.")
}

func createProject(cmd *cobra.Command, args []string) {
	if cmd.Flag(helpFlag).Changed {
		fmt.Println("Help for init command")
		_ = cmd.Help()
		return
	}

	moduleName := cmd.Flag(moduleNameFlag).Value.String()
	binName := cmd.Flag(binNameFlag).Value.String()
	description := cmd.Flag(descriptionFlag).Value.String()
	vendorName := cmd.Flag(vendorNameFlag).Value.String()

	if moduleName == "" {
		_ = cmd.Help()
		return
	}
	if binName == "" {
		_ = cmd.Help()
		return
	}
	if vendorName == "" {
		vendorName = "TODO: provide a vendor name"
	}
	if description == "" {
		description = fmt.Sprintf("TODO: provide a description for %s", binName)
	}

	fmt.Printf("module name: %s\nbin name: %s\nvendor name: %s\ndescription: %s\n", moduleName, binName, vendorName, description)
}
