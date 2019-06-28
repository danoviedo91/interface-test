package models

// Square ...
type Square struct {
	Side float64
}

// ShapeType ...
func (s Square) ShapeType() string {
	return "Square"
}

// Area ...
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// Perimeter ...
func (s Square) Perimeter() float64 {
	return s.Side * 4
}

// Rectangle ...
type Rectangle struct {
	Length float64
	Width  float64
}

// ShapeType ...
func (r Rectangle) ShapeType() string {
	return "Rectangle"
}

// Area ...
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Perimeter ...
func (r Rectangle) Perimeter() float64 {
	return (r.Length * 2) + (r.Width * 2)
}
