package gettingstarted

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
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
	params := gopter.DefaultTestParameters()
	params.MinSuccessfulTests = 100
	properties := gopter.NewProperties(params)

	properties.Property("Factorial(n) matches the product 1..n", prop.ForAll(
		func(n int) bool {
			expected := 1
			for i := 1; i <= n; i++ {
				expected *= i
			}
			return Factorial(n) == expected
		},
		gen.IntRange(0, 10),
	))

	properties.TestingRun(t)
}

func Test_Factorial2(t *testing.T) {
	params := gopter.DefaultTestParameters()
	params.MinSuccessfulTests = 100
	properties := gopter.NewProperties(params)

	properties.Property("Factorial(n) matches the product 1..n", prop.ForAll(
		func(n int) bool {
			expected := 1
			for i := 1; i <= n; i++ {
				expected *= i
			}
			return Factorial2(n) == expected
		},
		gen.IntRange(0, 10),
	))

	properties.TestingRun(t)
}
