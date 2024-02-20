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
	w4.DrawColors.SetPrimary(w4.Primary)
	w4.DrawText("Hello from TinyGo!", w4.Point{X: 10, Y: 10})
	if w4.Gamepad.X() {
		w4.DrawColors.SetPrimary(w4.Dark)
	}
	w4.Blit(smiley, w4.Point{X: 76, Y: 76}, w4.Size{Width: 8, Height: 8}, 0)
	w4.DrawText("Press X to blink", w4.Point{X: 16, Y: 90})
}
