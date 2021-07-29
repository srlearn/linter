// Code generated from ILPLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ILPLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseILPLangListener is a complete listener for a parse tree produced by ILPLangParser.
type BaseILPLangListener struct{}

var _ ILPLangListener = &BaseILPLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseILPLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseILPLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseILPLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseILPLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseILPLangListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseILPLangListener) ExitStart(ctx *StartContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseILPLangListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseILPLangListener) ExitTerm(ctx *TermContext) {}

// EnterLastterm is called when production lastterm is entered.
func (s *BaseILPLangListener) EnterLastterm(ctx *LasttermContext) {}

// ExitLastterm is called when production lastterm is exited.
func (s *BaseILPLangListener) ExitLastterm(ctx *LasttermContext) {}
