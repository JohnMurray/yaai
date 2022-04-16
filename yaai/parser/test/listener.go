package test

import (
	"fmt"

	"github.com/JohnMurray/yaii/yaai/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Test implementationf or a parse-tree and error listener
type testListener struct {
	*parser.BaseYaaiListener
	*antlr.DefaultErrorListener

	packageName string
	errors      []error
}

func createParser(input string) (*parser.YaaiParser, *testListener) {
	tl := &testListener{}

	// setup the input
	is := antlr.NewInputStream(input)

	// create the lexer
	lexer := parser.NewYaaiLexer(is)
	lexer.AddErrorListener(tl)
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
	fmt.Println("!!!!!!!!!!!!!!!!! Report Ambiguity [error listener]")
	fmt.Println("  --> This is new.... /shrug")
}
func (tl *testListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	fmt.Println("!!!!!!!!!!!!!!!!! Report Attempting Full Context [error listener]")
	fmt.Println("  --> This is new.... /shrug")
}
func (tl *testListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	fmt.Println("!!!!!!!!!!!!!!!!! Report Context Sensitivity [error listener]")
	fmt.Println("  --> This is new.... /shrug")
}
