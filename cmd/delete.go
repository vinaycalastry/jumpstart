package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an existing boiler plate code stored.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleted the project from db")
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
