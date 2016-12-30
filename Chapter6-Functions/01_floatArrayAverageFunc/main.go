package main

import (
	"fmt"
)

func average(xs []float64) float64 {

	var total float64          // By default, total is set to its "zero" value - which in this case IS 0
	for _, value := range xs { // _ replaces i - _ is a way to discard i's value
		total += value
	}
	return total / float64(len(xs)) // len is of type int, it must be converted to type float64 to work with the total variable which is of type float64
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

}
