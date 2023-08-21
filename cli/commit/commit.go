package commit

import (
	"errors"
	"fmt"
	"os"

	"github.com/pmezard/go-difflib/difflib"
)

func CommitFile(filePath string) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(bytes)

}

func getChanges(fromFile string, toFile string) string {
	a, err := os.ReadFile(fromFile)
	if err != nil {
		fmt.Println(err)
		// change this
		panic(errors.New("AAAAAAAAAAA"))
	}
	b, err := os.ReadFile(toFile)
	if err != nil {
		fmt.Println(err)
		// change this
		panic(errors.New("AAAAAAAAAAA"))

	}
	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(string(a)),
		B:        difflib.SplitLines(string(b)),
		FromFile: "Original",
		ToFile:   "Current",
		Context:  3,
	}
	changeString, _ := difflib.GetUnifiedDiffString(diff)
	return changeString

}
