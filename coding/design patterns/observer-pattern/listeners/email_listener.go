package listeners

import "fmt"

// EmailListener sends notifications via email.
type EmailListener struct {
	email string
}

// NewEmailListener creates a new EmailListener.
func NewEmailListener(email string) *EmailListener {
	return &EmailListener{email: email}
}

// Update sends an email notification when an event occurs.
func (el *EmailListener) Update(eventType, data string) {
	fmt.Printf("Email to %s: Event '%s' occurred with data '%s'\n", el.email, eventType, data)
}
