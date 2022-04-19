package test

import (
	"testing"

	"github.com/JohnMurray/yaii/yaai/parser"
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
		helloMsg = "Hello, World!"
	}

	receive (x int) {
		// Do some stuff
	}

	receive (data myStruct) {
		// handle the data
	}
}

`)))
}
