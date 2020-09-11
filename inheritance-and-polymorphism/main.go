package main

import "fmt"

type animal interface {
	makeSound()
}

type mammal struct {
	furColor string
}

func (m mammal) displayFurColor() {
	fmt.Printf("Fur color: %s.\n", m.furColor)
}

type dog struct {
	mammal
}

type cat struct {
	mammal
}

func (c cat) makeSound() {
	fmt.Println("Meow!")
}

func (d dog) makeSound() {
	fmt.Println("Woof!")
}

func wrapper(a animal) {
	a.makeSound()
}

// invalid: interfaces cannot be receivers
// func (a animal) doSomething() {
// 	fmt.Println("Invalid.")
// }

func main() {
	c := cat{mammal: mammal{furColor: "black"}}
	d := dog{mammal{furColor: "brown"}}
	var a animal

	c.displayFurColor()
	d.displayFurColor()

	a = c
	a.makeSound()

	a = d
	a.makeSound()

	wrapper(a)
}
