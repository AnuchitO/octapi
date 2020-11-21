package tdd

import "fmt"

func fizzbuzz(n int) string {
	if n == 5 || n == 10 {
		return "Buzz"
	}
	if n % 3 == 0 {
		return "Fizz"
	}

	return fmt.Sprint(n)
}
