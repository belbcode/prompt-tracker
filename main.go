package main

import (
	functions "github.com/belbcode/prompt-tracker/logic/functions"
)

func main() {
	// fmt.Println("Prompt Tracker :)")
	// P := read.BracketedText("The {color} fox jumped over the {adjective} mouse", [2]rune{'{', '}'})
	// fmt.Println(P)
	functions.Watch()
}
