package button

// WindowsButton is a concrete implementation of the Button interface for Windows.
type WindowsButton struct{}

// Render renders a Windows-style button.
func (w *WindowsButton) Render() string {
    return "Rendering a Windows button."
}

// OnClick defines the click behavior for a Windows button.
func (w *WindowsButton) OnClick() string {
    return "Windows button clicked!"
}
