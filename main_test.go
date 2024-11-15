package validbraces

import (
	"testing"

)

func TestIsValid(t *testing.T) {
	testCase := []struct {
		input    string
		expected bool
	}{
		{"(){}[]", true},
		{"([{}])", true},
		{"(}", false},
		{"[(])", false},
		{"[({})](]", false},
	}

	for _, tc := range testCase {
		res := ValidBraces(tc.input)
		
		if res != tc.expected {
			t.Errorf("Expected %v but got %v", tc.expected, res)
		}
	}
}
