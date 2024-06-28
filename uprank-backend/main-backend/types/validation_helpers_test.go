package types

import (
	"reflect"
	"testing"
)

func TestFindNilFields(t *testing.T) {
	type TestStruct struct {
		Field1 *string
		Field2 *int
		Field3 *float64
		Field4 *TestStruct
	}
	// Test case 1: All fields are nil
	field1 := "test"
	field2 := 42
	field3 := 3.14
	tests := []struct {
		input    interface{}
		expected []string
	}{
		{
			input: &TestStruct{
				Field1: nil,
				Field2: nil,
				Field3: nil,
				Field4: nil,
			},
			expected: []string{"Field1", "Field2", "Field3", "Field4"},
		},
		{
			input: &TestStruct{
				Field1: &field1,
				Field2: nil,
				Field3: nil,
				Field4: nil,
			},
			expected: []string{"Field2", "Field3", "Field4"},
		},
		{
			input: &TestStruct{
				Field1: &field1,
				Field2: &field2,
				Field3: nil,
				Field4: nil,
			},
			expected: []string{"Field3", "Field4"},
		},
		{
			input: &TestStruct{
				Field1: &field1,
				Field2: &field2,
				Field3: &field3,
				Field4: nil,
			},
			expected: []string{"Field4"},
		},
		{
			input: &TestStruct{
				Field1: &field1,
				Field2: &field2,
				Field3: &field3,
				Field4: &TestStruct{},
			},
			expected: []string{},
		},
	}

	for _, test := range tests {
		result := findNilFields(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestGetNumFields(t *testing.T) {
	type TestStruct struct {
		Field1 *string
		Field2 *int
		Field3 *float64
		Field4 *TestStruct
	}
	type EmptyTestStruct struct{}
	tests := []struct {
		input    interface{}
		expected int
	}{
		{
			input: &TestStruct{
				Field1: nil,
				Field2: nil,
				Field3: nil,
				Field4: nil,
			},
			expected: 4,
		},
		{
			input:    &EmptyTestStruct{},
			expected: 0,
		},
	}

	for _, test := range tests {
		result := getNumFields(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}

}
