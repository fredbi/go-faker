package person

import "github.com/fredbi/go-faker/pkg/generic"

type Faker struct {
	generic.Faker
}

func (f *Faker) Name() string            { return "" }
func (f *Faker) LastName() string        { return "" }
func (f *Faker) FirstName() string       { return "" }
func (f *Faker) Title() string           { return "" }
func (f *Faker) TitleMale() string       { return "" }
func (f *Faker) TitleFemale() string     { return "" }
func (f *Faker) FirstNameFemale() string { return "" }
func (f *Faker) FirstNameMale() string   { return "" }
func (f *Faker) Gender() string          { return "" }
func (f *Faker) Email() string           { return "" }
func (f *Faker) PhoneNumber() string     { return "" }
func (f *Faker) PhoneFormatted() string  { return "" }
func (f *Faker) SSN() string             { return "" }
