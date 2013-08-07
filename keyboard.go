package we

import (
	"fmt"
	"strings"
)

// Mod corresponds to a bitfield of modifier keys.
type Mod int

// Modifier key bitfield values.
const (
	// ModShift means one or more Shift keys were held down.
	ModShift Mod = 1 << iota
	// ModControl means one or more Control keys were held down.
	ModControl
	// ModAlt means one or more Alt keys were held down.
	ModAlt
	// ModSuper means one or more Super keys were held down.
	ModSuper
)

func (mod Mod) String() string {
	if mod == 0 {
		return ""
	}
	var mods []string
	if mod&ModControl != 0 {
		mods = append(mods, "control")
	}
	if mod&ModShift != 0 {
		mods = append(mods, "shift")
	}
	if mod&ModAlt != 0 {
		mods = append(mods, "alt")
	}
	if mod&ModSuper != 0 {
		mods = append(mods, "super")
	}
	return "[" + strings.Join(mods, "+") + "]"
}

// Key corresponds to a keyboard key.
type Key int

// Keyboard keys.
//
// These key codes are equivalent to those used in GLFW 3.0. They are inspired
// by the USB HID Usage Tables v1.12 [1] (p. 53-60), but re-arranged to map to
// 7-bit ASCII for printable keys (function keys are put in the 256+ range).
//
// [1]: http://www.usb.org/developers/devclass_docs/Hut1_12v2.pdf
const (
	// Printable keys.

	KeySpace        Key = 32
	KeyApostrophe   Key = 39 // '
	KeyComma        Key = 44 // ,
	KeyMinus        Key = 45 // -
	KeyPeriod       Key = 46 // .
	KeySlash        Key = 47 // /
	Key0            Key = 48
	Key1            Key = 49
	Key2            Key = 50
	Key3            Key = 51
	Key4            Key = 52
	Key5            Key = 53
	Key6            Key = 54
	Key7            Key = 55
	Key8            Key = 56
	Key9            Key = 57
	KeySemicolon    Key = 59 // ;
	KeyEqual        Key = 61 // =
	KeyA            Key = 65
	KeyB            Key = 66
	KeyC            Key = 67
	KeyD            Key = 68
	KeyE            Key = 69
	KeyF            Key = 70
	KeyG            Key = 71
	KeyH            Key = 72
	KeyI            Key = 73
	KeyJ            Key = 74
	KeyK            Key = 75
	KeyL            Key = 76
	KeyM            Key = 77
	KeyN            Key = 78
	KeyO            Key = 79
	KeyP            Key = 80
	KeyQ            Key = 81
	KeyR            Key = 82
	KeyS            Key = 83
	KeyT            Key = 84
	KeyU            Key = 85
	KeyV            Key = 86
	KeyW            Key = 87
	KeyX            Key = 88
	KeyY            Key = 89
	KeyZ            Key = 90
	KeyLeftBracket  Key = 91  // [
	KeyBackslash    Key = 92  // \
	KeyRightBracket Key = 93  // ]
	KeyGraveAccent  Key = 96  // `
	KeyWorld1       Key = 161 // non-US #1
	KeyWorld2       Key = 162 // non-US #2

	// Function keys.

	KeyEscape       Key = 256
	KeyEnter        Key = 257
	KeyTab          Key = 258
	KeyBackspace    Key = 259
	KeyInsert       Key = 260
	KeyDelete       Key = 261
	KeyRight        Key = 262
	KeyLeft         Key = 263
	KeyDown         Key = 264
	KeyUp           Key = 265
	KeyPageUp       Key = 266
	KeyPageDown     Key = 267
	KeyHome         Key = 268
	KeyEnd          Key = 269
	KeyCapsLock     Key = 280
	KeyScrollLock   Key = 281
	KeyNumLock      Key = 282
	KeyPrintScreen  Key = 283
	KeyPause        Key = 284
	KeyF1           Key = 290
	KeyF2           Key = 291
	KeyF3           Key = 292
	KeyF4           Key = 293
	KeyF5           Key = 294
	KeyF6           Key = 295
	KeyF7           Key = 296
	KeyF8           Key = 297
	KeyF9           Key = 298
	KeyF10          Key = 299
	KeyF11          Key = 300
	KeyF12          Key = 301
	KeyF13          Key = 302
	KeyF14          Key = 303
	KeyF15          Key = 304
	KeyF16          Key = 305
	KeyF17          Key = 306
	KeyF18          Key = 307
	KeyF19          Key = 308
	KeyF20          Key = 309
	KeyF21          Key = 310
	KeyF22          Key = 311
	KeyF23          Key = 312
	KeyF24          Key = 313
	KeyF25          Key = 314
	KeyKp0          Key = 320
	KeyKp1          Key = 321
	KeyKp2          Key = 322
	KeyKp3          Key = 323
	KeyKp4          Key = 324
	KeyKp5          Key = 325
	KeyKp6          Key = 326
	KeyKp7          Key = 327
	KeyKp8          Key = 328
	KeyKp9          Key = 329
	KeyKpDecimal    Key = 330
	KeyKpDivide     Key = 331
	KeyKpMultiply   Key = 332
	KeyKpSubtract   Key = 333
	KeyKpAdd        Key = 334
	KeyKpEnter      Key = 335
	KeyKpEqual      Key = 336
	KeyLeftShift    Key = 340
	KeyLeftControl  Key = 341
	KeyLeftAlt      Key = 342
	KeyLeftSuper    Key = 343
	KeyRightShift   Key = 344
	KeyRightControl Key = 345
	KeyRightAlt     Key = 346
	KeyRightSuper   Key = 347
	KeyMenu         Key = 348
)

// keyNames maps from Key value to string description.
var keyNames = map[Key]string{
	KeySpace:        "[space]",
	KeyWorld1:       "[world 1]",
	KeyWorld2:       "[world 2]",
	KeyEscape:       "[escape]",
	KeyEnter:        "[enter]",
	KeyTab:          "[tab]",
	KeyBackspace:    "[backspace]",
	KeyInsert:       "[insert]",
	KeyDelete:       "[delete]",
	KeyRight:        "[right]",
	KeyLeft:         "[left]",
	KeyDown:         "[down]",
	KeyUp:           "[up]",
	KeyPageUp:       "[page up]",
	KeyPageDown:     "[page down]",
	KeyHome:         "[home]",
	KeyEnd:          "[end]",
	KeyCapsLock:     "[caps lock]",
	KeyScrollLock:   "[scroll lock]",
	KeyNumLock:      "[num lock]",
	KeyPrintScreen:  "[print screen]",
	KeyPause:        "[pause]",
	KeyF1:           "[f1]",
	KeyF2:           "[f2]",
	KeyF3:           "[f3]",
	KeyF4:           "[f4]",
	KeyF5:           "[f5]",
	KeyF6:           "[f6]",
	KeyF7:           "[f7]",
	KeyF8:           "[f8]",
	KeyF9:           "[f9]",
	KeyF10:          "[f10]",
	KeyF11:          "[f11]",
	KeyF12:          "[f12]",
	KeyF13:          "[f13]",
	KeyF14:          "[f14]",
	KeyF15:          "[f15]",
	KeyF16:          "[f16]",
	KeyF17:          "[f17]",
	KeyF18:          "[f18]",
	KeyF19:          "[f19]",
	KeyF20:          "[f20]",
	KeyF21:          "[f21]",
	KeyF22:          "[f22]",
	KeyF23:          "[f23]",
	KeyF24:          "[f24]",
	KeyF25:          "[f25]",
	KeyKp0:          "[kp 0]",
	KeyKp1:          "[kp 1]",
	KeyKp2:          "[kp 2]",
	KeyKp3:          "[kp 3]",
	KeyKp4:          "[kp 4]",
	KeyKp5:          "[kp 5]",
	KeyKp6:          "[kp 6]",
	KeyKp7:          "[kp 7]",
	KeyKp8:          "[kp 8]",
	KeyKp9:          "[kp 9]",
	KeyKpDecimal:    "[kp decimal]",
	KeyKpDivide:     "[kp divide]",
	KeyKpMultiply:   "[kp multiply]",
	KeyKpSubtract:   "[kp subtract]",
	KeyKpAdd:        "[kp add]",
	KeyKpEnter:      "[kp enter]",
	KeyKpEqual:      "[kp equal]",
	KeyLeftShift:    "[left shift]",
	KeyLeftControl:  "[left control]",
	KeyLeftAlt:      "[left alt]",
	KeyLeftSuper:    "[left super]",
	KeyRightShift:   "[right shift]",
	KeyRightControl: "[right control]",
	KeyRightAlt:     "[right alt]",
	KeyRightSuper:   "[right super]",
	KeyMenu:         "[menu]",
}

func (key Key) String() string {
	s, ok := keyNames[key]
	if ok {
		return s
	}
	if key >= KeyApostrophe && key <= KeyGraveAccent {
		return string(key)
	}
	return fmt.Sprintf("[unknown key: %d]", int(key))
}

// KeyPress is triggered when a keyboard key is pressed.
type KeyPress struct {
	// Keyboard key.
	Key Key
	// Bitfield of modifier keys.
	Mod Mod
}

// KeyRelease is triggered when a keyboard key is released.
type KeyRelease struct {
	// Keyboard key.
	Key Key
	// Bitfield of modifier keys.
	Mod Mod
}

// KeyRepeat is triggered when a keyboard key was held down until it repeated.
type KeyRepeat struct {
	// Keyboard key.
	Key Key
	// Bitfield of modifier keys.
	Mod Mod
}

// KeyRune is triggered when a unicode character is typed on the keyboard. For
// instance if `a` and `shift` are held down on the keyboard KeyRune will
// correspond to 'A'.
type KeyRune rune
