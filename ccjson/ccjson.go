// At this point, I kinda get the idea.
// Todo: escaping
// Non-string values

// The file on json.org was getting flagged as a security risk,
// So I decided to call it at step 4 of 5.

package ccjson

import (
	"ccjp/ccjson/parse"
	"fmt"
	"os"
)

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

	result, _, err := parse.ParseObject(pointer+1, endPointer, stringified_data)

	if err != nil {
		return nil, fmt.Errorf("error: parsejson %v", err)
	}

	return result, nil
}
