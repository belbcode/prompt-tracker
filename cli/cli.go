package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/belbcode/prompt-tracker/utils"
)

var arguments []string = []string{"init", "commit", "help"}

const projectDirname string = ".pt"
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

func runCli() {

	// switch {
	// 	case
	// }
}
