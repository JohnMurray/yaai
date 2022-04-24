// Code generated from yaai/Yaai.g4 by ANTLR 4.9.3. DO NOT EDIT.

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

// EnterEos is called when production eos is entered.
func (s *BaseYaaiListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseYaaiListener) ExitEos(ctx *EosContext) {}
