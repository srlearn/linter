// Code generated from ILPLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ILPLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ILPLangListener is a complete listener for a parse tree produced by ILPLangParser.
type ILPLangListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterClassification_task is called when entering the classification_task production.
	EnterClassification_task(c *Classification_taskContext)

	// EnterRegression_task is called when entering the regression_task production.
	EnterRegression_task(c *Regression_taskContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterRegression_term is called when entering the regression_term production.
	EnterRegression_term(c *Regression_termContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitClassification_task is called when exiting the classification_task production.
	ExitClassification_task(c *Classification_taskContext)

	// ExitRegression_task is called when exiting the regression_task production.
	ExitRegression_task(c *Regression_taskContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitRegression_term is called when exiting the regression_term production.
	ExitRegression_term(c *Regression_termContext)
}
