package utils

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

type Config struct {
	InitTime int64
	RepoDir  string
	Objects  map[string]FileObject
}

type FileObject struct {
	SourceFile string
	RepoPath   string
	Properties FileInfo
	LastCommit string
}

type FileInfo struct {
	IsDir   bool
	Name    string
	Size    int64
	ModTime int64
}

func GetCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic("Critical Error, Unable to reach OS: " + err.Error())
	}
	return cwd
}

func IsExistingProject(cwd string) bool {
	fileInfo, err := os.Stat(cwd + "/.pt")
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func ReadConfig(cwd string) ([]byte, error) {
	var localConfigPath string = cwd + "/.pt/pt.config.json"
	bytes, err := os.ReadFile(localConfigPath)
	if err != nil {
		return make([]byte, 0), err
	}
	return bytes, nil
}

func ParseConfig(configJson []byte) Config {
	var parsedConfig Config
	err := json.Unmarshal(configJson, &parsedConfig)
	if err != nil {
		panic("Configuration file may have been corrupted. " + err.Error())
	}
	return parsedConfig
}

func GetConfig(cwd string) (Config, error) {
	var config Config
	if !IsExistingProject(cwd) {
		return config, errors.New("Cannot scaffold as project doesn't exist")
	}
	jsonConfig, err := ReadConfig(cwd)
	if err != nil {
		return config, err
	}
	config = ParseConfig(jsonConfig)
	return config, nil
}

func ConfigToJSON(config Config) []byte {
	jsonData, err := json.MarshalIndent(config, "\n", "	")
	if err != nil {
		panic(err)
	}
	return jsonData

}

func GetLatestTrackedFile(fileHash string) (latestCommit string, err error) {
	cwd := GetCwd()
	config, err := GetConfig(cwd)
	entries, err := os.ReadDir(config.Objects[fileHash].RepoPath)
	return filepath.Join(config.RepoDir, fileHash, entries[len(entries)-1].Name()), err
}

type FileInfoSlice []os.FileInfo

func (fis FileInfoSlice) Len() int           { return len(fis) }
func (fis FileInfoSlice) Less(i, j int) bool { return fis[i].ModTime().Before(fis[j].ModTime()) }
func (fis FileInfoSlice) Swap(i, j int)      { fis[i], fis[j] = fis[j], fis[i] }

func GetSortedFileInfos(dir string) ([]os.FileInfo, error) {
	dirEntries, err := getUnsortedDirEntries(dir)
	if err != nil {
		return nil, err
	}
	fileInfos := make([]os.FileInfo, len(dirEntries))
	for i, entry := range dirEntries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		fileInfos[i] = info
	}

	sort.Sort(FileInfoSlice(fileInfos))
	return fileInfos, nil
}

func getUnsortedDirEntries(dir string) ([]os.DirEntry, error) {
	fileInfos, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return fileInfos, nil
}

// func readDirModOrder(directory string) ([]os.DirEntry, error) {
// 	entries, err := os.ReadDir(directory)
// 	sort.Slice(entries, func(i, j int) bool {
// 		propertiesI, errI := entries[i].Info()
// 		propertiesJ, errJ := entries[j].Info()
// 		if errI != nil || errJ != nil {
// 			panic("Error accessing file properties" + errI.Error() + errJ.Error())
// 		}
// 		print(propertiesJ.ModTime().Unix() == propertiesI.ModTime().Unix())
// 		return propertiesI.ModTime().Unix() < propertiesJ.ModTime().Unix()
// 	})

// 	return entries, err
// }

func ObjectsFromFS(c Config) (map[string]FileObject, error) {
	trackedDirs, err := os.ReadDir(c.RepoDir)
	if err != nil {
		return nil, err
	}
	fileObjects := make(map[string]FileObject, len(trackedDirs))

	for _, dir := range trackedDirs {
		Object := c.Objects[dir.Name()]
		fileInfo, err := os.Stat(Object.SourceFile)
		if err != nil {
			return nil, err
		}
		allCommits, err := GetSortedFileInfos(Object.RepoPath)
		if err != nil {
			return nil, err
		}
		fo := FileObject{
			SourceFile: Object.SourceFile,
			RepoPath:   Object.RepoPath,
			Properties: ExtractFileInfo(fileInfo),
			LastCommit: allCommits[0].Name(),
		}
		fileObjects[dir.Name()] = fo
	}

	return fileObjects, nil
}

func UpdateConfig() (err error) {
	cwd := GetCwd()
	config, err := GetConfig(cwd)
	objectsFromFs, err := ObjectsFromFS(config)
	newConfig := Config{
		Objects:  objectsFromFs,
		RepoDir:  config.RepoDir,
		InitTime: config.InitTime,
	}
	jsonData := ConfigToJSON(newConfig)
	err = os.WriteFile(filepath.Join(config.RepoDir), jsonData, 0644)
	return err

}

func ExtractFileInfo(f fs.FileInfo) FileInfo {
	return FileInfo{
		IsDir:   f.IsDir(),
		Name:    f.Name(),
		Size:    f.Size(),
		ModTime: f.ModTime().Unix(),
	}
}
