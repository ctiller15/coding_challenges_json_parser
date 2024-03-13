package ccjson

import (
	"fmt"
	"os"
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
	var val string

	for i < endptr {

		if len(key) > 0 && len(val) > 0 {
			result[key] = val
			key = ""
			val = ""
		}

		// Can likely be extracted into a parseline func.
		if len(key) > 0 && data[i] == '"' {
			new_val, pos, err := extractValue(i+1, endptr, data)

			if err != nil {
				return result, i, fmt.Errorf("error: parseObject: %v", err)
			}

			val = new_val
			i = pos
		}

		if len(key) == 0 && data[i] == '"' {
			new_key, pos, err := extractKey(i+1, endptr, data)

			if err != nil {
				return result, i, fmt.Errorf("error: parseObject: %v", err)
			}

			key = new_key
			i = pos
		}

		if data[i] == '}' {
			return result, i, nil
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
