package state

type (
	State[S, A any] func(S) (A, S)
)

func Apply[S, A any](f func(S) (A, S)) State[S, A] {
	return f
}

func Unit[S, A any](a A) State[S, A] {
	return func(s S) (A, S) {
		return a, s
	}
}

func Run[S, A any](st State[S, A], s S) (A, S) {
	return st(s)
}

func FlatMap[S, A, B any](st State[S, A], f func(A) State[S, B]) State[S, B] {
	return func(s S) (B, S) {
		a, s1 := st(s)
		return f(a)(s1)
	}
}

func Map[S, A, B any](st State[S, A], f func(A) B) State[S, B] {
	return FlatMap[S, A, B](st, func(a A) State[S, B] {
		return Unit[S, B](f(a))
	})
}

func Map2[S, A, B, C any](st State[S, A], st2 State[S, B], f func(A, B) C) State[S, C] {
	return FlatMap[S, A, C](st, func(a A) State[S, C] {
		return FlatMap[S, B, C](st2, func(b B) State[S, C] {
			return Unit[S, C](f(a, b))
		})
	})
}
