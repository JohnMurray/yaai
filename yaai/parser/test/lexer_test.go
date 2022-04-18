package test

import (
	"testing"

	"github.com/JohnMurray/yaii/yaai/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/assert"
)

func TestLexerPackage(t *testing.T) {
	lexer := parser.NewYaaiLexer(antlr.NewInputStream(`package myPackage`))
	assert.NoError(t, lexerSnapshot(t, "package-simple", lexer))
}
