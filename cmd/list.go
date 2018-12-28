package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vinaycalastry/jumpstart/db"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all available boiler plate codes stored.",
	Run: func(cmd *cobra.Command, args []string) {
		links, err := db.GetLinks()

		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}

		if len(links) == 0 {
			fmt.Println("Nothing stored so far. Add a new project using jumpstart add <projname> <link>")
		}

		fmt.Println("The following links to quickstart codes are stored:")
		for i, link := range links {
			fmt.Printf("%d: %s (%s)\n", i+1, link.Key, link.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
