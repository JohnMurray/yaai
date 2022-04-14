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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 23, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2,
	2, 2, 19, 2, 8, 3, 2, 2, 2, 4, 11, 3, 2, 2, 2, 6, 16, 3, 2, 2, 2, 8, 9,
	5, 4, 3, 2, 9, 10, 7, 2, 2, 3, 10, 3, 3, 2, 2, 2, 11, 12, 7, 4, 2, 2, 12,
	13, 7, 8, 2, 2, 13, 14, 7, 5, 2, 2, 14, 15, 7, 9, 2, 2, 15, 5, 3, 2, 2,
	2, 16, 17, 7, 3, 2, 2, 17, 18, 7, 8, 2, 2, 18, 19, 7, 5, 2, 2, 19, 20,
	7, 6, 2, 2, 20, 21, 7, 7, 2, 2, 21, 7, 3, 2, 2, 2, 2,
}
var literalNames = []string{
	"", "'actor'", "'package'", "", "'{'", "'}'", "' '",
}
var symbolicNames = []string{
	"", "", "KEYWORD_PACKAGE", "IDENTIFIER", "L_BRACKET", "R_BRACKET", "SPACE",
	"EOL", "WHITESPACE",
}

var ruleNames = []string{
	"unit", "package_decl", "actor",
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
	YaaiParserEOF             = antlr.TokenEOF
	YaaiParserT__0            = 1
	YaaiParserKEYWORD_PACKAGE = 2
	YaaiParserIDENTIFIER      = 3
	YaaiParserL_BRACKET       = 4
	YaaiParserR_BRACKET       = 5
	YaaiParserSPACE           = 6
	YaaiParserEOL             = 7
	YaaiParserWHITESPACE      = 8
)

// YaaiParser rules.
const (
	YaaiParserRULE_unit         = 0
	YaaiParserRULE_package_decl = 1
	YaaiParserRULE_actor        = 2
)

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YaaiParserRULE_unit
	return p
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YaaiParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) Package_decl() IPackage_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackage_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackage_declContext)
}

func (s *UnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(YaaiParserEOF, 0)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.EnterUnit(s)
	}
}

func (s *UnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.ExitUnit(s)
	}
}

func (p *YaaiParser) Unit() (localctx IUnitContext) {
	this := p
	_ = this

	localctx = NewUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, YaaiParserRULE_unit)

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
		p.Package_decl()
	}
	{
		p.SetState(7)
		p.Match(YaaiParserEOF)
	}

	return localctx
}

// IPackage_declContext is an interface to support dynamic dispatch.
type IPackage_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackage_declContext differentiates from other interfaces.
	IsPackage_declContext()
}

type Package_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackage_declContext() *Package_declContext {
	var p = new(Package_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YaaiParserRULE_package_decl
	return p
}

func (*Package_declContext) IsPackage_declContext() {}

func NewPackage_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Package_declContext {
	var p = new(Package_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YaaiParserRULE_package_decl

	return p
}

func (s *Package_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Package_declContext) KEYWORD_PACKAGE() antlr.TerminalNode {
	return s.GetToken(YaaiParserKEYWORD_PACKAGE, 0)
}

func (s *Package_declContext) SPACE() antlr.TerminalNode {
	return s.GetToken(YaaiParserSPACE, 0)
}

func (s *Package_declContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(YaaiParserIDENTIFIER, 0)
}

func (s *Package_declContext) EOL() antlr.TerminalNode {
	return s.GetToken(YaaiParserEOL, 0)
}

func (s *Package_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Package_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Package_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.EnterPackage_decl(s)
	}
}

func (s *Package_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.ExitPackage_decl(s)
	}
}

func (p *YaaiParser) Package_decl() (localctx IPackage_declContext) {
	this := p
	_ = this

	localctx = NewPackage_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, YaaiParserRULE_package_decl)

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
		p.Match(YaaiParserKEYWORD_PACKAGE)
	}
	{
		p.SetState(10)
		p.Match(YaaiParserSPACE)
	}
	{
		p.SetState(11)
		p.Match(YaaiParserIDENTIFIER)
	}
	{
		p.SetState(12)
		p.Match(YaaiParserEOL)
	}

	return localctx
}

// IActorContext is an interface to support dynamic dispatch.
type IActorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKeyword returns the keyword token.
	GetKeyword() antlr.Token

	// SetKeyword sets the keyword token.
	SetKeyword(antlr.Token)

	// IsActorContext differentiates from other interfaces.
	IsActorContext()
}

type ActorContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	keyword antlr.Token
}

func NewEmptyActorContext() *ActorContext {
	var p = new(ActorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YaaiParserRULE_actor
	return p
}

func (*ActorContext) IsActorContext() {}

func NewActorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActorContext {
	var p = new(ActorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YaaiParserRULE_actor

	return p
}

func (s *ActorContext) GetParser() antlr.Parser { return s.parser }

func (s *ActorContext) GetKeyword() antlr.Token { return s.keyword }

func (s *ActorContext) SetKeyword(v antlr.Token) { s.keyword = v }

func (s *ActorContext) SPACE() antlr.TerminalNode {
	return s.GetToken(YaaiParserSPACE, 0)
}

func (s *ActorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(YaaiParserIDENTIFIER, 0)
}

func (s *ActorContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(YaaiParserL_BRACKET, 0)
}

func (s *ActorContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(YaaiParserR_BRACKET, 0)
}

func (s *ActorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.EnterActor(s)
	}
}

func (s *ActorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.ExitActor(s)
	}
}

func (p *YaaiParser) Actor() (localctx IActorContext) {
	this := p
	_ = this

	localctx = NewActorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, YaaiParserRULE_actor)

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
		p.SetState(14)

		var _m = p.Match(YaaiParserT__0)

		localctx.(*ActorContext).keyword = _m
	}
	{
		p.SetState(15)
		p.Match(YaaiParserSPACE)
	}
	{
		p.SetState(16)
		p.Match(YaaiParserIDENTIFIER)
	}
	{
		p.SetState(17)
		p.Match(YaaiParserL_BRACKET)
	}
	{
		p.SetState(18)
		p.Match(YaaiParserR_BRACKET)
	}

	return localctx
}
