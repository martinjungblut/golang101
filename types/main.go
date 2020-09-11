package main

import "fmt"

// sum takes two int, a and b, and returns an int
func sum(a int, b int) int {
	return a + b
}

// sum32 takes two int32, a and b, and returns an int32
func sum32(a int32, b int32) int32 {
	return a + b
}

// sumptr works essentially the same as sum(), but works on pointers
// using type inference here
func sumptr(a *int, b *int) *int {
	// this is the alternative, if not using type inference
	// var result int
	// result = *a + *b   <- notice the assignment operator is different

	result := *a + *b
	return &result
}

func exampleSum() {
	fmt.Printf("%d is the result of %d + %d.\n", sum(15, 2), 15, 2)

	a, b := 6, 3
	c, d := &a, &b
	fmt.Printf("%d is the result of %d + %d.\n", *sumptr(&a, &b), a, b)
	fmt.Printf("%d is the result of %d + %d.\n", *sumptr(c, d), *c, *d)

	fmt.Printf("%d is the result of %d + %d.\n", sum(2147483647, 1), 2147483647, 1)
	fmt.Printf("%d is the result of %d + %d.\n", sum32(2147483647, 1), 2147483647, 1)
}

type person struct {
	name string
	age  int
}

func hello(p person) {
	fmt.Printf("Hello, my name is %s and I'm %d years old.\n", p.name, p.age)
}

func exampleStruct() {
	// may also be written as...
	// martin := person{name: "Martin", age: 29}
	var martin person
	martin.name = "Martin"
	martin.age = 29

	hello(martin)

	martin.name = "Martin Jungblut"
	martin.age = 30
	hello(martin)
}

// arrays are values themselves (not pointers, as they are in C), and they are fully copied over
// to the new stack frames...
// there is no way to bypass this size-oriented limitation as of Go 1.14
func reverseStringArray(strings [3]string) [3]string {
	largestIndex := len(strings) - 1

	// could also be written as: var reversed [3]string
	reversed := [3]string{}

	for i := 0; i <= largestIndex; i++ {
		reversed[largestIndex-i] = strings[i]
	}

	return reversed
}

func exampleArrays() {
	// array literal, could also be written as: names := [3]string{"Martin", "John", "Robert"}
	var names [3]string = [...]string{"Martin", "John", "Robert"}

	for index, name := range names {
		fmt.Printf("Name at index %d is %s.\n", index, name)
	}
	fmt.Println("...")

	for index, name := range reverseStringArray(names) {
		fmt.Printf("Name at index %d is %s.\n", index, name)
	}
}

func reverseStringSlice(k []string) []string {
	m := make([]string, 0, len(k))

	for i := len(k) - 1; i >= 0; i-- {
		m = append(m, k[i])
	}

	return m
}

func reverseStringSliceInPlace(strings *[]string) {
	dstrings := *strings
	largestIndex := len(dstrings) - 1

	for i := 0; i <= largestIndex/2; i++ {
		dstrings[largestIndex-i], dstrings[i] = dstrings[i], dstrings[largestIndex-i]
	}
}

func exampleSlices() {
	// slice literal
	var names []string = []string{"Martin", "John", "Robert"}

	for index, name := range names {
		fmt.Printf("Name at index %d is %s.\n", index, name)
	}
	fmt.Println("...")

	for index, name := range reverseStringSlice(names) {
		fmt.Printf("Name at index %d is %s.\n", index, name)
	}
	fmt.Println("...")

	reverseStringSliceInPlace(&names)
	for index, name := range names {
		fmt.Printf("Name at index %d is %s.\n", index, name)
	}
}

func main() {
	exampleSum()
	fmt.Println("---")

	exampleStruct()
	fmt.Println("---")

	exampleArrays()
	fmt.Println("---")

	exampleSlices()
	fmt.Println("---")
}
