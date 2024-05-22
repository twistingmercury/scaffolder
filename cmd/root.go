package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scaffolder",
	Short: "Scaffolder is a CLI tool for Go project scaffolding",
	Long: `Scaffolder is a command-line tool that allows you to 
quickly create new Go projects based on templates 
hosted on GitHub. It simplifies the process of 
setting up a new project by cloning a template 
repository and replacing template tokens with 
user-provided values.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.scaffolder.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("help", "h", false, "Help for scaffolder command")
}
