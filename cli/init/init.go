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

	utils "github.com/belbcode/prompt-tracker/utils"
)

var arguments []string = []string{"init", "commit", "help"}

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
	const dirname string = ".pt"
	const ownerReadWrite = 0700 // 0700 sets read, write, and execute permissions for the owner only

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

type Config struct {
	NewDate      int64
	TrackedFiles map[string]FileObject
}

type FileObject struct {
	FullPath   string
	Properties fs.FileInfo
}

func CreateConfig(parentDirectory string) {

	trackedFiles := os.Args[2:]
	fileObjects := make(map[string]FileObject, len(trackedFiles))

	if len(trackedFiles) == 0 {
		panic(errors.New("No files specified to be tracked"))
	} else {
		for _, file := range trackedFiles {
			//may need to fix this
			fileInfo, err := os.Stat(file)
			if err != nil {
				//maybe create a tracked file if a certain flag is raised in future
				panic(errors.New("File to be tracked does not exist."))
			}
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fullpath := filepath.Join(cwd, file)
			uuid, err := utils.GenerateRandomID(16)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			println(fullpath, fileInfo)
			fileObjects[uuid] = FileObject{
				FullPath:   fullpath,
				Properties: fileInfo,
			}

		}
	}
	config := &Config{
		TrackedFiles: fileObjects,
		NewDate:      time.Now().Unix(),
	}

	jsonData, err := json.Marshal(config)
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

func Scaffold(parentDirectory string, trackedFiles []string) {
	// os.Mkdir(parentDirectory + "/")
}
