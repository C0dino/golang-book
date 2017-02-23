package main

import (
	"fmt"
)

type person struct {
	Name string
}

type android struct {
	person
	model string
}

func (p *person) talk() {
	fmt.Println("Hi my name is", p.Name)
}

func main() {
	p := person{"Steve"}
	p.talk()

	rTwo := android{p, "Droid"}
	rTwo.talk()
	fmt.Println(rTwo.model)

	a := new(android)
	a.person.talk()

	a = new(android)
	a.talk()
}
