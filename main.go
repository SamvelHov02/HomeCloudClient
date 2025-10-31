package main

import (
	"client/cli"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Hello from the Client side")
	var cmd *cli.Command

	fmt.Println(args[2])

	switch args[1] {
	case "-gt":
		cmd = cli.GetTreeCmd
	case "-g":
		cmd = cli.GetFile
	}

	// Need to get all flags before executing
	fmt.Println(cmd)
	cmd.Init(cmd.Name)
	cmd.Build(args[1:])
	cmd.Execute()
}
