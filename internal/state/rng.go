package state

type (
	RNG interface {
		NextInt() (int, RNG)
	}

	RNGSimple struct {
		seed uint64
	}
)

const (
	multiplier uint64 = 0x5DEECE66D
	addend     uint64 = 0xB
	mask       uint64 = (1 << 48) - 1 // 0xFFFFFFFFFFFF
)

func NewRNGSimple(seed uint64) RNG {
	return &RNGSimple{seed: seed}
}

func (r RNGSimple) NextInt() (int, RNG) {
	newSeed := (r.seed*multiplier + addend) & mask
	nextRNG := RNGSimple{seed: newSeed}
	n := int(newSeed >> 16)
	return n, nextRNG
}

func NonNegativeInt(rng RNG) (int, RNG) {
	i, r := rng.NextInt()
	if i < 0 {
		return -(i + 1), r
	}
	return i, r
}

func RNGBoolean(rng RNG) (bool, RNG) {
	i, r := rng.NextInt()
	return i%2 == 0, r
}
