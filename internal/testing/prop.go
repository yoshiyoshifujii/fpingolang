package testing

import (
	"fmt"
	"time"

	"yoshiyoshifujii/fpingolang/internal/state"
)

type (
	MaxSize      int
	TestCases    int
	SuccessCount int
	FailedCase   string
	ResultKind   string
	Result       struct {
		Kind      ResultKind
		Successes SuccessCount
		Failure   FailedCase
	}

	Prop func(MaxSize, TestCases, state.RNG) Result

	LazyList[A any] func() (A, LazyList[A])
)

var (
	ResultPassed    ResultKind = "Passed"
	ResultFalsified ResultKind = "Falsified"
	ResultProved    ResultKind = "Proved"
)

func (r Result) IsFalsified() bool {
	return r.Kind == ResultFalsified
}

func RandomLazyList[A any](g Gen[A], rng state.RNG) LazyList[A] {
	return func() (A, LazyList[A]) {
		value, next := state.Run[state.RNG, A](AsState(g), rng)
		return value, RandomLazyList[A](g, next)
	}
}

func ForAll[A any](as Gen[A], f func(A) bool) Prop {
	return func(maxSize MaxSize, testCases TestCases, rng state.RNG) Result {
		list := RandomLazyList[A](as, rng)
		for i := 0; i < int(testCases); i++ {
			value, next := list()
			list = next
			ok := true
			failure := ""
			func() {
				defer func() {
					if recovered := recover(); recovered != nil {
						ok = false
						failure = fmt.Sprintf("panic: %v", recovered)
					}
				}()
				ok = f(value)
				if !ok {
					failure = fmt.Sprintf("%v", value)
				}
			}()
			if !ok {
				return Result{
					Kind:      ResultFalsified,
					Successes: SuccessCount(i),
					Failure:   FailedCase(failure),
				}
			}
		}
		return Result{
			Kind:      ResultPassed,
			Successes: SuccessCount(testCases),
		}
	}
}

func (p Prop) Check() Result {
	return p(100, 100, state.NewRNGSimple(uint64(time.Now().UnixMilli())))
}
