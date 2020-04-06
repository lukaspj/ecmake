package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func GetVersionCommand(config Config) *cobra.Command {
	return &cobra.Command{
		Use: "version",
		Short: "Print the version of ECMAke",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(fmt.Sprintf("%v, commit %v, built at %v", config.Version, config.Commit, config.Date))
		},
	}
}