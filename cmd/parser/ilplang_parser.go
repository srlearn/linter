// Code generated from ILPLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ILPLang

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 29, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 7, 2, 8, 10, 2, 12, 2, 14, 2, 11, 11, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 20, 10, 3, 12, 3, 14, 3, 23,
	11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 4, 2, 4, 2, 2, 2, 28, 2, 9,
	3, 2, 2, 2, 4, 14, 3, 2, 2, 2, 6, 8, 5, 4, 3, 2, 7, 6, 3, 2, 2, 2, 8, 11,
	3, 2, 2, 2, 9, 7, 3, 2, 2, 2, 9, 10, 3, 2, 2, 2, 10, 12, 3, 2, 2, 2, 11,
	9, 3, 2, 2, 2, 12, 13, 7, 2, 2, 3, 13, 3, 3, 2, 2, 2, 14, 15, 7, 8, 2,
	2, 15, 16, 7, 3, 2, 2, 16, 21, 7, 8, 2, 2, 17, 18, 7, 5, 2, 2, 18, 20,
	7, 8, 2, 2, 19, 17, 3, 2, 2, 2, 20, 23, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2,
	21, 22, 3, 2, 2, 2, 22, 24, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 24, 25, 7,
	4, 2, 2, 25, 26, 7, 6, 2, 2, 26, 27, 7, 7, 2, 2, 27, 5, 3, 2, 2, 2, 4,
	9, 21,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "','", "'.'",
}
var symbolicNames = []string{
	"", "LPAREN", "RPAREN", "COMMA", "PERIOD", "NEWLINE", "OBJECT",
}

var ruleNames = []string{
	"start", "term",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ILPLangParser struct {
	*antlr.BaseParser
}

func NewILPLangParser(input antlr.TokenStream) *ILPLangParser {
	this := new(ILPLangParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ILPLang.g4"

	return this
}

// ILPLangParser tokens.
const (
	ILPLangParserEOF     = antlr.TokenEOF
	ILPLangParserLPAREN  = 1
	ILPLangParserRPAREN  = 2
	ILPLangParserCOMMA   = 3
	ILPLangParserPERIOD  = 4
	ILPLangParserNEWLINE = 5
	ILPLangParserOBJECT  = 6
)

// ILPLangParser rules.
const (
	ILPLangParserRULE_start = 0
	ILPLangParserRULE_term  = 1
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ILPLangParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ILPLangParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(ILPLangParserEOF, 0)
}

func (s *StartContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *StartContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ILPLangListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ILPLangListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *ILPLangParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ILPLangParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(7)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ILPLangParserOBJECT {
		{
			p.SetState(4)
			p.Term()
		}

		p.SetState(9)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(10)
		p.Match(ILPLangParserEOF)
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ILPLangParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ILPLangParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllOBJECT() []antlr.TerminalNode {
	return s.GetTokens(ILPLangParserOBJECT)
}

func (s *TermContext) OBJECT(i int) antlr.TerminalNode {
	return s.GetToken(ILPLangParserOBJECT, i)
}

func (s *TermContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ILPLangParserLPAREN, 0)
}

func (s *TermContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ILPLangParserRPAREN, 0)
}

func (s *TermContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(ILPLangParserPERIOD, 0)
}

func (s *TermContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ILPLangParserNEWLINE, 0)
}

func (s *TermContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ILPLangParserCOMMA)
}

func (s *TermContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ILPLangParserCOMMA, i)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ILPLangListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ILPLangListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *ILPLangParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ILPLangParserRULE_term)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.Match(ILPLangParserOBJECT)
	}
	{
		p.SetState(13)
		p.Match(ILPLangParserLPAREN)
	}
	{
		p.SetState(14)
		p.Match(ILPLangParserOBJECT)
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ILPLangParserCOMMA {
		{
			p.SetState(15)
			p.Match(ILPLangParserCOMMA)
		}
		{
			p.SetState(16)
			p.Match(ILPLangParserOBJECT)
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(22)
		p.Match(ILPLangParserRPAREN)
	}
	{
		p.SetState(23)
		p.Match(ILPLangParserPERIOD)
	}
	{
		p.SetState(24)
		p.Match(ILPLangParserNEWLINE)
	}

	return localctx
}
