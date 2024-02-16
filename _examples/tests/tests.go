package main

import "github.com/orsinium-labs/wasm4go/w4"

func main() {}

func init() {
	w4.Start = start
	w4.Update = update
}

func assert(x bool) {
	if !x {
		panic("assertion failed")
	}
}

func start() {
	testMemory()
}

func testMemory() {
	// palette
	w4.Palette.Set(w4.Light, w4.Color{R: 45, G: 56, B: 67})
	assert(w4.Palette.Get(w4.Light) == w4.Color{R: 45, G: 56, B: 67})
	w4.Palette.Set(w4.Dark, w4.Color{R: 3, G: 4, B: 5})
	assert(w4.Palette.Get(w4.Dark) == w4.Color{R: 3, G: 4, B: 5})
	w4.Palette.Set(w4.Primary, w4.Color{R: 7, G: 8, B: 9})
	assert(w4.Palette.Get(w4.Primary) == w4.Color{R: 7, G: 8, B: 9})
	w4.Palette.Set(w4.Secondary, w4.Color{R: 11, G: 12, B: 13})
	assert(w4.Palette.Get(w4.Secondary) == w4.Color{R: 11, G: 12, B: 13})

	// netplay
	assert(!w4.NetPlay.Active())
	assert(w4.NetPlay.Player() == 0)

	// system flags
	w4.SystemFlags.HideGamepadOverlay(true)
	w4.SystemFlags.HideGamepadOverlay(false)
	w4.SystemFlags.PreserveFrameBuffer(true)
	w4.SystemFlags.PreserveFrameBuffer(false)

	// mouse
	assert(!w4.Mouse.Left())
	assert(!w4.Mouse.Right())
	assert(!w4.Mouse.Middle())
	assert(w4.Mouse.X() == 0)
	assert(w4.Mouse.Y() == 0)

	// gamepad
	for _, g := range w4.Gamepads {
		assert(!g.X())
		assert(!g.Z())
		assert(!g.Left())
		assert(!g.Right())
		assert(!g.Up())
		assert(!g.Down())
	}

	// draw colors
	assert(w4.DrawColors.First() == w4.Secondary)
	assert(w4.DrawColors.Second() == w4.Transparent)
	w4.DrawColors.SetFirst(w4.Dark)
	assert(w4.DrawColors.First() == w4.Dark)
	assert(w4.DrawColors.Second() == w4.Transparent)
	w4.DrawColors.SetSecond(w4.Light)
	assert(w4.DrawColors.Second() == w4.Light)
}

func update() {
}
