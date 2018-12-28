package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "Replaces an existing code stored with another link.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Here is the list")
	},
}

func init() {
	RootCmd.AddCommand(replaceCmd)
}
