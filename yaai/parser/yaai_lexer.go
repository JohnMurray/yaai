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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 25, 177,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 7, 15, 132, 10, 15, 12, 15, 14,
	15, 135, 11, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 21, 6, 21, 149, 10, 21, 13, 21, 14, 21, 150, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23, 161, 10, 23,
	12, 23, 14, 23, 164, 11, 23, 3, 23, 3, 23, 3, 24, 6, 24, 169, 10, 24, 13,
	24, 14, 24, 170, 3, 24, 5, 24, 174, 10, 24, 3, 24, 3, 24, 2, 2, 25, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 47, 25, 3, 2, 8, 4, 2, 67, 92, 99, 124, 6, 2, 50, 59, 67,
	92, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12,
	15, 15, 3, 3, 61, 61, 2, 181, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7,
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2,
	15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2,
	2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2,
	2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2,
	2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3,
	2, 2, 2, 2, 47, 3, 2, 2, 2, 3, 49, 3, 2, 2, 2, 5, 55, 3, 2, 2, 2, 7, 63,
	3, 2, 2, 2, 9, 70, 3, 2, 2, 2, 11, 78, 3, 2, 2, 2, 13, 83, 3, 2, 2, 2,
	15, 87, 3, 2, 2, 2, 17, 91, 3, 2, 2, 2, 19, 97, 3, 2, 2, 2, 21, 103, 3,
	2, 2, 2, 23, 108, 3, 2, 2, 2, 25, 115, 3, 2, 2, 2, 27, 122, 3, 2, 2, 2,
	29, 129, 3, 2, 2, 2, 31, 136, 3, 2, 2, 2, 33, 139, 3, 2, 2, 2, 35, 141,
	3, 2, 2, 2, 37, 143, 3, 2, 2, 2, 39, 145, 3, 2, 2, 2, 41, 148, 3, 2, 2,
	2, 43, 152, 3, 2, 2, 2, 45, 156, 3, 2, 2, 2, 47, 173, 3, 2, 2, 2, 49, 50,
	7, 99, 2, 2, 50, 51, 7, 101, 2, 2, 51, 52, 7, 118, 2, 2, 52, 53, 7, 113,
	2, 2, 53, 54, 7, 116, 2, 2, 54, 4, 3, 2, 2, 2, 55, 56, 7, 114, 2, 2, 56,
	57, 7, 99, 2, 2, 57, 58, 7, 101, 2, 2, 58, 59, 7, 109, 2, 2, 59, 60, 7,
	99, 2, 2, 60, 61, 7, 105, 2, 2, 61, 62, 7, 103, 2, 2, 62, 6, 3, 2, 2, 2,
	63, 64, 7, 117, 2, 2, 64, 65, 7, 118, 2, 2, 65, 66, 7, 116, 2, 2, 66, 67,
	7, 119, 2, 2, 67, 68, 7, 101, 2, 2, 68, 69, 7, 118, 2, 2, 69, 8, 3, 2,
	2, 2, 70, 71, 7, 116, 2, 2, 71, 72, 7, 103, 2, 2, 72, 73, 7, 101, 2, 2,
	73, 74, 7, 103, 2, 2, 74, 75, 7, 107, 2, 2, 75, 76, 7, 120, 2, 2, 76, 77,
	7, 103, 2, 2, 77, 10, 3, 2, 2, 2, 78, 79, 7, 118, 2, 2, 79, 80, 7, 123,
	2, 2, 80, 81, 7, 114, 2, 2, 81, 82, 7, 103, 2, 2, 82, 12, 3, 2, 2, 2, 83,
	84, 7, 120, 2, 2, 84, 85, 7, 99, 2, 2, 85, 86, 7, 116, 2, 2, 86, 14, 3,
	2, 2, 2, 87, 88, 7, 107, 2, 2, 88, 89, 7, 112, 2, 2, 89, 90, 7, 118, 2,
	2, 90, 16, 3, 2, 2, 2, 91, 92, 7, 107, 2, 2, 92, 93, 7, 112, 2, 2, 93,
	94, 7, 118, 2, 2, 94, 95, 7, 53, 2, 2, 95, 96, 7, 52, 2, 2, 96, 18, 3,
	2, 2, 2, 97, 98, 7, 107, 2, 2, 98, 99, 7, 112, 2, 2, 99, 100, 7, 118, 2,
	2, 100, 101, 7, 56, 2, 2, 101, 102, 7, 54, 2, 2, 102, 20, 3, 2, 2, 2, 103,
	104, 7, 119, 2, 2, 104, 105, 7, 107, 2, 2, 105, 106, 7, 112, 2, 2, 106,
	107, 7, 118, 2, 2, 107, 22, 3, 2, 2, 2, 108, 109, 7, 119, 2, 2, 109, 110,
	7, 107, 2, 2, 110, 111, 7, 112, 2, 2, 111, 112, 7, 118, 2, 2, 112, 113,
	7, 53, 2, 2, 113, 114, 7, 52, 2, 2, 114, 24, 3, 2, 2, 2, 115, 116, 7, 119,
	2, 2, 116, 117, 7, 107, 2, 2, 117, 118, 7, 112, 2, 2, 118, 119, 7, 118,
	2, 2, 119, 120, 7, 56, 2, 2, 120, 121, 7, 54, 2, 2, 121, 26, 3, 2, 2, 2,
	122, 123, 7, 117, 2, 2, 123, 124, 7, 118, 2, 2, 124, 125, 7, 116, 2, 2,
	125, 126, 7, 107, 2, 2, 126, 127, 7, 112, 2, 2, 127, 128, 7, 105, 2, 2,
	128, 28, 3, 2, 2, 2, 129, 133, 9, 2, 2, 2, 130, 132, 9, 3, 2, 2, 131, 130,
	3, 2, 2, 2, 132, 135, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 134, 3, 2,
	2, 2, 134, 30, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 136, 137, 7, 60, 2, 2,
	137, 138, 7, 63, 2, 2, 138, 32, 3, 2, 2, 2, 139, 140, 7, 125, 2, 2, 140,
	34, 3, 2, 2, 2, 141, 142, 7, 127, 2, 2, 142, 36, 3, 2, 2, 2, 143, 144,
	7, 42, 2, 2, 144, 38, 3, 2, 2, 2, 145, 146, 7, 43, 2, 2, 146, 40, 3, 2,
	2, 2, 147, 149, 9, 4, 2, 2, 148, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2,
	150, 148, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 42, 3, 2, 2, 2, 152, 153,
	9, 5, 2, 2, 153, 154, 3, 2, 2, 2, 154, 155, 8, 22, 2, 2, 155, 44, 3, 2,
	2, 2, 156, 157, 7, 49, 2, 2, 157, 158, 7, 49, 2, 2, 158, 162, 3, 2, 2,
	2, 159, 161, 10, 6, 2, 2, 160, 159, 3, 2, 2, 2, 161, 164, 3, 2, 2, 2, 162,
	160, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 165, 3, 2, 2, 2, 164, 162,
	3, 2, 2, 2, 165, 166, 8, 23, 3, 2, 166, 46, 3, 2, 2, 2, 167, 169, 9, 6,
	2, 2, 168, 167, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2,
	170, 171, 3, 2, 2, 2, 171, 174, 3, 2, 2, 2, 172, 174, 9, 7, 2, 2, 173,
	168, 3, 2, 2, 2, 173, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 176,
	8, 24, 4, 2, 176, 48, 3, 2, 2, 2, 8, 2, 133, 150, 162, 170, 173, 5, 8,
	2, 2, 2, 3, 2, 4, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'actor'", "'package'", "'struct'", "'receive'", "'type'", "'var'",
	"'int'", "'int32'", "'int64'", "'uint'", "'uint32'", "'uint64'", "'string'",
	"", "':='", "'{'", "'}'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "ACTOR", "PACKAGE", "STRUCT", "RECEIVE", "TYPE", "VAR", "T_INT", "T_INT32",
	"T_INT64", "T_UINT", "T_UINT32", "T_UINT64", "T_STRING", "IDENTIFIER",
	"VAR_INITIALIZER", "L_BRACKET", "R_BRACKET", "L_PAREN", "R_PAREN", "NUMERIC_LITERAL",
	"NB_WS", "LINE_COMMENT", "EOS",
}

var lexerRuleNames = []string{
	"ACTOR", "PACKAGE", "STRUCT", "RECEIVE", "TYPE", "VAR", "T_INT", "T_INT32",
	"T_INT64", "T_UINT", "T_UINT32", "T_UINT64", "T_STRING", "IDENTIFIER",
	"VAR_INITIALIZER", "L_BRACKET", "R_BRACKET", "L_PAREN", "R_PAREN", "NUMERIC_LITERAL",
	"NB_WS", "LINE_COMMENT", "EOS",
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
	YaaiLexerACTOR           = 1
	YaaiLexerPACKAGE         = 2
	YaaiLexerSTRUCT          = 3
	YaaiLexerRECEIVE         = 4
	YaaiLexerTYPE            = 5
	YaaiLexerVAR             = 6
	YaaiLexerT_INT           = 7
	YaaiLexerT_INT32         = 8
	YaaiLexerT_INT64         = 9
	YaaiLexerT_UINT          = 10
	YaaiLexerT_UINT32        = 11
	YaaiLexerT_UINT64        = 12
	YaaiLexerT_STRING        = 13
	YaaiLexerIDENTIFIER      = 14
	YaaiLexerVAR_INITIALIZER = 15
	YaaiLexerL_BRACKET       = 16
	YaaiLexerR_BRACKET       = 17
	YaaiLexerL_PAREN         = 18
	YaaiLexerR_PAREN         = 19
	YaaiLexerNUMERIC_LITERAL = 20
	YaaiLexerNB_WS           = 21
	YaaiLexerLINE_COMMENT    = 22
	YaaiLexerEOS             = 23
)
