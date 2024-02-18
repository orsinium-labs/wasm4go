package main

import (
	"math/rand"

	"github.com/orsinium-labs/wasm4go/w4"
)

// The length of a side of one snake segment, in pixels.
const size = 8

var (
	snake = Snake{}

	// The global counter of frames. Used to skip some because wasm-4 has 60 FPS
	// but that's too fast.
	frameCount = 0

	// The fruit position. Place the first one in the center of the screen.
	fruit = w4.Point{X: 80, Y: 80}

	// The raw binary content of a sprite for the fruit.
	// A nicer solution would be to use go:embed but it doesn't work with wasm target.
	fruitSprite = []byte{0x00, 0xa0, 0x02, 0x00, 0x0e, 0xf0, 0x36, 0x5c, 0xd6, 0x57, 0xd5, 0x57, 0x35, 0x5c, 0x0f, 0xf0}

	// Randomization function for placing the fruit.
	// We can't use a global random instance because there is no way to make a global
	// random seed in pure wasm.
	randInt = rand.New(rand.NewSource(1)).Intn
)

func main() {
	// Set the function to be called on start.
	w4.Start = start

	// Set the function to be called on each frame update.
	w4.Update = update
}

// Update the snake direction based on the buttons pressed on the gamepad.
func input() {
	g := w4.Gamepads[0]
	if g.Up() {
		snake.Up()
	}
	if g.Down() {
		snake.Down()
	}
	if g.Left() {
		snake.Left()
	}
	if g.Right() {
		snake.Right()
	}
}

func start() {
	// https://wasm4.org/docs/tutorials/snake/setting-color-palette
	w4.Palette.Set(
		// White color for the background.
		w4.Color{R: 0xfb, G: 0xf7, B: 0xf3},
		// A soft orange color for the fruit body.
		w4.Color{R: 0xe5, G: 0xb0, B: 0x83},
		// Dark green color for the snake's segments and fruit leaf.
		w4.Color{R: 0x42, G: 0x6e, B: 0x5d},
		// Dark blue color for the snake's body.
		w4.Color{R: 0x20, G: 0x28, B: 0x3d},
	)
	snake.Reset()
}

func update() {
	input()
	frameCount++
	// Skip every 10th frame so that the snake doesn't move too fast
	if frameCount%10 == 0 {
		snake.Update()
		// If snake eats itself, reset it to the initial state.
		if snake.IsDead() {
			snake.Reset()
		}
		// If the snake's head is on the tile with the fruit,
		// increase the snake's length and update the fruit's position.
		if snake.Body[0] == fruit {
			snake.Body = append(snake.Body, snake.Body[len(snake.Body)-1])
			fruit.X = uint8(randInt(20) * size)
			fruit.Y = uint8(randInt(20) * size)
		}
	}
	snake.Draw()
	// Draw the fruit.
	w4.DrawColors.Set(w4.Light, w4.Primary, w4.Secondary, w4.Dark)
	w4.Blit(fruitSprite, fruit, w4.Size{Width: size, Height: size}, w4.TwoBPP)
}

type Direction struct {
	X int
	Y int
}

type Snake struct {
	Body      []w4.Point
	Direction Direction
}

// Place the snake at the start with 3-segment length and moving to the right.
func (s *Snake) Reset() {
	s.Body = []w4.Point{
		{X: size * 2, Y: 0},
		{X: size, Y: 0},
		{X: 0, Y: 0},
	}
	s.Direction = Direction{X: size, Y: 0}
}

// Draw the snake's body
func (s *Snake) Draw() {
	// Draw green rectangles with blue outline for body segments.
	w4.DrawColors.SetFirst(w4.Secondary)
	w4.DrawColors.SetSecond(w4.Dark)
	rsize := w4.Size{Width: size, Height: size}
	for _, part := range s.Body {
		w4.DrawRect(part, rsize)
	}

	// Draw blue rectangle for the head.
	w4.DrawColors.SetFirst(w4.Dark)
	w4.DrawColors.SetSecond(w4.Transparent)
	head := s.Body[0]
	w4.DrawRect(head, rsize)
}

func (s *Snake) Update() {
	// Shift the snake's segments
	for i := len(s.Body) - 1; i > 0; i-- {
		s.Body[i] = s.Body[i-1]
	}

	// Shift the snake's head in the movement direction,
	// wrapping it around the screen if necessary.
	s.Body[0].X = uint8((int(s.Body[0].X) + s.Direction.X) % 160)
	s.Body[0].Y = uint8((int(s.Body[0].Y) + s.Direction.Y) % 160)
	// It is more than 160 if the integer overflows.
	if s.Body[0].X > 160 {
		s.Body[0].X = 160 - size
	}
	if s.Body[0].Y > 160 {
		s.Body[0].Y = 160 - size
	}
}

func (s *Snake) Up() {
	if s.Direction.Y == 0 {
		s.Direction.X, s.Direction.Y = 0, -size
	}
}

func (s *Snake) Down() {
	if s.Direction.Y == 0 {
		s.Direction.X, s.Direction.Y = 0, size
	}
}

func (s *Snake) Left() {
	if s.Direction.X == 0 {
		s.Direction.X, s.Direction.Y = -size, 0
	}
}

func (s *Snake) Right() {
	if s.Direction.X == 0 {
		s.Direction.X, s.Direction.Y = size, 0
	}
}

// Check if the snake's head is on the same position as one of its segments.
func (s *Snake) IsDead() bool {
	for index := 1; index < len(s.Body); index++ {
		if s.Body[0] == s.Body[index] {
			return true
		}
	}
	return false
}
