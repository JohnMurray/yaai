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

// EnterStart is called when production start is entered.
func (s *BaseyaaiListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseyaaiListener) ExitStart(ctx *StartContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseyaaiListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseyaaiListener) ExitNumber(ctx *NumberContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseyaaiListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseyaaiListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseyaaiListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseyaaiListener) ExitAddSub(ctx *AddSubContext) {}
