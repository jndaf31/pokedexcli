package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{input: " hello world ", expected: []string{"hello", "world"}},
		{input: "  hello   WORLD   ", expected: []string{"hello", "world"}},
		{input: "HELLO WoRLd", expected: []string{"hello", "world"}},
		{input: "hello the world", expected: []string{"hello", "the", "world"}},
		{input: "helloworld", expected: []string{"helloworld"}},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length doesn't match. expected: %v actual: %v", len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			wrd := actual[i]
			expWrd := c.expected[i]

			if wrd != expWrd {
				t.Errorf("results don't match. expected: %v actual: %v", expWrd, wrd)
			}
		}
	}
}
