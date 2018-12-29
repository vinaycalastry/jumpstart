package cmd

import (
	"fmt"

	"github.com/vinaycalastry/jumpstart/db"

	"github.com/spf13/cobra"
)

var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "Replaces an existing code stored with another link.",
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		newURL := args[1]

		err := db.DeleteLink(title)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = db.AddLink(title, newURL)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Replaced existing link for %s with %s", title, newURL)
	},
}

func init() {
	RootCmd.AddCommand(replaceCmd)
}
