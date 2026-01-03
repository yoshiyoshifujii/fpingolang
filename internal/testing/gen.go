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

func GenUnit[A any](a A) Gen[A] {
	return AsGen(state.Unit[state.RNG, A](a))
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

func Sequence[A any](actions []Gen[A]) Gen[[]A] {
	var ss []state.State[state.RNG, A]
	for _, action := range actions {
		ss = append(ss, AsState(action))
	}
	return AsGen(state.Sequence[state.RNG, A](ss))
}

func ListOfN[A any](n int, g Gen[A]) Gen[[]A] {
	var gs []Gen[A]
	for i := 0; i < n; i++ {
		gs = append(gs, g)
	}
	return Sequence[A](gs)
}

var (
	SmallInt = Choose(-10, 10)
	GenInt   = AsGen(state.State[state.RNG, int](state.RandInt))
)
