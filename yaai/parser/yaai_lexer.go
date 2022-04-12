// Code generated from yaai/yaai.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 54, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 8, 6, 8, 43, 10, 8, 13, 8, 14, 8, 44, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 2, 2, 12, 3, 3, 5, 4, 7, 5, 9,
	6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 3, 2, 5, 3, 2, 50, 59,
	5, 2, 11, 12, 15, 15, 34, 34, 3, 2, 99, 124, 2, 54, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2,
	2, 21, 3, 2, 2, 2, 3, 23, 3, 2, 2, 2, 5, 31, 3, 2, 2, 2, 7, 33, 3, 2, 2,
	2, 9, 35, 3, 2, 2, 2, 11, 37, 3, 2, 2, 2, 13, 39, 3, 2, 2, 2, 15, 42, 3,
	2, 2, 2, 17, 48, 3, 2, 2, 2, 19, 50, 3, 2, 2, 2, 21, 52, 3, 2, 2, 2, 23,
	24, 7, 114, 2, 2, 24, 25, 7, 99, 2, 2, 25, 26, 7, 101, 2, 2, 26, 27, 7,
	109, 2, 2, 27, 28, 7, 99, 2, 2, 28, 29, 7, 105, 2, 2, 29, 30, 7, 103, 2,
	2, 30, 4, 3, 2, 2, 2, 31, 32, 7, 44, 2, 2, 32, 6, 3, 2, 2, 2, 33, 34, 7,
	49, 2, 2, 34, 8, 3, 2, 2, 2, 35, 36, 7, 45, 2, 2, 36, 10, 3, 2, 2, 2, 37,
	38, 7, 47, 2, 2, 38, 12, 3, 2, 2, 2, 39, 40, 9, 2, 2, 2, 40, 14, 3, 2,
	2, 2, 41, 43, 9, 3, 2, 2, 42, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 42,
	3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 47, 8, 8, 2, 2,
	47, 16, 3, 2, 2, 2, 48, 49, 7, 34, 2, 2, 49, 18, 3, 2, 2, 2, 50, 51, 9,
	4, 2, 2, 51, 20, 3, 2, 2, 2, 52, 53, 7, 97, 2, 2, 53, 22, 3, 2, 2, 2, 4,
	2, 44, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'package'", "'*'", "'/'", "'+'", "'-'", "", "", "' '", "", "'_'",
}

var lexerSymbolicNames = []string{
	"", "", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE", "SPACE", "LETTER",
	"UNDERSCORE",
}

var lexerRuleNames = []string{
	"T__0", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE", "SPACE", "LETTER",
	"UNDERSCORE",
}

type yaaiLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewyaaiLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *yaaiLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewyaaiLexer(input antlr.CharStream) *yaaiLexer {
	l := new(yaaiLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "yaai.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// yaaiLexer tokens.
const (
	yaaiLexerT__0       = 1
	yaaiLexerMUL        = 2
	yaaiLexerDIV        = 3
	yaaiLexerADD        = 4
	yaaiLexerSUB        = 5
	yaaiLexerNUMBER     = 6
	yaaiLexerWHITESPACE = 7
	yaaiLexerSPACE      = 8
	yaaiLexerLETTER     = 9
	yaaiLexerUNDERSCORE = 10
)
