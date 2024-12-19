package product

// IShoe defines the interface for all shoe products.
type IShoe interface {
    SetLogo(logo string)
    SetSize(size int)
    GetLogo() string
    GetSize() int
}

// Shoe is the base struct for all shoes.
type Shoe struct {
    Logo string
    Size int
}

// SetLogo sets the logo for the shoe.
func (s *Shoe) SetLogo(logo string) {
    s.Logo = logo
}

// GetLogo returns the logo of the shoe.
func (s *Shoe) GetLogo() string {
    return s.Logo
}

// SetSize sets the size for the shoe.
func (s *Shoe) SetSize(size int) {
    s.Size = size
}

// GetSize returns the size of the shoe.
func (s *Shoe) GetSize() int {
    return s.Size
}

// AdidasShoe represents an Adidas shoe.
type AdidasShoe struct {
    Shoe
}

// NikeShoe represents a Nike shoe.
type NikeShoe struct {
    Shoe
}
