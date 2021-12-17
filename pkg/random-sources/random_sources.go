package sources

import "math/rand"

type (
	// Rand is a concurrent-safe random source based on math.Rand standard library
	Rand struct {
		*options
		source rand.Source64
	}

	// CryptoRand is a concurrent-safe random source based on crypto/rand standard library
	CryptoRand struct {
		*options
		*cryptoRand
	}

	sourceType uint8
)

const (
	mathBased sourceType = iota
	cryptoBased
)

// NewRand builds a new concurrent-safe random source based on math.Rand
func NewRand(opts ...Option) *Rand {
	o := defaultOptions()

	var source rand.Source64
	switch o.sourceType {
	case mathBased:
		source = &concurrentSource{}
	case cryptoBased:
		source = &cryptoRand{}
	default:
		panic("unsupport random source type")
	}

	source.Seed(o.seed)

	return &Rand{
		options: o,
		source:  source,
	}
}
