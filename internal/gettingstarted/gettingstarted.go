package gettingstarted

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func formatAbs(x int) string {
	msg := "The absolute value of %d is %d"
	return fmt.Sprintf(msg, x, abs(x))
}
