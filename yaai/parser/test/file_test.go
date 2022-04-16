package test

/// import (
/// 	"testing"
///
/// 	"github.com/antlr/antlr4/runtime/Go/antlr"
/// 	"github.com/stretchr/testify/assert"
/// )
///
/// // Just a smoke tests that tests the various combinations of components that
/// // compose a "file". Rigurous testing of each individual component should be
/// // covered by other tests.
/// func TestValidFile(t *testing.T) {
/// 	p, tl := createParser(`package TestPackage
///
/// actor TheActor{
///
/// }`)
/// 	antlr.ParseTreeWalkerDefault.Walk(tl, p.File())
/// 	assert.Empty(t, tl.errors)
/// }
