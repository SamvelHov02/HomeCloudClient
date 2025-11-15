package Node

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	httphelper "github.com/SamvelHov02/HomeCloudHTTP"
)

const VaultPath = "/Users/samvelhovhannisyan/Documents/dev/Personal/HomeCloud/client/"

/*
Function starts a client process
method : int
- 1 = GET
- 2 = POST
- 3 = PUT
- 4 = DELETE
*/
func Start(method string, resource string) ([]byte, httphelper.Status) {
	fmt.Println("Starting client process...")
	// conn, err := net.Dial("tcp", "192.168.50.132:8080")
	var h httphelper.Header
	h.Add("Accept", "application/json")

	body := httphelper.Body{}

	switch method {
	case "post":
		Info, err := os.Stat(VaultPath + resource)

		if err != nil {
			log.Fatal(err)
		}

		if !Info.IsDir() {
			bodyData, err := os.ReadFile(VaultPath + resource)
			if err != nil {
				log.Fatal(err)
			}

			body.Data = string(bodyData)
			h.Add("Content-Type", "application/json")
		} else {
			resource = "dir/" + resource
		}
	case "put":
		file, err := os.ReadFile(VaultPath + resource)

		if err != nil {
			log.Fatal(err)
		}

		body.Data = string(file)
		h.Add("Content-Type", "application/json")
	case "delete":
		Info, err := os.Stat(VaultPath + resource)

		if err != nil {
			log.Fatal(err)
		}

		if Info.IsDir() {
			resource = "dir/" + resource
		}
		h.Add("Content-Type", "application/json")
		h.Add("Content-Length", "0")
	default:
		h.Add("Content-Length", "0")
	}

	request := httphelper.WriteRequest(method, "api/"+method+"/"+resource, h, body)

	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Error connecting to server:", err)
		log.Fatal(err)
	}
	defer conn.Close()


	_, err = conn.Write(request)

	if err != nil {
		log.Fatal(err)
	}

	response, err := io.ReadAll(conn)

	if err != nil {
		log.Fatal(err)
	}

	responseData, _, Status := httphelper.ReadResponse(response)

	return responseData, Status
}

func UpdateFile(response httphelper.Body, resource string) {
	f, err := os.Create(VaultPath + resource)

	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(f)

	_, err = w.Write([]byte(response.Data))

	if err != nil {
		log.Fatal(err)
	}

	w.Flush()
	f.Close()
}

func GetLocalMerkle() httphelper.Leaf {
	tree := &httphelper.Tree{}
	tree.Init(VaultPath + "Vault")
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
				} else {
					difference = append(difference, sChild.Name)
				}
			}
		}
	}
	return difference
}

func StatusResult(Status httphelper.Status) {
	switch Status.Code {
	case 204:
		fmt.Println("Operation Successful!")
	case 404:
		fmt.Println("Resource wasn't found")
	case 409:
		fmt.Println("File already exists on the server")
	case 400:
		fmt.Println("Bad request")
	}
}
