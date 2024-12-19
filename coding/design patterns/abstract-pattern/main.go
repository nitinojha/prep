package main

import (
	"abstract-pattern/factory"
	"abstract-pattern/product"
	"fmt"
)

// printShoeDetails prints the details of a shoe.
func printShoeDetails(s product.IShoe) {
	fmt.Printf("Shoe Logo: %s\n", s.GetLogo())
	fmt.Printf("Shoe Size: %d\n", s.GetSize())
	fmt.Println()
}

// printShirtDetails prints the details of a shirt.
func printShirtDetails(s product.IShirt) {
	fmt.Printf("Shirt Logo: %s\n", s.GetLogo())
	fmt.Printf("Shirt Size: %d\n", s.GetSize())
	fmt.Println()
}

func main() {
	// Create Adidas factory
	adidasFactory, _ := factory.GetSportsFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	// Create Nike factory
	nikeFactory, _ := factory.GetSportsFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	// Print details
	fmt.Println("Adidas Products:")
	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

	fmt.Println("Nike Products:")
	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)
}
