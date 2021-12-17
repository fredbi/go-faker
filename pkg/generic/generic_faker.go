package generic

// Faker is a generic faker that provides standard random sources
// and facilities to all fake data providers.
type Faker struct {
	*options
}

// New generic faker
func New(opts ...Option) *Faker {
	o := defaultOptions(opts...)

	return &Faker{
		options: o,
	}
}
