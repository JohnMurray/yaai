package test

import (
	"fmt"
	"testing"

	"github.com/JohnMurray/yaii/yaai/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/assert"
)

func createParser(input string) (*parser.YaaiParser, *testListener) {
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
	p := parser.NewYaaiParser(stream)
	tl := &testListener{}
	p.AddErrorListener(tl)
	return p, tl
}

type testListener struct {
	*parser.BaseYaaiListener
	*antlr.DefaultErrorListener

	packageName string
	errors      []error
}

func (tl *testListener) SyntaxError(recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException) {
	tl.errors = append(tl.errors, fmt.Errorf(
		"[%d:%d] (%v) %s",
		line, column, offendingSymbol, msg,
	))
}

func (tl *testListener) ExitPackage_name(c *parser.Package_nameContext) {
	tl.packageName = c.GetText()
}

func TestParsePackageDecl(t *testing.T) {
	p, tl := createParser("package Stonks\n")
	antlr.ParseTreeWalkerDefault.Walk(tl, p.Package_decl())

	assert.Equal(t, "Stonks", tl.packageName)
	assert.Empty(t, tl.errors)
}
