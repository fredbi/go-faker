package text

import (
	"github.com/fredbi/go-faker/pkg/generic"
)

// Faker generates fake text data like sentences, paragraphs, lorem ipsum etc.
type Faker struct {
	*options
	*generic.Faker
}

// New text faker
func New(opts ...Option) *Faker {
	o := defaultOptions(opts...)

	return &Faker{
		options: o,
		Faker:   generic.New(o.genericOptions...),
	}
}

// Key uniquely identifying this "text" faker
func (f *Faker) Key() string {
	return "text"
}

// Register this faker with a generic Faker
func (f *Faker) Register(gen *generic.Faker) {
	f.Faker = gen
}

// Paragraph constructs a fake paragraph of text, made of a few sentences.
func (f *Faker) Paragraph() string           { return "" }
func (f *Faker) Sentence() string            { return "" }
func (f *Faker) Phrase() string              { return "" }
func (f *Faker) Word() string                { return "" }
func (f *Faker) Noun() string                { return "" }
func (f *Faker) Verb() string                { return "" }
func (f *Faker) Preposition() string         { return "" }
func (f *Faker) Adjective() string           { return "" }
func (f *Faker) LoremIpsumWord() string      { return "" }
func (f *Faker) LoremIpsumSentence() string  { return "" }
func (f *Faker) LoremIpsumParagraph() string { return "" }
func (f *Faker) Question() string            { return "" }
func (f *Faker) Quote() string               { return "" }
