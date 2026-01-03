package datastructures

import (
	"testing"
	"yoshiyoshifujii/fpingolang/internal/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func arrayToList[A any](a []A) List[A] {
	return ListApply[A](a...)
}

func TestListTail(t *testing.T) {
	common.PropSuiteTest(t, "List.Tail", common.GenIntList(), func(arr []int) {
		switch a := arrayToList(arr).(type) {
		case ListNil:
			assert.Panics(t, func() {
				ListTail[int](a)
			})
		case ListCons[int]:
			x, xs := a.Head, a.Tail
			assert.Equal(t, xs, ListTail[int](Cons[int](x, xs)))
		default:
			require.Fail(t, "unexpected type", a)
		}
	})
}
