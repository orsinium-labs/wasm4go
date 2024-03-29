package main

import "github.com/orsinium-labs/wasm4go/w4"

func init() {
	w4.Update = update
}

func update() {
	// Set the palette to black-red-green-blue
	w4.Palette.Set(
		w4.Color{R: 0x00, G: 0x00, B: 0x00},
		w4.Color{R: 0xff, G: 0x00, B: 0x00},
		w4.Color{R: 0x00, G: 0xff, B: 0x00},
		w4.Color{R: 0x00, G: 0x00, B: 0xff},
	)

	// disable outline color
	w4.DrawColors.SetSecondary(w4.Transparent)

	size := w4.Size{Width: 30, Height: 30}

	// draw background-colored squary (should not be visible, naturally)
	// and letter K
	w4.DrawColors.SetPrimary(1)
	p := w4.Point{X: 50, Y: 40}
	w4.DrawRect(p, size)
	w4.DrawColors.SetPrimary(4)
	w4.DrawText("K", p)

	// draw red square and letter R
	w4.DrawColors.SetPrimary(2)
	p = w4.Point{X: 50, Y: 70}
	w4.DrawRect(p, size)
	w4.DrawColors.SetPrimary(1)
	w4.DrawText("R", p)

	// draw green square and letter G
	w4.DrawColors.SetPrimary(3)
	p = w4.Point{X: 80, Y: 40}
	w4.DrawRect(p, size)
	w4.DrawColors.SetPrimary(1)
	w4.DrawText("G", p)

	// draw blue square and letter B
	w4.DrawColors.SetPrimary(4)
	p = w4.Point{X: 80, Y: 70}
	w4.DrawRect(p, size)
	w4.DrawColors.SetPrimary(1)
	w4.DrawText("B", p)

}
