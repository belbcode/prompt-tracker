package commit

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/belbcode/prompt-tracker/utils"
	"github.com/pmezard/go-difflib/difflib"
)

const ownerReadWrite = 0700 // 0700 sets read, write, and execute permissions for the owner only

func CommitFile(currentFile string) {

	hashed := utils.HashName(currentFile)
	latestFile, err := utils.GetLatestTrackedFile(hashed)
	diff := GetDiff(latestFile, currentFile)

	bytes, err := os.ReadFile(currentFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	cwd := utils.GetCwd()
	config, err := utils.GetConfig(cwd)
	name, err := utils.GenerateRandomID(16)
	path := filepath.Join(config.TrackedFiles[hashed].RepoPath, name)

	file, err := os.Create(path)
	defer file.Close()
	// Write data using the Writer interface

	unifiedDiffString, err := difflib.GetUnifiedDiffString(diff)
	file.Write([]byte(unifiedDiffString))
	file.Seek(int64(len([]byte(unifiedDiffString))), 0)
	file.Write(bytes)
}

func GetDiff(fromFile string, toFile string) difflib.UnifiedDiff {
	println(fromFile, toFile)
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
	// changeString, _ := difflib.GetUnifiedDiffString(diff)
	return diff

}

// func WriteCommit(originalPath string, ) {
// 	cwd := utils.GetCwd()
// 	config, err := utils.GetConfig(cwd)

// 	content, err := os.ReadFile(originalPath)
// 	fmt.Println(string(content))
// 	name, err := utils.GenerateRandomID(16)
// 	path := filepath.Join(, name)
// 	os.Create(path)
// 	os.WriteFile(path, content, ownerReadWrite)
// }
