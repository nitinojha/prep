package strategy

// RouteStrategy defines the interface for route strategies.
type RouteStrategy interface {
	CalculateRoute(start, end string) string
}
