package cli

import (
	Node "client/node"
	"encoding/json"
	"fmt"
	"log"

	httphelper "github.com/SamvelHov02/HomeCloudHTTP"
)

const VaultPath = "/home/samo/dev/HomeCloud/client/"

var GetTreeCmd = &Command{
	Name: "Get Tree",

	Run: func(cmd *Command) {
		resp, _ := Node.Start("get", "tree")
		tree := Node.GetLocalMerkle()
		serverTree := &httphelper.Tree{}

		err := json.Unmarshal(resp, serverTree)

		if err != nil {
			log.Fatal(err)
		}

		l := httphelper.Leaf{
			Category: "dir",
			Hash:     serverTree.RootHash,
			Children: serverTree.Children,
			Name:     serverTree.Root,
		}

		fmt.Println(serverTree.Children)
		fmt.Println("-----------------------")
		fmt.Println(tree.Children)

		differences := Node.CompareTrees(l, tree)

		// -r resolver, gets all the updated files
		if _, ok := cmd.FlagsParam["-r"]; ok {
			for _, file := range differences {
				var body httphelper.Body
				resp, Status := Node.Start("get", file)
				json.Unmarshal(resp, &body)
				Node.UpdateFile(body, file)
				Node.StatusResult(Status)
			}
		} else {
			fmt.Println("Printing only the differences : ", differences)
		}
	},
}

var GetFile = &Command{
	Name:        "Get File",
	Description: "Fetches a file from the server",
	Run: func(cmd *Command) {
		resp, Status := Node.Start("get", cmd.FlagsParam["-g"])
		body := httphelper.Body{}
		err := json.Unmarshal(resp, &body)

		if err != nil {
			log.Fatal(err)
		}
		Node.UpdateFile(body, cmd.FlagsParam["-g"])
		Node.StatusResult(Status)
	},
}

// Refactor ? Could aggregate POST to be like the DELTE method.
var PostFile = &Command{
	Name:        "Create File",
	Description: "Creates a new File received from the client",
	Run: func(cmd *Command) {
		Node.Start("post", cmd.FlagsParam["-p"])
	},
}

var PostDir = &Command{
	Name:        "Create Directory",
	Description: "Creates a local directory on the server",
	Run: func(cmd *Command) {
		Node.Start("post", cmd.FlagsParam["-pd"])
	},
}

var PutFile = &Command{
	Name:        "Update file",
	Description: "Updates a file on the server",
	Run: func(cmd *Command) {
		Node.Start("put", cmd.FlagsParam["-u"])
	},
}

var DeleteResource = &Command{
	Name:        "Delete Resource",
	Description: "Deletes a resource, file or directory from the server",
	Run: func(cmd *Command) {
		Node.Start("delete", cmd.FlagsParam["-d"])
	},
}
