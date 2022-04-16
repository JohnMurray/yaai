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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 8, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 2, 2, 3, 2, 2, 2, 2, 6, 2, 4, 3, 2, 2,
	2, 4, 5, 7, 3, 2, 2, 5, 6, 7, 14, 2, 2, 6, 3, 3, 2, 2, 2, 2,
}
var literalNames = []string{
	"", "'package'", "'actor'", "'type'", "'struct'", "'int'", "'int32'", "'int64'",
	"'uint'", "'uint32'", "'uint64'", "'string'", "", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "PACKAGE", "ACTOR", "TYPE", "STRUCT", "T_INT", "T_INT32", "T_INT64",
	"T_UINT", "T_UINT32", "T_UINT64", "T_STRING", "IDENTIFIER", "L_BRACKET",
	"R_BRACKET", "NB_WS", "EOS",
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
	YaaiParserEOF        = antlr.TokenEOF
	YaaiParserPACKAGE    = 1
	YaaiParserACTOR      = 2
	YaaiParserTYPE       = 3
	YaaiParserSTRUCT     = 4
	YaaiParserT_INT      = 5
	YaaiParserT_INT32    = 6
	YaaiParserT_INT64    = 7
	YaaiParserT_UINT     = 8
	YaaiParserT_UINT32   = 9
	YaaiParserT_UINT64   = 10
	YaaiParserT_STRING   = 11
	YaaiParserIDENTIFIER = 12
	YaaiParserL_BRACKET  = 13
	YaaiParserR_BRACKET  = 14
	YaaiParserNB_WS      = 15
	YaaiParserEOS        = 16
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
