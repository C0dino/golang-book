package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func totalArea(circles []Circle, rectangles []Rectangle) float64 {
	var total float64
	for _, c := range circles {
		total += c.area()
	}
	return total
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y2)
	return l * w
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

}
