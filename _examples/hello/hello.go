package main

import "github.com/orsinium-labs/wasm4go/w4"

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

func main() {
	w4.Update = update
}

func update() {
	w4.DrawColors.SetFirst(w4.Primary)
	w4.DrawText("Hello from TinyGo!", w4.Point{10, 10})
	gamepad := w4.Gamepads[0]
	if gamepad.X() {
		w4.DrawColors.SetFirst(w4.Dark)
	}
	w4.Blit(smiley, w4.Point{X: 76, Y: 76}, w4.Size{Width: 8, Height: 8}, 0)
	w4.DrawText("Press X to blink", w4.Point{16, 90})
}
