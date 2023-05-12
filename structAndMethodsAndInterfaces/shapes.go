package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Triangle struct {
	Base float64
	Height float64
}

type Circles struct {
	Redius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circles) Area() float64 {
	return c.Redius * c.Redius * math.Pi
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}