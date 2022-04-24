package main

import (
	"fmt"
	"strings"

	"github.com/JohnMurray/yaai/parser"
)

func main() {
	tree := parser.ParseFile("/tmp/blah.yaai")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println(tree.Print(0))
	fmt.Println(strings.Repeat("-", 80))
}
