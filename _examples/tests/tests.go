package main

import "github.com/orsinium-labs/wasm4go/w4"

func main() {
	w4.Start = start
	w4.Update = update
}

func assert(x bool) {
	if !x {
		panic("assertion failed")
	}
}

func start() {
	testDraw()
	testMemory()
}

func testMemory() {
	// palette
	w4.Palette.Set(
		w4.Color{R: 45, G: 56, B: 67},
		w4.Color{R: 7, G: 8, B: 9},
		w4.Color{R: 11, G: 12, B: 13},
		w4.Color{R: 3, G: 4, B: 5},
	)
	assert(w4.Palette.Get(w4.Light) == w4.Color{R: 45, G: 56, B: 67})
	assert(w4.Palette.Get(w4.Primary) == w4.Color{R: 7, G: 8, B: 9})
	assert(w4.Palette.Get(w4.Secondary) == w4.Color{R: 11, G: 12, B: 13})
	assert(w4.Palette.Get(w4.Dark) == w4.Color{R: 3, G: 4, B: 5})

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

func testDraw() {
	for _, b := range w4.FrameBuffer[:500] {
		assert(b == 0)
	}

	// chack that horizontal line modifies the framebuffer as expected
	w4.DrawHorLine(w4.Point{X: 0, Y: 0}, 10)
	assert(w4.FrameBuffer[0] == 0b1010_1010)
	assert(w4.FrameBuffer[1] == 0b1010_1010)
	assert(w4.FrameBuffer[2] == 0b0000_1010)
	assert(w4.FrameBuffer[3] == 0b0000_0000)

	// smoke test the rest of the drawing functions
	w4.DrawHorLine(w4.Point{X: 10, Y: 20}, 30)
	w4.DrawVertLine(w4.Point{X: 10, Y: 20}, 30)
	w4.DrawLine(w4.Point{X: 10, Y: 20}, w4.Point{X: 30, Y: 40})
	w4.DrawEllipse(w4.Point{X: 50, Y: 60}, w4.Size{Width: 10, Height: 20})
	w4.DrawRect(w4.Point{X: 50, Y: 60}, w4.Size{Width: 10, Height: 20})
	w4.DrawText("Test me!", w4.Point{X: 50, Y: 60})
	w4.Trace("all tests have passed")
}

func update() {
}
