package gettingstarted

import (
	"testing"
	"yoshiyoshifujii/fpingolang/internal/common"

	"github.com/stretchr/testify/assert"
)

func Test_FormatAbs(t *testing.T) {
	// given
	in := -42

	// when
	actual := FormatAbs(in)

	// then
	assert.Equal(t, "The absolute value of -42 is 42", actual)
}

func Test_Factorial(t *testing.T) {
	// given
	in := 5

	// when
	actual := Factorial(in)

	// then
	assert.Equal(t, 120, actual)
}

func Test_FactorialProperties(t *testing.T) {
	for n := 0; n <= 10; n++ {
		expected := 1
		for i := 1; i <= n; i++ {
			expected *= i
		}
		assert.Equal(t, expected, Factorial(n))
	}
}

func Test_Factorial2(t *testing.T) {
	for n := 0; n <= 10; n++ {
		expected := 1
		for i := 1; i <= n; i++ {
			expected *= i
		}
		assert.Equal(t, expected, Factorial2(n))
	}
}

func Test_Fib(t *testing.T) {
	for i, expected := range common.TheFirst21FibonacciNumbers {
		assert.Equal(t, expected, Fib(i))
	}
}
