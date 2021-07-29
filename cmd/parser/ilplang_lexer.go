// Code generated from ILPLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 8, 47, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	5, 9, 39, 10, 9, 3, 10, 3, 10, 7, 10, 43, 10, 10, 12, 10, 14, 10, 46, 11,
	10, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 2, 15, 2, 17, 2, 19, 8,
	3, 2, 5, 3, 2, 12, 12, 3, 2, 50, 59, 3, 2, 99, 124, 2, 46, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 19, 3, 2, 2, 2, 3, 21, 3, 2, 2, 2, 5, 23, 3, 2, 2, 2, 7, 25, 3,
	2, 2, 2, 9, 27, 3, 2, 2, 2, 11, 29, 3, 2, 2, 2, 13, 31, 3, 2, 2, 2, 15,
	33, 3, 2, 2, 2, 17, 38, 3, 2, 2, 2, 19, 40, 3, 2, 2, 2, 21, 22, 7, 42,
	2, 2, 22, 4, 3, 2, 2, 2, 23, 24, 7, 43, 2, 2, 24, 6, 3, 2, 2, 2, 25, 26,
	7, 46, 2, 2, 26, 8, 3, 2, 2, 2, 27, 28, 7, 48, 2, 2, 28, 10, 3, 2, 2, 2,
	29, 30, 9, 2, 2, 2, 30, 12, 3, 2, 2, 2, 31, 32, 9, 3, 2, 2, 32, 14, 3,
	2, 2, 2, 33, 34, 9, 4, 2, 2, 34, 16, 3, 2, 2, 2, 35, 39, 5, 15, 8, 2, 36,
	39, 5, 13, 7, 2, 37, 39, 7, 97, 2, 2, 38, 35, 3, 2, 2, 2, 38, 36, 3, 2,
	2, 2, 38, 37, 3, 2, 2, 2, 39, 18, 3, 2, 2, 2, 40, 44, 5, 15, 8, 2, 41,
	43, 5, 17, 9, 2, 42, 41, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2, 44, 42, 3, 2,
	2, 2, 44, 45, 3, 2, 2, 2, 45, 20, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 5, 2,
	38, 44, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "','", "'.'",
}

var lexerSymbolicNames = []string{
	"", "LPAREN", "RPAREN", "COMMA", "PERIOD", "NEWLINE", "OBJECT",
}

var lexerRuleNames = []string{
	"LPAREN", "RPAREN", "COMMA", "PERIOD", "NEWLINE", "DIGIT", "SMALL_LETTER",
	"ALPHANUMERIC", "OBJECT",
}

type ILPLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewILPLangLexer(input antlr.CharStream) *ILPLangLexer {

	l := new(ILPLangLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ILPLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ILPLangLexer tokens.
const (
	ILPLangLexerLPAREN  = 1
	ILPLangLexerRPAREN  = 2
	ILPLangLexerCOMMA   = 3
	ILPLangLexerPERIOD  = 4
	ILPLangLexerNEWLINE = 5
	ILPLangLexerOBJECT  = 6
)
