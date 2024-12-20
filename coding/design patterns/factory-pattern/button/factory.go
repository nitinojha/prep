package button

import "fmt"

// ButtonFactory is the factory interface for creating buttons.
type ButtonFactory interface {
    CreateButton() Button
}

// WindowsFactory is a factory for creating Windows buttons.
type WindowsFactory struct{}

// CreateButton creates a Windows button.
func (w *WindowsFactory) CreateButton() Button {
    return &WindowsButton{}
}

// HTMLFactory is a factory for creating HTML buttons.
type HTMLFactory struct{}

// CreateButton creates an HTML button.
func (h *HTMLFactory) CreateButton() Button {
    return &HTMLButton{}
}

// GetFactory returns the appropriate factory based on the type.
func GetFactory(factoryType string) (ButtonFactory, error) {
    switch factoryType {
    case "windows":
        return &WindowsFactory{}, nil
    case "html":
        return &HTMLFactory{}, nil
    default:
        return nil, fmt.Errorf("unknown factory type: %s", factoryType)
    }
}
