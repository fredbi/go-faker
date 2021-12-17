package faker

import (
	"sync"

	"github.com/fredbi/go-faker/pkg/generic"
)

type (
	// Option to configure a faker
	Option func(*options)

	options struct {
		providers      sync.Map
		genericOptions []generic.Option
	}
)

func defaultOptions(opts ...Option) *options {
	o := &options{} // TODO: add some default providers

	for _, apply := range opts {
		apply(o)
	}

	return o
}

// WithGenericOptions sets the options for the generic random generation source for fakers.
func WithGenericOptions(genericOptions ...generic.Option) Option {
	return func(o *options) {
		o.genericOptions = genericOptions
	}
}

// WithProviders registers providers for a faker.
//
// The providers may later be recalled by the key they have registered with.
//
// WithProviders may be called several times and will append providers.
//
// The default provider may be overridden with a provider which identifies itself with the "default" key.
func WithProviders(providers ...Provider) Option {
	return func(o *options) {
		for _, provider := range providers {
			o.providers.Store(provider.Key(), provider)
		}
	}
}
