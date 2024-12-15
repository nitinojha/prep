package strategy

import "fmt"

// WalkingStrategy implements RouteStrategy for walking navigation.
type WalkingStrategy struct{}

func (w *WalkingStrategy) CalculateRoute(start, end string) string {
	return fmt.Sprintf("Calculating walking route from %s to %s", start, end)
}
