package auth

import (
	"github.com/fredbi/go-faker/pkg/generic"
)

type Faker struct {
	generic.Faker
}

func (f *Faker) Username() string { return "" }
func (f *Faker) Password() string { return "" }
