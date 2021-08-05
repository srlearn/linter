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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 57, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 5,
	2, 15, 10, 2, 3, 2, 3, 2, 3, 3, 7, 3, 20, 10, 3, 12, 3, 14, 3, 23, 11,
	3, 3, 4, 7, 4, 26, 10, 4, 12, 4, 14, 4, 29, 11, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 7, 5, 36, 10, 5, 12, 5, 14, 5, 39, 11, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 4, 3, 2, 10, 11, 3, 2, 9, 10, 2, 55,
	2, 14, 3, 2, 2, 2, 4, 21, 3, 2, 2, 2, 6, 27, 3, 2, 2, 2, 8, 30, 3, 2, 2,
	2, 10, 44, 3, 2, 2, 2, 12, 15, 5, 4, 3, 2, 13, 15, 5, 6, 4, 2, 14, 12,
	3, 2, 2, 2, 14, 13, 3, 2, 2, 2, 15, 16, 3, 2, 2, 2, 16, 17, 7, 2, 2, 3,
	17, 3, 3, 2, 2, 2, 18, 20, 5, 8, 5, 2, 19, 18, 3, 2, 2, 2, 20, 23, 3, 2,
	2, 2, 21, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 5, 3, 2, 2, 2, 23, 21,
	3, 2, 2, 2, 24, 26, 5, 10, 6, 2, 25, 24, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2,
	27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 7, 3, 2, 2, 2, 29, 27, 3, 2,
	2, 2, 30, 31, 7, 11, 2, 2, 31, 32, 7, 4, 2, 2, 32, 37, 7, 11, 2, 2, 33,
	34, 7, 6, 2, 2, 34, 36, 9, 2, 2, 2, 35, 33, 3, 2, 2, 2, 36, 39, 3, 2, 2,
	2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 40, 3, 2, 2, 2, 39, 37,
	3, 2, 2, 2, 40, 41, 7, 5, 2, 2, 41, 42, 7, 7, 2, 2, 42, 43, 7, 8, 2, 2,
	43, 9, 3, 2, 2, 2, 44, 45, 7, 3, 2, 2, 45, 46, 7, 4, 2, 2, 46, 47, 7, 11,
	2, 2, 47, 48, 7, 4, 2, 2, 48, 49, 7, 11, 2, 2, 49, 50, 7, 5, 2, 2, 50,
	51, 7, 6, 2, 2, 51, 52, 9, 3, 2, 2, 52, 53, 7, 5, 2, 2, 53, 54, 7, 7, 2,
	2, 54, 55, 7, 8, 2, 2, 55, 11, 3, 2, 2, 2, 6, 14, 21, 27, 37,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'regressionExample'", "'('", "')'", "','", "'.'",
}
var symbolicNames = []string{
	"", "", "LPAREN", "RPAREN", "COMMA", "PERIOD", "NEWLINE", "FLOAT", "INTEGER",
	"OBJECT",
}

var ruleNames = []string{
	"start", "classification_task", "regression_task", "term", "regression_term",
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
	ILPLangParserT__0    = 1
	ILPLangParserLPAREN  = 2
	ILPLangParserRPAREN  = 3
	ILPLangParserCOMMA   = 4
	ILPLangParserPERIOD  = 5
	ILPLangParserNEWLINE = 6
	ILPLangParserFLOAT   = 7
	ILPLangParserINTEGER = 8
	ILPLangParserOBJECT  = 9
)

// ILPLangParser rules.
const (
	ILPLangParserRULE_start               = 0
	ILPLangParserRULE_classification_task = 1
	ILPLangParserRULE_regression_task     = 2
	ILPLangParserRULE_term                = 3
	ILPLangParserRULE_regression_term     = 4
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

func (s *StartContext) Classification_task() IClassification_taskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassification_taskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassification_taskContext)
}

func (s *StartContext) Regression_task() IRegression_taskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegression_taskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegression_taskContext)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(12)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(10)
			p.Classification_task()
		}

	case 2:
		{
			p.SetState(11)
			p.Regression_task()
		}

	}
	{
		p.SetState(14)
		p.Match(ILPLangParserEOF)
	}

	return localctx
}

// IClassification_taskContext is an interface to support dynamic dispatch.
type IClassification_taskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassification_taskContext differentiates from other interfaces.
	IsClassification_taskContext()
}

type Classification_taskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassification_taskContext() *Classification_taskContext {
	var p = new(Classification_taskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ILPLangParserRULE_classification_task
	return p
}

func (*Classification_taskContext) IsClassification_taskContext() {}

func NewClassification_taskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Classification_taskContext {
	var p = new(Classification_taskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ILPLangParserRULE_classification_task

	return p
}

func (s *Classification_taskContext) GetParser() antlr.Parser { return s.parser }

func (s *Classification_taskContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *Classification_taskContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Classification_taskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Classification_taskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Classification_taskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ILPLangListener); ok {
		listenerT.EnterClassification_task(s)
	}
}

func (s *Classification_taskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ILPLangListener); ok {
		listenerT.ExitClassification_task(s)
	}
}

func (p *ILPLangParser) Classification_task() (localctx IClassification_taskContext) {
	localctx = NewClassification_taskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ILPLangParserRULE_classification_task)
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ILPLangParserOBJECT {
		{
			p.SetState(16)
			p.Term()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRegression_taskContext is an interface to support dynamic dispatch.
type IRegression_taskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegression_taskContext differentiates from other interfaces.
	IsRegression_taskContext()
}

type Regression_taskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegression_taskContext() *Regression_taskContext {
	var p = new(Regression_taskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ILPLangParserRULE_regression_task
	return p
}

func (*Regression_taskContext) IsRegression_taskContext() {}

func NewRegression_taskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Regression_taskContext {
	var p = new(Regression_taskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ILPLangParserRULE_regression_task

	return p
}

func (s *Regression_taskContext) GetParser() antlr.Parser { return s.parser }

func (s *Regression_taskContext) AllRegression_term() []IRegression_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegression_termContext)(nil)).Elem())
	var tst = make([]IRegression_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegression_termContext)
		}
	}

	return tst
}

func (s *Regression_taskContext) Regression_term(i int) IRegression_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegression_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegression_termContext)
}

func (s *Regression_taskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Regression_taskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Regression_taskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ILPLangListener); ok {
		listenerT.EnterRegression_task(s)
	}
}

func (s *Regression_taskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ILPLangListener); ok {
		listenerT.ExitRegression_task(s)
	}
}

func (p *ILPLangParser) Regression_task() (localctx IRegression_taskContext) {
	localctx = NewRegression_taskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ILPLangParserRULE_regression_task)
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
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ILPLangParserT__0 {
		{
			p.SetState(22)
			p.Regression_term()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *TermContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(ILPLangParserINTEGER)
}

func (s *TermContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(ILPLangParserINTEGER, i)
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
	p.EnterRule(localctx, 6, ILPLangParserRULE_term)
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
		p.SetState(28)
		p.Match(ILPLangParserOBJECT)
	}
	{
		p.SetState(29)
		p.Match(ILPLangParserLPAREN)
	}
	{
		p.SetState(30)
		p.Match(ILPLangParserOBJECT)
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ILPLangParserCOMMA {
		{
			p.SetState(31)
			p.Match(ILPLangParserCOMMA)
		}
		{
			p.SetState(32)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ILPLangParserINTEGER || _la == ILPLangParserOBJECT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
		p.Match(ILPLangParserRPAREN)
	}
	{
		p.SetState(39)
		p.Match(ILPLangParserPERIOD)
	}
	{
		p.SetState(40)
		p.Match(ILPLangParserNEWLINE)
	}

	return localctx
}

// IRegression_termContext is an interface to support dynamic dispatch.
type IRegression_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegression_termContext differentiates from other interfaces.
	IsRegression_termContext()
}

type Regression_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegression_termContext() *Regression_termContext {
	var p = new(Regression_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ILPLangParserRULE_regression_term
	return p
}

func (*Regression_termContext) IsRegression_termContext() {}

func NewRegression_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Regression_termContext {
	var p = new(Regression_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ILPLangParserRULE_regression_term

	return p
}

func (s *Regression_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Regression_termContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(ILPLangParserLPAREN)
}

func (s *Regression_termContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(ILPLangParserLPAREN, i)
}

func (s *Regression_termContext) AllOBJECT() []antlr.TerminalNode {
	return s.GetTokens(ILPLangParserOBJECT)
}

func (s *Regression_termContext) OBJECT(i int) antlr.TerminalNode {
	return s.GetToken(ILPLangParserOBJECT, i)
}

func (s *Regression_termContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(ILPLangParserRPAREN)
}

func (s *Regression_termContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(ILPLangParserRPAREN, i)
}

func (s *Regression_termContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ILPLangParserCOMMA, 0)
}

func (s *Regression_termContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(ILPLangParserPERIOD, 0)
}

func (s *Regression_termContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ILPLangParserNEWLINE, 0)
}

func (s *Regression_termContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ILPLangParserFLOAT, 0)
}

func (s *Regression_termContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ILPLangParserINTEGER, 0)
}

func (s *Regression_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Regression_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Regression_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ILPLangListener); ok {
		listenerT.EnterRegression_term(s)
	}
}

func (s *Regression_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ILPLangListener); ok {
		listenerT.ExitRegression_term(s)
	}
}

func (p *ILPLangParser) Regression_term() (localctx IRegression_termContext) {
	localctx = NewRegression_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ILPLangParserRULE_regression_term)
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
		p.SetState(42)
		p.Match(ILPLangParserT__0)
	}
	{
		p.SetState(43)
		p.Match(ILPLangParserLPAREN)
	}
	{
		p.SetState(44)
		p.Match(ILPLangParserOBJECT)
	}
	{
		p.SetState(45)
		p.Match(ILPLangParserLPAREN)
	}
	{
		p.SetState(46)
		p.Match(ILPLangParserOBJECT)
	}
	{
		p.SetState(47)
		p.Match(ILPLangParserRPAREN)
	}
	{
		p.SetState(48)
		p.Match(ILPLangParserCOMMA)
	}
	{
		p.SetState(49)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ILPLangParserFLOAT || _la == ILPLangParserINTEGER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(50)
		p.Match(ILPLangParserRPAREN)
	}
	{
		p.SetState(51)
		p.Match(ILPLangParserPERIOD)
	}
	{
		p.SetState(52)
		p.Match(ILPLangParserNEWLINE)
	}

	return localctx
}
