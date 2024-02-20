package w4

type u8 = uint8

// A point on the plane.
type Point struct {
	X u8
	Y u8
}

// A point with equal x value and y set to 0.
func (p Point) XAxis() Point {
	p.Y = 0
	return p
}

// A point with equal y value and x set to 0.
func (p Point) YAxis() Point {
	p.X = 0
	return p
}

// The componentwise minimum of two Points.
func (p Point) ComponentMin(other Point) Point {
	if other.X < p.X {
		p.X = other.X
	}
	if other.Y < p.Y {
		p.Y = other.Y
	}
	return p
}

// The componentwise maximum of two Points.
func (p Point) ComponentMax(other Point) Point {
	if other.X > p.X {
		p.X = other.X
	}
	if other.Y > p.Y {
		p.Y = other.Y
	}
	return p
}

// Convert the Point to a Size with width equal to x and height equal to y.
func (p Point) AsSize() Size {
	return Size{Width: p.X, Height: p.Y}
}

// Add size width to the point x and size height to the point y.
func (p Point) AddSize(s Size) Point {
	p.X += s.Width
	p.Y += s.Height
	return p
}

// If the point is outside of the screen, wrap it around to fit on the screen.
func (p Point) Wrap() Point {
	p.X %= 160
	p.Y %= 160
	return p
}

// Size of a 2D shape.
type Size struct {
	Width  u8
	Height u8
}

// Add two sizes together componentwise.
func (s Size) AddSize(other Size) Size {
	s.Width += other.Width
	s.Height += other.Height
	return s
}

// Convert Size to a Point with x set to width and y set to height.
func (s Size) AsPoint() Point {
	return Point{X: s.Width, Y: s.Height}
}

// The componentwise minimum of two Sizes.
func (s Size) ComponentMin(other Size) Size {
	if other.Width < s.Width {
		s.Width = other.Width
	}
	if other.Height < s.Height {
		s.Height = other.Height
	}
	return s
}

// The componentwise maximum of two Sizes.
func (s Size) ComponentMax(other Size) Size {
	if other.Width > s.Width {
		s.Width = other.Width
	}
	if other.Height > s.Height {
		s.Height = other.Height
	}
	return s
}
