package w4_test

import "github.com/orsinium-labs/wasm4go/w4"

func ExampleBlit() {
	var smiley = []byte{
		0b11000011,
		0b10000001,
		0b00100100,
		0b00100100,
		0b00000000,
		0b00100100,
		0b10011001,
		0b11000011,
	}
	size := w4.Size{Width: 8, Height: 8}
	w4.Blit(smiley, w4.Point{X: 76, Y: 76}, size, 0)
}

func ExampleDrawLine() {
	w4.DrawLine(
		w4.Point{X: 10, Y: 20},
		w4.Point{X: 20, Y: 10},
	)
}

func ExampleDrawHorLine() {
	w4.DrawHorLine(w4.Point{X: 10, Y: 20}, 30)
}

func ExampleDrawVertLine() {
	w4.DrawVertLine(w4.Point{X: 40, Y: 50}, 30)
}

func ExampleDrawEllipse() {
	// Draw a circle with 30 pixels diameter.
	w4.DrawEllipse(w4.Point{X: 10, Y: 10}, w4.Size{Width: 30, Height: 30})
}

func ExampleDrawRect() {
	// Draw a square 30 pixels side length.
	w4.DrawRect(w4.Point{X: 10, Y: 10}, w4.Size{Width: 30, Height: 30})
}

func ExampleDrawText() {
	w4.DrawText("Hello world!", w4.Point{X: 10, Y: 10})
}

func ExampleTrace() {
	w4.Trace("Some message")
}

func ExampleSave() {
	data := []byte("some data")
	w4.Save(data)
}

func ExampleLoad() {
	data := make([]byte, 1024)
	w4.Load(data)
}

func ExampleColor() {
	w4.Palette.Set(
		w4.Color{R: 0xfb, G: 0xf7, B: 0xf3},
		w4.Color{R: 0xe5, G: 0xb0, B: 0x83},
		w4.Color{R: 0x42, G: 0x6e, B: 0x5d},
		w4.Color{R: 0x20, G: 0x28, B: 0x3d},
	)
}

func ExamplePalette() {
	w4.Palette.Set(
		w4.Color{R: 0xfb, G: 0xf7, B: 0xf3},
		w4.Color{R: 0xe5, G: 0xb0, B: 0x83},
		w4.Color{R: 0x42, G: 0x6e, B: 0x5d},
		w4.Color{R: 0x20, G: 0x28, B: 0x3d},
	)
}

func ExampleGamepads() {
	g := w4.Gamepads[0]
	if g.X() {
		// the X button is pressed
	}
}

func ExampleDrawColors() {
	w4.DrawColors.SetPrimary(w4.Secondary)
	w4.DrawColors.SetSecondary(w4.Dark)
	size := w4.Size{Width: 20, Height: 20}
	w4.DrawRect(w4.Point{X: 10, Y: 10}, size)
}
