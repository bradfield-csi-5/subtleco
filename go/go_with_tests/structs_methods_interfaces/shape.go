package shapes

import "math"

type Rectangle struct {
  Width float64
  Height float64
}

func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
  return 2 * (r.Height + r.Width)
}

type Circle struct {
  Radius float64
}

func (c Circle) Area() float64 {
  return math.Pi * math.Pow(c.Radius, 2)
}

type Triangle struct {
  Base float64
  Height float64
}

type Shape interface {
  Area() float64
}

func (t Triangle) Area() float64 {
  return (t.Base * t.Height) / 2
}