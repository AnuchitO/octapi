package tdd

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("should return 1 when input 1", func(t *testing.T) {
		r := fizzbuzz(1)

		if r != "1" {
			t.Errorf("exected %s but got %s", "1", r)
		}
	})

	t.Run("should return 2 when input 2", func(t *testing.T) {
		r := fizzbuzz(2)

		if r != "2" {
			t.Errorf("exected %s but got %s", "2", r)
		}
	})

	t.Run("should return Fizz when input 3", func(t *testing.T) {
		r := fizzbuzz(3)

		if r != "Fizz" {
			t.Errorf("exected %s but got %s", "Fizz", r)
		}
	})

	t.Run("should return 4 when input 4", func(t *testing.T) {
		r := fizzbuzz(4)

		if r != "4" {
			t.Errorf("exected %s but got %s", "4", r)
		}
	})

	t.Run("should return Buzz when input 5", func(t *testing.T) {
		r := fizzbuzz(5)

		if r != "Buzz" {
			t.Errorf("exected %s but got %s", "Buzz", r)
		}
	})
}
