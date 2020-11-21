package tdd

import "fmt"

func fizzbuzz(n int) string {
	if n == 3 {
		return "Fizz"
	}

	return fmt.Sprint(n)
}
