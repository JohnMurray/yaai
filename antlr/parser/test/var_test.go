package test

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/assert"
)

func TestVarDecl(t *testing.T) {
	testInputs := []string{
		"var boop string = \"hello there\"",
		"var boop int = 32",
		"var Bloop float64 = 1.2345",
	}
	for _, test := range testInputs {
		t.Run(test, func(t *testing.T) {
			p, tl := createParser(t, test)
			antlr.ParseTreeWalkerDefault.Walk(tl, p.VarDecl())
			assert.Empty(t, tl.errors)
		})
	}
}

func TestVarAssignment(t *testing.T) {

}

func TestVarInit(t *testing.T) {

}
