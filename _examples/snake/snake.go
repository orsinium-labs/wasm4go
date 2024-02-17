package main

import "github.com/orsinium-labs/wasm4go/w4"

var (
	snake       = &Snake{}
	frameCount  = 0
	fruit       = w4.Point{X: 10, Y: 10}
	fruitSprite = []byte{0x00, 0xa0, 0x02, 0x00, 0x0e, 0xf0, 0x36, 0x5c, 0xd6, 0x57, 0xd5, 0x57, 0x35, 0x5c, 0x0f, 0xf0}
)

func main() {
	w4.Start = start
	w4.Update = update
}

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
	w4.Palette.Set(1, w4.Color{R: 0xf3, G: 0xf7, B: 0xfb})
	w4.Palette.Set(2, w4.Color{R: 0x83, G: 0xb0, B: 0xe5})
	w4.Palette.Set(3, w4.Color{R: 0x5d, G: 0x6e, B: 0x42})
	w4.Palette.Set(4, w4.Color{R: 0x3d, G: 0x28, B: 0x20})

	snake.Reset()
}

func update() {
	frameCount++

	input()

	if frameCount%10 == 0 {
		snake.Update()

		if snake.IsDead() {
			snake.Reset()
		}

		if snake.Body[0] == fruit {
			snake.Body = append(snake.Body, snake.Body[len(snake.Body)-1])
			fruit.X = 20
			fruit.Y = 20
		}
	}
	snake.Draw()

	w4.DrawColors.Set(w4.Transparent, w4.Primary, w4.Secondary, w4.Dark)
	w4.Blit(fruitSprite, w4.Point{X: fruit.X * 8, Y: fruit.Y * 8}, w4.Size{Width: 8, Height: 8}, w4.TwoBPP)
}

type Direction struct {
	X int
	Y int
}

type Snake struct {
	Body      []w4.Point
	Direction Direction
}

func (s *Snake) Reset() {
	s.Body = []w4.Point{
		{X: 2, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: 0},
	}
	s.Direction = Direction{X: 1, Y: 0}
}

func (s *Snake) Draw() {
	// *w4.DRAW_COLORS = 0x0043
	w4.DrawColors.SetFirst(w4.Secondary)
	w4.DrawColors.SetSecond(w4.Dark)
	size := w4.Size{Width: 8, Height: 8}
	for _, part := range s.Body {
		w4.DrawRect(w4.Point{X: part.X * 8, Y: part.Y * 8}, size)
	}

	w4.DrawColors.SetFirst(w4.Dark)
	w4.DrawColors.SetSecond(w4.Transparent)
	head := s.Body[0]
	w4.DrawRect(w4.Point{X: head.X * 8, Y: head.Y * 8}, size)
}

func (s *Snake) Update() {
	for i := len(s.Body) - 1; i > 0; i-- {
		s.Body[i] = s.Body[i-1]
	}

	s.Body[0].X = uint8((int(s.Body[0].X) + s.Direction.X) % 20)
	s.Body[0].Y = uint8((int(s.Body[0].Y) + s.Direction.Y) % 20)
	if s.Body[0].X < 0 {
		s.Body[0].X = 19
	}
	if s.Body[0].Y < 0 {
		s.Body[0].Y = 19
	}
}

func (s *Snake) Up() {
	if s.Direction.Y == 0 {
		s.Direction.X, s.Direction.Y = 0, -1
	}
}

func (s *Snake) Down() {
	if s.Direction.Y == 0 {
		s.Direction.X, s.Direction.Y = 0, 1
	}
}

func (s *Snake) Left() {
	if s.Direction.X == 0 {
		s.Direction.X, s.Direction.Y = -1, 0
	}
}

func (s *Snake) Right() {
	if s.Direction.X == 0 {
		s.Direction.X, s.Direction.Y = 1, 0
	}
}

func (s *Snake) IsDead() bool {
	for index := 1; index < len(s.Body)-1; index++ {
		if s.Body[0] == s.Body[index] {
			return true
		}
	}
	return false
}
