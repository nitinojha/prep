package editor

// Editor represents an object that triggers events.
type Editor struct {
	EventManager *EventManager
	content      string
}

// NewEditor creates a new Editor with an associated EventManager.
func NewEditor() *Editor {
	return &Editor{
		EventManager: NewEventManager(),
	}
}

// Open simulates opening a file and triggers an "open" event.
func (e *Editor) Open(fileName string) {
	e.content = fileName
	e.EventManager.Notify("open", fileName)
}

// Save simulates saving a file and triggers a "save" event.
func (e *Editor) Save() {
	e.EventManager.Notify("save", e.content)
}
