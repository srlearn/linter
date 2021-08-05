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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 86, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 5, 10, 63, 10, 10, 3, 11, 7, 11, 66, 10, 11, 12, 11, 14, 11, 69,
	11, 11, 3, 11, 3, 11, 6, 11, 73, 10, 11, 13, 11, 14, 11, 74, 3, 12, 6,
	12, 78, 10, 12, 13, 12, 14, 12, 79, 3, 13, 6, 13, 83, 10, 13, 13, 13, 14,
	13, 84, 2, 2, 14, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 2, 17, 2, 19,
	2, 21, 9, 23, 10, 25, 11, 3, 2, 5, 3, 2, 12, 12, 3, 2, 50, 59, 3, 2, 99,
	124, 2, 88, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 3, 27, 3, 2, 2, 2, 5, 45, 3, 2, 2, 2,
	7, 47, 3, 2, 2, 2, 9, 49, 3, 2, 2, 2, 11, 51, 3, 2, 2, 2, 13, 53, 3, 2,
	2, 2, 15, 55, 3, 2, 2, 2, 17, 57, 3, 2, 2, 2, 19, 62, 3, 2, 2, 2, 21, 67,
	3, 2, 2, 2, 23, 77, 3, 2, 2, 2, 25, 82, 3, 2, 2, 2, 27, 28, 7, 116, 2,
	2, 28, 29, 7, 103, 2, 2, 29, 30, 7, 105, 2, 2, 30, 31, 7, 116, 2, 2, 31,
	32, 7, 103, 2, 2, 32, 33, 7, 117, 2, 2, 33, 34, 7, 117, 2, 2, 34, 35, 7,
	107, 2, 2, 35, 36, 7, 113, 2, 2, 36, 37, 7, 112, 2, 2, 37, 38, 7, 71, 2,
	2, 38, 39, 7, 122, 2, 2, 39, 40, 7, 99, 2, 2, 40, 41, 7, 111, 2, 2, 41,
	42, 7, 114, 2, 2, 42, 43, 7, 110, 2, 2, 43, 44, 7, 103, 2, 2, 44, 4, 3,
	2, 2, 2, 45, 46, 7, 42, 2, 2, 46, 6, 3, 2, 2, 2, 47, 48, 7, 43, 2, 2, 48,
	8, 3, 2, 2, 2, 49, 50, 7, 46, 2, 2, 50, 10, 3, 2, 2, 2, 51, 52, 7, 48,
	2, 2, 52, 12, 3, 2, 2, 2, 53, 54, 9, 2, 2, 2, 54, 14, 3, 2, 2, 2, 55, 56,
	9, 3, 2, 2, 56, 16, 3, 2, 2, 2, 57, 58, 9, 4, 2, 2, 58, 18, 3, 2, 2, 2,
	59, 63, 5, 17, 9, 2, 60, 63, 5, 15, 8, 2, 61, 63, 7, 97, 2, 2, 62, 59,
	3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 61, 3, 2, 2, 2, 63, 20, 3, 2, 2, 2,
	64, 66, 5, 15, 8, 2, 65, 64, 3, 2, 2, 2, 66, 69, 3, 2, 2, 2, 67, 65, 3,
	2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 70, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70,
	72, 5, 11, 6, 2, 71, 73, 5, 15, 8, 2, 72, 71, 3, 2, 2, 2, 73, 74, 3, 2,
	2, 2, 74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 22, 3, 2, 2, 2, 76, 78,
	5, 15, 8, 2, 77, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2,
	79, 80, 3, 2, 2, 2, 80, 24, 3, 2, 2, 2, 81, 83, 5, 19, 10, 2, 82, 81, 3,
	2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85,
	26, 3, 2, 2, 2, 8, 2, 62, 67, 74, 79, 84, 2,
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
	"", "'regressionExample'", "'('", "')'", "','", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "LPAREN", "RPAREN", "COMMA", "PERIOD", "NEWLINE", "FLOAT", "INTEGER",
	"OBJECT",
}

var lexerRuleNames = []string{
	"T__0", "LPAREN", "RPAREN", "COMMA", "PERIOD", "NEWLINE", "DIGIT", "SMALL_LETTER",
	"ALPHANUMERIC", "FLOAT", "INTEGER", "OBJECT",
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
	ILPLangLexerT__0    = 1
	ILPLangLexerLPAREN  = 2
	ILPLangLexerRPAREN  = 3
	ILPLangLexerCOMMA   = 4
	ILPLangLexerPERIOD  = 5
	ILPLangLexerNEWLINE = 6
	ILPLangLexerFLOAT   = 7
	ILPLangLexerINTEGER = 8
	ILPLangLexerOBJECT  = 9
)
