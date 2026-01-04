package parallelism

type (
	Executor interface {
		Submit(task func())
	}

	Future[A any] func(cb func(A))

	Par[A any] func(es Executor) Future[A]
)

func (p Par[A]) Run(es Executor) A {
	result := make(chan A, 1)
	p(es)(func(a A) {
		result <- a
	})
	return <-result
}

func ParUnit[A any](a A) Par[A] {
	return func(es Executor) Future[A] {
		return func(cb func(A)) {
			cb(a)
		}
	}
}

func ParEval(es Executor, r func()) {
	es.Submit(r)
}

func ParChoice[A any](cond Par[bool], t Par[A], f Par[A]) Par[A] {
	return func(es Executor) Future[A] {
		return func(cb func(A)) {
			cond(es)(func(b bool) {
				if b {
					ParEval(es, func() {
						t(es)(cb)
					})
					return
				}
				ParEval(es, func() {
					f(es)(cb)
				})
			})
		}
	}
}
