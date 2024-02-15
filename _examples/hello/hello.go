package main

import "github.com/orsinium-labs/wasm4go/w4"

func main() {}

func init() {
	w4.Update = update
}

func update() {
	w4.DrawText("Hello from TinyGo!", w4.Point{10, 10})
}
