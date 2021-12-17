package address

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/fredbi/go-faker/pkg/generic"
)

type Faker struct {
	generic.Faker
}

type AddressInfo gofakeit.AddressInfo

func (f *Faker) Address() *AddressInfo { return nil }
func (f *Faker) AddressString() string { return "" }
func (f *Faker) City() string          { return "" }
func (f *Faker) Country() string       { return "" }
func (f *Faker) Zip() string           { return "" }
func (f *Faker) State() string         { return "" }
func (f *Faker) Street() string        { return "" }
func (f *Faker) StreetNumber() string  { return "" }
