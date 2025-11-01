package Node

import (
	"log"
	"os"
	"testing"

	httphelper "github.com/SamvelHov02/HomeCloudHTTP"
)

func TestUpdateFile(t *testing.T) {
	resp := httphelper.Body{Data: "# Test file 1\n\nA little more data"}
	UpdateFile(resp, "test1.md")
	actual, err := os.ReadFile("/Users/samvelhovhannisyan/Documents/dev/Personal/HomeCloud/client/Vault/test1.md")

	if err != nil {
		log.Fatal(err)
	}

	if string(actual) != resp.Data {
		t.Errorf("Expected %s got %s", resp.Data, string(actual))
	}
}

func TestUpdatefileNoFile(t *testing.T) {
	resp := httphelper.Body{Data: "# Test file 1\n\nA little more data"}
	UpdateFile(resp, "/test3.md")
	actual, err := os.ReadFile("/Users/samvelhovhannisyan/Documents/dev/Personal/HomeCloud/client/Vault/test3.md")

	if err != nil {
		log.Fatal(err)
	}

	if string(actual) != resp.Data {
		t.Errorf("Expected %s got %s", resp.Data, string(actual))
	} else {
		os.Remove("/Users/samvelhovhannisyan/Documents/dev/Personal/HomeCloud/client/Vault/test3.md")
	}
}
