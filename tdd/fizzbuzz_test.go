package tdd

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	var testcases = []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{9, "Fizz"},
	}

	for _, c := range testcases {
		name := fmt.Sprintf("should return %s when input %d", c.expected, c.input)
		t.Run(name, func(t *testing.T) {
			r := fizzbuzz(c.input)

			if r != c.expected {
				t.Errorf("exected %s but got %s", c.expected, r)
			}
		})
	}
}
