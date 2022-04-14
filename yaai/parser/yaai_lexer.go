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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 50, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 6, 4, 34, 10, 4, 13, 4, 14, 4, 35,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 6, 8, 47, 10, 8,
	13, 8, 14, 8, 48, 2, 2, 9, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9,
	3, 2, 5, 4, 2, 67, 92, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124,
	4, 2, 12, 12, 15, 15, 2, 51, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7,
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2,
	15, 3, 2, 2, 2, 3, 17, 3, 2, 2, 2, 5, 23, 3, 2, 2, 2, 7, 31, 3, 2, 2, 2,
	9, 37, 3, 2, 2, 2, 11, 39, 3, 2, 2, 2, 13, 41, 3, 2, 2, 2, 15, 46, 3, 2,
	2, 2, 17, 18, 7, 99, 2, 2, 18, 19, 7, 101, 2, 2, 19, 20, 7, 118, 2, 2,
	20, 21, 7, 113, 2, 2, 21, 22, 7, 116, 2, 2, 22, 4, 3, 2, 2, 2, 23, 24,
	7, 114, 2, 2, 24, 25, 7, 99, 2, 2, 25, 26, 7, 101, 2, 2, 26, 27, 7, 109,
	2, 2, 27, 28, 7, 99, 2, 2, 28, 29, 7, 105, 2, 2, 29, 30, 7, 103, 2, 2,
	30, 6, 3, 2, 2, 2, 31, 33, 9, 2, 2, 2, 32, 34, 9, 3, 2, 2, 33, 32, 3, 2,
	2, 2, 34, 35, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 8,
	3, 2, 2, 2, 37, 38, 7, 125, 2, 2, 38, 10, 3, 2, 2, 2, 39, 40, 7, 127, 2,
	2, 40, 12, 3, 2, 2, 2, 41, 42, 7, 34, 2, 2, 42, 43, 3, 2, 2, 2, 43, 44,
	8, 7, 2, 2, 44, 14, 3, 2, 2, 2, 45, 47, 9, 4, 2, 2, 46, 45, 3, 2, 2, 2,
	47, 48, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 16, 3,
	2, 2, 2, 5, 2, 35, 48, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'actor'", "'package'", "", "'{'", "'}'", "' '",
}

var lexerSymbolicNames = []string{
	"", "", "KEYWORD_PACKAGE", "IDENTIFIER", "L_BRACKET", "R_BRACKET", "SPACE",
	"EOL",
}

var lexerRuleNames = []string{
	"T__0", "KEYWORD_PACKAGE", "IDENTIFIER", "L_BRACKET", "R_BRACKET", "SPACE",
	"EOL",
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
	YaaiLexerT__0            = 1
	YaaiLexerKEYWORD_PACKAGE = 2
	YaaiLexerIDENTIFIER      = 3
	YaaiLexerL_BRACKET       = 4
	YaaiLexerR_BRACKET       = 5
	YaaiLexerSPACE           = 6
	YaaiLexerEOL             = 7
)
