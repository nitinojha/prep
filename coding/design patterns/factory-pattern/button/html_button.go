package button

// HTMLButton is a concrete implementation of the Button interface for HTML.
type HTMLButton struct{}

// Render renders an HTML-style button.
func (h *HTMLButton) Render() string {
    return `<button>HTML Button</button>`
}

// OnClick defines the click behavior for an HTML button.
func (h *HTMLButton) OnClick() string {
    return "HTML button clicked!"
}
