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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 84, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 6, 5, 54, 10, 5, 13, 5, 14, 5, 55, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 8, 6, 8, 63, 10, 8, 13, 8, 14, 8, 64, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 6, 13, 79,
	10, 13, 13, 13, 14, 13, 80, 3, 13, 3, 13, 3, 80, 2, 14, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 2, 19, 2, 21, 2, 23, 2, 25, 10, 3, 2,
	5, 4, 2, 67, 92, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 4, 2,
	12, 12, 15, 15, 2, 85, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2,
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3,
	2, 2, 2, 2, 25, 3, 2, 2, 2, 3, 27, 3, 2, 2, 2, 5, 36, 3, 2, 2, 2, 7, 43,
	3, 2, 2, 2, 9, 51, 3, 2, 2, 2, 11, 57, 3, 2, 2, 2, 13, 59, 3, 2, 2, 2,
	15, 62, 3, 2, 2, 2, 17, 66, 3, 2, 2, 2, 19, 68, 3, 2, 2, 2, 21, 70, 3,
	2, 2, 2, 23, 72, 3, 2, 2, 2, 25, 78, 3, 2, 2, 2, 27, 28, 7, 114, 2, 2,
	28, 29, 7, 99, 2, 2, 29, 30, 7, 101, 2, 2, 30, 31, 7, 109, 2, 2, 31, 32,
	7, 99, 2, 2, 32, 33, 7, 105, 2, 2, 33, 34, 7, 103, 2, 2, 34, 35, 7, 34,
	2, 2, 35, 4, 3, 2, 2, 2, 36, 37, 7, 99, 2, 2, 37, 38, 7, 101, 2, 2, 38,
	39, 7, 118, 2, 2, 39, 40, 7, 113, 2, 2, 40, 41, 7, 116, 2, 2, 41, 42, 7,
	34, 2, 2, 42, 6, 3, 2, 2, 2, 43, 44, 7, 114, 2, 2, 44, 45, 7, 99, 2, 2,
	45, 46, 7, 101, 2, 2, 46, 47, 7, 109, 2, 2, 47, 48, 7, 99, 2, 2, 48, 49,
	7, 105, 2, 2, 49, 50, 7, 103, 2, 2, 50, 8, 3, 2, 2, 2, 51, 53, 9, 2, 2,
	2, 52, 54, 9, 3, 2, 2, 53, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 53,
	3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 10, 3, 2, 2, 2, 57, 58, 7, 125, 2,
	2, 58, 12, 3, 2, 2, 2, 59, 60, 7, 127, 2, 2, 60, 14, 3, 2, 2, 2, 61, 63,
	9, 4, 2, 2, 62, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2,
	64, 65, 3, 2, 2, 2, 65, 16, 3, 2, 2, 2, 66, 67, 7, 34, 2, 2, 67, 18, 3,
	2, 2, 2, 68, 69, 7, 12, 2, 2, 69, 20, 3, 2, 2, 2, 70, 71, 7, 15, 2, 2,
	71, 22, 3, 2, 2, 2, 72, 73, 7, 11, 2, 2, 73, 24, 3, 2, 2, 2, 74, 79, 5,
	17, 9, 2, 75, 79, 5, 19, 10, 2, 76, 79, 5, 21, 11, 2, 77, 79, 5, 23, 12,
	2, 78, 74, 3, 2, 2, 2, 78, 75, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 77,
	3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2,
	81, 82, 3, 2, 2, 2, 82, 83, 8, 13, 2, 2, 83, 26, 3, 2, 2, 2, 7, 2, 55,
	64, 78, 80, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'package '", "'actor '", "'package'", "", "'{'", "'}'",
}

var lexerSymbolicNames = []string{
	"", "", "", "KEYWORD_PACKAGE", "IDENTIFIER", "L_BRACKET", "R_BRACKET",
	"EOL", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "KEYWORD_PACKAGE", "IDENTIFIER", "L_BRACKET", "R_BRACKET",
	"EOL", "Space", "NewLine", "CariageReturn", "Tab", "WS",
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
	YaaiLexerT__1            = 2
	YaaiLexerKEYWORD_PACKAGE = 3
	YaaiLexerIDENTIFIER      = 4
	YaaiLexerL_BRACKET       = 5
	YaaiLexerR_BRACKET       = 6
	YaaiLexerEOL             = 7
	YaaiLexerWS              = 8
)
