package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " PIKAchu, charmander, Bulbasaur ",
			expected: []string{"pikachu", "charmander", "bulbasaur"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "Pikachu, Pikachu, Pikachu",
			expected: []string{"pikachu", "pikachu", "pikachu"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) = %v; expected %v", c.input, actual, c.expected)
			t.Fail()
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				if word != expectedWord {
					t.Errorf("cleanInput(%q)[%d] = %q; expected %q", c.input, i, word, expectedWord)
				}
			}
		}
	}
}
