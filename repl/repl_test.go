package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	testCases := []struct{
		input		string
		expected	[]string
	}{
		{
			input: "  hello world    ", expected: []string{"hello", "world"},
		},
		{
			input: "Why here? ", expected: []string{"why", "here?"},
		},
		{
			input: "Uhhhhh!     Not now.", expected: []string{"uhhhhh!", "not", "now."},
		},
		{
			input: "	", expected: []string{},
		},
	}

	for _, tc := range testCases {
		actual := cleanInput(tc.input)

		if len(actual) != len(tc.expected) {
			t.Fatalf("expected %#v, got %#v", tc.expected, actual)
		}

		for i := range actual {
			if actual[i] != tc.expected[i] {
				t.Errorf("Expected: %#v, Got: %#v\n", tc.expected, actual)
			}
		}
	}
}