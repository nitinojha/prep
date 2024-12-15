package strategy

import "fmt"

// RoadStrategy implements RouteStrategy for road navigation.
type RoadStrategy struct{}

func (r *RoadStrategy) CalculateRoute(start, end string) string {
	return fmt.Sprintf("Calculating road route from %s to %s", start, end)
}
