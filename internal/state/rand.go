package state

type (
	Rand[A any] func(RNG) (A, RNG)
)

var (
	RandInt Rand[int] = func(rng RNG) (int, RNG) {
		return rng.NextInt()
	}
)
