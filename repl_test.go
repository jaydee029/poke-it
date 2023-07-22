package main

import (
	"testing"
)

func TestClean(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{
			input:  "hello world",
			output: []string{"hello", "world"},
		},
		{
			input:  "Hello world",
			output: []string{"hello", "world"},
		},
		{
			input:  "helloworld",
			output: []string{"helloworld"},
		},
	}

	for _, cse := range cases {
		actual := cleanText(cse.input)
		if len(actual) != len(cse.output) {
			t.Errorf("Lengths %v and %v are not equal", len(actual), len(cse.output))
			continue
		}

		for i := range actual {
			if actual[i] != cse.output[i] {
				t.Errorf("The words %s and %s are not same", actual[i], cse.output[i])
			}
		}
	}

}
