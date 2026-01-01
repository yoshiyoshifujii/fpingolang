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

func goFunc(n, acc int) int {
	if n <= 0 {
		return acc
	}
	return goFunc(n-1, n*acc)
}

func Factorial(n int) int {
	return goFunc(n, 1)
}
