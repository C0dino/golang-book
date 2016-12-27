package main

import (
	"fmt"
)

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxgyen"
	elements["F"] = "Flourine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])

	name, ok := elements["Un"] // Comma OK idiom
	fmt.Println(name, ok)      // false - No key "Un"

	name, ok = elements["C"]
	fmt.Println(name, ok)
}
