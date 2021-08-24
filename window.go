package we

// A Close event is triggered when the window is closed.
type Close struct{}

// A Resize event is triggered when the window is resized.
type Resize struct {
	// New window dimensions.
	Width, Height int
}
