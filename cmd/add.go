package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vinaycalastry/jumpstart/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a link to a new project and can be used whenever required.",
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		link := args[1]

		err := db.AddLink(title, link)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("New project added to store.")

	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
