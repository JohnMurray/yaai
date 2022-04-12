// Code generated from yaai/yaai.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // yaai

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 27, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	3, 3, 3, 6, 3, 17, 10, 3, 13, 3, 14, 3, 18, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 2, 2, 6, 2, 4, 6, 8, 2, 3, 4, 2, 8, 8, 11, 12, 2, 23, 2,
	10, 3, 2, 2, 2, 4, 14, 3, 2, 2, 2, 6, 22, 3, 2, 2, 2, 8, 24, 3, 2, 2, 2,
	10, 11, 7, 3, 2, 2, 11, 12, 7, 10, 2, 2, 12, 13, 5, 4, 3, 2, 13, 3, 3,
	2, 2, 2, 14, 16, 5, 6, 4, 2, 15, 17, 5, 8, 5, 2, 16, 15, 3, 2, 2, 2, 17,
	18, 3, 2, 2, 2, 18, 16, 3, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19, 20, 3, 2, 2,
	2, 20, 21, 7, 13, 2, 2, 21, 5, 3, 2, 2, 2, 22, 23, 7, 11, 2, 2, 23, 7,
	3, 2, 2, 2, 24, 25, 9, 2, 2, 2, 25, 9, 3, 2, 2, 2, 3, 18,
}
var literalNames = []string{
	"", "'package'", "'*'", "'/'", "'+'", "'-'", "", "", "' '", "", "'_'",
}
var symbolicNames = []string{
	"", "", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE", "SPACE", "LETTER",
	"UNDERSCORE", "EOL",
}

var ruleNames = []string{
	"package_decl", "package_name", "package_name_start", "package_name_end",
}

type yaaiParser struct {
	*antlr.BaseParser
}

// NewyaaiParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *yaaiParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewyaaiParser(input antlr.TokenStream) *yaaiParser {
	this := new(yaaiParser)
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
	this.GrammarFileName = "yaai.g4"

	return this
}

// yaaiParser tokens.
const (
	yaaiParserEOF        = antlr.TokenEOF
	yaaiParserT__0       = 1
	yaaiParserMUL        = 2
	yaaiParserDIV        = 3
	yaaiParserADD        = 4
	yaaiParserSUB        = 5
	yaaiParserNUMBER     = 6
	yaaiParserWHITESPACE = 7
	yaaiParserSPACE      = 8
	yaaiParserLETTER     = 9
	yaaiParserUNDERSCORE = 10
	yaaiParserEOL        = 11
)

// yaaiParser rules.
const (
	yaaiParserRULE_package_decl       = 0
	yaaiParserRULE_package_name       = 1
	yaaiParserRULE_package_name_start = 2
	yaaiParserRULE_package_name_end   = 3
)

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
	p.RuleIndex = yaaiParserRULE_package_decl
	return p
}

func (*Package_declContext) IsPackage_declContext() {}

func NewPackage_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Package_declContext {
	var p = new(Package_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = yaaiParserRULE_package_decl

	return p
}

func (s *Package_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Package_declContext) SPACE() antlr.TerminalNode {
	return s.GetToken(yaaiParserSPACE, 0)
}

func (s *Package_declContext) Package_name() IPackage_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackage_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackage_nameContext)
}

func (s *Package_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Package_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Package_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(yaaiListener); ok {
		listenerT.EnterPackage_decl(s)
	}
}

func (s *Package_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(yaaiListener); ok {
		listenerT.ExitPackage_decl(s)
	}
}

func (p *yaaiParser) Package_decl() (localctx IPackage_declContext) {
	this := p
	_ = this

	localctx = NewPackage_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, yaaiParserRULE_package_decl)

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
		p.SetState(8)
		p.Match(yaaiParserT__0)
	}
	{
		p.SetState(9)
		p.Match(yaaiParserSPACE)
	}
	{
		p.SetState(10)
		p.Package_name()
	}

	return localctx
}

// IPackage_nameContext is an interface to support dynamic dispatch.
type IPackage_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackage_nameContext differentiates from other interfaces.
	IsPackage_nameContext()
}

type Package_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackage_nameContext() *Package_nameContext {
	var p = new(Package_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = yaaiParserRULE_package_name
	return p
}

func (*Package_nameContext) IsPackage_nameContext() {}

func NewPackage_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Package_nameContext {
	var p = new(Package_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = yaaiParserRULE_package_name

	return p
}

func (s *Package_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Package_nameContext) Package_name_start() IPackage_name_startContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackage_name_startContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackage_name_startContext)
}

func (s *Package_nameContext) EOL() antlr.TerminalNode {
	return s.GetToken(yaaiParserEOL, 0)
}

func (s *Package_nameContext) AllPackage_name_end() []IPackage_name_endContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPackage_name_endContext)(nil)).Elem())
	var tst = make([]IPackage_name_endContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPackage_name_endContext)
		}
	}

	return tst
}

func (s *Package_nameContext) Package_name_end(i int) IPackage_name_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackage_name_endContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPackage_name_endContext)
}

func (s *Package_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Package_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Package_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(yaaiListener); ok {
		listenerT.EnterPackage_name(s)
	}
}

func (s *Package_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(yaaiListener); ok {
		listenerT.ExitPackage_name(s)
	}
}

func (p *yaaiParser) Package_name() (localctx IPackage_nameContext) {
	this := p
	_ = this

	localctx = NewPackage_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, yaaiParserRULE_package_name)
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
		p.SetState(12)
		p.Package_name_start()
	}
	p.SetState(14)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<yaaiParserNUMBER)|(1<<yaaiParserLETTER)|(1<<yaaiParserUNDERSCORE))) != 0) {
		{
			p.SetState(13)
			p.Package_name_end()
		}

		p.SetState(16)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(18)
		p.Match(yaaiParserEOL)
	}

	return localctx
}

// IPackage_name_startContext is an interface to support dynamic dispatch.
type IPackage_name_startContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackage_name_startContext differentiates from other interfaces.
	IsPackage_name_startContext()
}

type Package_name_startContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackage_name_startContext() *Package_name_startContext {
	var p = new(Package_name_startContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = yaaiParserRULE_package_name_start
	return p
}

func (*Package_name_startContext) IsPackage_name_startContext() {}

func NewPackage_name_startContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Package_name_startContext {
	var p = new(Package_name_startContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = yaaiParserRULE_package_name_start

	return p
}

func (s *Package_name_startContext) GetParser() antlr.Parser { return s.parser }

func (s *Package_name_startContext) LETTER() antlr.TerminalNode {
	return s.GetToken(yaaiParserLETTER, 0)
}

func (s *Package_name_startContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Package_name_startContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Package_name_startContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(yaaiListener); ok {
		listenerT.EnterPackage_name_start(s)
	}
}

func (s *Package_name_startContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(yaaiListener); ok {
		listenerT.ExitPackage_name_start(s)
	}
}

func (p *yaaiParser) Package_name_start() (localctx IPackage_name_startContext) {
	this := p
	_ = this

	localctx = NewPackage_name_startContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, yaaiParserRULE_package_name_start)

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
		p.SetState(20)
		p.Match(yaaiParserLETTER)
	}

	return localctx
}

// IPackage_name_endContext is an interface to support dynamic dispatch.
type IPackage_name_endContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackage_name_endContext differentiates from other interfaces.
	IsPackage_name_endContext()
}

type Package_name_endContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackage_name_endContext() *Package_name_endContext {
	var p = new(Package_name_endContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = yaaiParserRULE_package_name_end
	return p
}

func (*Package_name_endContext) IsPackage_name_endContext() {}

func NewPackage_name_endContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Package_name_endContext {
	var p = new(Package_name_endContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = yaaiParserRULE_package_name_end

	return p
}

func (s *Package_name_endContext) GetParser() antlr.Parser { return s.parser }

func (s *Package_name_endContext) LETTER() antlr.TerminalNode {
	return s.GetToken(yaaiParserLETTER, 0)
}

func (s *Package_name_endContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(yaaiParserNUMBER, 0)
}

func (s *Package_name_endContext) UNDERSCORE() antlr.TerminalNode {
	return s.GetToken(yaaiParserUNDERSCORE, 0)
}

func (s *Package_name_endContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Package_name_endContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Package_name_endContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(yaaiListener); ok {
		listenerT.EnterPackage_name_end(s)
	}
}

func (s *Package_name_endContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(yaaiListener); ok {
		listenerT.ExitPackage_name_end(s)
	}
}

func (p *yaaiParser) Package_name_end() (localctx IPackage_name_endContext) {
	this := p
	_ = this

	localctx = NewPackage_name_endContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, yaaiParserRULE_package_name_end)
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
		p.SetState(22)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<yaaiParserNUMBER)|(1<<yaaiParserLETTER)|(1<<yaaiParserUNDERSCORE))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
