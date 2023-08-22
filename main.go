package main

import (
	"fmt"
	"log"

	initRuntime "github.com/belbcode/prompt-tracker/cli/init"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
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
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			initRuntime.Init(args)
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
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
