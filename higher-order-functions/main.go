package main

import "fmt"

func currySum(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func applySquared(i int, c func(int) int) int {
	return c(i * i)
}

func fhalve(f float64) float64 {
	return f / 2
}

func wrapper(c func(float64) float64) func(int) int {
	return func(i int) int {
		return int(c(float64(i)))
	}
}

func main() {
	add3 := currySum(3)
	fmt.Printf("%d\n", add3(2))

	fmt.Printf("%d\n", applySquared(4, add3))
	fmt.Printf("%d\n", applySquared(4, currySum(5)))

	fmt.Printf("%f\n", fhalve(16))

	fmt.Printf("%d\n", applySquared(4, wrapper(fhalve)))
}
