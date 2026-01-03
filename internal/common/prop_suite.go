package common

import (
	"testing"

	proptesting "yoshiyoshifujii/fpingolang/internal/testing"
)

func PropSuiteTest[A any](t *testing.T, name string, a proptesting.Gen[A], f func(A)) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		t.Helper()
		g := func(a A) bool {
			f(a)
			return true
		}
		prop := proptesting.ForAll[A](a, g)
		result := prop.Check()
		if result.IsFalsified() {
			t.Fatalf("%s falsified after %d successes: %s", name, result.Successes, result.Failure)
		}
	})
}
