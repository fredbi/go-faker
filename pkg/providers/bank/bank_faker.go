package bank

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/fredbi/go-faker/pkg/generic"
)

type Faker struct {
	generic.Faker
}

type CreditCardInfo gofakeit.CreditCardInfo

func (f *Faker) CreditCard() *CreditCardInfo { return nil }
func (f *Faker) CreditCardNumber() string    { return "" }
func (f *Faker) CreditCardType() string      { return "" }
func (f *Faker) CreditCardCVV() string       { return "" }
func (f *Faker) AmountWithCurrency() string  { return "" }
func (f *Faker) Currency() string            { return "" }
func (f *Faker) CurrencyLong() string        { return "" }
func (f *Faker) Amount() float64             { return 0 }
func (f *Faker) Price() float64              { return 0 }
func (f *Faker) ACHRouting() string          { return "" }
func (f *Faker) ACHAccount() string          { return "" }
func (f *Faker) BitcoinAddress() string      { return "" }
func (f *Faker) BitcoinPrivateKey() string   { return "" }
