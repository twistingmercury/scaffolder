package cmd

import "github.com/spf13/cobra"

var (
	rootCmd    *cobra.Command
	initCmd    *cobra.Command
	versionCmd = &cobra.Command{}
)

func Initialize() {
	rootCmd = NewRootCmd()
	versionCmd = NewVersionCmd()
	initCmd = NewInitCommand()
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(versionCmd)
}

func RootCmd() *cobra.Command {
	return rootCmd
}
