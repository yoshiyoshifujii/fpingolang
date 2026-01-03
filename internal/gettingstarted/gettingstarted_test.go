package gettingstarted

import (
	"testing"
	"yoshiyoshifujii/fpingolang/internal/common"
	testing2 "yoshiyoshifujii/fpingolang/internal/testing"

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
	common.PropSuiteTest(t, "Factorial matches the product 1..n", testing2.SmallInt, func(n int) {
		expected := 1
		for i := 1; i <= n; i++ {
			expected *= i
		}
		assert.Equal(t, expected, Factorial(n))
	})
}

func Test_Factorial2(t *testing.T) {
	common.PropSuiteTest(t, "Factorial matches the product 1..n", testing2.SmallInt, func(n int) {
		expected := 1
		for i := 1; i <= n; i++ {
			expected *= i
		}
		assert.Equal(t, expected, Factorial2(n))
	})
}

func Test_Fib(t *testing.T) {
	common.PropSuiteTest(t, "Fib", common.GenLengthOfFibonacciSeq(), func(n int) {
		expected := common.TheFirst21FibonacciNumbers[n]
		assert.Equal(t, expected, Fib(n))
	})
}
