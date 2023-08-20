package functions

import logic "github.com/belbcode/prompt-tracker/logic"

func commit(promptText string) {
	prompt := logic.FromText(promptText)
	prompt.variables

}
