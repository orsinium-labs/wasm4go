package w4

import "unsafe"

// The wasm-4-controlled region of memory, from zero to the user data beginning.
//
// https://wasm4.org/docs/reference/memory#memory-map
var memory = (*[6560]byte)(unsafe.Pointer(uintptr(0)))

type Color uint32

// 	return uint(c.R) | (uint(c.G) << 8) | (uint(c.B) << 24)

// An array of 4 colors, each represented by a 32 bit integer.
type palette struct{}

var Palette = palette{}

func (palette) Get(d DrawColor) Color {
	if d == Transparent {
		panic("can't get Transparent color")
	}
	// return Color(memory[0x04+uint(d-1)])
}

func (palette) Set(d DrawColor, c Color) {
	if d == Transparent {
		panic("can't change Transparent color")
	}
	// memory[0x04+uint(d-1)] = c
}

type DrawColor u8

const (
	// Do not draw this color
	Transparent DrawColor = 0

	// The 1st color in the palette. Usually a very light, almost white, color.
	Light DrawColor = 1

	// The 2nd color in the palette. Usually a bright and intensive accent color.
	Primary DrawColor = 2

	// The 3rd color in the palette. Usually a distinc color but darker than primary.
	Secondary DrawColor = 3

	// The 4th color in the palette. A very dark color, used for night or contrast.
	Dark DrawColor = 4
)

// Indexes into the color palette used by all drawing functions.
type drawColors struct{}

// The primary draw color.
//
// Used for fill color of shapes or the main color of text or line.
func (drawColors) First() DrawColor { return DrawColor(memory[0x14] & 0x0f) }

// Set the primary draw color.
//
// Used for fill color of shapes or the main color of text or line.
func (drawColors) SetFirst(c DrawColor) { memory[0x14] = (memory[0x14] & 0x0f) | byte(c) }

// The secondary draw color.
//
// Used for stroke color of shapes or the background color of text.
func (drawColors) Second() DrawColor { return DrawColor(memory[0x14] & 0xf0 >> 4) }

// Set the secondary draw color.
//
// Used for stroke color of shapes or the background color of text.
func (drawColors) SetSecond(c DrawColor) { memory[0x14] = (memory[0x14] & 0xf0) | byte(c<<4) }

type gamepad uint

// Check if X button is currently pressent on the gamepad.
func (g gamepad) X() bool { return memory[g]&1 != 0 }

// Check if Z button is currently pressent on the gamepad.
func (g gamepad) Z() bool { return memory[g]&2 != 0 }

// Check if left button is currently pressent on the gamepad.
func (g gamepad) Left() bool { return memory[g]&16 != 0 }

// Check if right button is currently pressent on the gamepad.
func (g gamepad) Right() bool { return memory[g]&32 != 0 }

// Check if up button is currently pressent on the gamepad.
func (g gamepad) Up() bool { return memory[g]&64 != 0 }

// Check if down button is currently pressent on the gamepad.
func (g gamepad) Down() bool { return memory[g]&128 != 0 }

// X     bool
// Y     bool
// Left  bool
// Right bool
// Up    bool
// Down  bool

// 4 gamepads, with each gamepad represented by a single byte.
type gamepads [4]gamepad

var Gamepads = gamepads{0x16, 0x17, 0x18, 0x19}

// Byte containing the mouse position and mouse buttons state.
type mouse struct{}

var Mouse = mouse{}

func (mouse) X() u8 { return u8(memory[0x1a]) }

func (mouse) Y() u8 { return u8(memory[0x1c]) }

func (mouse) Left() bool { return u8(memory[0x1e])&1 != 0 }

func (mouse) Right() bool { return u8(memory[0x1e])&2 != 0 }

func (mouse) Middle() bool { return u8(memory[0x1e])&4 != 0 }

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
