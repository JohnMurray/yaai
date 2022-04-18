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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 23, 164,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 7, 14, 124, 10,
	14, 12, 14, 14, 14, 127, 11, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 21, 3, 21, 7, 21, 148, 10, 21, 12, 21, 14, 21, 151, 11, 21, 3,
	21, 3, 21, 3, 22, 6, 22, 156, 10, 22, 13, 22, 14, 22, 157, 3, 22, 5, 22,
	161, 10, 22, 3, 22, 3, 22, 2, 2, 23, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 3, 2, 7, 4, 2, 67, 92,
	99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 4, 2, 11, 11, 34, 34, 4,
	2, 12, 12, 15, 15, 3, 3, 61, 61, 2, 167, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2,
	29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2,
	2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2,
	2, 3, 45, 3, 2, 2, 2, 5, 51, 3, 2, 2, 2, 7, 59, 3, 2, 2, 2, 9, 66, 3, 2,
	2, 2, 11, 74, 3, 2, 2, 2, 13, 79, 3, 2, 2, 2, 15, 83, 3, 2, 2, 2, 17, 89,
	3, 2, 2, 2, 19, 95, 3, 2, 2, 2, 21, 100, 3, 2, 2, 2, 23, 107, 3, 2, 2,
	2, 25, 114, 3, 2, 2, 2, 27, 121, 3, 2, 2, 2, 29, 128, 3, 2, 2, 2, 31, 130,
	3, 2, 2, 2, 33, 132, 3, 2, 2, 2, 35, 134, 3, 2, 2, 2, 37, 136, 3, 2, 2,
	2, 39, 139, 3, 2, 2, 2, 41, 143, 3, 2, 2, 2, 43, 160, 3, 2, 2, 2, 45, 46,
	7, 99, 2, 2, 46, 47, 7, 101, 2, 2, 47, 48, 7, 118, 2, 2, 48, 49, 7, 113,
	2, 2, 49, 50, 7, 116, 2, 2, 50, 4, 3, 2, 2, 2, 51, 52, 7, 114, 2, 2, 52,
	53, 7, 99, 2, 2, 53, 54, 7, 101, 2, 2, 54, 55, 7, 109, 2, 2, 55, 56, 7,
	99, 2, 2, 56, 57, 7, 105, 2, 2, 57, 58, 7, 103, 2, 2, 58, 6, 3, 2, 2, 2,
	59, 60, 7, 117, 2, 2, 60, 61, 7, 118, 2, 2, 61, 62, 7, 116, 2, 2, 62, 63,
	7, 119, 2, 2, 63, 64, 7, 101, 2, 2, 64, 65, 7, 118, 2, 2, 65, 8, 3, 2,
	2, 2, 66, 67, 7, 116, 2, 2, 67, 68, 7, 103, 2, 2, 68, 69, 7, 101, 2, 2,
	69, 70, 7, 103, 2, 2, 70, 71, 7, 107, 2, 2, 71, 72, 7, 120, 2, 2, 72, 73,
	7, 103, 2, 2, 73, 10, 3, 2, 2, 2, 74, 75, 7, 118, 2, 2, 75, 76, 7, 123,
	2, 2, 76, 77, 7, 114, 2, 2, 77, 78, 7, 103, 2, 2, 78, 12, 3, 2, 2, 2, 79,
	80, 7, 107, 2, 2, 80, 81, 7, 112, 2, 2, 81, 82, 7, 118, 2, 2, 82, 14, 3,
	2, 2, 2, 83, 84, 7, 107, 2, 2, 84, 85, 7, 112, 2, 2, 85, 86, 7, 118, 2,
	2, 86, 87, 7, 53, 2, 2, 87, 88, 7, 52, 2, 2, 88, 16, 3, 2, 2, 2, 89, 90,
	7, 107, 2, 2, 90, 91, 7, 112, 2, 2, 91, 92, 7, 118, 2, 2, 92, 93, 7, 56,
	2, 2, 93, 94, 7, 54, 2, 2, 94, 18, 3, 2, 2, 2, 95, 96, 7, 119, 2, 2, 96,
	97, 7, 107, 2, 2, 97, 98, 7, 112, 2, 2, 98, 99, 7, 118, 2, 2, 99, 20, 3,
	2, 2, 2, 100, 101, 7, 119, 2, 2, 101, 102, 7, 107, 2, 2, 102, 103, 7, 112,
	2, 2, 103, 104, 7, 118, 2, 2, 104, 105, 7, 53, 2, 2, 105, 106, 7, 52, 2,
	2, 106, 22, 3, 2, 2, 2, 107, 108, 7, 119, 2, 2, 108, 109, 7, 107, 2, 2,
	109, 110, 7, 112, 2, 2, 110, 111, 7, 118, 2, 2, 111, 112, 7, 56, 2, 2,
	112, 113, 7, 54, 2, 2, 113, 24, 3, 2, 2, 2, 114, 115, 7, 117, 2, 2, 115,
	116, 7, 118, 2, 2, 116, 117, 7, 116, 2, 2, 117, 118, 7, 107, 2, 2, 118,
	119, 7, 112, 2, 2, 119, 120, 7, 105, 2, 2, 120, 26, 3, 2, 2, 2, 121, 125,
	9, 2, 2, 2, 122, 124, 9, 3, 2, 2, 123, 122, 3, 2, 2, 2, 124, 127, 3, 2,
	2, 2, 125, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 28, 3, 2, 2, 2,
	127, 125, 3, 2, 2, 2, 128, 129, 7, 125, 2, 2, 129, 30, 3, 2, 2, 2, 130,
	131, 7, 127, 2, 2, 131, 32, 3, 2, 2, 2, 132, 133, 7, 42, 2, 2, 133, 34,
	3, 2, 2, 2, 134, 135, 7, 43, 2, 2, 135, 36, 3, 2, 2, 2, 136, 137, 7, 47,
	2, 2, 137, 138, 7, 64, 2, 2, 138, 38, 3, 2, 2, 2, 139, 140, 9, 4, 2, 2,
	140, 141, 3, 2, 2, 2, 141, 142, 8, 20, 2, 2, 142, 40, 3, 2, 2, 2, 143,
	144, 7, 49, 2, 2, 144, 145, 7, 49, 2, 2, 145, 149, 3, 2, 2, 2, 146, 148,
	10, 5, 2, 2, 147, 146, 3, 2, 2, 2, 148, 151, 3, 2, 2, 2, 149, 147, 3, 2,
	2, 2, 149, 150, 3, 2, 2, 2, 150, 152, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2,
	152, 153, 8, 21, 3, 2, 153, 42, 3, 2, 2, 2, 154, 156, 9, 5, 2, 2, 155,
	154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158,
	3, 2, 2, 2, 158, 161, 3, 2, 2, 2, 159, 161, 9, 6, 2, 2, 160, 155, 3, 2,
	2, 2, 160, 159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 163, 8, 22, 4, 2,
	163, 44, 3, 2, 2, 2, 7, 2, 125, 149, 157, 160, 5, 8, 2, 2, 2, 3, 2, 4,
	2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'actor'", "'package'", "'struct'", "'receive'", "'type'", "'int'",
	"'int32'", "'int64'", "'uint'", "'uint32'", "'uint64'", "'string'", "",
	"'{'", "'}'", "'('", "')'", "'->'",
}

var lexerSymbolicNames = []string{
	"", "ACTOR", "PACKAGE", "STRUCT", "RECEIVE", "TYPE", "T_INT", "T_INT32",
	"T_INT64", "T_UINT", "T_UINT32", "T_UINT64", "T_STRING", "IDENTIFIER",
	"L_BRACKET", "R_BRACKET", "L_PAREN", "R_PAREN", "LAMBDA_ARROW", "NB_WS",
	"LINE_COMMENT", "EOS",
}

var lexerRuleNames = []string{
	"ACTOR", "PACKAGE", "STRUCT", "RECEIVE", "TYPE", "T_INT", "T_INT32", "T_INT64",
	"T_UINT", "T_UINT32", "T_UINT64", "T_STRING", "IDENTIFIER", "L_BRACKET",
	"R_BRACKET", "L_PAREN", "R_PAREN", "LAMBDA_ARROW", "NB_WS", "LINE_COMMENT",
	"EOS",
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
	YaaiLexerACTOR        = 1
	YaaiLexerPACKAGE      = 2
	YaaiLexerSTRUCT       = 3
	YaaiLexerRECEIVE      = 4
	YaaiLexerTYPE         = 5
	YaaiLexerT_INT        = 6
	YaaiLexerT_INT32      = 7
	YaaiLexerT_INT64      = 8
	YaaiLexerT_UINT       = 9
	YaaiLexerT_UINT32     = 10
	YaaiLexerT_UINT64     = 11
	YaaiLexerT_STRING     = 12
	YaaiLexerIDENTIFIER   = 13
	YaaiLexerL_BRACKET    = 14
	YaaiLexerR_BRACKET    = 15
	YaaiLexerL_PAREN      = 16
	YaaiLexerR_PAREN      = 17
	YaaiLexerLAMBDA_ARROW = 18
	YaaiLexerNB_WS        = 19
	YaaiLexerLINE_COMMENT = 20
	YaaiLexerEOS          = 21
)
