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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 52, 52, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 33, 10, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 5, 3,
	2, 15, 23, 3, 2, 25, 27, 3, 3, 52, 52, 2, 43, 2, 20, 3, 2, 2, 2, 4, 23,
	3, 2, 2, 2, 6, 27, 3, 2, 2, 2, 8, 34, 3, 2, 2, 2, 10, 38, 3, 2, 2, 2, 12,
	43, 3, 2, 2, 2, 14, 45, 3, 2, 2, 2, 16, 47, 3, 2, 2, 2, 18, 49, 3, 2, 2,
	2, 20, 21, 5, 4, 3, 2, 21, 22, 7, 2, 2, 3, 22, 3, 3, 2, 2, 2, 23, 24, 7,
	8, 2, 2, 24, 25, 7, 24, 2, 2, 25, 26, 5, 18, 10, 2, 26, 5, 3, 2, 2, 2,
	27, 28, 7, 14, 2, 2, 28, 29, 7, 24, 2, 2, 29, 32, 5, 12, 7, 2, 30, 33,
	5, 18, 10, 2, 31, 33, 5, 8, 5, 2, 32, 30, 3, 2, 2, 2, 32, 31, 3, 2, 2,
	2, 33, 7, 3, 2, 2, 2, 34, 35, 7, 29, 2, 2, 35, 36, 5, 16, 9, 2, 36, 37,
	5, 18, 10, 2, 37, 9, 3, 2, 2, 2, 38, 39, 7, 24, 2, 2, 39, 40, 7, 28, 2,
	2, 40, 41, 5, 16, 9, 2, 41, 42, 5, 18, 10, 2, 42, 11, 3, 2, 2, 2, 43, 44,
	5, 14, 8, 2, 44, 13, 3, 2, 2, 2, 45, 46, 9, 2, 2, 2, 46, 15, 3, 2, 2, 2,
	47, 48, 9, 3, 2, 2, 48, 17, 3, 2, 2, 2, 49, 50, 9, 4, 2, 2, 50, 19, 3,
	2, 2, 2, 3, 32,
}
var literalNames = []string{
	"", "'actor'", "'for'", "'func'", "'init'", "'interface'", "'package'",
	"'receive'", "'self'", "'struct'", "'type'", "'unhandled'", "'var'", "'int'",
	"'int32'", "'int64'", "'uint'", "'uint32'", "'uint64'", "'float32'", "'float64'",
	"'string'", "", "", "", "", "':='", "'='", "'++'", "'--'", "'+='", "'-='",
	"'=='", "'!='", "'<='", "'>='", "'<'", "'>'", "'{'", "'}'", "'('", "')'",
	"'.'", "'+'", "'-'", "'*'", "'/'", "','",
}
var symbolicNames = []string{
	"", "ACTOR", "FOR", "FUNC", "INIT", "INTERFACE", "PACKAGE", "RECEIVE",
	"SELF", "STRUCT", "TYPE", "UNHANDLED", "VAR", "T_INT", "T_INT32", "T_INT64",
	"T_UINT", "T_UINT32", "T_UINT64", "T_FLOAT32", "T_FLOAT64", "T_STRING",
	"IDENTIFIER", "STRING_LITERAL", "NUMERIC_LITERAL", "FLOAT_LITERAL", "VAR_INITIALIZER",
	"ASSIGNMENT", "INCR", "DECR", "PLUS_EQUAL", "MINUS_EQUAL", "EQUAL_EQUAL",
	"NOT_EQUAL", "LESS_THAN_EQUAL", "GREATER_THAN_EQUAL", "LESS_THAN", "GREATER_THAN",
	"L_BRACKET", "R_BRACKET", "L_PAREN", "R_PAREN", "DOT", "PLUS", "MINUS",
	"STAR", "F_SLASH", "COMMA", "NB_WS", "LINE_COMMENT", "EOS",
}

var ruleNames = []string{
	"file", "packageClause", "varDecl", "varAssignment", "varInit", "type_",
	"builtinType", "expression", "eos",
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
	YaaiT_FLOAT32          = 19
	YaaiT_FLOAT64          = 20
	YaaiT_STRING           = 21
	YaaiIDENTIFIER         = 22
	YaaiSTRING_LITERAL     = 23
	YaaiNUMERIC_LITERAL    = 24
	YaaiFLOAT_LITERAL      = 25
	YaaiVAR_INITIALIZER    = 26
	YaaiASSIGNMENT         = 27
	YaaiINCR               = 28
	YaaiDECR               = 29
	YaaiPLUS_EQUAL         = 30
	YaaiMINUS_EQUAL        = 31
	YaaiEQUAL_EQUAL        = 32
	YaaiNOT_EQUAL          = 33
	YaaiLESS_THAN_EQUAL    = 34
	YaaiGREATER_THAN_EQUAL = 35
	YaaiLESS_THAN          = 36
	YaaiGREATER_THAN       = 37
	YaaiL_BRACKET          = 38
	YaaiR_BRACKET          = 39
	YaaiL_PAREN            = 40
	YaaiR_PAREN            = 41
	YaaiDOT                = 42
	YaaiPLUS               = 43
	YaaiMINUS              = 44
	YaaiSTAR               = 45
	YaaiF_SLASH            = 46
	YaaiCOMMA              = 47
	YaaiNB_WS              = 48
	YaaiLINE_COMMENT       = 49
	YaaiEOS                = 50
)

// Yaai rules.
const (
	YaaiRULE_file          = 0
	YaaiRULE_packageClause = 1
	YaaiRULE_varDecl       = 2
	YaaiRULE_varAssignment = 3
	YaaiRULE_varInit       = 4
	YaaiRULE_type_         = 5
	YaaiRULE_builtinType   = 6
	YaaiRULE_expression    = 7
	YaaiRULE_eos           = 8
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
		p.SetState(18)
		p.PackageClause()
	}
	{
		p.SetState(19)
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
		p.SetState(21)
		p.Match(YaaiPACKAGE)
	}
	{
		p.SetState(22)

		var _m = p.Match(YaaiIDENTIFIER)

		localctx.(*PackageClauseContext).packageName = _m
	}
	{
		p.SetState(23)
		p.Eos()
	}

	return localctx
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YaaiRULE_varDecl
	return p
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YaaiRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) VAR() antlr.TerminalNode {
	return s.GetToken(YaaiVAR, 0)
}

func (s *VarDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(YaaiIDENTIFIER, 0)
}

func (s *VarDeclContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VarDeclContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *VarDeclContext) VarAssignment() IVarAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarAssignmentContext)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (p *Yaai) VarDecl() (localctx IVarDeclContext) {
	this := p
	_ = this

	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, YaaiRULE_varDecl)

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
		p.SetState(25)
		p.Match(YaaiVAR)
	}
	{
		p.SetState(26)
		p.Match(YaaiIDENTIFIER)
	}
	{
		p.SetState(27)
		p.Type_()
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case YaaiEOF, YaaiEOS:
		{
			p.SetState(28)
			p.Eos()
		}

	case YaaiASSIGNMENT:
		{
			p.SetState(29)
			p.VarAssignment()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVarAssignmentContext is an interface to support dynamic dispatch.
type IVarAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarAssignmentContext differentiates from other interfaces.
	IsVarAssignmentContext()
}

type VarAssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarAssignmentContext() *VarAssignmentContext {
	var p = new(VarAssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YaaiRULE_varAssignment
	return p
}

func (*VarAssignmentContext) IsVarAssignmentContext() {}

func NewVarAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarAssignmentContext {
	var p = new(VarAssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YaaiRULE_varAssignment

	return p
}

func (s *VarAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *VarAssignmentContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(YaaiASSIGNMENT, 0)
}

func (s *VarAssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VarAssignmentContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *VarAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.EnterVarAssignment(s)
	}
}

func (s *VarAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.ExitVarAssignment(s)
	}
}

func (p *Yaai) VarAssignment() (localctx IVarAssignmentContext) {
	this := p
	_ = this

	localctx = NewVarAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, YaaiRULE_varAssignment)

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
		p.SetState(32)
		p.Match(YaaiASSIGNMENT)
	}
	{
		p.SetState(33)
		p.Expression()
	}
	{
		p.SetState(34)
		p.Eos()
	}

	return localctx
}

// IVarInitContext is an interface to support dynamic dispatch.
type IVarInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarInitContext differentiates from other interfaces.
	IsVarInitContext()
}

type VarInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarInitContext() *VarInitContext {
	var p = new(VarInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YaaiRULE_varInit
	return p
}

func (*VarInitContext) IsVarInitContext() {}

func NewVarInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarInitContext {
	var p = new(VarInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YaaiRULE_varInit

	return p
}

func (s *VarInitContext) GetParser() antlr.Parser { return s.parser }

func (s *VarInitContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(YaaiIDENTIFIER, 0)
}

func (s *VarInitContext) VAR_INITIALIZER() antlr.TerminalNode {
	return s.GetToken(YaaiVAR_INITIALIZER, 0)
}

func (s *VarInitContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VarInitContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *VarInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.EnterVarInit(s)
	}
}

func (s *VarInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.ExitVarInit(s)
	}
}

func (p *Yaai) VarInit() (localctx IVarInitContext) {
	this := p
	_ = this

	localctx = NewVarInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, YaaiRULE_varInit)

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
		p.SetState(36)
		p.Match(YaaiIDENTIFIER)
	}
	{
		p.SetState(37)
		p.Match(YaaiVAR_INITIALIZER)
	}
	{
		p.SetState(38)
		p.Expression()
	}
	{
		p.SetState(39)
		p.Eos()
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YaaiRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YaaiRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) BuiltinType() IBuiltinTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltinTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltinTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *Yaai) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, YaaiRULE_type_)

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
		p.SetState(41)
		p.BuiltinType()
	}

	return localctx
}

// IBuiltinTypeContext is an interface to support dynamic dispatch.
type IBuiltinTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuiltinTypeContext differentiates from other interfaces.
	IsBuiltinTypeContext()
}

type BuiltinTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltinTypeContext() *BuiltinTypeContext {
	var p = new(BuiltinTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YaaiRULE_builtinType
	return p
}

func (*BuiltinTypeContext) IsBuiltinTypeContext() {}

func NewBuiltinTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltinTypeContext {
	var p = new(BuiltinTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YaaiRULE_builtinType

	return p
}

func (s *BuiltinTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BuiltinTypeContext) T_INT() antlr.TerminalNode {
	return s.GetToken(YaaiT_INT, 0)
}

func (s *BuiltinTypeContext) T_INT32() antlr.TerminalNode {
	return s.GetToken(YaaiT_INT32, 0)
}

func (s *BuiltinTypeContext) T_INT64() antlr.TerminalNode {
	return s.GetToken(YaaiT_INT64, 0)
}

func (s *BuiltinTypeContext) T_UINT() antlr.TerminalNode {
	return s.GetToken(YaaiT_UINT, 0)
}

func (s *BuiltinTypeContext) T_UINT32() antlr.TerminalNode {
	return s.GetToken(YaaiT_UINT32, 0)
}

func (s *BuiltinTypeContext) T_UINT64() antlr.TerminalNode {
	return s.GetToken(YaaiT_UINT64, 0)
}

func (s *BuiltinTypeContext) T_FLOAT32() antlr.TerminalNode {
	return s.GetToken(YaaiT_FLOAT32, 0)
}

func (s *BuiltinTypeContext) T_FLOAT64() antlr.TerminalNode {
	return s.GetToken(YaaiT_FLOAT64, 0)
}

func (s *BuiltinTypeContext) T_STRING() antlr.TerminalNode {
	return s.GetToken(YaaiT_STRING, 0)
}

func (s *BuiltinTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuiltinTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.EnterBuiltinType(s)
	}
}

func (s *BuiltinTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.ExitBuiltinType(s)
	}
}

func (p *Yaai) BuiltinType() (localctx IBuiltinTypeContext) {
	this := p
	_ = this

	localctx = NewBuiltinTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, YaaiRULE_builtinType)
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
		p.SetState(43)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<YaaiT_INT)|(1<<YaaiT_INT32)|(1<<YaaiT_INT64)|(1<<YaaiT_UINT)|(1<<YaaiT_UINT32)|(1<<YaaiT_UINT64)|(1<<YaaiT_FLOAT32)|(1<<YaaiT_FLOAT64)|(1<<YaaiT_STRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YaaiRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YaaiRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(YaaiSTRING_LITERAL, 0)
}

func (s *ExpressionContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(YaaiNUMERIC_LITERAL, 0)
}

func (s *ExpressionContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(YaaiFLOAT_LITERAL, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YaaiListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Yaai) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, YaaiRULE_expression)
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
		p.SetState(45)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<YaaiSTRING_LITERAL)|(1<<YaaiNUMERIC_LITERAL)|(1<<YaaiFLOAT_LITERAL))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.EnterRule(localctx, 16, YaaiRULE_eos)
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
		p.SetState(47)
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
