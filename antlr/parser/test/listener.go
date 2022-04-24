package test

//lint:file-ignore U1000 Methods used by antlr's visitor pattern

import (
	"fmt"
	"testing"

	"github.com/JohnMurray/yaii/antlr/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/assert"
)

// Test implementationf or a parse-tree and error listener
type testListener struct {
	*parser.BaseYaaiListener
	*antlr.DefaultErrorListener

	t           *testing.T
	packageName string
	errors      []error
}

func createParser(t *testing.T, input string) (*parser.Yaai, *testListener) {
	tl := &testListener{t: t}

	// setup the input
	is := antlr.NewInputStream(input)

	// create the lexer
	lexer := parser.NewYaaiLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(tl)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// create the parser
	p := parser.NewYaai(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(tl)
	return p, tl
}

func (tl *testListener) ExitPackageClause(c *parser.PackageClauseContext) {
	tl.packageName = c.GetPackageName().GetText()
}

// --------------------------------------------------------------------------------
// antlr.ErrorListener (parser & lexer)

// Populate the listener with error encountered. Try and form
// something that's somewhat readable to troubleshoot.
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

func (tl *testListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	assert.Fail(tl.t, "Report Ambiguity [error listener] --> This is new... /shrug")
}
func (tl *testListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	assert.Fail(tl.t, "Report Attempting Full Context [error listener] --> This is new... /shrug")
}
func (tl *testListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	assert.Fail(tl.t, "Report Context Sensitivity [error listener] --> This is new... /shrug")
}
