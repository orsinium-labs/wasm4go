package main

import "github.com/orsinium-labs/wasm4go/w4"

func init() {
	w4.Update = update
}

func update() {
	var offsetX uint8 = 60
	for i, g := range w4.Gamepads {
		var offsetY uint8 = uint8(72 + i*10)
		w4.DrawText(itoa(i), w4.Point{X: offsetX, Y: offsetY})
		if g.X() {
			w4.DrawText("\x80", w4.Point{X: offsetX + 8, Y: offsetY})
		}
		if g.Z() {
			w4.DrawText("\x81", w4.Point{X: offsetX + 16, Y: offsetY})
		}
		if g.Left() {
			w4.DrawText("\x84", w4.Point{X: offsetX + 24, Y: offsetY})
		}
		if g.Right() {
			w4.DrawText("\x85", w4.Point{X: offsetX + 32, Y: offsetY})
		}
		if g.Up() {
			w4.DrawText("\x86", w4.Point{X: offsetX + 38, Y: offsetY})
		}
		if g.Down() {
			w4.DrawText("\x87", w4.Point{X: offsetX + 46, Y: offsetY})
		}
	}
}

func itoa(i int) string {
	switch i {
	case 0:
		return "1"
	case 1:
		return "2"
	case 2:
		return "3"
	case 3:
		return "4"
	}
	panic("gamepad out of range")
}
