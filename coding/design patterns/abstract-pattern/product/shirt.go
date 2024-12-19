package product

// IShirt defines the interface for all shirt products.
type IShirt interface {
    SetLogo(logo string)
    SetSize(size int)
    GetLogo() string
    GetSize() int
}

// Shirt is the base struct for all shirts.
type Shirt struct {
    Logo string
    Size int
}

// SetLogo sets the logo for the shirt.
func (s *Shirt) SetLogo(logo string) {
    s.Logo = logo
}

// GetLogo returns the logo of the shirt.
func (s *Shirt) GetLogo() string {
    return s.Logo
}

// SetSize sets the size for the shirt.
func (s *Shirt) SetSize(size int) {
    s.Size = size
}

// GetSize returns the size of the shirt.
func (s *Shirt) GetSize() int {
    return s.Size
}

// AdidasShirt represents an Adidas shirt.
type AdidasShirt struct {
    Shirt
}

// NikeShirt represents a Nike shirt.
type NikeShirt struct {
    Shirt
}
