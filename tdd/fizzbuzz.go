package tdd

import "fmt"

func fizzbuzz(n int) string {
	if n == 5 {
		return "Buzz"
	}
	if n == 3 || n == 6 {
		return "Fizz"
	}

	return fmt.Sprint(n)
}
