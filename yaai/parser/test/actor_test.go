package test

/// import (
/// 	"testing"
///
/// 	"github.com/antlr/antlr4/runtime/Go/antlr"
/// 	"github.com/stretchr/testify/assert"
/// )
///
/// func TestActorDecls(t *testing.T) {
/// 	testInputs := []string{
/// 		"actor test {}",
/// 		"actor  Test {}",              // additional spaces
/// 		"actor test{}",                // no spaces before braces
/// 		"actor test { }",              // spaces between braces
/// 		"actor test { \n\n\r\r\t\t }", // mixed spaces between
/// 		"actor test {\n\n}",           // more mixed spacing
/// 	}
///
/// 	for _, test := range testInputs {
/// 		t.Run(test, func(t *testing.T) {
/// 			p, tl := createParser(test)
/// 			antlr.ParseTreeWalkerDefault.Walk(tl, p.Actor())
/// 			assert.Empty(t, tl.errors)
/// 		})
/// 	}
/// }
///
/// func TestBadActorDecls(t *testing.T) {
/// 	testInputs := []string{
/// 		"actorTest {}", // no space
/// 		"actor {}",     // no actor name
/// 	}
///
/// 	for _, test := range testInputs {
/// 		t.Run(test, func(t *testing.T) {
/// 			p, tl := createParser(test)
/// 			antlr.ParseTreeWalkerDefault.Walk(tl, p.Actor())
/// 			assert.NotEmpty(t, tl.errors)
/// 		})
/// 	}
/// }
///
