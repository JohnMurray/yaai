package main

import (
	"fmt"

	"github.com/JohnMurray/yaai/parser"
)

func main() {
	tree := parser.ParseFile("/tmp/blah.yaai")
	fmt.Println(tree.Print(0))
}
