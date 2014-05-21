// Package we specifies commonly used types and constants for window events.
package we

// An Event represents a window, keyboard or mouse event.
//
// Window events
//
// The following window events are defined by this package.
//    http://godoc.org/github.com/mewmew/we#Close
//    http://godoc.org/github.com/mewmew/we#Resize
//
// Keyboard events
//
// The following keyboard events are defined by this package.
//    http://godoc.org/github.com/mewmew/we#KeyPress
//    http://godoc.org/github.com/mewmew/we#KeyRelease
//    http://godoc.org/github.com/mewmew/we#KeyRepeat
//    http://godoc.org/github.com/mewmew/we#KeyRune
//
// Mouse events
//
// The following mouse events are defined by this package.
//    http://godoc.org/github.com/mewmew/we#MousePress
//    http://godoc.org/github.com/mewmew/we#MouseRelease
//    http://godoc.org/github.com/mewmew/we#MouseMove
//    http://godoc.org/github.com/mewmew/we#MouseDrag
//    http://godoc.org/github.com/mewmew/we#MouseEnter
//    http://godoc.org/github.com/mewmew/we#ScrollX
//    http://godoc.org/github.com/mewmew/we#ScrollY
type Event interface{}
