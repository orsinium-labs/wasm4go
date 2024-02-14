package w4

import "unsafe"

// The wasm-4-controlled region of memory, from zero to the user data beginning.
//
// https://wasm4.org/docs/reference/memory#memory-map
var memory = (*[6560]byte)(unsafe.Pointer(uintptr(0)))

type Color struct {
	R u8
	G u8
	B u8
}

// Serialize the color into uint number as it is represented in wasm-4 memory.
func (c Color) Serialize() uint {
	return uint(c.R) | (uint(c.G) << 8) | (uint(c.B) << 24)
}

// An array of 4 colors, each represented by a 32 bit integer.
type Palette struct {
	Light     Color
	Primary   Color
	Secondary Color
	Dark      Color
}

type DrawColor u8

const (
	Transparent DrawColor = 0
	Light       DrawColor = 1
	Primary     DrawColor = 2
	Secondary   DrawColor = 3
	Dark        DrawColor = 4
)

// Indexes into the color palette used by all drawing functions.
type DrawColors struct {
	Fill   DrawColor
	Stroke DrawColor
}

type Gamepad struct {
	X     bool
	Y     bool
	Left  bool
	Right bool
	Up    bool
	Down  bool
}

// 4 gamepads, with each gamepad represented by a single byte.
type Gamepads [4]Gamepad

// Byte containing the mouse position and mouse buttons state.
type Mouse struct {
	X      int16
	Y      int16
	Left   bool
	Right  bool
	Middle bool
}

// Byte containing flags that modify WASM-4's operation. By default all flags are off.
type systemFlags struct{}

var SystemFlags = systemFlags{}

func (systemFlags) PreserveFrameBuffer(v bool) {
	if v {
		memory[0x1f] |= 0b1
	} else {
		memory[0x1f] &= 0b0
	}
}

func (systemFlags) HideGamepadOverlay(v bool) {
	if v {
		memory[0x1f] |= 0b10
	} else {
		memory[0x1f] &= 0b01
	}
}

// Byte containing netplay multiplayer state.
type netplay struct{}

var NetPlay = netplay{}

// Local player index (0 to 3).
func (netplay) Player() u8 {
	n := memory[0x20]
	return u8(n & 0b11)
}

// True if netplay is currently active.
func (netplay) Active() bool {
	n := memory[0x20]
	return n&0b100 != 0
}

// Array of 160x160 pixels, with each pixel packed into 2 bits (colors 0 to 3).
type frameBuffer [6400]byte

var FrameBuffer = memory[0xa0:]
