package main

import (
	"client/cli"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Hello from the Client side", args)
	var cmd *cli.Command

	switch args[1] {
	case "-gt":
		cmd = cli.GetTreeCmd
	case "-g":
		cmd = cli.GetFile
	case "-p":
		cmd = cli.PostFile
	case "-pd":
		cmd = cli.PostDir
	case "-u":
		cmd = cli.PutFile
	}

	// Need to get all flags before executing
	cmd.Init(cmd.Name)
	cmd.Build(args[1:])
	cmd.Execute()
}
