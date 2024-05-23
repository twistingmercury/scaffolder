package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	Usage     = "scaffolder [commands] [flags]"
	ShortDesc = "Scaffolder is a CLI tool for Go project scaffolding"
	LongDesc  = `Scaffolder is a command-line tool that allows you to 
quickly create new Go projects based on templates 
hosted on GitHub. It simplifies the process of 
setting up a new project by cloning a template 
repository and replacing template tokens with 
user-provided values.`
)

func NewRootCmd() *cobra.Command {
	rCmd := cobra.Command{
		Use:   Usage,
		Short: ShortDesc,
		Long:  LongDesc,
	}
	rCmd.SetHelpFunc(help)
	return &rCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if rootCmd.Flags().Changed("help") {
		_ = rootCmd.Help()
		return
	}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func help(cmd *cobra.Command, args []string) {
	fmt.Fprintln(cmd.OutOrStdout(), cmd.UsageString())
}
