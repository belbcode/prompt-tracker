package initRuntime

import (
	"fmt"
	"os"
	"testing"
)

func TestExtractMetaData(t *testing.T) {
	name := "test.txt"
	f, err := os.Stat(name)
	if err != nil {
		fmt.Println("f-off")
	}
	res := ExtractMetaData(f)

	t.Error(res)
}
