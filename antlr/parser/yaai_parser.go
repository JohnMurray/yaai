// Code generated from /home/john/Dropbox/projects/yaai//antlr/Yaai.g4 by ANTLR 4.9.3. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 18, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 3, 3, 3, 49, 49, 2, 14, 2, 8,
	3, 2, 2, 2, 4, 11, 3, 2, 2, 2, 6, 15, 3, 2, 2, 2, 8, 9, 5, 4, 3, 2, 9,
	10, 7, 2, 2, 3, 10, 3, 3, 2, 2, 2, 11, 12, 7, 8, 2, 2, 12, 13, 7, 22, 2,
	2, 13, 14, 5, 6, 4, 2, 14, 5, 3, 2, 2, 2, 15, 16, 9, 2, 2, 2, 16, 7, 3,
	2, 2, 2, 2,
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
	"file", "packageClause", "eos",
}

type Yaai struct {
	*antlr.BaseParser
}

// NewYaai produces a new parser instance for the optional input antlr.TokenStream.
//
// The *Yaai instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewYaai(input antlr.TokenStream) *Yaai {
	this := new(Yaai)
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

// Yaai tokens.
const (
	YaaiEOF                = antlr.TokenEOF
	YaaiACTOR              = 1
	YaaiFOR                = 2
	YaaiFUNC               = 3
	YaaiINIT               = 4
	YaaiINTERFACE          = 5
	YaaiPACKAGE            = 6
	YaaiRECEIVE            = 7
	YaaiSELF               = 8
	YaaiSTRUCT             = 9
	YaaiTYPE               = 10
	YaaiUNHANDLED          = 11
	YaaiVAR                = 12
	YaaiT_INT              = 13
	YaaiT_INT32            = 14
	YaaiT_INT64            = 15
	YaaiT_UINT             = 16
	YaaiT_UINT32           = 17
	YaaiT_UINT64           = 18
	YaaiT_STRING           = 19
	YaaiIDENTIFIER         = 20
	YaaiSTRING_LITERAL     = 21
	YaaiNUMERIC_LITERAL    = 22
	YaaiVAR_INITIALIZER    = 23
	YaaiASSIGNMENT         = 24
	YaaiINCR               = 25
	YaaiDECR               = 26
	YaaiPLUS_EQUAL         = 27
	YaaiMINUS_EQUAL        = 28
	YaaiEQUAL_EQUAL        = 29
	YaaiNOT_EQUAL          = 30
	YaaiLESS_THAN_EQUAL    = 31
	YaaiGREATER_THAN_EQUAL = 32
	YaaiLESS_THAN          = 33
	YaaiGREATER_THAN       = 34
	YaaiL_BRACKET          = 35
	YaaiR_BRACKET          = 36
	YaaiL_PAREN            = 37
	YaaiR_PAREN            = 38
	YaaiDOT                = 39
	YaaiPLUS               = 40
	YaaiMINUS              = 41
	YaaiSTAR               = 42
	YaaiF_SLASH            = 43
	YaaiCOMMA              = 44
	YaaiNB_WS              = 45
	YaaiLINE_COMMENT       = 46
	YaaiEOS                = 47
)

// Yaai rules.
const (
	YaaiRULE_file          = 0
	YaaiRULE_packageClause = 1
	YaaiRULE_eos           = 2
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YaaiRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YaaiRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) PackageClause() IPackageClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackageClauseContext)
}

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(YaaiEOF, 0)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *Yaai) File() (localctx IFileContext) {
	this := p
	_ = this

	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, YaaiRULE_file)

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
		p.SetState(6)
		p.PackageClause()
	}
	{
		p.SetState(7)
		p.Match(YaaiEOF)
	}

	return localctx
}

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
	p.RuleIndex = YaaiRULE_packageClause
	return p
}

func (*PackageClauseContext) IsPackageClauseContext() {}

func NewPackageClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageClauseContext {
	var p = new(PackageClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YaaiRULE_packageClause

	return p
}

func (s *PackageClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageClauseContext) GetPackageName() antlr.Token { return s.packageName }

func (s *PackageClauseContext) SetPackageName(v antlr.Token) { s.packageName = v }

func (s *PackageClauseContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(YaaiPACKAGE, 0)
}

func (s *PackageClauseContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *PackageClauseContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(YaaiIDENTIFIER, 0)
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

func (p *Yaai) PackageClause() (localctx IPackageClauseContext) {
	this := p
	_ = this

	localctx = NewPackageClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, YaaiRULE_packageClause)

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
		p.SetState(9)
		p.Match(YaaiPACKAGE)
	}
	{
		p.SetState(10)

		var _m = p.Match(YaaiIDENTIFIER)

		localctx.(*PackageClauseContext).packageName = _m
	}
	{
		p.SetState(11)
		p.Eos()
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YaaiRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YaaiRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) EOS() antlr.TerminalNode {
	return s.GetToken(YaaiEOS, 0)
}

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(YaaiEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.ExitEos(s)
	}
}

func (p *Yaai) Eos() (localctx IEosContext) {
	this := p
	_ = this

	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, YaaiRULE_eos)
	var _la int

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
		p.SetState(13)
		_la = p.GetTokenStream().LA(1)

		if !(_la == YaaiEOF || _la == YaaiEOS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
