package initRuntime

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"

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

func inLocalDir(args []string) (err error) {
	for _, fileName := range args {
		//maybe include functionality that ignore ex: if flag == "generous"
		_, err = os.Stat(fileName)
		return err
	}
	return nil
}

func createConfig(args []string) utils.Config {
	const repoPath string = "/.pt/"
	cwd := utils.GetCwd()

	//get tracked files
	fileObjects := make(map[string]utils.FileObject, len(args))

	for _, fileName := range args {
		fileInfo, err := os.Stat(fileName)
		if err != nil {
			log.Fatal(err)
			//maybe with flags this can be different //ex: if flag == "generous" : just ignore file
		}

		ogPath := filepath.Join(cwd, fileName)
		hashedName := utils.HashString(ogPath)

		fileObjects[hashedName] = utils.FileObject{
			OriginalPath: ogPath,
			RepoPath:     filepath.Join(cwd, repoPath, hashedName),
			Properties:   ExtractFileInfo(fileInfo),
		}

	}
	repoDir := filepath.Join(cwd, repoPath)

	config := &utils.Config{
		TrackedFiles: fileObjects,
		InitTime:     time.Now().Unix(),
		RepoDir:      repoDir,
	}
	return *config
}

func configToJSON(config utils.Config) []byte {
	jsonData, err := json.MarshalIndent(config, "\n", "	")
	if err != nil {
		panic(err)
	}
	return jsonData

}

func scaffold(config utils.Config) (err error) {
	configFilename := "/pt.config.json"
	jsonData := configToJSON(config)
	err = os.Mkdir(config.RepoDir, ownerReadWrite)
	err = os.WriteFile(filepath.Join(config.RepoDir, configFilename), jsonData, 0644)
	for _, file := range config.TrackedFiles {
		dirName := utils.HashString(file.OriginalPath)
		err = os.Mkdir(filepath.Join(config.RepoDir, dirName), ownerReadWrite)
	}
	return err
}

func Init(args []string) {
	config := createConfig(args)
	err := scaffold(config)

	if err != nil {
		log.Fatal("Failed to initialize PromptTracker: ", err)
	}

	fmt.Println("Initialization successful, created repo @ ", config.RepoDir)

}

func ExtractFileInfo(f fs.FileInfo) utils.FileInfo {
	return utils.FileInfo{
		IsDir:   f.IsDir(),
		Name:    f.Name(),
		Size:    f.Size(),
		ModTime: f.ModTime().Unix(),
	}
}
