package factory

import "abstract-pattern/product"

// Adidas is the concrete factory for Adidas products.
type Adidas struct{}

// MakeShoe creates an Adidas shoe.
func (a *Adidas) MakeShoe() product.IShoe {
	return &product.AdidasShoe{
		Shoe: product.Shoe{
			Logo: "adidas",
			Size: 14,
		},
	}
}

// MakeShirt creates an Adidas shirt.
func (a *Adidas) MakeShirt() product.IShirt {
	return &product.AdidasShirt{
		Shirt: product.Shirt{
			Logo: "adidas",
			Size: 14,
		},
	}
}
