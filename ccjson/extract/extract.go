package extract

import (
	"fmt"
	"slices"
	"strconv"
	"unicode"
)

func ExtractKey(ptr int, endptr int, data []rune) (string, int, error) {
	i := ptr
	var builtKey []rune

	for i < endptr {
		if data[i] == '"' {
			return string(builtKey), i, nil
		}

		builtKey = append(builtKey, data[i])
		i++
	}

	// If it hits the end and does not return, error.
	return string(builtKey), i, fmt.Errorf("invalid json, no key terminator")
}

func ExtractValueDigit(ptr int, endptr int, data []rune) (int, int, error) {
	i := ptr
	var builtKey []rune
	var terminators = []rune{' ', ',', '}', '\n'}

	for i < endptr {
		if unicode.IsDigit(data[i]) {
			builtKey = append(builtKey, data[i])
		} else if slices.Contains(terminators, data[i]) {
			num, err := strconv.Atoi(string(builtKey))

			if err != nil {
				return 0, i, fmt.Errorf("extractValueDigit: %v", err)
			}

			return num, i, nil
		}

		i++
	}

	return 0, i, fmt.Errorf("invalid value")
}

func ExtractValueNull(ptr int, endptr int, data []rune) (bool, int, error) {
	i := ptr
	var builtValue []rune

	characters := []rune{'n', 'u', 'l'}
	nullSlice := []rune{'n', 'u', 'l', 'l'}

	for i < endptr {
		if !slices.Contains(characters, data[i]) {
			return false, i, fmt.Errorf("invalid value, must be null")
		}

		builtValue = append(builtValue, data[i])
		if slices.Compare(builtValue, nullSlice) == 0 {
			return true, i, nil
		}

		i++
	}

	return false, i, fmt.Errorf("invalid value")
}

func ExtractValueBool(ptr int, endptr int, data []rune) (bool, int, error) {
	i := ptr
	var builtValue []rune

	characters := []rune{'t', 'r', 'u', 'e', 'f', 'a', 'l', 's'}
	trueSlice := []rune{'t', 'r', 'u', 'e'}
	falseSlice := []rune{'f', 'a', 'l', 's', 'e'}

	for i < endptr {

		if !slices.Contains(characters, data[i]) {
			return false, i, fmt.Errorf("invalid value, must be true or false")
		}

		builtValue = append(builtValue, data[i])
		if slices.Compare(builtValue, trueSlice) == 0 {
			return true, i, nil
		}

		if slices.Compare(builtValue, falseSlice) == 0 {
			return false, i, nil
		}

		i++
	}

	return false, i, fmt.Errorf("invalid value")
}

// Start by extracting just string values. We can handle more complex json
// At a later step.
func ExtractValue(ptr int, endptr int, data []rune) (string, int, error) {
	i := ptr
	var builtValue []rune

	for i < endptr {
		if data[i] == '"' {
			return string(builtValue), i, nil
		}

		builtValue = append(builtValue, data[i])
		i++
	}

	// If it hits the end and does not return, error.
	return string(builtValue), i, fmt.Errorf("invalid json, no value terminator")
}
