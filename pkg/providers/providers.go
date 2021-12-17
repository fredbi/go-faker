package providers

import (
	"github.com/fredbi/go-faker/pkg/generic"
	"github.com/fredbi/go-faker/pkg/providers/address"
	"github.com/fredbi/go-faker/pkg/providers/animal"
	"github.com/fredbi/go-faker/pkg/providers/auth"
	"github.com/fredbi/go-faker/pkg/providers/bank"
	"github.com/fredbi/go-faker/pkg/providers/color"
	"github.com/fredbi/go-faker/pkg/providers/company"
	"github.com/fredbi/go-faker/pkg/providers/network"
	"github.com/fredbi/go-faker/pkg/providers/person"
	tags "github.com/fredbi/go-faker/pkg/providers/struct-tags"
	"github.com/fredbi/go-faker/pkg/providers/text"
	"github.com/fredbi/go-faker/pkg/providers/time"
	"github.com/fredbi/go-faker/pkg/providers/types"
)

type (
	// Default provider of fake data, bundling the following fakers:
	//
	// * address.Faker
	// * person.Faker
	// * tags.Faker
	// * text.Faker
	// * time.Faker
	// * types.Faker
	Default struct {
		*generic.Faker

		address *address.Faker
		person  *person.Faker
		tags    *tags.Faker
		text    *text.Faker
		time    *time.Faker
		types   *types.Faker
	}

	// Contrib provider of fake data, bundling the following fakers:
	//
	// * animal.Faker
	// * auth.Faker
	// * bank.Faker
	// * color.Faker
	// * company.Faker
	// * network.Faker
	Contrib struct {
		*generic.Faker

		animal  *animal.Faker
		auth    *auth.Faker
		bank    *bank.Faker
		color   *color.Faker
		company *company.Faker
		// file     *file.Faker
		// geom     *geom.Faker
		// grpc     *grpc.Faker
		// language *language.Faker
		network *network.Faker
		// postgres *postgres.Faker
		// web *web.Faker
	}

	// All providers bundled in a single one
	All struct {
		Default
		Contrib
	}
)

// NewDefault builds a new default fake data provider.
func NewDefault() *Default {
	return &Default{}
}

// Key uniquely identifies this provider with key "default".
func (bp *Default) Key() string {
	return "default"
}

// Register the default provider
func (bp *Default) Register(*generic.Faker) {
}

// NewContrib builds a new Contrib provider
func NewContrib() *Contrib {
	return &Contrib{}
}

// Key uniquely identifies this provider with key "contrib".
func (cp *Contrib) Key() string {
	return "contrib"
}

// Register the contrib provider
func (cp *Contrib) Register(*generic.Faker) {
}
