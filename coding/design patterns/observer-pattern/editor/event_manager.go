package editor

import "observer-pattern/listeners"

// EventManager handles event subscriptions and notifications.
type EventManager struct {
	listeners map[string][]listeners.EventListener
}

// NewEventManager creates a new EventManager.
func NewEventManager() *EventManager {
	return &EventManager{
		listeners: make(map[string][]listeners.EventListener),
	}
}

// Subscribe adds a listener for a specific event type.
func (em *EventManager) Subscribe(eventType string, listener listeners.EventListener) {
	em.listeners[eventType] = append(em.listeners[eventType], listener)
}

// Unsubscribe removes a listener for a specific event type.
func (em *EventManager) Unsubscribe(eventType string, listener listeners.EventListener) {
	if _, found := em.listeners[eventType]; found {
		newListeners := []listeners.EventListener{}
		for _, l := range em.listeners[eventType] {
			if l != listener {
				newListeners = append(newListeners, l)
			}
		}
		em.listeners[eventType] = newListeners
	}
}

// Notify sends updates to all listeners of a specific event type.
func (em *EventManager) Notify(eventType, data string) {
	if listeners, found := em.listeners[eventType]; found {
		for _, listener := range listeners {
			listener.Update(eventType, data)
		}
	}
}
