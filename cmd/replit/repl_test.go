package main

import (
	"testing"
)

func TestCleanInput(t *testing.T){
	cases := []struct{
		input string
		expected [] string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "string one",
			expected: []string{"string", "one"},
		},
		{
			input: "string   two",
			expected: []string{"string", "two"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("Expected: %d words, Got: %d words", len(c.expected), len(actual))
			t.FailNow()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected: %s, got: %s", expectedWord, word)
				t.FailNow()
			}

		}
	}
}