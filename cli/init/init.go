package initRuntime

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	utils "github.com/belbcode/prompt-tracker/utils"
)

var arguments []string = []string{"init", "commit", "help"}

const dirname string = ".pt"
const ownerReadWrite = 0700 // 0700 sets read, write, and execute permissions for the owner only

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
			SourceFile: ogPath,
			RepoPath:   filepath.Join(cwd, repoPath, hashedName),
			Properties: utils.ExtractFileInfo(fileInfo),
			LastCommit: utils.HashString(""), //The first file to be initialized will be empty
		}

	}
	repoDir := filepath.Join(cwd, repoPath)

	config := &utils.Config{
		Objects:  fileObjects,
		InitTime: time.Now().Unix(),
		RepoDir:  repoDir,
	}
	return *config
}

func scaffold(config utils.Config) (err error) {
	configFilename := "/pt.config.json"

	//write config file
	jsonData := utils.ConfigToJSON(config)
	err = os.Mkdir(config.RepoDir, ownerReadWrite)
	err = os.WriteFile(filepath.Join(config.RepoDir, configFilename), jsonData, 0644)

	//scaffold repo
	for _, file := range config.Objects {
		dirName := utils.HashString(file.SourceFile)
		err = os.Mkdir(filepath.Join(config.RepoDir, dirName), ownerReadWrite)
		data := ""
		err = os.WriteFile(filepath.Join(config.RepoDir, dirName, utils.HashString(data)), []byte(data), ownerReadWrite)
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
