package Node

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"

	httphelper "github.com/SamvelHov02/HomeCloudHTTP"
)

const VaultPath = "/Users/samvelhovhannisyan/Documents/dev/Personal/HomeCloud/client/Vault"

/*
Function starts a client process
method : int
- 1 = GET
- 2 = POST
- 3 = PUT
- 4 = DELETE
*/
func Start(method int, resource string) httphelper.ResponseBody {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var h httphelper.Header
	h.Add("Accept", "application/json")

	request := httphelper.WriteRequest(method, resource, h)

	_, err = conn.Write(request)

	if err != nil {
		log.Fatal(err)
	}

	response, err := io.ReadAll(conn)

	if err != nil {
		log.Fatal(err)
	}
	// TODO
	responseData, _, _ := httphelper.ReadResponse(response)
	return responseData
}

func UpdateFile(response httphelper.ResponseBody, resource string) {
	f, err := os.Create(VaultPath + resource)

	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(f)

	w.Write([]byte(response.Data))
	w.Flush()
	f.Close()
}

func GetLocalMerkle() httphelper.Leaf {
	tree := &httphelper.Tree{}
	tree.Init(VaultPath)
	tree.Build()
	tree.ComputeHash()

	leaf := httphelper.Leaf{
		Hash:     tree.RootHash,
		Name:     tree.Root,
		Children: tree.Children,
		Category: "dir",
	}

	return leaf
}

func CompareTrees(serverLeaf httphelper.Leaf, clientLeaf httphelper.Leaf) []string {
	var difference []string

	// If there are no changes then the Root hash will be the same
	if serverLeaf.Equal(clientLeaf) {
		return difference
	} else {
		for _, sChild := range serverLeaf.Children {
			for _, cChild := range clientLeaf.Children {
				/*
					1. Same file and same hash
					2. Same file and different hash
					3. Different files
				*/

				if sChild.Equal(*cChild) && cChild.Hash != sChild.Hash && cChild.Category == "file" {
					difference = append(difference, cChild.Name)
				} else if sChild.Equal(*cChild) && cChild.Hash != sChild.Hash && cChild.Category == "dir" {
					difference = append(difference, CompareTrees(*sChild, *cChild)...)
				}
			}
		}
	}
	return difference
}
