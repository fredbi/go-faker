package generic

import (
	mathrand "math/rand"
	"time"
)

type (
	// Option to configure a generic faker
	Option func(*options)

	options struct {
		seed                  int64
		source64              mathrand.Source64
		withUniqueSliceValues bool
		withNilEmptySlices    bool
		withNilEmptyMaps      bool
		withNoZeroValues      bool
	}
)

// TODO: rand source based on crypto/rand
// TODO: common gofakeit helpers
// TODO: ability to fetch values in configured maps

func defaultOptions(opts ...Option) *options {
	seed := time.Now().UnixNano()
	o := &options{
		seed:     seed,
		source64: nil, // TODO: goroutine-safe math rand source
	}

	for _, apply := range opts {
		apply(o)
	}

	return o
}

// WithSeed applies a seed to the random source, so as to make it reproducible.
func WithSeed(seed int64) Option {
	return func(o *options) {
		o.seed = seed
	}
}

// WithSource64 configures a random source
func WithSource64(source64 mathrand.Source64) Option {
	return func(o *options) {
		o.source64 = source64
	}
}

func WithUniqueSliceValues(enabled bool) Option {
	return func(o *options) {
		o.withUniqueSliceValues = enabled
	}
}

func WithNilEmptySlices(enabled bool) Option {
	return func(o *options) {
		o.withNilEmptySlices = enabled
	}
}

func WithNilEmptyMaps(enabled bool) Option {
	return func(o *options) {
		o.withNilEmptyMaps = enabled
	}
}

// WithNoZeroValues instructs the random ... TODO
func WithNoZeroValues(enabled bool) Option {
	return func(o *options) {
		o.withNoZeroValues = enabled
	}
}
