package test

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/assert"
)

func TestActorDecls(t *testing.T) {
	testInputs := []string{
		"actor test {}",
	}

	for _, test := range testInputs {
		t.Run(test, func(t *testing.T) {
			p, tl := createParser(test)
			antlr.ParseTreeWalkerDefault.Walk(tl, p.Actor())
			assert.Empty(t, tl.errors)
		})
	}
}
