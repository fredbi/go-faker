package text

import "github.com/fredbi/go-faker/pkg/generic"

type (
	// Option to configure a text data faker
	Option func(*options)

	options struct {
		genericOptions []generic.Option
	}
)

func defaultOptions(opts ...Option) *options {
	o := &options{}

	for _, apply := range opts {
		apply(o)
	}

	return o
}

// WithGenericOptions sets the options for the generic random generation source for fakers.
//
// This can be used when the text faker is used standalone.
func WithGenericOptions(genericOptions ...generic.Option) Option {
	return func(o *options) {
		o.genericOptions = genericOptions
	}
}
