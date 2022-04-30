// Code generated from /home/john/Dropbox/projects/yaai//antlr/Yaai.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Yaai
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseYaaiListener is a complete listener for a parse tree produced by Yaai.
type BaseYaaiListener struct{}

var _ YaaiListener = &BaseYaaiListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseYaaiListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseYaaiListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseYaaiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseYaaiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseYaaiListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseYaaiListener) ExitFile(ctx *FileContext) {}

// EnterPackageClause is called when production packageClause is entered.
func (s *BaseYaaiListener) EnterPackageClause(ctx *PackageClauseContext) {}

// ExitPackageClause is called when production packageClause is exited.
func (s *BaseYaaiListener) ExitPackageClause(ctx *PackageClauseContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseYaaiListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseYaaiListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterVarAssignment is called when production varAssignment is entered.
func (s *BaseYaaiListener) EnterVarAssignment(ctx *VarAssignmentContext) {}

// ExitVarAssignment is called when production varAssignment is exited.
func (s *BaseYaaiListener) ExitVarAssignment(ctx *VarAssignmentContext) {}

// EnterVarInit is called when production varInit is entered.
func (s *BaseYaaiListener) EnterVarInit(ctx *VarInitContext) {}

// ExitVarInit is called when production varInit is exited.
func (s *BaseYaaiListener) ExitVarInit(ctx *VarInitContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseYaaiListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseYaaiListener) ExitType_(ctx *Type_Context) {}

// EnterBuiltinType is called when production builtinType is entered.
func (s *BaseYaaiListener) EnterBuiltinType(ctx *BuiltinTypeContext) {}

// ExitBuiltinType is called when production builtinType is exited.
func (s *BaseYaaiListener) ExitBuiltinType(ctx *BuiltinTypeContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseYaaiListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseYaaiListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseYaaiListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseYaaiListener) ExitEos(ctx *EosContext) {}
