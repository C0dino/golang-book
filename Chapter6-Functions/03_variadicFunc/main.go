package main

import (
	"fmt"
)

func add(args ...int) int {
	var total int
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	fmt.Println(add(1, 2, 3))
}
