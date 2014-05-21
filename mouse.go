package we

import (
	"fmt"
	"image"
	"strings"
)

// Button represent a bitfield of mouse buttons.
type Button int

// Mouse button bitfield values.
const (
	Button1 Button = 1 << iota
	Button2
	Button3
	Button4
	Button5
	Button6
	Button7
	Button8

	// Aliases.
	ButtonLeft   = Button1
	ButtonRight  = Button2
	ButtonMiddle = Button3
)

// buttonNames maps from Button values to string descriptions.
var buttonNames = map[Button]string{
	ButtonLeft:   "[left button]",
	ButtonMiddle: "[middle button]",
	ButtonRight:  "[right button]",
}

func (button Button) String() string {
	if button == 0 {
		return ""
	}

	// Exit early if only one mouse button is held down.
	s, ok := buttonNames[button]
	if ok {
		return s
	}

	var buttons []string
	if button&ButtonLeft != 0 {
		buttons = append(buttons, "left button")
	}
	if button&ButtonMiddle != 0 {
		buttons = append(buttons, "middle button")
	}
	if button&ButtonRight != 0 {
		buttons = append(buttons, "right button")
	}
	for mask := Button4; mask <= Button8; mask <<= 1 {
		if button&mask != 0 {
			buttons = append(buttons, fmt.Sprintf("button %d", mask))
		}
	}
	return "[" + strings.Join(buttons, "+") + "]"
}

// A MousePress event is triggered when one or more mouse buttons are pressed.
type MousePress struct {
	// Coordinates of the mouse cursor.
	image.Point
	// Bitfield of pressed mouse buttons.
	Button Button
	// Bitfield of key modifiers.
	Mod Mod
}

func (e MousePress) String() string {
	// Override the embedded String method of image.Point.
	return fmt.Sprintf("{%v %v %v}", e.Point, e.Button, e.Mod)
}

// A MouseRelease event is triggered when one or more mouse buttons are released.
type MouseRelease struct {
	// Coordinates of the mouse cursor.
	image.Point
	// Bitfield of released mouse buttons.
	Button Button
	// Bitfield of key modifiers.
	Mod Mod
}

func (e MouseRelease) String() string {
	// Override the embedded String method of image.Point.
	return fmt.Sprintf("{%v %v %v}", e.Point, e.Button, e.Mod)
}

// A MouseMove event is triggered when the mouse is moved from one location to
// another.
type MouseMove struct {
	// Coordinates of the mouse cursor.
	image.Point
	// Coordinates of the mouse cursor at the beginning of the move event.
	From image.Point
}

func (e MouseMove) String() string {
	// Override the embedded String method of image.Point.
	return fmt.Sprintf("{%v %v}", e.Point, e.From)
}

// A MouseDrag event is triggered when the mouse is moved from one location to
// another while one or more mouse buttons are held down.
type MouseDrag struct {
	// Coordinates of the mouse cursor.
	image.Point
	// Coordinates of the mouse cursor at the beginning of the drag event.
	From image.Point
	// Bitfield of mouse buttons that were held down.
	Button Button
	// Bitfield of key modifiers.
	Mod Mod
}

func (e MouseDrag) String() string {
	// Override the embedded String method of image.Point.
	return fmt.Sprintf("{%v %v %v %v}", e.Point, e.From, e.Button, e.Mod)
}

// A MouseEnter event is triggered when the mouse enters or leaves the window.
// On mouse enter the value is true, otherwise it is false.
type MouseEnter bool

func (enter MouseEnter) String() string {
	if enter {
		return "enter"
	}
	return "leave"
}

// A ScrollX event is triggered when the mouse wheel is scrolled on the
// horizontal axis.
type ScrollX struct {
	// Coordinates of the mouse cursor.
	image.Point
	// Horizontal scroll offset.
	Off int
	// Bitfield of key modifiers.
	Mod Mod
}

func (e ScrollX) String() string {
	// Override the embedded String method of image.Point.
	return fmt.Sprintf("{%v %v %v}", e.Point, e.Off, e.Mod)
}

// A ScrollY event is triggered when the mouse wheel is scrolled on the vertical
// axis.
type ScrollY struct {
	// Coordinates of the mouse cursor.
	image.Point
	// Vertical scroll offset.
	Off int
	// Bitfield of key modifiers.
	Mod Mod
}

func (e ScrollY) String() string {
	// Override the embedded String method of image.Point.
	return fmt.Sprintf("{%v %v %v}", e.Point, e.Off, e.Mod)
}
