package factory

import "abstract-pattern/product"

// Nike is the concrete factory for Nike products.
type Nike struct{}

// MakeShoe creates a Nike shoe.
func (n *Nike) MakeShoe() product.IShoe {
	return &product.NikeShoe{
		Shoe: product.Shoe{
			Logo: "nike",
			Size: 14,
		},
	}
}

// MakeShirt creates a Nike shirt.
func (n *Nike) MakeShirt() product.IShirt {
	return &product.NikeShirt{
		Shirt: product.Shirt{
			Logo: "nike",
			Size: 14,
		},
	}
}
