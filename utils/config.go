package utils

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	InitTime     int64
	RepoDir      string
	TrackedFiles map[string]FileObject
}

type FileObject struct {
	OriginalPath string
	RepoPath     string
	Properties   FileInfo
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

func GetLatestTrackedFile(fileHash string) (latestCommit string, err error) {
	cwd := GetCwd()
	config, err := GetConfig(cwd)
	entries, err := os.ReadDir(config.TrackedFiles[fileHash].RepoPath)
	return filepath.Join(config.RepoDir, fileHash, entries[len(entries)-1].Name()), err
}
