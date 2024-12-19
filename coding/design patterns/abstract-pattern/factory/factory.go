package factory

import (
	"abstract-pattern/product"
	"fmt"
)

// ISportsFactory defines the interface for sports factories.
type ISportsFactory interface {
	MakeShoe() product.IShoe
	MakeShirt() product.IShirt
}

// GetSportsFactory returns a factory instance based on the brand name.
func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
