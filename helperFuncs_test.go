package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello   ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello   ",
			expected: []string{"hello"},
		},
		{
			input:    "",
			expected: []string{""},
		},
		{
			input:    "hello, word, osiem, siedem",
			expected: []string{"hello,", "word,", "osiem,", "siedem"},
		},
		{
			input:    "BIG, small, small, BIG",
			expected: []string{"big,", "small,", "small,", "big"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Fatalf("Length of slices do not match. Expected: %d, got: %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Fatalf("Values in slices do not match. Expected: %s, got: %s", expectedWord, word)
			}
		}
	}
}
