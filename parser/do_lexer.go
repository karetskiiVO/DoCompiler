// Code generated from Do.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type DoLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var DoLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dolexerLexerInit() {
	staticData := &DoLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'var'", "'act'", "'('", "')'", "'{'", "'}'", "'with'", "'*'", "'pipe'",
		"'struct'", "';'", "'behavour'", "','", "'glob'", "'<'", "'>'", "'='",
		"'_'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "BOOL", "STRING", "COMPARETOKEN", "NUMBER", "NAME", "EMPTY",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "BOOL", "STRING", "COMPARETOKEN", "NUMBER", "NAME",
		"EMPTY",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 167, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 3, 19, 124, 8, 19, 1, 20, 1, 20, 5, 20, 128, 8, 20,
		10, 20, 12, 20, 131, 9, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 144, 8, 21, 1, 22, 3, 22, 147, 8,
		22, 1, 22, 4, 22, 150, 8, 22, 11, 22, 12, 22, 151, 1, 23, 1, 23, 5, 23,
		156, 8, 23, 10, 23, 12, 23, 159, 9, 23, 1, 24, 4, 24, 162, 8, 24, 11, 24,
		12, 24, 163, 1, 24, 1, 24, 1, 129, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47,
		24, 49, 25, 1, 0, 6, 2, 0, 60, 60, 62, 62, 2, 0, 43, 43, 45, 45, 1, 0,
		48, 57, 2, 0, 65, 90, 97, 122, 3, 0, 48, 57, 65, 90, 97, 122, 3, 0, 9,
		10, 13, 13, 32, 32, 176, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 1, 51, 1, 0,
		0, 0, 3, 55, 1, 0, 0, 0, 5, 59, 1, 0, 0, 0, 7, 61, 1, 0, 0, 0, 9, 63, 1,
		0, 0, 0, 11, 65, 1, 0, 0, 0, 13, 67, 1, 0, 0, 0, 15, 72, 1, 0, 0, 0, 17,
		74, 1, 0, 0, 0, 19, 79, 1, 0, 0, 0, 21, 86, 1, 0, 0, 0, 23, 88, 1, 0, 0,
		0, 25, 97, 1, 0, 0, 0, 27, 99, 1, 0, 0, 0, 29, 104, 1, 0, 0, 0, 31, 106,
		1, 0, 0, 0, 33, 108, 1, 0, 0, 0, 35, 110, 1, 0, 0, 0, 37, 112, 1, 0, 0,
		0, 39, 123, 1, 0, 0, 0, 41, 125, 1, 0, 0, 0, 43, 143, 1, 0, 0, 0, 45, 146,
		1, 0, 0, 0, 47, 153, 1, 0, 0, 0, 49, 161, 1, 0, 0, 0, 51, 52, 5, 118, 0,
		0, 52, 53, 5, 97, 0, 0, 53, 54, 5, 114, 0, 0, 54, 2, 1, 0, 0, 0, 55, 56,
		5, 97, 0, 0, 56, 57, 5, 99, 0, 0, 57, 58, 5, 116, 0, 0, 58, 4, 1, 0, 0,
		0, 59, 60, 5, 40, 0, 0, 60, 6, 1, 0, 0, 0, 61, 62, 5, 41, 0, 0, 62, 8,
		1, 0, 0, 0, 63, 64, 5, 123, 0, 0, 64, 10, 1, 0, 0, 0, 65, 66, 5, 125, 0,
		0, 66, 12, 1, 0, 0, 0, 67, 68, 5, 119, 0, 0, 68, 69, 5, 105, 0, 0, 69,
		70, 5, 116, 0, 0, 70, 71, 5, 104, 0, 0, 71, 14, 1, 0, 0, 0, 72, 73, 5,
		42, 0, 0, 73, 16, 1, 0, 0, 0, 74, 75, 5, 112, 0, 0, 75, 76, 5, 105, 0,
		0, 76, 77, 5, 112, 0, 0, 77, 78, 5, 101, 0, 0, 78, 18, 1, 0, 0, 0, 79,
		80, 5, 115, 0, 0, 80, 81, 5, 116, 0, 0, 81, 82, 5, 114, 0, 0, 82, 83, 5,
		117, 0, 0, 83, 84, 5, 99, 0, 0, 84, 85, 5, 116, 0, 0, 85, 20, 1, 0, 0,
		0, 86, 87, 5, 59, 0, 0, 87, 22, 1, 0, 0, 0, 88, 89, 5, 98, 0, 0, 89, 90,
		5, 101, 0, 0, 90, 91, 5, 104, 0, 0, 91, 92, 5, 97, 0, 0, 92, 93, 5, 118,
		0, 0, 93, 94, 5, 111, 0, 0, 94, 95, 5, 117, 0, 0, 95, 96, 5, 114, 0, 0,
		96, 24, 1, 0, 0, 0, 97, 98, 5, 44, 0, 0, 98, 26, 1, 0, 0, 0, 99, 100, 5,
		103, 0, 0, 100, 101, 5, 108, 0, 0, 101, 102, 5, 111, 0, 0, 102, 103, 5,
		98, 0, 0, 103, 28, 1, 0, 0, 0, 104, 105, 5, 60, 0, 0, 105, 30, 1, 0, 0,
		0, 106, 107, 5, 62, 0, 0, 107, 32, 1, 0, 0, 0, 108, 109, 5, 61, 0, 0, 109,
		34, 1, 0, 0, 0, 110, 111, 5, 95, 0, 0, 111, 36, 1, 0, 0, 0, 112, 113, 5,
		46, 0, 0, 113, 38, 1, 0, 0, 0, 114, 115, 5, 116, 0, 0, 115, 116, 5, 114,
		0, 0, 116, 117, 5, 117, 0, 0, 117, 124, 5, 101, 0, 0, 118, 119, 5, 102,
		0, 0, 119, 120, 5, 97, 0, 0, 120, 121, 5, 108, 0, 0, 121, 122, 5, 115,
		0, 0, 122, 124, 5, 101, 0, 0, 123, 114, 1, 0, 0, 0, 123, 118, 1, 0, 0,
		0, 124, 40, 1, 0, 0, 0, 125, 129, 5, 34, 0, 0, 126, 128, 9, 0, 0, 0, 127,
		126, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 129, 127,
		1, 0, 0, 0, 130, 132, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 132, 133, 5, 34,
		0, 0, 133, 42, 1, 0, 0, 0, 134, 135, 5, 61, 0, 0, 135, 144, 5, 61, 0, 0,
		136, 137, 5, 60, 0, 0, 137, 144, 5, 61, 0, 0, 138, 139, 5, 62, 0, 0, 139,
		144, 5, 61, 0, 0, 140, 144, 7, 0, 0, 0, 141, 142, 5, 33, 0, 0, 142, 144,
		5, 61, 0, 0, 143, 134, 1, 0, 0, 0, 143, 136, 1, 0, 0, 0, 143, 138, 1, 0,
		0, 0, 143, 140, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 44, 1, 0, 0, 0,
		145, 147, 7, 1, 0, 0, 146, 145, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147,
		149, 1, 0, 0, 0, 148, 150, 7, 2, 0, 0, 149, 148, 1, 0, 0, 0, 150, 151,
		1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 46, 1, 0,
		0, 0, 153, 157, 7, 3, 0, 0, 154, 156, 7, 4, 0, 0, 155, 154, 1, 0, 0, 0,
		156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158,
		48, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 162, 7, 5, 0, 0, 161, 160, 1,
		0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0,
		0, 164, 165, 1, 0, 0, 0, 165, 166, 6, 24, 0, 0, 166, 50, 1, 0, 0, 0, 8,
		0, 123, 129, 143, 146, 151, 157, 163, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// DoLexerInit initializes any static state used to implement DoLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewDoLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func DoLexerInit() {
	staticData := &DoLexerLexerStaticData
	staticData.once.Do(dolexerLexerInit)
}

// NewDoLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewDoLexer(input antlr.CharStream) *DoLexer {
	DoLexerInit()
	l := new(DoLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &DoLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Do.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DoLexer tokens.
const (
	DoLexerT__0         = 1
	DoLexerT__1         = 2
	DoLexerT__2         = 3
	DoLexerT__3         = 4
	DoLexerT__4         = 5
	DoLexerT__5         = 6
	DoLexerT__6         = 7
	DoLexerT__7         = 8
	DoLexerT__8         = 9
	DoLexerT__9         = 10
	DoLexerT__10        = 11
	DoLexerT__11        = 12
	DoLexerT__12        = 13
	DoLexerT__13        = 14
	DoLexerT__14        = 15
	DoLexerT__15        = 16
	DoLexerT__16        = 17
	DoLexerT__17        = 18
	DoLexerT__18        = 19
	DoLexerBOOL         = 20
	DoLexerSTRING       = 21
	DoLexerCOMPARETOKEN = 22
	DoLexerNUMBER       = 23
	DoLexerNAME         = 24
	DoLexerEMPTY        = 25
)
