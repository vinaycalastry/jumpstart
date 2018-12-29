package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vinaycalastry/jumpstart/db"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an existing boiler plate code stored.",
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		err := db.DeleteLink(title)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s has been removed from storage\n", title)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
