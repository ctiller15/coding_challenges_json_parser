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

func Test_step2_invalid_01(t *testing.T) {
	file_name := "./test_data/step2/invalid.json"
	_, err := ccjson.ParseJson(file_name)

	if err == nil {
		t.Errorf("No error detected.")
	}
}

func Test_step2_invalid_02(t *testing.T) {
	file_name := "./test_data/step2/invalid2.json"
	_, err := ccjson.ParseJson(file_name)

	if err == nil {
		t.Errorf("No error detected.")
	}
}

func Test_step3_invalid(t *testing.T) {
	file_name := "./test_data/step3/invalid.json"
	_, err := ccjson.ParseJson(file_name)

	if err == nil {
		t.Errorf("No error detected.")
	}
}

func Test_step_3_valid(t *testing.T) {
	file_name := "./test_data/step3/valid.json"
	result, err := ccjson.ParseJson(file_name)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	expected := map[string]interface{}{
		"key1": true,
		"key2": false,
		"key3": nil,
		"key4": "value",
		"key5": 101,
	}

	eq := reflect.DeepEqual(result, expected)

	if !eq {
		t.Errorf("result does not match expected json.")
	}
}

func Test_step_4_invalid(t *testing.T) {
	file_name := "./test_data/step4/invalid.json"
	_, err := ccjson.ParseJson(file_name)

	if err == nil {
		t.Errorf("Expected error.")
	}
}

func Test_step_4_valid(t *testing.T) {
	file_name := "./test_data/step4/valid.json"
	result, err := ccjson.ParseJson(file_name)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	expected := map[string]interface{}{
		"key":   "value",
		"key-n": 101,
		"key-o": map[string]interface{}{},
		"key-l": []string{},
	}

	eq := reflect.DeepEqual(result, expected)

	if !eq {
		t.Errorf("result does not match expected json.")
	}
}

func Test_step_4_valid2(t *testing.T) {
	file_name := "./test_data/step4/valid2.json"
	result, err := ccjson.ParseJson(file_name)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	expected := map[string]interface{}{
		"key":   "value",
		"key-n": 101,
		"key-o": map[string]interface{}{
			"inner key": "inner value",
		},
		"key-l": []string{"list value"},
	}

	eq := reflect.DeepEqual(result, expected)

	if !eq {
		t.Errorf("result does not match expected json.")
	}
}
