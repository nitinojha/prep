package button

// Button interface defines the common methods that all buttons must implement.
type Button interface {
    Render() string
    OnClick() string
}
