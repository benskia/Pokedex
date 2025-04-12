package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "HELLO",
			expected: []string{"hello"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		result := cleanInput(c.input)

		if len(result) != len(c.expected) {
			t.Errorf("length of %+v != %+v", result, c.expected)
			return
		}

		for i := range result {
			resultWord := result[i]
			expectWord := c.expected[i]
			if resultWord != expectWord {
				t.Errorf("%s != %s", resultWord, expectWord)
			}
		}
	}
}
