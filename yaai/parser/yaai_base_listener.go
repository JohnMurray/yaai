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

// EnterPackageClause is called when production packageClause is entered.
func (s *BaseYaaiListener) EnterPackageClause(ctx *PackageClauseContext) {}

// ExitPackageClause is called when production packageClause is exited.
func (s *BaseYaaiListener) ExitPackageClause(ctx *PackageClauseContext) {}
