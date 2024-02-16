package main

import "github.com/orsinium-labs/wasm4go/w4"

func main() {
	w4.Update = update
}

func update() {
	var offsetX uint8 = 60
	var offsetY uint8 = 72
	if w4.Gamepads[0].X() {
		w4.DrawText("\x80", w4.Point{X: offsetX, Y: offsetY})
	}
	if w4.Gamepads[0].Z() {
		w4.DrawText("\x81", w4.Point{X: offsetX + 8, Y: offsetY})
	}
	if w4.Gamepads[0].Left() {
		w4.DrawText("\x84", w4.Point{X: offsetX + 16, Y: offsetY})
	}
	if w4.Gamepads[0].Right() {
		w4.DrawText("\x85", w4.Point{X: offsetX + 24, Y: offsetY})
	}
	if w4.Gamepads[0].Up() {
		w4.DrawText("\x86", w4.Point{X: offsetX + 32, Y: offsetY})
	}
	if w4.Gamepads[0].Down() {
		w4.DrawText("\x87", w4.Point{X: offsetX + 38, Y: offsetY})
	}
}
