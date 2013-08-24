// Package we specifies the types and constants commonly used for window events.
package we

// Close is triggered when the window is closed.
type Close struct{}

// Resize is triggered when the window is resized.
type Resize struct {
	Width  int
	Height int
}
