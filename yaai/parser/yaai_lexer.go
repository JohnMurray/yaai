// Code generated from yaai/Yaai.g4 by ANTLR 4.9.3. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 45, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 6,
	6, 27, 10, 6, 13, 6, 14, 6, 28, 3, 7, 6, 7, 32, 10, 7, 13, 7, 14, 7, 33,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 2, 2, 9, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 3, 2, 6, 3, 2, 50, 59, 4, 2,
	67, 92, 99, 124, 4, 2, 12, 12, 15, 15, 5, 2, 11, 12, 15, 15, 34, 34, 2,
	46, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2,
	2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 3, 17, 3, 2,
	2, 2, 5, 19, 3, 2, 2, 2, 7, 21, 3, 2, 2, 2, 9, 23, 3, 2, 2, 2, 11, 26,
	3, 2, 2, 2, 13, 31, 3, 2, 2, 2, 15, 37, 3, 2, 2, 2, 17, 18, 9, 2, 2, 2,
	18, 4, 3, 2, 2, 2, 19, 20, 7, 34, 2, 2, 20, 6, 3, 2, 2, 2, 21, 22, 9, 3,
	2, 2, 22, 8, 3, 2, 2, 2, 23, 24, 7, 97, 2, 2, 24, 10, 3, 2, 2, 2, 25, 27,
	9, 4, 2, 2, 26, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2,
	28, 29, 3, 2, 2, 2, 29, 12, 3, 2, 2, 2, 30, 32, 9, 5, 2, 2, 31, 30, 3,
	2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34,
	35, 3, 2, 2, 2, 35, 36, 8, 7, 2, 2, 36, 14, 3, 2, 2, 2, 37, 38, 7, 114,
	2, 2, 38, 39, 7, 99, 2, 2, 39, 40, 7, 101, 2, 2, 40, 41, 7, 109, 2, 2,
	41, 42, 7, 99, 2, 2, 42, 43, 7, 105, 2, 2, 43, 44, 7, 103, 2, 2, 44, 16,
	3, 2, 2, 2, 5, 2, 28, 33, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "' '", "", "'_'", "", "", "'package'",
}

var lexerSymbolicNames = []string{
	"", "NUMBER", "SPACE", "LETTER", "UNDERSCORE", "EOL", "WHITESPACE", "KEYWORD_PACKAGE",
}

var lexerRuleNames = []string{
	"NUMBER", "SPACE", "LETTER", "UNDERSCORE", "EOL", "WHITESPACE", "KEYWORD_PACKAGE",
}

type YaaiLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewYaaiLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *YaaiLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewYaaiLexer(input antlr.CharStream) *YaaiLexer {
	l := new(YaaiLexer)
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
	l.GrammarFileName = "Yaai.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// YaaiLexer tokens.
const (
	YaaiLexerNUMBER          = 1
	YaaiLexerSPACE           = 2
	YaaiLexerLETTER          = 3
	YaaiLexerUNDERSCORE      = 4
	YaaiLexerEOL             = 5
	YaaiLexerWHITESPACE      = 6
	YaaiLexerKEYWORD_PACKAGE = 7
)
