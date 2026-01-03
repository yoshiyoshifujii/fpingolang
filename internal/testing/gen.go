package testing

import "yoshiyoshifujii/fpingolang/internal/state"

type (
	Gen[A any] state.State[state.RNG, A]
)

func AsState[A any](g Gen[A]) state.State[state.RNG, A] {
	return state.State[state.RNG, A](g)
}

func AsGen[A any](st state.State[state.RNG, A]) Gen[A] {
	return Gen[A](st)
}

func NonNegativeInt() Gen[int] {
	return AsGen(state.Apply[state.RNG, int](state.NonNegativeInt))
}

func Choose(start, stopExclusive int) Gen[int] {
	return Map(NonNegativeInt(), func(n int) int {
		return start + n%(stopExclusive-start)
	})
}

func Map[A, B any](self Gen[A], f func(a A) B) Gen[B] {
	return AsGen(state.Map[state.RNG, A, B](AsState(self), f))
}

func Map2[A, B, C any](self Gen[A], other Gen[B], f func(a A, b B) C) Gen[C] {
	return AsGen(state.Map2[state.RNG, A, B, C](AsState(self), AsState(other), f))
}

func FlatMap[A, B any](self Gen[A], f func(a A) Gen[B]) Gen[B] {
	return AsGen(state.FlatMap[state.RNG, A, B](AsState(self), func(a A) state.State[state.RNG, B] {
		return AsState(f(a))
	}))
}

var (
	SmallInt Gen[int] = Choose(-10, 10)
)
