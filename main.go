package main

import (
	Node "client/node"
	"fmt"
)

func main() {
	fmt.Println("Hello from the Client side")
	method := 1
	resp := Node.Start(method, "/test1.md")

	/*
		For which operations does the Client need to change ITS local Vault
		GET 	: YES, needs to update since it's used for syncing
		POST	: NO, tells server to sync with the changes client has made
		PUT 	: NO, tells server to sync with the changes client has made
		DELETE 	: NO, tells server to sync with the changes client has made
	*/
	if method == 1 {
		Node.UpdateFile(resp, "/test1.md")
	}
}
