// Package we specifies commonly used types and constants for window events. An
// overview of the various event types is given at
// http://pkg.go.dev/github.com/mewspring/we#Event
package we

// An Event represents a window, keyboard or mouse event.
//
// Window events
//
// The following window events are defined by this package.
//
//    http://pkg.go.dev/github.com/mewspring/we#Close
//    http://pkg.go.dev/github.com/mewspring/we#Resize
//
// Keyboard events
//
// The following keyboard events are defined by this package.
//
//    http://pkg.go.dev/github.com/mewspring/we#KeyPress
//    http://pkg.go.dev/github.com/mewspring/we#KeyRelease
//    http://pkg.go.dev/github.com/mewspring/we#KeyRepeat
//    http://pkg.go.dev/github.com/mewspring/we#KeyRune
//
// Mouse events
//
// The following mouse events are defined by this package.
//
//    http://pkg.go.dev/github.com/mewspring/we#MousePress
//    http://pkg.go.dev/github.com/mewspring/we#MouseRelease
//    http://pkg.go.dev/github.com/mewspring/we#MouseMove
//    http://pkg.go.dev/github.com/mewspring/we#MouseDrag
//    http://pkg.go.dev/github.com/mewspring/we#MouseEnter
//    http://pkg.go.dev/github.com/mewspring/we#ScrollX
//    http://pkg.go.dev/github.com/mewspring/we#ScrollY
type Event interface{}
