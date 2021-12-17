package network

import (
	"net"

	"github.com/fredbi/go-faker/pkg/generic"
)

type Faker struct {
	generic.Faker
}

func (f *Faker) Email() string      { return "" }
func (f *Faker) IP() net.IP         { return net.IP{} }
func (f *Faker) IPv4() string       { return "" }
func (f *Faker) IPv6() string       { return "" }
func (f *Faker) MacAddress() string { return "" }
func (f *Faker) URL() string        { return "" }
func (f *Faker) DomainName() string { return "" }
