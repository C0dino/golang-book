package main

import (
	"fmt"
)

func main() {
	var f float64

	fmt.Print("enter temperature in farenheight: ")
	fmt.Scan(&f)
	c := (f - 32) * 5 / 9
	fmt.Println("the equivalent temperature in celsius is: ", c)

}
