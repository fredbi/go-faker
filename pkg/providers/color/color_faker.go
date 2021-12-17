package color

import (
	"github.com/fredbi/go-faker/pkg/generic"
)

type Faker struct {
	generic.Faker
}

func (f *Faker) RGB() string  { return "" }
func (f *Faker) HTML() string { return "" }
