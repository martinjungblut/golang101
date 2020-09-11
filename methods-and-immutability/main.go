package main

import "fmt"

type person struct {
	name string
}

func (p person) greet() {
	fmt.Printf("Hi, my name is %s.\n", p.name)
}

func (p person) setName(name string) {
	p.name = name
}

func (p *person) setNamePtr(name string) {
	p.name = name
}

func main() {
	martin := person{name: "Martin"}
	martin.greet()

	martin.setName("Brandon")
	martin.greet()

	martin.setNamePtr("Martin Jungblut")
	martin.greet()
}
