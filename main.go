package main

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/vinaycalastry/jumpstart/cmd"
	"github.com/vinaycalastry/jumpstart/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "links.db")
	err := db.Init(dbPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd.Execute()
}
