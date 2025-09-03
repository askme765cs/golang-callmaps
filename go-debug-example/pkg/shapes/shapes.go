package shapes

import "math"

// Shape 是一个接口
type Shape interface {
	Area() float64
}

// Circle 实现了 Shape 接口
type Circle struct{ Radius float64 }
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

// Rectangle 实现了 Shape 接口
type Rectangle struct{ Width, Height float64 }
func (r Rectangle) Area() float64 { return r.Width * r.Height }