package main

import (
	"fmt"
)

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 // By default, total is set to its "zero" value - which in this case IS 0
	for i, value := range x { //i is declared but not used
		total += value
	}
	fmt.Println(total / float64(len(x))) // len is of type int, it must be converted to type float64 to work with the total variable which is of type float64
}
