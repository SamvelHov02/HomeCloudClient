package cli

import (
	Node "client/node"
	"encoding/json"
	"fmt"
	"log"

	httphelper "github.com/SamvelHov02/HomeCloudHTTP"
)

const VaultPath = "/Users/samvelhovhannisyan/Documents/dev/Personal/HomeCloud/client/Vault"

var GetTreeCmd = &Command{
	Name: "Get Tree",

	Run: func(cmd *Command) {
		resp := Node.Start("get", "tree")
		tree := Node.GetLocalMerkle()
		serverTree := &httphelper.Tree{}
		dataRaw, err := json.Marshal(resp.Data)

		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(dataRaw, serverTree)

		l := httphelper.Leaf{
			Category: "dir",
			Hash:     serverTree.RootHash,
			Children: serverTree.Children,
			Name:     serverTree.Root,
		}

		differences := Node.CompareTrees(l, tree)

		// -r resolver, gets all the updated files
		if _, ok := cmd.FlagsParam["-r"]; ok {
			for _, file := range differences {
				resp := Node.Start("get", file)
				Node.UpdateFile(resp, file)
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
		resp := Node.Start("get", cmd.FlagsParam["-g"])
		Node.UpdateFile(resp, cmd.FlagsParam["-g"])
	},
}
