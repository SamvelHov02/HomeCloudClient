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

	switch args[1] {
	case "-gt":
		cmd = cli.GetTreeCmd
	case "-g":
		cmd = cli.GetFile
	}

	cmd.Execute()
}
