package main

import (
	"C"
	"fmt"
)

//export SharedThing
func SharedThing(msg string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(msg)
	}
}

func main() {}
