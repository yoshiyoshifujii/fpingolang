package datastructures

type (
	List[A any] interface {
		list()
	}

	ListNil         struct{}
	ListCons[A any] struct {
		Head A
		Tail List[A]
	}
)

func (n ListNil) list() {
	panic("implement me")
}

func Cons[A any](head A, tail List[A]) List[A] {
	return ListCons[A]{Head: head, Tail: tail}
}

func (c ListCons[A]) list() {
	panic("implement me")
}

func ListApply[A any](as ...A) List[A] {
	if len(as) == 0 {
		return ListNil{}
	}
	return Cons[A](as[0], ListApply[A](as[1:]...))
}

func ListAppend[A any](a1 List[A], a2 List[A]) List[A] {
	switch a := a1.(type) {
	case ListNil:
		return a2
	case ListCons[A]:
		h, t := a.Head, a.Tail
		return Cons[A](h, ListAppend[A](t, a2))
	default:
		panic("unreachable")
	}
}

func FoldRight[A, B any](as List[A], z B, f func(A, B) B) B {
	switch a := as.(type) {
	case ListNil:
		return z
	case ListCons[A]:
		x, xs := a.Head, a.Tail
		return f(x, FoldRight[A, B](xs, z, f))
	default:
		panic("unreachable")
	}
}

func ListTail[A any](l List[A]) List[A] {
	switch a := l.(type) {
	case ListNil:
		panic("tail of empty list")
	case ListCons[A]:
		return a.Tail
	default:
		panic("unreachable")
	}
}
