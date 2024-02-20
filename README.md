# wasm4go

Framework for making [WASM-4](https://wasm4.org/) games with Go (and [TinyGo](https://tinygo.org/)).

Unlike the Go bindings that WASM-4 generates by default, this library is idiomatic Go with type safety, no direct memory manipulation and no byte operations. It provides a friendly API that is easy to read and hard to misuse.

Features:

* **Type-safe API** that is easy to use and hard to misuse.
* **Zero-cost abstraction** in most of the cases, thanks to the crazy optimizations that TinyGo does.
* **Helpful documentation** with links and examples.

## Installation

```bash
go get github.com/orsinium-labs/wasm4go
```

## Usage

```go
package main

import "github.com/orsinium-labs/wasm4go/w4"

func main() {
   w4.Update = update
}

func update() {
    w4.DrawColors.SetPrimary(w4.Primary)
    w4.DrawText("Hello from TinyGo!", w4.Point{X: 10, Y: 10})
    if w4.Gamepad.X() {
        w4.DrawColors.SetPrimary(w4.Dark)
    }
    w4.DrawText("Press X to blink", w4.Point{X: 16, Y: 90})
}
```

See [examples](./_examples) for some real code and [documentation](https://pkg.go.dev/github.com/orsinium-labs/wasm4go/w4) to learn what's inside.
