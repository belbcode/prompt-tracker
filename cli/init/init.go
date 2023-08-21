package initRuntime

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/belbcode/prompt-tracker/cli/commit"
	utils "github.com/belbcode/prompt-tracker/utils"
)

var arguments []string = []string{"init", "commit", "help"}

const dirname string = ".pt"
const ownerReadWrite = 0700 // 0700 sets read, write, and execute permissions for the owner only

func ParseArgs() (string, error) {
	if len(os.Args) < 2 {
		fmt.Println()
		return "", errors.New("Provide argument for pt. Example: pt init <file-name>")
	}

	if utils.Includes(arguments, os.Args[1]) {
		return os.Args[1], nil
	}
	return "", errors.New(os.Args[1] + "is not an argument use --help to see list of valid arguments") //Do this
}

func Init() (string, error) {

	fileInfo, dirErr := os.Stat(dirname)
	if dirErr == nil {
		if fileInfo.IsDir() {
			fmt.Println("You already initialized pt in this directory") //edge case
		} else {
			fmt.Println("Path exists, but it's not a directory.")
		}
		return "", errors.New("") //CHANGE
	} else if os.IsNotExist(dirErr) {
		fmt.Println("Directory does not exist.")
		err := os.Mkdir(dirname, ownerReadWrite)
		if err != nil {
			fmt.Println("There was an error initializing the app. May the lord have mercy on your os.", err)
			return "", errors.New("") //CHANGE
		}
		return dirname, nil

	} else {
		fmt.Println("Error:", dirErr)
		return "", errors.New("") //CHANGE
	}

}

func ExtractFileInfo(f fs.FileInfo) utils.FileInfo {
	return utils.FileInfo{
		IsDir:   f.IsDir(),
		Name:    f.Name(),
		Size:    f.Size(),
		ModTime: f.ModTime().Unix(),
	}
}

func CreateConfig(parentDirectory string) {

	trackedFiles := os.Args[2:]
	fileObjects := make(map[string]utils.FileObject, len(trackedFiles))
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(trackedFiles) == 0 {
		panic(errors.New("No files specified to be tracked"))
	} else {
		for _, file := range trackedFiles {
			//may need to fix this
			fileInfo, err := os.Stat(file)
			if err != nil {
				//maybe create a tracked file if a certain flag is raised in future
				panic(errors.New("Error trying to locate file, may not exist" + err.Error()))
			}
			fi := ExtractFileInfo(fileInfo)

			originalPath := filepath.Join(cwd, file)
			uuid := utils.HashName(originalPath)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			repoPath := filepath.Join(cwd, "/.pt/", uuid)

			fileObjects[uuid] = utils.FileObject{
				OriginalPath: originalPath,
				RepoPath:     repoPath,
				Properties:   fi,
			}

		}
	}
	trackedPath := filepath.Join(cwd, "/.pt/")
	config := &utils.Config{
		TrackedFiles: fileObjects,
		InitTime:     time.Now().Unix(),
		TrackedPath:  trackedPath,
	}

	jsonData, err := json.MarshalIndent(config, "\n", "	")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonString := string(jsonData)

	err = ioutil.WriteFile(parentDirectory+"/pt.config.json", []byte(jsonString), 0644)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File pt.config.json created successfully.")
	}
}

func Scaffold(config utils.Config) error {
	// config, err := utils.GetConfig(cwd)
	// if err != nil {
	// 	panic(err.Error())
	// }
	for _, fileObject := range config.TrackedFiles {
		err := os.Mkdir(fileObject.RepoPath, ownerReadWrite)
		if err != nil {
			fmt.Println("Unable to track: ", fileObject.Properties.Name)
			//Remove file from config
		}
		// initialize two files to be commited, one empty and the original
		empty := make([]byte, 0)
		name, err := utils.GenerateRandomID(16)
		path := filepath.Join(fileObject.RepoPath, name)
		os.Create(path)
		os.WriteFile(path, empty, ownerReadWrite)

		commit.CommitFile(fileObject.OriginalPath)

		// content, err := os.ReadFile(fileObject.OriginalPath)
		// fmt.Println(string(content))
		// name, err = utils.GenerateRandomID(16)
		// path = filepath.Join(fileObject.RepoPath, name)
		// os.Create(path)
		// os.WriteFile(path, content, ownerReadWrite)

	}
	return nil
}
