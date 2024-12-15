package listeners

// EventListener defines the interface for listeners.
type EventListener interface {
	Update(eventType, data string)
}
