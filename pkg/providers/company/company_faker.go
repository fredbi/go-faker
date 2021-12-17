package company

import (
	"github.com/fredbi/go-faker/pkg/generic"
)

type Faker struct {
	generic.Faker
}

func (f *Faker) Company() string               { return "" }
func (f *Faker) CompanyRegistrationID() string { return "" }
