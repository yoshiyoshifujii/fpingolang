package gettingstarted

import (
	"sort"
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

func genSortedArray() testing2.Gen[[]int] {
	a := common.GenList[int](common.GenShortNumber())
	return testing2.Map[[]int, []int](a, func(a2 []int) []int {
		sort.Ints(a2)
		return a2
	})
}

func genUnsortedArray() testing2.Gen[[]int] {
	return testing2.FlatMap[int, []int](testing2.Choose(2, 20), func(n int) testing2.Gen[[]int] {
		return testing2.FlatMap[int, []int](testing2.Choose(1, 21), func(a int) testing2.Gen[[]int] {
			return testing2.FlatMap[int, []int](testing2.Choose(0, a), func(b int) testing2.Gen[[]int] {
				return testing2.FlatMap(testing2.ListOfN(n-2, common.GenShortNumber()), func(rest []int) testing2.Gen[[]int] {
					arr := make([]int, 0, n)
					arr = append(arr, a, b)
					arr = append(arr, rest...)
					return testing2.GenUnit(arr)
				})
			})
		})
	})
}

func TestIsSorted(t *testing.T) {
	common.PropSuiteTest(t, "IsSorted for sorted array case", genSortedArray(), func(ints []int) {
		actual := IsSorted[int](ints, func(i int, i2 int) bool {
			return i > i2
		})
		assert.True(t, actual)
	})
	common.PropSuiteTest(t, "IsSorted for unsorted array case", genUnsortedArray(), func(ints []int) {
		actual := IsSorted[int](ints, func(i int, i2 int) bool {
			return i > i2
		})
		assert.False(t, actual)
	})
}

func TestCurry(t *testing.T) {
	mulCurry := func(a int) func(int) int {
		return Curry[int, int, int](func(a1, b1 int) int {
			return a1 * b1
		})(a)
	}
	common.PropSuiteTest(t, "Curry", testing2.GenProduct(testing2.GenInt, testing2.GenInt), func(p testing2.Pair[int, int]) {
		n, m := p.First, p.Second
		expected := mulCurry(n)(m)
		assert.Equal(t, expected, n*m)
	})
}

func TestUncurry(t *testing.T) {
	mulUncurry := func(a, b int) int {
		return Uncurry[int, int, int](func(a1 int) func(int) int {
			return func(b1 int) int {
				return a1 * b1
			}
		})(a, b)
	}
	common.PropSuiteTest(t, "Uncurry", testing2.GenProduct(testing2.GenInt, testing2.GenInt), func(p testing2.Pair[int, int]) {
		n, m := p.First, p.Second
		expected := mulUncurry(n, m)
		assert.Equal(t, expected, n*m)
	})
}
