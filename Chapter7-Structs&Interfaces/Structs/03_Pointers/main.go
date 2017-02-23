package main

import (
	"fmt"
	"math"
)

type circle struct {
	x, y, r float64
	// or
	// x float64
	// y float64
	// r float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(c *circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {

	c := circle{0, 0, 5}
	//or
	//	c := circle{x: 0, y: 0, r: 5}

	fmt.Println(circleArea(&c))

	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))

}
