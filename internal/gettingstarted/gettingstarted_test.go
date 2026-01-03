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
		return testing2.FlatMap(testing2.ListOfN(n, common.GenShortNumber()), func(list []int) testing2.Gen[[]int] {
			var arr []int
			for num, i := range list {
				if i%2 == 0 {
					arr = append(arr, num+100)
				}
				arr = append(arr, num-100)
			}
			return testing2.GenUnit(arr)
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
