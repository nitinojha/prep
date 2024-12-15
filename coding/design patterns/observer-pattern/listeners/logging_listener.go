package listeners

import "fmt"

// LoggingListener logs events to a file or console.
type LoggingListener struct {
	logFile string
}

// NewLoggingListener creates a new LoggingListener.
func NewLoggingListener(logFile string) *LoggingListener {
	return &LoggingListener{logFile: logFile}
}

// Update logs the event when it occurs.
func (ll *LoggingListener) Update(eventType, data string) {
	fmt.Printf("Log to %s: Event '%s' occurred with data '%s'\n", ll.logFile, eventType, data)
}
