package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/twistingmercury/scaffolder/conf"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "scaffolder version",
		Short: "Returns the current scaffolder version",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprintf(cmd.OutOrStdout(), "scaffolder version %s\n\n", conf.BuildVersion)
			return err
		},
	}
}
