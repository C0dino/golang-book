package main

import (
	"fmt"
)

func main() {
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxgyen",
		"F":  "Flourine",
		"Ne": "Neon",
	}
	fmt.Println(elements["Li"])
}
