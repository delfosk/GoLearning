package main

import "math"

//Perimeter function
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

//Circle struct
type Circle struct {
	Radius float64
}

//Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

//Area of rectangle
func (r Rectangle) Area() float64 {

	return r.Width * r.Height
}

//Area of triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

//Area of circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Shape interface
type Shape interface {
	Area() float64
}

func main() {}
