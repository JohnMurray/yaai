package test

import (
	"testing"

	"github.com/JohnMurray/yaai/antlr/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestLexerGeneral(t *testing.T) {
	lexerSnapshot(t, "general", parser.NewYaaiLexer(antlr.NewInputStream(`package myPackage

type myStruct struct {
	count uint32
	msg   string
}

actor myActor {
	var helloMsg string
	var count := 0

	init {
		helloMsg = "Hello, World! \"{}()\"--++"
	}

	receive (x int) {
		// Do some stuff
	}

	receive (data myStruct) {
		// handle the data
	}

	func incCount(by int, mult int) {
		self.count += (by * mult)
	}

	unhandled (ti TypeInfo, msg interface{}) {
		// handled unkown/unsupported messages
	}
}

func doThing(strct *myStruct) (myStruct, error) {
	// Do the thing... (ya know)
	for i := 0; i < (*strct).count; i++ {
		fmt.Println((*strct).msg)
	}
}

`)))
}
