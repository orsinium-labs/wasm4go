# wasm4go examples

This directory contains examples of using [wasm4go](../). Some are small and useless, some are almost real games.

Available examples:

* [hello](./hello/hello.go): the same code example as wasm-4 CLI generates by default.
* [buttons](./buttons/buttons.go): displays on the screen the button you press.
* [snake](./snake/snake.go): re-implementation of the [Snake tutorial](https://wasm4.org/docs/tutorials/snake/goal) from thw official WASM-4 documentation.
* [tests](./tests/tests.go): integration tests for the project.

## Running locally

1. [Install the latest dev version of TinyGo](https://tinygo.org/getting-started/install/linux/#development-builds).
1. [Install WASM-4 CLI](https://wasm4.org/docs/getting-started/setup)
1. [Install task](https://taskfile.dev/installation/)
1. Run `task run -- ./snake` to run the snake demo.
