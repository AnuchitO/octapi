package tdd

import (
	"fmt"
	"testing"
)

var tests = []struct {
	input    int
	expected string
}{
	{1, "1"},
	{2, "2"},
	{3, "Fizz"},
	{4, "4"},
	{5, "Buzz"},
	{6, "Fizz"},
}

func TestFizzBuzz(t *testing.T) {
	for _, test := range tests {
		name := fmt.Sprintf("should return %s when input %d", test.expected, test.input)
		t.Run(name, func(t *testing.T) {
			r := fizzbuzz(test.input)

			if r != test.expected {
				t.Errorf("exected %s but got %s", test.expected, r)
			}
		})
	}
}
