package strategy

import "fmt"

// PublicTransportStrategy implements RouteStrategy for public transport navigation.
type PublicTransportStrategy struct{}

func (p *PublicTransportStrategy) CalculateRoute(start, end string) string {
	return fmt.Sprintf("Calculating public transport route from %s to %s", start, end)
}
