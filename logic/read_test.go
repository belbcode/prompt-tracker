package logic_test

import (
	"testing"

	read "github.com/belbcode/prompt-tracker/logic"
)

func TestBracketedText(t *testing.T) {
	test := "The {color} fox jumped over the {adjective} mouse"
	brackets := [2]rune{'{', '}'}
	results := read.BracketedText(test, brackets)
	correctAnswer := [2]string{"color", "adjective"}
	for i := range results {
		if results[i] != correctAnswer[i] {
			t.Error("expected ", correctAnswer[i], " got ", results[i])
		}
	}
}

func TestValidateVariableIdentfiers(t *testing.T) {
	tests := [7]string{"greeting", "input", "  sd   a", "{sdsda}", "aas{aaa", ":&8229", "test"}
	corrects := [7]bool{true, true, false, false, false, false, true}
	for index, test := range tests {
		result := read.ValidateVariableIdentifier(test)
		if result != corrects[index] {
			t.Error("expected ", corrects[index], " got ", result)
		}
	}
}
