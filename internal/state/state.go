package state

type (
	State[S any, A any] func(S) (A, S)

	stateImpl[S any, A any] struct {
		underlying State[S, A]
	}
)

func FlatMap[S, A, B any](st State[S, A], f func(A) State[S, B]) State[S, B] {
	return func(s S) (B, S) {
		a, s1 := st(s)
		return f(a)(s1)
	}
}
