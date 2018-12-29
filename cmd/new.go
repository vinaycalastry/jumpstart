package cmd

import (
	"fmt"
	"os/exec"

	"github.com/vinaycalastry/jumpstart/db"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Downloads the new boiler plate code for the mentioned project",
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		downloadURL, err := db.FindLink(title)
		downloadURLString := string(downloadURL)
		if err != nil {
			fmt.Println(err)
			return
		}
		command := exec.Command("git", "clone", downloadURLString)
		err = command.Run()
		if err != nil {
			fmt.Println("Cloning the quickstart repo failed")
			return
		}
		fmt.Printf("Quickstart code for %s has been downloaded to current directory\n", title)
	},
}

func init() {
	RootCmd.AddCommand(newCmd)
}
