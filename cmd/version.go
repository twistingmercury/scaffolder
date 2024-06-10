package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/twistingmercury/scaffolder/conf"
)

const logo = `
 o-o            o-o o-o    o    o         
|               |   |      |    |         
 o-o   o-o  oo -O- -O- o-o |  o-O o-o o-o 
    | |    | |  |   |  | | | |  | |-' |   
o--o   o-o o-o- o   o  o-o o  o-o o-o o `

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Returns the current scaffolder version",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprintf(cmd.OutOrStdout(),
				"%s\nversion: %s, commit date: %s\n\n",
				logo,
				conf.BuildVersion(),
				conf.BuildDate())
			return err
		},
	}
}
