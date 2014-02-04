// Package we specifies the types and constants commonly used for window events.
package we

// A Close event is triggered when the window is closed.
type Close struct{}

// A Resize event is triggered when the window is resized.
type Resize struct {
	Width  int
	Height int
}
