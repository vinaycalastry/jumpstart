package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "jumpstart",
	Short: "Jumpstart is a very simple CLI tool to download boiler plate code.",
	Long:  `Jumpstart is a simple CLI tool to download starting code for various projects including electron, vue, react, nodejs, angular, golang etc.`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
