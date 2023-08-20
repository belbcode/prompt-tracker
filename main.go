package main

import (
	logic "github.com/belbcode/prompt-tracker/logic"
)

func main() {
	// fmt.Println("Prompt Tracker :)")
	// P := read.BracketedText("The {color} fox jumped over the {adjective} mouse", [2]rune{'{', '}'})
	// fmt.Println(P)
	logic.Watch()
}
