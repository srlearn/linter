// Code generated from ILPLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ILPLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ILPLangListener is a complete listener for a parse tree produced by ILPLangParser.
type ILPLangListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)
}
