package ccjson

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
)

func extractKey(ptr int, endptr int, data []rune) (string, int, error) {
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

func extractValueDigit(ptr int, endptr int, data []rune) (int, int, error) {
	i := ptr
	var builtKey []rune
	var terminators = []rune{' ', ',', '}', '\n'}

	for i < endptr {
		if unicode.IsDigit(data[i]) {
			builtKey = append(builtKey, data[i])
		} else if slices.Contains(terminators, data[i]) {
			fmt.Printf("%c\n", data[i])
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

func extractValueNull(ptr int, endptr int, data []rune) (bool, int, error) {
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

func extractValueBool(ptr int, endptr int, data []rune) (bool, int, error) {
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
func extractValue(ptr int, endptr int, data []rune) (string, int, error) {
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

func parseObject(ptr int, endptr int, data []rune) (map[string]interface{}, int, error) {
	result := map[string]interface{}{}

	i := ptr

	var key string
	var val interface{}
	var commaActive bool
	var pairCount int
	var val_is_nil bool

	for i < endptr {
		if data[i] == ',' {
			commaActive = true
		}

		if len(key) > 0 && val != nil {
			result[key] = val
			key = ""
			val = nil
			pairCount++
		} else if len(key) > 0 && val_is_nil {
			result[key] = nil
			key = ""
			val = nil
			pairCount++
			val_is_nil = false
		}

		// Can likely be extracted into a parseline func.
		if len(key) > 0 {
			if data[i] == '"' {
				new_val, pos, err := extractValue(i+1, endptr, data)

				if err != nil {
					return result, i, fmt.Errorf("error: parseObject: %v", err)
				}

				val = new_val
				i = pos
			} else if data[i] == 't' || data[i] == 'f' {
				new_val, pos, err := extractValueBool(i, endptr, data)

				if err != nil {
					return result, i, fmt.Errorf("error: parseObject: %v", err)
				}

				val = new_val
				i = pos
			} else if data[i] == 'n' {
				is_nil, pos, err := extractValueNull(i, endptr, data)

				if err != nil {
					return result, i, fmt.Errorf("error: parseObject: %v", err)
				}

				val_is_nil = is_nil

				i = pos
			} else if unicode.IsDigit(data[i]) {
				digit, pos, err := extractValueDigit(i, endptr, data)

				if err != nil {
					return result, i, fmt.Errorf("error: parseObject: %v", err)
				}

				val = digit
				i = pos
			}
		}

		if len(key) == 0 && unicode.IsLetter(data[i]) {
			return result, i, fmt.Errorf("property keys must be doublequoted")
		}

		if len(key) == 0 && data[i] == '"' {
			// Starting new key value pair
			commaActive = false
			new_key, pos, err := extractKey(i+1, endptr, data)

			if err != nil {
				return result, i, fmt.Errorf("error: parseObject: %v", err)
			}

			key = new_key
			i = pos
		}

		if data[i] == '}' {
			if commaActive {
				return result, i, fmt.Errorf("invalid json. trailing comma")
			} else {
				return result, i, nil
			}
		}

		i++
	}

	return result, i, fmt.Errorf("invalid json. no closing brace")
}

func ParseJson(fileName string) (map[string]interface{}, error) {
	file_data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	stringified_data := []rune(string(file_data))

	if len(stringified_data) == 0 {
		return nil, fmt.Errorf("invalid json. Must contain more than 0 characters")
	}

	var pointer int
	endPointer := len(stringified_data)

	if stringified_data[pointer] != '{' {
		return nil, fmt.Errorf("json must begin with a curly brace \\'{\\'")
	}

	result, index, err := parseObject(pointer+1, endPointer, stringified_data)

	fmt.Println(index)

	if err != nil {
		return nil, fmt.Errorf("error: parsejson %v", err)
	}

	return result, nil
}
