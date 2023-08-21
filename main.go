package main

import (
	"fmt"

	initRuntime "github.com/belbcode/prompt-tracker/cli/init"
)

func main() {
	arg, parseErr := initRuntime.ParseArgs()
	if parseErr != nil {
		fmt.Println(parseErr)
	}
	if arg == "init" {
		parentDir, initErr := initRuntime.Init()
		if initErr != nil {
			fmt.Println(initErr)
			return
		}
		initRuntime.CreateConfig(parentDir)

		// defer func() {
		// 	err := os.Remove(parentDir)
		// 	if err != nil {
		// 		fmt.Println("You may need to manually remove the .pt directory as there was an error in initialization and removing the dir.")
		// 	}
		// }()

		return
	}

	fmt.Println("Not yet implemented :(")
}
