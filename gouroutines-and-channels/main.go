package main

import (
	"fmt"
	"os"
)

func displayPid() {
	fmt.Printf("PID: %d\n", os.Getpid())
}

func chanSum(i int, ch chan int) {
	ch <- i + <-ch
}

func chanFeed(i int, ch chan int) {
	ch <- i
}

func main() {
	ch := make(chan int)
	defer close(ch)

	displayPid()
	go displayPid()

	go chanFeed(10, ch)
	go chanSum(5, ch)

	output := <-ch
	fmt.Printf("Channel's output: %v.\n", output)
}
