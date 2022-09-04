package usecase

import (
	"errors"
	"unicode"
)

func UnpackString(input string) (string, error) {
	var result []rune
	normalizedInput := []rune(input)

	hasPreviousSymbInit := false
	var previousSymb rune
	for _, symb := range normalizedInput {
		if unicode.IsDigit(symb) {
			if !hasPreviousSymbInit {
				return "", errors.New("Incorrect string")
			}
			for idx := 0; idx < int(symb-'0')-1; idx++ {
				result = append(result, previousSymb)
			}
			continue
		}

		result = append(result, symb)
		previousSymb = symb
		hasPreviousSymbInit = true
	}
	return string(result), nil
}
