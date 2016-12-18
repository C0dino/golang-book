package main

import (
	"fmt"
)

func main() {
	fmt.Println("1 + 1 =", 1+1)     // go infers the type int based on the 1
	fmt.Println("1 + 1 =", 1.0+1.0) // go infers floating-point numbers based on the .0 in 1.0
}
