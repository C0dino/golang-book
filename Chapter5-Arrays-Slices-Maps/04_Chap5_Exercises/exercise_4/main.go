package main

import (
	"fmt"
	"sort"
)

// var n, smallest, bigest int

func main() {
	numbers := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	sliceNumbers := sort.IntSlice(numbers)
	sort.Sort(sliceNumbers)
	fmt.Printf("Smallest integer is %d\n", sliceNumbers[0])
	fmt.Printf("Largest integer is %d", sliceNumbers[len(sliceNumbers)-1])

	// ALTERNATE SOLUTION
	// 	for _, v := range x {
	// 		if v < n {
	// 			fmt.Println(n)
	// 			fmt.Println(v, ">", n)
	// 		} else {
	// 			fmt.Println(v, "<", n)
	// 			n = v
	// 			bigest = n
	// 		}

	// 	}
	// 	fmt.Println(bigest, "Is the biggest number")

	// 	for _, v := range x {
	// 		if v > n {
	// 			fmt.Println(v, ">", n)
	// 		} else {
	// 			fmt.Println(v, "<", n)
	// 			n = v
	// 			smallest = n
	// 		}

	// 	}
	// 	fmt.Println(smallest, "Is the smallest number")
}
