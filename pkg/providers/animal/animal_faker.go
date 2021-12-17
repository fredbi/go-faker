package animal

import (
	"github.com/fredbi/go-faker/pkg/generic"
)

type Faker struct {
	generic.Faker
}

func (f *Faker) Pet() string { return "" }
