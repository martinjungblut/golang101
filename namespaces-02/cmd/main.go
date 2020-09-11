package main

import (
	"namespaces-02/first"
	"namespaces-02/first/subfirst"
	"namespaces-02/second"
	"namespaces-02/second/subsecond"
)

func main() {
	first.Hello()
	subfirst.Hello()
	second.Hello()
	subsecond.Hello()
}
