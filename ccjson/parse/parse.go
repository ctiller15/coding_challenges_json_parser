package parse

import (
	"ccjp/ccjson/extract"
	"fmt"
	"unicode"
)

func parseArray(ptr int, endptr int, data []rune) ([]string, int, error) {
	// Start with string arrays. Can get more complex later.
	result := []string{}

	i := ptr
	var val string

	for i < endptr {
		if len(val) > 0 {
			result = append(result, val)
			val = ""
		}

		if len(val) == 0 {
			if data[i] == '\'' {
				return result, i, fmt.Errorf("invalid value")
			} else if data[i] == '"' {
				new_val, pos, err := extract.ExtractValue(i+1, endptr, data)

				if err != nil {
					return result, i, fmt.Errorf("parseArray: %v", err)
				}

				val = new_val
				i = pos
			}
		}

		if data[i] == ']' {
			return result, i, nil
		}

		i++
	}

	return result, i, fmt.Errorf("invalid value")
}

func ParseObject(ptr int, endptr int, data []rune) (map[string]interface{}, int, error) {
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
				new_val, pos, err := extract.ExtractValue(i+1, endptr, data)

				if err != nil {
					return result, i, fmt.Errorf("error: parseObject: %v", err)
				}

				val = new_val
				i = pos
			} else if data[i] == 't' || data[i] == 'f' {
				new_val, pos, err := extract.ExtractValueBool(i, endptr, data)

				if err != nil {
					return result, i, fmt.Errorf("error: parseObject: %v", err)
				}

				val = new_val
				i = pos
			} else if data[i] == 'n' {
				is_nil, pos, err := extract.ExtractValueNull(i, endptr, data)

				if err != nil {
					return result, i, fmt.Errorf("error: parseObject: %v", err)
				}

				val_is_nil = is_nil

				i = pos
			} else if unicode.IsDigit(data[i]) {
				digit, pos, err := extract.ExtractValueDigit(i, endptr, data)

				if err != nil {
					return result, i, fmt.Errorf("error: parseObject: %v", err)
				}

				val = digit
				i = pos
			} else if data[i] == '[' {
				arr, pos, err := parseArray(i, endptr, data)

				if err != nil {
					return result, i, fmt.Errorf("error: parseObject %v", err)
				}

				val = arr
				i = pos
			} else if data[i] == '{' {
				obj, pos, err := ParseObject(i, endptr, data)

				if err != nil {
					return result, i, fmt.Errorf("error: parseObject %v", err)
				}
				val = obj
				i = pos
			}

		}

		if len(key) == 0 && unicode.IsLetter(data[i]) {
			return result, i, fmt.Errorf("property keys must be doublequoted")
		}

		if len(key) == 0 && data[i] == '"' {
			// Starting new key value pair
			commaActive = false
			new_key, pos, err := extract.ExtractKey(i+1, endptr, data)

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
				return result, i + 1, nil
			}
		}

		i++
	}

	return result, i, fmt.Errorf("invalid json. no closing brace")
}
