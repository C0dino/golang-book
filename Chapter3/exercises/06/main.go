package main

import "fmt"

func main() {

	var feet float64
	const meters = 0.3048

	fmt.Print("enter feet: ")
	fmt.Scan(&feet)
	fmt.Println(feet, "feet =", feet*meters, " meters")
}
