package main

import (
	"ccjp/ccjson"
	"testing"
)

func Test_step1_invalid(t *testing.T) {
	file_name := "./test_data/step1/invalid.json"
	result, err := ccjson.ParseJson(file_name)

	if err == nil {
		t.Errorf("no error found. Expected error.")
	}

	if result {
		t.Errorf("parsing result = %t; want false", result)
	}
}

func Test_step1_valid(t *testing.T) {
	file_name := "./test_data/step1/valid.json"
	result, err := ccjson.ParseJson(file_name)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	if !result {
		t.Errorf("invalid parsing result = %t; want true", result)
	}
}
