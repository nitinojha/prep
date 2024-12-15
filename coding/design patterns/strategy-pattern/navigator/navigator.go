package navigator

import "strategy-pattern/strategy"

// Navigator is the context class that uses a RouteStrategy.
type Navigator struct {
	strategy strategy.RouteStrategy
}

// SetStrategy allows changing the navigation strategy dynamically.
func (n *Navigator) SetStrategy(strategy strategy.RouteStrategy) {
	n.strategy = strategy
}

// Navigate uses the current strategy to calculate a route.
func (n *Navigator) Navigate(start, end string) string {
	if n.strategy == nil {
		return "No strategy set for navigation"
	}
	return n.strategy.CalculateRoute(start, end)
}
