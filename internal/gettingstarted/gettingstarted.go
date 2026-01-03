package gettingstarted

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func FormatAbs(x int) string {
	msg := "The absolute value of %d is %d"
	return fmt.Sprintf(msg, x, abs(x))
}

func goFactorial(n, acc int) int {
	if n <= 0 {
		return acc
	}
	return goFactorial(n-1, n*acc)
}

func Factorial(n int) int {
	return goFactorial(n, 1)
}

func Factorial2(n int) int {
	acc := 1
	i := n
	for i > 0 {
		acc *= i
		i = i - 1
	}
	return acc
}

// Exercise 1: Write a function to compute the nth fibonacci number

// 0 and 1 are the first two numbers in the sequence,
// so we start the accumulators with those.
// At every iteration, we add the two numbers to get the next one.

func goFib(n, current, next int) int {
	if n <= 0 {
		return current
	}
	return goFib(n-1, next, current+next)
}

func Fib(n int) int {
	return goFib(n, 0, 1)
}
