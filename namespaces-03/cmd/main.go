package main

import (
	"namespaces-03/first"
	"namespaces-03/second"
	"namespaces-03/second/subsecond"
)

func main() {
	first.Hello()
	first.Goodbye()
	second.Hello()
	subsecond.Hello()
}
