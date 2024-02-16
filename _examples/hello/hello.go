package main

import "github.com/orsinium-labs/wasm4go/w4"

func main() {}

func init() {
	w4.Update = update
}

func update() {
	w4.Palette.Set(w4.Light, w4.Color{B: 255})
	w4.DrawText("Hello from TinyGo!", w4.Point{10, 10})
	if !w4.NetPlay.Active() {
		w4.DrawHorLine(w4.Point{20, 30}, 40)
	}
}
