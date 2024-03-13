package main

import (
	"ccjp/ccjson"
	"reflect"
	"testing"
)

func Test_step1_invalid(t *testing.T) {
	file_name := "./test_data/step1/invalid.json"
	result, err := ccjson.ParseJson(file_name)

	if err == nil {
		t.Errorf("no error found. Expected error.")
	}

	if result != nil {
		t.Errorf("parsing result = %v; want nil", result)
	}
}

func Test_step1_valid(t *testing.T) {
	file_name := "./test_data/step1/valid.json"
	result, err := ccjson.ParseJson(file_name)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	if result == nil {
		t.Errorf("invalid parsing result = %v; want object", result)
	}
}

func Test_step2_valid_01(t *testing.T) {
	file_name := "./test_data/step2/valid.json"
	result, err := ccjson.ParseJson(file_name)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	expected := map[string]interface{}{
		"key": "value",
	}

	eq := reflect.DeepEqual(result, expected)

	if !eq {
		t.Errorf("result does not match expected json.")
	}
}

func Test_step2_valid_02(t *testing.T) {
	file_name := "./test_data/step2/valid2.json"
	result, err := ccjson.ParseJson(file_name)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	expected := map[string]interface{}{
		"key":  "value",
		"key2": "value",
	}

	eq := reflect.DeepEqual(result, expected)

	if !eq {
		t.Errorf("result does not match expected json.")
	}
}
