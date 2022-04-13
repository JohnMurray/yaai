package test

import (
	"testing"

	"github.com/JohnMurray/yaii/yaai/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func createParser(input string) *parser.YaaiParser {
	// setup the input
	is := antlr.NewInputStream(input)

	// create the lexer
	lexer := parser.NewYaaiLexer(is)
	// for {
	// 	t := lexer.NextToken()
	// 	if t.GetTokenType() == antlr.TokenEOF {
	// 		fmt.Println("EOF token")
	// 		break
	// 	}
	// 	fmt.Printf("%s (%v)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	// }
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// create the parser
	return parser.NewYaaiParser(stream)
}

type testListener struct {
	*parser.BaseYaaiListener
}

func TestParsePackageDecl(t *testing.T) {
	p := createParser("package Stonks\n\nl")
	antlr.ParseTreeWalkerDefault.Walk(&testListener{}, p.Package_decl())
}
