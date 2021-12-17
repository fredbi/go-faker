package faker

import (
	"github.com/fredbi/go-faker/pkg/generic"
	"github.com/fredbi/go-faker/pkg/providers"
)

type (

	// Provider knows how to extend a Faker with some fake data generation methods.
	//
	// A provider is uniquely identified by a key.
	//
	// Although fake data providers may be used standalone, registering them under
	// a single Faker host object allows for sharing generic settings for the random source.
	Provider interface {
		Key() string
		Register(*generic.Faker)
	}

	// Faker parent struct to generate fake data.
	//
	// A Faker provides some infrastructure to manipulate multiple fake data providers.
	Faker struct {
		*options

		*generic.Faker
	}
)

var (
	_ Provider = &providers.Default{}
	_ Provider = &providers.Contrib{}
)

// NewFaker builds a new faker.
//
// By default, the faker supports the default provider: provider.Default.
func NewFaker(opts ...Option) *Faker {
	o := defaultOptions(opts...)

	f := &Faker{
		options: o,
		Faker:   generic.New(o.genericOptions...),
	}

	f.options.providers.Range(func(key, value interface{}) bool {
		p := value.(Provider)

		p.Register(f.Faker)

		return true
	})

	return f
}

// Provider returns the provider that is registered for a given key.
//
// To use the returned provider, a type assertion would be required.
func (f *Faker) Provider(key string) Provider {
	v, ok := f.providers.Load(key)
	if !ok {
		return nil
	}

	return v.(Provider)
}

// Default provider, ready for use.
func (f *Faker) Default() *providers.Default {
	v, ok := f.providers.Load("default")
	if !ok {
		return nil
	}

	d, ok := v.(*providers.Default)
	if !ok {
		return nil
	}

	return d
}

// Contrib provider, ready for use
func (f *Faker) Contrib() *providers.Contrib {
	v, ok := f.providers.Load("contrib")
	if !ok {
		return nil
	}

	d, ok := v.(*providers.Contrib)
	if !ok {
		return nil
	}

	return d
}

// TODO(fred): add more convenience methods
