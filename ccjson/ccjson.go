package ccjson

import (
	"fmt"
	"os"
)

func ParseJson(fileName string) (bool, error) {
	file_data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	stringified_data := string(file_data)

	if len(stringified_data) == 0 {
		return false, fmt.Errorf("invalid json. Must contain more than 0 characters")
	}

	var brace_count int

	for i, s := range stringified_data {
		if i == 0 {
			// Ensure first character is a brace.
			if s != '{' {
				return false, fmt.Errorf("json must begin with a curly brace \\'{\\'")
			}

			brace_count++
		}

		if s == '}' {
			brace_count--
		}
		fmt.Println(s)
	}

	return true, nil
}
