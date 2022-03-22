package main

import "fmt"

func main() {
	builder := NewAddressBuilder().
	WithDoorNumber("1").
	WithPostcode("SR2 7LT")

	fmt.Println(builder.GetAddress())
}

type AddressBuilder struct {
	address string
}

func NewAddressBuilder() *AddressBuilder {
	return &AddressBuilder{}
}

func (a *AddressBuilder) WithPostcode(p string) *AddressBuilder {
	a.address = a.address + p

	return a
}

func (a *AddressBuilder) WithDoorNumber(d string) *AddressBuilder {
	a.address = a.address + d

	return a
}

func (a *AddressBuilder) GetAddress() string {
	return a.address
}
