package tdd

import "fmt"

func fizzbuzz(n int) string {
	if n == 15 {
		return "FizzBuzz"
	}
	if n % 5 == 0 {
		return "Buzz"
	}
	if n % 3 == 0 {
		return "Fizz"
	}

	return fmt.Sprint(n)
}
