package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
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
		fmt.Printf("Found \"%s\". Dowloading quick-start code now.\n", title)
		if err != nil {
			fmt.Println(err)
			return
		}

		//Execute git clone
		var stdoutBuf, stderrBuf bytes.Buffer
		command := exec.Command("git", "clone", downloadURLString)

		stdoutIn, _ := command.StdoutPipe()
		stderrIn, _ := command.StderrPipe()

		var errStdout, errStderr error
		stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
		stderr := io.MultiWriter(os.Stderr, &stderrBuf)
		err = command.Start()

		if err != nil {
			fmt.Println("Cloning the quickstart repo failed")
			return
		}
		go func() {
			_, errStdout = io.Copy(stdout, stdoutIn)
		}()

		go func() {
			_, errStderr = io.Copy(stderr, stderrIn)
		}()

		err = command.Wait()
		if err != nil || errStdout != nil || errStderr != nil {
			fmt.Println("Cloning the quickstart repo failed")
			return
		}
		fmt.Printf("Quickstart code for %s has been downloaded to current directory\n", title)
	},
}

func init() {
	RootCmd.AddCommand(newCmd)
}
