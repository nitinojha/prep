package main

import (
	"factory-pattern/button"
	"fmt"
)

func main() {
	// Get a Windows button factory
	windowsFactory, err := button.GetFactory("windows")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a Windows button
	windowsButton := windowsFactory.CreateButton()
	fmt.Println(windowsButton.Render())
	fmt.Println(windowsButton.OnClick())

	// Get an HTML button factory
	htmlFactory, err := button.GetFactory("html")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create an HTML button
	htmlButton := htmlFactory.CreateButton()
	fmt.Println(htmlButton.Render())
	fmt.Println(htmlButton.OnClick())
}
