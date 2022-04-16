// Code generated from yaai/YaaiLexer.g4 by ANTLR 4.9.3. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 18, 127,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3,
	13, 6, 13, 106, 10, 13, 13, 13, 14, 13, 107, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 6, 17, 119, 10, 17, 13, 17, 14, 17,
	120, 3, 17, 5, 17, 124, 10, 17, 3, 17, 3, 17, 2, 2, 18, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 3, 2, 7, 4, 2, 67, 92, 99, 124, 6, 2, 50, 59,
	67, 92, 97, 97, 99, 124, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12, 15, 15, 3,
	3, 61, 61, 2, 129, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2,
	2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2,
	2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2,
	2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3,
	2, 2, 2, 2, 33, 3, 2, 2, 2, 3, 35, 3, 2, 2, 2, 5, 43, 3, 2, 2, 2, 7, 49,
	3, 2, 2, 2, 9, 54, 3, 2, 2, 2, 11, 61, 3, 2, 2, 2, 13, 65, 3, 2, 2, 2,
	15, 71, 3, 2, 2, 2, 17, 77, 3, 2, 2, 2, 19, 82, 3, 2, 2, 2, 21, 89, 3,
	2, 2, 2, 23, 96, 3, 2, 2, 2, 25, 103, 3, 2, 2, 2, 27, 109, 3, 2, 2, 2,
	29, 111, 3, 2, 2, 2, 31, 113, 3, 2, 2, 2, 33, 123, 3, 2, 2, 2, 35, 36,
	7, 114, 2, 2, 36, 37, 7, 99, 2, 2, 37, 38, 7, 101, 2, 2, 38, 39, 7, 109,
	2, 2, 39, 40, 7, 99, 2, 2, 40, 41, 7, 105, 2, 2, 41, 42, 7, 103, 2, 2,
	42, 4, 3, 2, 2, 2, 43, 44, 7, 99, 2, 2, 44, 45, 7, 101, 2, 2, 45, 46, 7,
	118, 2, 2, 46, 47, 7, 113, 2, 2, 47, 48, 7, 116, 2, 2, 48, 6, 3, 2, 2,
	2, 49, 50, 7, 118, 2, 2, 50, 51, 7, 123, 2, 2, 51, 52, 7, 114, 2, 2, 52,
	53, 7, 103, 2, 2, 53, 8, 3, 2, 2, 2, 54, 55, 7, 117, 2, 2, 55, 56, 7, 118,
	2, 2, 56, 57, 7, 116, 2, 2, 57, 58, 7, 119, 2, 2, 58, 59, 7, 101, 2, 2,
	59, 60, 7, 118, 2, 2, 60, 10, 3, 2, 2, 2, 61, 62, 7, 107, 2, 2, 62, 63,
	7, 112, 2, 2, 63, 64, 7, 118, 2, 2, 64, 12, 3, 2, 2, 2, 65, 66, 7, 107,
	2, 2, 66, 67, 7, 112, 2, 2, 67, 68, 7, 118, 2, 2, 68, 69, 7, 53, 2, 2,
	69, 70, 7, 52, 2, 2, 70, 14, 3, 2, 2, 2, 71, 72, 7, 107, 2, 2, 72, 73,
	7, 112, 2, 2, 73, 74, 7, 118, 2, 2, 74, 75, 7, 56, 2, 2, 75, 76, 7, 54,
	2, 2, 76, 16, 3, 2, 2, 2, 77, 78, 7, 119, 2, 2, 78, 79, 7, 107, 2, 2, 79,
	80, 7, 112, 2, 2, 80, 81, 7, 118, 2, 2, 81, 18, 3, 2, 2, 2, 82, 83, 7,
	119, 2, 2, 83, 84, 7, 107, 2, 2, 84, 85, 7, 112, 2, 2, 85, 86, 7, 118,
	2, 2, 86, 87, 7, 53, 2, 2, 87, 88, 7, 52, 2, 2, 88, 20, 3, 2, 2, 2, 89,
	90, 7, 119, 2, 2, 90, 91, 7, 107, 2, 2, 91, 92, 7, 112, 2, 2, 92, 93, 7,
	118, 2, 2, 93, 94, 7, 56, 2, 2, 94, 95, 7, 54, 2, 2, 95, 22, 3, 2, 2, 2,
	96, 97, 7, 117, 2, 2, 97, 98, 7, 118, 2, 2, 98, 99, 7, 116, 2, 2, 99, 100,
	7, 107, 2, 2, 100, 101, 7, 112, 2, 2, 101, 102, 7, 105, 2, 2, 102, 24,
	3, 2, 2, 2, 103, 105, 9, 2, 2, 2, 104, 106, 9, 3, 2, 2, 105, 104, 3, 2,
	2, 2, 106, 107, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2,
	108, 26, 3, 2, 2, 2, 109, 110, 7, 125, 2, 2, 110, 28, 3, 2, 2, 2, 111,
	112, 7, 127, 2, 2, 112, 30, 3, 2, 2, 2, 113, 114, 9, 4, 2, 2, 114, 115,
	3, 2, 2, 2, 115, 116, 8, 16, 2, 2, 116, 32, 3, 2, 2, 2, 117, 119, 9, 5,
	2, 2, 118, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2,
	120, 121, 3, 2, 2, 2, 121, 124, 3, 2, 2, 2, 122, 124, 9, 6, 2, 2, 123,
	118, 3, 2, 2, 2, 123, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 126,
	8, 17, 3, 2, 126, 34, 3, 2, 2, 2, 6, 2, 107, 120, 123, 4, 8, 2, 2, 4, 2,
	2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'package'", "'actor'", "'type'", "'struct'", "'int'", "'int32'", "'int64'",
	"'uint'", "'uint32'", "'uint64'", "'string'", "", "'{'", "'}'",
}

var lexerSymbolicNames = []string{
	"", "PACKAGE", "ACTOR", "TYPE", "STRUCT", "T_INT", "T_INT32", "T_INT64",
	"T_UINT", "T_UINT32", "T_UINT64", "T_STRING", "IDENTIFIER", "L_BRACKET",
	"R_BRACKET", "NB_WS", "EOS",
}

var lexerRuleNames = []string{
	"PACKAGE", "ACTOR", "TYPE", "STRUCT", "T_INT", "T_INT32", "T_INT64", "T_UINT",
	"T_UINT32", "T_UINT64", "T_STRING", "IDENTIFIER", "L_BRACKET", "R_BRACKET",
	"NB_WS", "EOS",
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
	l.GrammarFileName = "YaaiLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// YaaiLexer tokens.
const (
	YaaiLexerPACKAGE    = 1
	YaaiLexerACTOR      = 2
	YaaiLexerTYPE       = 3
	YaaiLexerSTRUCT     = 4
	YaaiLexerT_INT      = 5
	YaaiLexerT_INT32    = 6
	YaaiLexerT_INT64    = 7
	YaaiLexerT_UINT     = 8
	YaaiLexerT_UINT32   = 9
	YaaiLexerT_UINT64   = 10
	YaaiLexerT_STRING   = 11
	YaaiLexerIDENTIFIER = 12
	YaaiLexerL_BRACKET  = 13
	YaaiLexerR_BRACKET  = 14
	YaaiLexerNB_WS      = 15
	YaaiLexerEOS        = 16
)
