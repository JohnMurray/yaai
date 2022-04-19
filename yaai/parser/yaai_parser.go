// Code generated from yaai/Yaai.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Yaai
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 8, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 2, 2, 3, 2, 2, 2, 2, 6, 2, 4, 3, 2, 2,
	2, 4, 5, 7, 8, 2, 2, 5, 6, 7, 22, 2, 2, 6, 3, 3, 2, 2, 2, 2,
}
var literalNames = []string{
	"", "'actor'", "'for'", "'func'", "'init'", "'interface'", "'package'",
	"'receive'", "'self'", "'struct'", "'type'", "'unhandled'", "'var'", "'int'",
	"'int32'", "'int64'", "'uint'", "'uint32'", "'uint64'", "'string'", "",
	"", "", "':='", "'='", "'++'", "'--'", "'+='", "'-='", "'=='", "'!='",
	"'<='", "'>='", "'<'", "'>'", "'{'", "'}'", "'('", "')'", "'.'", "'+'",
	"'-'", "'*'", "'/'", "','",
}
var symbolicNames = []string{
	"", "ACTOR", "FOR", "FUNC", "INIT", "INTERFACE", "PACKAGE", "RECEIVE",
	"SELF", "STRUCT", "TYPE", "UNHANDLED", "VAR", "T_INT", "T_INT32", "T_INT64",
	"T_UINT", "T_UINT32", "T_UINT64", "T_STRING", "IDENTIFIER", "STRING_LITERAL",
	"NUMERIC_LITERAL", "VAR_INITIALIZER", "ASSIGNMENT", "INCR", "DECR", "PLUS_EQUAL",
	"MINUS_EQUAL", "EQUAL_EQUAL", "NOT_EQUAL", "LESS_THAN_EQUAL", "GREATER_THAN_EQUAL",
	"LESS_THAN", "GREATER_THAN", "L_BRACKET", "R_BRACKET", "L_PAREN", "R_PAREN",
	"DOT", "PLUS", "MINUS", "STAR", "F_SLASH", "COMMA", "NB_WS", "LINE_COMMENT",
	"EOS",
}

var ruleNames = []string{
	"packageClause",
}

type YaaiParser struct {
	*antlr.BaseParser
}

// NewYaaiParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *YaaiParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewYaaiParser(input antlr.TokenStream) *YaaiParser {
	this := new(YaaiParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Yaai.g4"

	return this
}

// YaaiParser tokens.
const (
	YaaiParserEOF                = antlr.TokenEOF
	YaaiParserACTOR              = 1
	YaaiParserFOR                = 2
	YaaiParserFUNC               = 3
	YaaiParserINIT               = 4
	YaaiParserINTERFACE          = 5
	YaaiParserPACKAGE            = 6
	YaaiParserRECEIVE            = 7
	YaaiParserSELF               = 8
	YaaiParserSTRUCT             = 9
	YaaiParserTYPE               = 10
	YaaiParserUNHANDLED          = 11
	YaaiParserVAR                = 12
	YaaiParserT_INT              = 13
	YaaiParserT_INT32            = 14
	YaaiParserT_INT64            = 15
	YaaiParserT_UINT             = 16
	YaaiParserT_UINT32           = 17
	YaaiParserT_UINT64           = 18
	YaaiParserT_STRING           = 19
	YaaiParserIDENTIFIER         = 20
	YaaiParserSTRING_LITERAL     = 21
	YaaiParserNUMERIC_LITERAL    = 22
	YaaiParserVAR_INITIALIZER    = 23
	YaaiParserASSIGNMENT         = 24
	YaaiParserINCR               = 25
	YaaiParserDECR               = 26
	YaaiParserPLUS_EQUAL         = 27
	YaaiParserMINUS_EQUAL        = 28
	YaaiParserEQUAL_EQUAL        = 29
	YaaiParserNOT_EQUAL          = 30
	YaaiParserLESS_THAN_EQUAL    = 31
	YaaiParserGREATER_THAN_EQUAL = 32
	YaaiParserLESS_THAN          = 33
	YaaiParserGREATER_THAN       = 34
	YaaiParserL_BRACKET          = 35
	YaaiParserR_BRACKET          = 36
	YaaiParserL_PAREN            = 37
	YaaiParserR_PAREN            = 38
	YaaiParserDOT                = 39
	YaaiParserPLUS               = 40
	YaaiParserMINUS              = 41
	YaaiParserSTAR               = 42
	YaaiParserF_SLASH            = 43
	YaaiParserCOMMA              = 44
	YaaiParserNB_WS              = 45
	YaaiParserLINE_COMMENT       = 46
	YaaiParserEOS                = 47
)

// YaaiParserRULE_packageClause is the YaaiParser rule.
const YaaiParserRULE_packageClause = 0

// IPackageClauseContext is an interface to support dynamic dispatch.
type IPackageClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPackageName returns the packageName token.
	GetPackageName() antlr.Token

	// SetPackageName sets the packageName token.
	SetPackageName(antlr.Token)

	// IsPackageClauseContext differentiates from other interfaces.
	IsPackageClauseContext()
}

type PackageClauseContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	packageName antlr.Token
}

func NewEmptyPackageClauseContext() *PackageClauseContext {
	var p = new(PackageClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YaaiParserRULE_packageClause
	return p
}

func (*PackageClauseContext) IsPackageClauseContext() {}

func NewPackageClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageClauseContext {
	var p = new(PackageClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YaaiParserRULE_packageClause

	return p
}

func (s *PackageClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageClauseContext) GetPackageName() antlr.Token { return s.packageName }

func (s *PackageClauseContext) SetPackageName(v antlr.Token) { s.packageName = v }

func (s *PackageClauseContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(YaaiParserPACKAGE, 0)
}

func (s *PackageClauseContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(YaaiParserIDENTIFIER, 0)
}

func (s *PackageClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.EnterPackageClause(s)
	}
}

func (s *PackageClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.ExitPackageClause(s)
	}
}

func (p *YaaiParser) PackageClause() (localctx IPackageClauseContext) {
	this := p
	_ = this

	localctx = NewPackageClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, YaaiParserRULE_packageClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(2)
		p.Match(YaaiParserPACKAGE)
	}
	{
		p.SetState(3)

		var _m = p.Match(YaaiParserIDENTIFIER)

		localctx.(*PackageClauseContext).packageName = _m
	}

	return localctx
}
