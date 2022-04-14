// Code generated from yaai/Yaai.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Yaai

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseYaaiListener is a complete listener for a parse tree produced by YaaiParser.
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

// EnterUnit is called when production unit is entered.
func (s *BaseYaaiListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseYaaiListener) ExitUnit(ctx *UnitContext) {}

// EnterPackage_decl is called when production package_decl is entered.
func (s *BaseYaaiListener) EnterPackage_decl(ctx *Package_declContext) {}

// ExitPackage_decl is called when production package_decl is exited.
func (s *BaseYaaiListener) ExitPackage_decl(ctx *Package_declContext) {}

// EnterActor is called when production actor is entered.
func (s *BaseYaaiListener) EnterActor(ctx *ActorContext) {}

// ExitActor is called when production actor is exited.
func (s *BaseYaaiListener) ExitActor(ctx *ActorContext) {}
