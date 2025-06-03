package main

import "testing"

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
		input:    "no outer whitespace",
		expected: []string{"no", "outer", "whitespace"},
	},
	{
		input:    " ",
		expected: []string{},
	},
	{
		input:    "TEST TEST",
		expected: []string{"test", "test"},
	},
}
for _, c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("Whoops looks like there was an error with the expexted: %v and actual: %v", actual, c.expected)
	} 

	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord{
			t.Errorf("Whoops looks like word: %v and expected word: %v do not match", word, expectedWord)
		}
	}
}
}