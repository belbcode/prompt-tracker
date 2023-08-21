package functions

import (
	"fmt"

	logic "github.com/belbcode/prompt-tracker/logic"
)

func commit(promptText string) {
	prompt := logic.FromText(promptText)
	fmt.Println(prompt)
	// prompt.variables
}
