package main

import (
	"fmt"
	"strategy-pattern/navigator"
	"strategy-pattern/strategy"
)

func main() {
	n := &navigator.Navigator{}

	// Use Road Strategy
	n.SetStrategy(&strategy.RoadStrategy{})
	fmt.Println(n.Navigate("Point A", "Point B"))

	// Use Public Transport Strategy
	n.SetStrategy(&strategy.PublicTransportStrategy{})
	fmt.Println(n.Navigate("Point A", "Point B"))

	// Use Walking Strategy
	n.SetStrategy(&strategy.WalkingStrategy{})
	fmt.Println(n.Navigate("Point A", "Point B"))
}
