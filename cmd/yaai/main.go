package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
)

var (
	path *string = kingpin.Flag("path", "path to shared object").
		Required().
		String()
)

func main() {
	kingpin.Parse()

	fmt.Println(*path)
}
