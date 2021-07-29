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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 42, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 5, 2, 12, 10, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 19, 10, 3, 12, 3, 14, 3, 22, 11, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 33, 10, 4, 12,
	4, 14, 4, 36, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2,
	2, 2, 42, 2, 11, 3, 2, 2, 2, 4, 13, 3, 2, 2, 2, 6, 27, 3, 2, 2, 2, 8, 12,
	5, 4, 3, 2, 9, 12, 5, 6, 4, 2, 10, 12, 7, 2, 2, 3, 11, 8, 3, 2, 2, 2, 11,
	9, 3, 2, 2, 2, 11, 10, 3, 2, 2, 2, 12, 3, 3, 2, 2, 2, 13, 14, 7, 8, 2,
	2, 14, 15, 7, 3, 2, 2, 15, 20, 7, 8, 2, 2, 16, 17, 7, 5, 2, 2, 17, 19,
	7, 8, 2, 2, 18, 16, 3, 2, 2, 2, 19, 22, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2,
	20, 21, 3, 2, 2, 2, 21, 23, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 23, 24, 7,
	4, 2, 2, 24, 25, 7, 6, 2, 2, 25, 26, 7, 7, 2, 2, 26, 5, 3, 2, 2, 2, 27,
	28, 7, 8, 2, 2, 28, 29, 7, 3, 2, 2, 29, 34, 7, 8, 2, 2, 30, 31, 7, 5, 2,
	2, 31, 33, 7, 8, 2, 2, 32, 30, 3, 2, 2, 2, 33, 36, 3, 2, 2, 2, 34, 32,
	3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 37, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2,
	37, 38, 7, 4, 2, 2, 38, 39, 7, 6, 2, 2, 39, 40, 7, 2, 2, 3, 40, 7, 3, 2,
	2, 2, 5, 11, 20, 34,
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
	"start", "term", "lastterm",
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
	ILPLangParserRULE_start    = 0
	ILPLangParserRULE_term     = 1
	ILPLangParserRULE_lastterm = 2
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

func (s *StartContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *StartContext) Lastterm() ILasttermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILasttermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILasttermContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(ILPLangParserEOF, 0)
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

	p.SetState(9)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(6)
			p.Term()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(7)
			p.Lastterm()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(8)
			p.Match(ILPLangParserEOF)
		}

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
		p.SetState(11)
		p.Match(ILPLangParserOBJECT)
	}
	{
		p.SetState(12)
		p.Match(ILPLangParserLPAREN)
	}
	{
		p.SetState(13)
		p.Match(ILPLangParserOBJECT)
	}
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ILPLangParserCOMMA {
		{
			p.SetState(14)
			p.Match(ILPLangParserCOMMA)
		}
		{
			p.SetState(15)
			p.Match(ILPLangParserOBJECT)
		}

		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(21)
		p.Match(ILPLangParserRPAREN)
	}
	{
		p.SetState(22)
		p.Match(ILPLangParserPERIOD)
	}
	{
		p.SetState(23)
		p.Match(ILPLangParserNEWLINE)
	}

	return localctx
}

// ILasttermContext is an interface to support dynamic dispatch.
type ILasttermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLasttermContext differentiates from other interfaces.
	IsLasttermContext()
}

type LasttermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLasttermContext() *LasttermContext {
	var p = new(LasttermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ILPLangParserRULE_lastterm
	return p
}

func (*LasttermContext) IsLasttermContext() {}

func NewLasttermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LasttermContext {
	var p = new(LasttermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ILPLangParserRULE_lastterm

	return p
}

func (s *LasttermContext) GetParser() antlr.Parser { return s.parser }

func (s *LasttermContext) AllOBJECT() []antlr.TerminalNode {
	return s.GetTokens(ILPLangParserOBJECT)
}

func (s *LasttermContext) OBJECT(i int) antlr.TerminalNode {
	return s.GetToken(ILPLangParserOBJECT, i)
}

func (s *LasttermContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ILPLangParserLPAREN, 0)
}

func (s *LasttermContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ILPLangParserRPAREN, 0)
}

func (s *LasttermContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(ILPLangParserPERIOD, 0)
}

func (s *LasttermContext) EOF() antlr.TerminalNode {
	return s.GetToken(ILPLangParserEOF, 0)
}

func (s *LasttermContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ILPLangParserCOMMA)
}

func (s *LasttermContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ILPLangParserCOMMA, i)
}

func (s *LasttermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LasttermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LasttermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ILPLangListener); ok {
		listenerT.EnterLastterm(s)
	}
}

func (s *LasttermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ILPLangListener); ok {
		listenerT.ExitLastterm(s)
	}
}

func (p *ILPLangParser) Lastterm() (localctx ILasttermContext) {
	localctx = NewLasttermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ILPLangParserRULE_lastterm)
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
		p.SetState(25)
		p.Match(ILPLangParserOBJECT)
	}
	{
		p.SetState(26)
		p.Match(ILPLangParserLPAREN)
	}
	{
		p.SetState(27)
		p.Match(ILPLangParserOBJECT)
	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ILPLangParserCOMMA {
		{
			p.SetState(28)
			p.Match(ILPLangParserCOMMA)
		}
		{
			p.SetState(29)
			p.Match(ILPLangParserOBJECT)
		}

		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(35)
		p.Match(ILPLangParserRPAREN)
	}
	{
		p.SetState(36)
		p.Match(ILPLangParserPERIOD)
	}
	{
		p.SetState(37)
		p.Match(ILPLangParserEOF)
	}

	return localctx
}
