package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/belbcode/prompt-tracker/cli/commit"
	initRuntime "github.com/belbcode/prompt-tracker/cli/init"
	"github.com/belbcode/prompt-tracker/utils"
	"github.com/spf13/cobra"
)

var (
	projectFlag bool
	targetFlag  bool
	rootCmd     = &cobra.Command{
		Use:   "pt",
		Short: "A prompt tracking program.",
		Long: `This program tracks relevant metadata of text files containing prompts.
It supports development of complex prompt-chains.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to PromptTracker!\nFeel free to ask for --help.\nOr jump right into it if you know what you're doing.")
		},
	}
	initCmd = &cobra.Command{
		Use:   "init [files to track]",
		Short: "Initializes Project.",
		Run: func(cmd *cobra.Command, args []string) {
			var cwd string = utils.GetCwd()
			if targetFlag {
				cwd, err := filepath.Abs(args[0])
				print(cwd, args[0])
				if err != nil {
					panic(err)
				}
				fi, err := os.Stat(cwd)
				if err != nil {
					panic(err)
				}
				if !fi.IsDir() {
					panic("Need specify a directory")
				}

			}
			if projectFlag {
				entries, err := os.ReadDir(cwd)
				if err != nil {
					panic("Failed to access directory, may lack permissions.")
				}
				args = utils.MapToString[fs.DirEntry](entries, func(entry fs.DirEntry) string {
					return entry.Name()
				})
			}
			initRuntime.Init(args)
		},
	}
	commitCmd = &cobra.Command{
		Use:   "commit [files to commit]",
		Short: "Commits changes into repo.",
		Run: func(cmd *cobra.Command, args []string) {
			commit.CommitFiles(args)
		},
	}
)

func init() {
	rootCmd.Flags().BoolVarP(&projectFlag, "projectFlag", "p", true, "set arguments to all files in the project directory")
	initCmd.Flags().BoolVarP(&targetFlag, "targetFlag", "t", false, "set project directory")
	rootCmd.AddCommand(initCmd, commitCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	// arg, parseErr := initRuntime.ParseArgs()
	// if parseErr != nil {
	// 	fmt.Println(parseErr)
	// }
	// if arg == "init" {
	// 	parentDir, initErr := initRuntime.Init()
	// 	if initErr != nil {
	// 		fmt.Println(initErr)
	// 		return
	// 	}
	// 	initRuntime.CreateConfig(parentDir)
	// 	cwd := utils.GetCwd()
	// 	config, err := utils.GetConfig(cwd)
	// 	if err != nil {
	// 		fmt.Println(":(")
	// 	}
	// 	initRuntime.Scaffold(config)

	// 	// defer func() {
	// 	// 	err := os.Remove(parentDir)
	// 	// 	if err != nil {
	// 	// 		fmt.Println("You may need to manually remove the .pt directory as there was an error in initialization and removing the dir.")
	// 	// 	}
	// 	// }()

	// 	return
	// }
	// if arg == "commit" {
	// 	cwd := utils.GetCwd()
	// 	config, err := utils.GetConfig(cwd)
	// 	if err != nil {
	// 		fmt.Println("There was an error commiting the file:", err)
	// 		return
	// 	}
	// 	for _, file := range config.TrackedFiles {
	// 		commit.CommitFile(file.OriginalPath)
	// 	}

	// }

	// fmt.Println("Not yet implemented :(")
}
