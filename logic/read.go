package logic

import (
	"unicode"
)

type Prompt struct {
	text      string
	variables []string
	brackets  [2]rune
}

func (p *Prompt) FindVariables() {
	variables := BracketedText(p.text, p.brackets)
	for _, v := range variables {
		ValidateVariableIdentifier(v)
	}
	p.variables = variables
}

func BracketedText(promptText string, brackets [2]rune) []string {
	opener, closer := brackets[0], brackets[1]
	var variableArray []string = make([]string, 0) //different Data-Structure for future optimization

	for index := 0; index < len(promptText); index++ {

		element := rune(promptText[index])
		if element == opener {
			var end int
			for j := index + 1; j < len(promptText); j++ {
				subElement := rune(promptText[j])
				if subElement == closer {
					variableArray = append(variableArray, promptText[index+1:j])
					end = j + 1
					break
				}
			}
			index = end
		}
	}
	return variableArray
}

type Validator struct {
	message  string
	function func(rune) bool
}

func ValidateVariableIdentifier(variable string) bool {
	//figure out a better way to import validators
	a := Validator{"inappropriate character", unicode.IsLetter}
	validators := make([]Validator, 0)
	validators = append(validators, a)
	//
	for _, c := range variable {
		for _, validator := range validators {
			if !validator.function(c) {
				//some logging function x.log(validator.message)
				return false
			}
		}
	}
	return true

	//Future implementation of variable "strictness"
	//no-white-space
	//no-bracket-identifiers
}

func FromText(promptText string) Prompt {
	prompt := Prompt{
		text:     promptText,
		brackets: [2]rune{'{', '}'},
	}
	prompt.FindVariables()
	return prompt
}

// func encodePrompt(p *Prompt) string {

// }

// func decodePrompt(encodedString string) Prompt {

// }
