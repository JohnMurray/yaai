// Code generated from yaai/yaai.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // yaai

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseyaaiListener is a complete listener for a parse tree produced by yaaiParser.
type BaseyaaiListener struct{}

var _ yaaiListener = &BaseyaaiListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseyaaiListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseyaaiListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseyaaiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseyaaiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPackage_decl is called when production package_decl is entered.
func (s *BaseyaaiListener) EnterPackage_decl(ctx *Package_declContext) {}

// ExitPackage_decl is called when production package_decl is exited.
func (s *BaseyaaiListener) ExitPackage_decl(ctx *Package_declContext) {}

// EnterPackage_name is called when production package_name is entered.
func (s *BaseyaaiListener) EnterPackage_name(ctx *Package_nameContext) {}

// ExitPackage_name is called when production package_name is exited.
func (s *BaseyaaiListener) ExitPackage_name(ctx *Package_nameContext) {}

// EnterPackage_name_start is called when production package_name_start is entered.
func (s *BaseyaaiListener) EnterPackage_name_start(ctx *Package_name_startContext) {}

// ExitPackage_name_start is called when production package_name_start is exited.
func (s *BaseyaaiListener) ExitPackage_name_start(ctx *Package_name_startContext) {}

// EnterPackage_name_end is called when production package_name_end is entered.
func (s *BaseyaaiListener) EnterPackage_name_end(ctx *Package_name_endContext) {}

// ExitPackage_name_end is called when production package_name_end is exited.
func (s *BaseyaaiListener) ExitPackage_name_end(ctx *Package_name_endContext) {}
