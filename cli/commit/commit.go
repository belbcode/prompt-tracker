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

func getRelevantCommitPaths(arg string, config utils.Config) (sourcePath string, commitPath string) {
	hashedFileKey := utils.HashString(arg)
	Object, exists := config.Objects[hashedFileKey]
	if !exists {
		panic("file specified in commit: " + arg + " has not been initialized") //Create ADD Option
	}
	sourcePath = Object.SourceFile
	commitPath = Object.LastCommit
	return sourcePath, commitPath
}

func isSameContent(bytesSource []byte, bytesLastCommit []byte) bool {
	//true: no change, false: yes change
	//this can be optimized with utils.SoftCompare and technically the SHA1 of bytesLastCommit is already encoded within its name
	//you can also check if the file had been modified since the last commit, a lot of optimizations in this function
	if len(bytesSource) == len(bytesLastCommit) {
		if utils.SoftCompare(bytesSource, bytesLastCommit) {
			return utils.HardCompare(bytesSource, bytesLastCommit) //
		}
	}
	return false //change detected
}

func getDiff(bytesSource []byte, bytesLastCommit []byte) difflib.UnifiedDiff {
	return difflib.UnifiedDiff{
		A:        difflib.SplitLines(string(bytesLastCommit)),
		B:        difflib.SplitLines(string(bytesSource)),
		FromFile: "Original",
		//FromDate: FileInfo.ModTime().Unix()
		//ToDate: time.Now().Unix()
		ToFile:  "Current",
		Context: 3,
	}
}

func writeDiff(commitPath string, diff difflib.UnifiedDiff) error {
	file, err := os.Create(commitPath)
	defer file.Close()
	err = difflib.WriteUnifiedDiff(file, diff)
	return err
}

// func writeConfig

func CommitFiles(args []string) error {
	cwd := utils.GetCwd()
	config, err := utils.GetConfig(cwd)
	if err != nil {
		return err
	}

	for _, arg := range args {

		sourceFilePath, commitFilePath := getRelevantCommitPaths(arg, config)
		sourceBytes, err := os.ReadFile(sourceFilePath)
		commitBytes, err := os.ReadFile(commitFilePath)

		diff := getDiff(sourceBytes, commitBytes)

		newCommitName := utils.HashString(string(sourceBytes))
		hashedFileName := utils.HashString(arg)
		newCommitPath := filepath.Join(config.Objects[hashedFileName].RepoPath, newCommitName)

		err = writeDiff(newCommitPath, diff)
		if err != nil {
			return err
		}

	}

	return utils.UpdateConfig()
	// hashed := utils.HashString(currentFile)
	// latestFile, err := utils.GetLatestTrackedFile(hashed)
	// diff := GetDiff(latestFile, currentFile)

	// bytes, err := os.ReadFile(currentFile)
	// if err != nil {
	// 	return err
	// }

	// cwd := utils.GetCwd()
	// config, err := utils.GetConfig(cwd)
	// name, err := utils.GenerateRandomID(16)
	// path := filepath.Join(config.TrackedFiles[hashed].RepoPath, name)

	// file, err := os.Create(path)
	// defer file.Close()
	// // Write data using the Writer interface

	// unifiedDiffString, err := difflib.GetUnifiedDiffString(diff)
	// file.Write([]byte(unifiedDiffString))
	// file.Seek(int64(len([]byte(unifiedDiffString))), 0)
	// file.Write(bytes)
	// return nil
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
