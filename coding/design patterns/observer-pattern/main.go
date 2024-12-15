package main

import (
	"observer-pattern/editor"
	"observer-pattern/listeners"
)

func main() {
	// Create an editor and event listeners.
	e := editor.NewEditor()

	emailListener := listeners.NewEmailListener("user@example.com")
	loggingListener := listeners.NewLoggingListener("app.log")

	// Subscribe listeners to events.
	e.EventManager.Subscribe("open", emailListener)
	e.EventManager.Subscribe("save", loggingListener)

	// Perform editor actions.
	e.Open("test_file.txt")
	e.Save()
}
