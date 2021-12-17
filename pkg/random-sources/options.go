package sources

type (
	// Option configures a random source
	Option func(*options)

	options struct {
		seed       int64
		sourceType sourceType
	}
)

func defaultOptions(opts ...Option) *options {
	o := &options{
		seed:       1,
		sourceType: mathBased,
	}

	for _, apply := range opts {
		apply(o)
	}

	return o
}

// WithSeed configures the seed for the random generator.
//
// When not configured, the generator gets its default seed value 1.
func WithSeed(seed int64) Option {
	return func(o *options) {
		o.seed = seed
	}
}
