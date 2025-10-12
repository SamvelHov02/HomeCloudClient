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
