package main

import (
	"fmt"

	"github.com/JohnMurray/yaii/yaai/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type testListener struct {
	*parser.BaseYaaiListener
	*antlr.DefaultErrorListener
}

func main() {

	tl := &testListener{}

	// setup the input
	is := antlr.NewInputStream(`package myPackage

type MyStruct struct {
	thing      int
	otherThing string
}

actor MyActor {
	receive (x int) -> {
		// this is a comment
		// actor MyCommentActor {}
	}
}
`)

	// create the lexer
	lexer := parser.NewYaaiLexer(is)
	lexer.AddErrorListener(tl)
	fmt.Println("-------------------------------------------------------------------------------------------------------")
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("%s(%v) -> ", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
	fmt.Println("-------------------------------------------------------------------------------------------------------")
}
