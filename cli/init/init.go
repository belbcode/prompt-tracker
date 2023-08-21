package init

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	utils "github.com/belbcode/prompt-tracker/utils"
)

var arguments []string = []string{"init", "commit", "help"}

func ParseArgs() (string, error) {
	if len(os.Args) < 1 {
		fmt.Println()
		return "", errors.New("Provide argument for pt. Example: pt init <file-name>")
	}

	if utils.Includes(arguments, os.Args[0]) {
		return os.Args[0], nil
	}
	return "", errors.New(os.Args[0] + "is not an argument use --help to see list of valid arguments") //Do this
}

func Init() {
	const dirname string = ".pt"
	const ownerReadWrite = 0700 // 0700 sets read, write, and execute permissions for the owner only

	fileInfo, dirErr := os.Stat(dirname)
	if dirErr == nil {
		if fileInfo.IsDir() {
			fmt.Println("You already initialized pt in this directory") //edge case
		} else {
			fmt.Println("Path exists, but it's not a directory.")
		}
		return
	} else if os.IsNotExist(dirErr) {
		fmt.Println("Directory does not exist.")
		err := os.Mkdir(dirname, ownerReadWrite)
		if err != nil {
			fmt.Println("There was an error initializing the app. May the lord have mercy on your os.", err)
			return
		}

	} else {
		fmt.Println("Error:", dirErr)
	}

}

type Config struct {
	date         int64
	trackedFiles []string
}

func createConfig() {
	trackedFiles := os.Args[1:]
	config := Config{
		trackedFiles: trackedFiles,
		date:         time.Now().Unix(),
	}
	jsonData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = ioutil.WriteFile("pt.config.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File pt.config.json created successfully.")
	}
}
