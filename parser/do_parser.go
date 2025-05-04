// Code generated from Do.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Do
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type DoParser struct {
	*antlr.BaseParser
}

var DoParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func doParserInit() {
	staticData := &DoParserStaticData
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
		"program", "definition", "globalVariableDefinition", "functionDefinition",
		"typeDefinition", "type", "structdefinition", "behavourdefinition",
		"typetuple", "arglist", "argsublist", "basetypefild", "varfield", "globalvarfield",
		"fieldname", "argname", "typename", "genericparamslist", "genericarglist",
		"statement", "assign", "expressiontuple", "expression", "functioncall",
		"expressiontuplelhv", "expressiontuplerhv", "variableuse", "constantuse",
		"emptyexpression", "dividedname", "bool", "string", "number",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 276, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 5, 0, 68, 8, 0, 10, 0, 12, 0, 71, 9, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 3, 1, 78, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 3, 3, 87, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 95, 8,
		3, 10, 3, 12, 3, 98, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 105, 8,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 116, 8,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 123, 8, 6, 1, 6, 1, 6, 5, 6, 127,
		8, 6, 10, 6, 12, 6, 130, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		3, 8, 139, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 145, 8, 8, 10, 8, 12, 8,
		148, 9, 8, 1, 8, 1, 8, 3, 8, 152, 8, 8, 1, 9, 3, 9, 155, 8, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 3, 9, 161, 8, 9, 1, 10, 1, 10, 1, 10, 5, 10, 166, 8, 10,
		10, 10, 12, 10, 169, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16,
		1, 16, 3, 16, 189, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 195, 8, 17,
		10, 17, 12, 17, 198, 9, 17, 3, 17, 200, 8, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 5, 18, 208, 8, 18, 10, 18, 12, 18, 211, 9, 18, 3, 18,
		213, 8, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 3, 20, 222,
		8, 20, 1, 21, 1, 21, 1, 21, 5, 21, 227, 8, 21, 10, 21, 12, 21, 230, 9,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 240,
		8, 22, 1, 23, 1, 23, 1, 23, 3, 23, 245, 8, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 3, 27, 258, 8, 27,
		1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 5, 29, 265, 8, 29, 10, 29, 12, 29, 268,
		9, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 0, 0, 33, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 0, 0, 276, 0, 69, 1, 0,
		0, 0, 2, 77, 1, 0, 0, 0, 4, 79, 1, 0, 0, 0, 6, 83, 1, 0, 0, 0, 8, 101,
		1, 0, 0, 0, 10, 115, 1, 0, 0, 0, 12, 117, 1, 0, 0, 0, 14, 133, 1, 0, 0,
		0, 16, 151, 1, 0, 0, 0, 18, 160, 1, 0, 0, 0, 20, 162, 1, 0, 0, 0, 22, 172,
		1, 0, 0, 0, 24, 175, 1, 0, 0, 0, 26, 178, 1, 0, 0, 0, 28, 182, 1, 0, 0,
		0, 30, 184, 1, 0, 0, 0, 32, 186, 1, 0, 0, 0, 34, 190, 1, 0, 0, 0, 36, 203,
		1, 0, 0, 0, 38, 216, 1, 0, 0, 0, 40, 218, 1, 0, 0, 0, 42, 223, 1, 0, 0,
		0, 44, 239, 1, 0, 0, 0, 46, 241, 1, 0, 0, 0, 48, 248, 1, 0, 0, 0, 50, 250,
		1, 0, 0, 0, 52, 252, 1, 0, 0, 0, 54, 257, 1, 0, 0, 0, 56, 259, 1, 0, 0,
		0, 58, 261, 1, 0, 0, 0, 60, 269, 1, 0, 0, 0, 62, 271, 1, 0, 0, 0, 64, 273,
		1, 0, 0, 0, 66, 68, 3, 2, 1, 0, 67, 66, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0,
		69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 72, 1, 0, 0, 0, 71, 69, 1,
		0, 0, 0, 72, 73, 5, 0, 0, 1, 73, 1, 1, 0, 0, 0, 74, 78, 3, 6, 3, 0, 75,
		78, 3, 8, 4, 0, 76, 78, 3, 4, 2, 0, 77, 74, 1, 0, 0, 0, 77, 75, 1, 0, 0,
		0, 77, 76, 1, 0, 0, 0, 78, 3, 1, 0, 0, 0, 79, 80, 5, 1, 0, 0, 80, 81, 5,
		24, 0, 0, 81, 82, 3, 32, 16, 0, 82, 5, 1, 0, 0, 0, 83, 84, 5, 2, 0, 0,
		84, 86, 5, 24, 0, 0, 85, 87, 3, 34, 17, 0, 86, 85, 1, 0, 0, 0, 86, 87,
		1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 5, 3, 0, 0, 89, 90, 3, 18, 9, 0,
		90, 91, 5, 4, 0, 0, 91, 92, 3, 16, 8, 0, 92, 96, 5, 5, 0, 0, 93, 95, 3,
		38, 19, 0, 94, 93, 1, 0, 0, 0, 95, 98, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0,
		96, 97, 1, 0, 0, 0, 97, 99, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 99, 100, 5,
		6, 0, 0, 100, 7, 1, 0, 0, 0, 101, 102, 5, 7, 0, 0, 102, 104, 5, 24, 0,
		0, 103, 105, 3, 34, 17, 0, 104, 103, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0,
		105, 106, 1, 0, 0, 0, 106, 107, 3, 10, 5, 0, 107, 9, 1, 0, 0, 0, 108, 109,
		5, 8, 0, 0, 109, 116, 3, 10, 5, 0, 110, 111, 5, 9, 0, 0, 111, 116, 3, 10,
		5, 0, 112, 116, 3, 32, 16, 0, 113, 116, 3, 12, 6, 0, 114, 116, 3, 14, 7,
		0, 115, 108, 1, 0, 0, 0, 115, 110, 1, 0, 0, 0, 115, 112, 1, 0, 0, 0, 115,
		113, 1, 0, 0, 0, 115, 114, 1, 0, 0, 0, 116, 11, 1, 0, 0, 0, 117, 118, 5,
		10, 0, 0, 118, 128, 5, 5, 0, 0, 119, 123, 3, 22, 11, 0, 120, 123, 3, 24,
		12, 0, 121, 123, 3, 26, 13, 0, 122, 119, 1, 0, 0, 0, 122, 120, 1, 0, 0,
		0, 122, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 5, 11, 0, 0, 125,
		127, 1, 0, 0, 0, 126, 122, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126,
		1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130, 128, 1, 0,
		0, 0, 131, 132, 5, 6, 0, 0, 132, 13, 1, 0, 0, 0, 133, 134, 5, 12, 0, 0,
		134, 135, 5, 5, 0, 0, 135, 136, 5, 6, 0, 0, 136, 15, 1, 0, 0, 0, 137, 139,
		3, 10, 5, 0, 138, 137, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 152, 1, 0,
		0, 0, 140, 141, 5, 3, 0, 0, 141, 146, 3, 10, 5, 0, 142, 143, 5, 13, 0,
		0, 143, 145, 3, 10, 5, 0, 144, 142, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146,
		144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 149, 1, 0, 0, 0, 148, 146,
		1, 0, 0, 0, 149, 150, 5, 4, 0, 0, 150, 152, 1, 0, 0, 0, 151, 138, 1, 0,
		0, 0, 151, 140, 1, 0, 0, 0, 152, 17, 1, 0, 0, 0, 153, 155, 3, 20, 10, 0,
		154, 153, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 161, 1, 0, 0, 0, 156,
		157, 3, 20, 10, 0, 157, 158, 5, 13, 0, 0, 158, 159, 3, 20, 10, 0, 159,
		161, 1, 0, 0, 0, 160, 154, 1, 0, 0, 0, 160, 156, 1, 0, 0, 0, 161, 19, 1,
		0, 0, 0, 162, 167, 3, 30, 15, 0, 163, 164, 5, 13, 0, 0, 164, 166, 3, 30,
		15, 0, 165, 163, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0,
		167, 168, 1, 0, 0, 0, 168, 170, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170,
		171, 3, 10, 5, 0, 171, 21, 1, 0, 0, 0, 172, 173, 3, 10, 5, 0, 173, 174,
		3, 36, 18, 0, 174, 23, 1, 0, 0, 0, 175, 176, 3, 28, 14, 0, 176, 177, 3,
		10, 5, 0, 177, 25, 1, 0, 0, 0, 178, 179, 3, 28, 14, 0, 179, 180, 5, 14,
		0, 0, 180, 181, 3, 10, 5, 0, 181, 27, 1, 0, 0, 0, 182, 183, 5, 24, 0, 0,
		183, 29, 1, 0, 0, 0, 184, 185, 5, 24, 0, 0, 185, 31, 1, 0, 0, 0, 186, 188,
		3, 58, 29, 0, 187, 189, 3, 34, 17, 0, 188, 187, 1, 0, 0, 0, 188, 189, 1,
		0, 0, 0, 189, 33, 1, 0, 0, 0, 190, 199, 5, 15, 0, 0, 191, 196, 5, 24, 0,
		0, 192, 193, 5, 13, 0, 0, 193, 195, 5, 24, 0, 0, 194, 192, 1, 0, 0, 0,
		195, 198, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197,
		200, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 191, 1, 0, 0, 0, 199, 200,
		1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 5, 16, 0, 0, 202, 35, 1, 0,
		0, 0, 203, 212, 5, 15, 0, 0, 204, 209, 3, 10, 5, 0, 205, 206, 5, 13, 0,
		0, 206, 208, 3, 10, 5, 0, 207, 205, 1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209,
		207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 213, 1, 0, 0, 0, 211, 209,
		1, 0, 0, 0, 212, 204, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214, 1, 0,
		0, 0, 214, 215, 5, 16, 0, 0, 215, 37, 1, 0, 0, 0, 216, 217, 3, 40, 20,
		0, 217, 39, 1, 0, 0, 0, 218, 221, 3, 48, 24, 0, 219, 220, 5, 17, 0, 0,
		220, 222, 3, 50, 25, 0, 221, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222,
		41, 1, 0, 0, 0, 223, 228, 3, 44, 22, 0, 224, 225, 5, 13, 0, 0, 225, 227,
		3, 44, 22, 0, 226, 224, 1, 0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1,
		0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 43, 1, 0, 0, 0, 230, 228, 1, 0, 0,
		0, 231, 240, 3, 56, 28, 0, 232, 240, 3, 52, 26, 0, 233, 234, 5, 3, 0, 0,
		234, 235, 3, 42, 21, 0, 235, 236, 5, 4, 0, 0, 236, 240, 1, 0, 0, 0, 237,
		240, 3, 54, 27, 0, 238, 240, 3, 46, 23, 0, 239, 231, 1, 0, 0, 0, 239, 232,
		1, 0, 0, 0, 239, 233, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 239, 238, 1, 0,
		0, 0, 240, 45, 1, 0, 0, 0, 241, 242, 3, 58, 29, 0, 242, 244, 5, 3, 0, 0,
		243, 245, 3, 42, 21, 0, 244, 243, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245,
		246, 1, 0, 0, 0, 246, 247, 5, 4, 0, 0, 247, 47, 1, 0, 0, 0, 248, 249, 3,
		42, 21, 0, 249, 49, 1, 0, 0, 0, 250, 251, 3, 42, 21, 0, 251, 51, 1, 0,
		0, 0, 252, 253, 3, 58, 29, 0, 253, 53, 1, 0, 0, 0, 254, 258, 3, 60, 30,
		0, 255, 258, 3, 62, 31, 0, 256, 258, 3, 64, 32, 0, 257, 254, 1, 0, 0, 0,
		257, 255, 1, 0, 0, 0, 257, 256, 1, 0, 0, 0, 258, 55, 1, 0, 0, 0, 259, 260,
		5, 18, 0, 0, 260, 57, 1, 0, 0, 0, 261, 266, 5, 24, 0, 0, 262, 263, 5, 19,
		0, 0, 263, 265, 5, 24, 0, 0, 264, 262, 1, 0, 0, 0, 265, 268, 1, 0, 0, 0,
		266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 59, 1, 0, 0, 0, 268, 266,
		1, 0, 0, 0, 269, 270, 5, 20, 0, 0, 270, 61, 1, 0, 0, 0, 271, 272, 5, 21,
		0, 0, 272, 63, 1, 0, 0, 0, 273, 274, 5, 23, 0, 0, 274, 65, 1, 0, 0, 0,
		25, 69, 77, 86, 96, 104, 115, 122, 128, 138, 146, 151, 154, 160, 167, 188,
		196, 199, 209, 212, 221, 228, 239, 244, 257, 266,
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

// DoParserInit initializes any static state used to implement DoParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DoParserInit() {
	staticData := &DoParserStaticData
	staticData.once.Do(doParserInit)
}

// NewDoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDoParser(input antlr.TokenStream) *DoParser {
	DoParserInit()
	this := new(DoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &DoParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Do.g4"

	return this
}

// DoParser tokens.
const (
	DoParserEOF          = antlr.TokenEOF
	DoParserT__0         = 1
	DoParserT__1         = 2
	DoParserT__2         = 3
	DoParserT__3         = 4
	DoParserT__4         = 5
	DoParserT__5         = 6
	DoParserT__6         = 7
	DoParserT__7         = 8
	DoParserT__8         = 9
	DoParserT__9         = 10
	DoParserT__10        = 11
	DoParserT__11        = 12
	DoParserT__12        = 13
	DoParserT__13        = 14
	DoParserT__14        = 15
	DoParserT__15        = 16
	DoParserT__16        = 17
	DoParserT__17        = 18
	DoParserT__18        = 19
	DoParserBOOL         = 20
	DoParserSTRING       = 21
	DoParserCOMPARETOKEN = 22
	DoParserNUMBER       = 23
	DoParserNAME         = 24
	DoParserEMPTY        = 25
)

// DoParser rules.
const (
	DoParserRULE_program                  = 0
	DoParserRULE_definition               = 1
	DoParserRULE_globalVariableDefinition = 2
	DoParserRULE_functionDefinition       = 3
	DoParserRULE_typeDefinition           = 4
	DoParserRULE_type                     = 5
	DoParserRULE_structdefinition         = 6
	DoParserRULE_behavourdefinition       = 7
	DoParserRULE_typetuple                = 8
	DoParserRULE_arglist                  = 9
	DoParserRULE_argsublist               = 10
	DoParserRULE_basetypefild             = 11
	DoParserRULE_varfield                 = 12
	DoParserRULE_globalvarfield           = 13
	DoParserRULE_fieldname                = 14
	DoParserRULE_argname                  = 15
	DoParserRULE_typename                 = 16
	DoParserRULE_genericparamslist        = 17
	DoParserRULE_genericarglist           = 18
	DoParserRULE_statement                = 19
	DoParserRULE_assign                   = 20
	DoParserRULE_expressiontuple          = 21
	DoParserRULE_expression               = 22
	DoParserRULE_functioncall             = 23
	DoParserRULE_expressiontuplelhv       = 24
	DoParserRULE_expressiontuplerhv       = 25
	DoParserRULE_variableuse              = 26
	DoParserRULE_constantuse              = 27
	DoParserRULE_emptyexpression          = 28
	DoParserRULE_dividedname              = 29
	DoParserRULE_bool                     = 30
	DoParserRULE_string                   = 31
	DoParserRULE_number                   = 32
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllDefinition() []IDefinitionContext
	Definition(i int) IDefinitionContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(DoParserEOF, 0)
}

func (s *ProgramContext) AllDefinition() []IDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefinitionContext); ok {
			tst[i] = t.(IDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Definition(i int) IDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *DoParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DoParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&134) != 0 {
		{
			p.SetState(66)
			p.Definition()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
		p.Match(DoParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionDefinition() IFunctionDefinitionContext
	TypeDefinition() ITypeDefinitionContext
	GlobalVariableDefinition() IGlobalVariableDefinitionContext

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_definition
	return p
}

func InitEmptyDefinitionContext(p *DefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_definition
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) FunctionDefinition() IFunctionDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionContext)
}

func (s *DefinitionContext) TypeDefinition() ITypeDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefinitionContext)
}

func (s *DefinitionContext) GlobalVariableDefinition() IGlobalVariableDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalVariableDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobalVariableDefinitionContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *DoParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DoParserRULE_definition)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DoParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.FunctionDefinition()
		}

	case DoParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.TypeDefinition()
		}

	case DoParserT__0:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(76)
			p.GlobalVariableDefinition()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGlobalVariableDefinitionContext is an interface to support dynamic dispatch.
type IGlobalVariableDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	Typename() ITypenameContext

	// IsGlobalVariableDefinitionContext differentiates from other interfaces.
	IsGlobalVariableDefinitionContext()
}

type GlobalVariableDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalVariableDefinitionContext() *GlobalVariableDefinitionContext {
	var p = new(GlobalVariableDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_globalVariableDefinition
	return p
}

func InitEmptyGlobalVariableDefinitionContext(p *GlobalVariableDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_globalVariableDefinition
}

func (*GlobalVariableDefinitionContext) IsGlobalVariableDefinitionContext() {}

func NewGlobalVariableDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalVariableDefinitionContext {
	var p = new(GlobalVariableDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_globalVariableDefinition

	return p
}

func (s *GlobalVariableDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalVariableDefinitionContext) NAME() antlr.TerminalNode {
	return s.GetToken(DoParserNAME, 0)
}

func (s *GlobalVariableDefinitionContext) Typename() ITypenameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypenameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypenameContext)
}

func (s *GlobalVariableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalVariableDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlobalVariableDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterGlobalVariableDefinition(s)
	}
}

func (s *GlobalVariableDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitGlobalVariableDefinition(s)
	}
}

func (p *DoParser) GlobalVariableDefinition() (localctx IGlobalVariableDefinitionContext) {
	localctx = NewGlobalVariableDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DoParserRULE_globalVariableDefinition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(DoParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Match(DoParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Typename()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDefinitionContext is an interface to support dynamic dispatch.
type IFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	Arglist() IArglistContext
	Typetuple() ITypetupleContext
	Genericparamslist() IGenericparamslistContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefinitionContext() *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_functionDefinition
	return p
}

func InitEmptyFunctionDefinitionContext(p *FunctionDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_functionDefinition
}

func (*FunctionDefinitionContext) IsFunctionDefinitionContext() {}

func NewFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_functionDefinition

	return p
}

func (s *FunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionContext) NAME() antlr.TerminalNode {
	return s.GetToken(DoParserNAME, 0)
}

func (s *FunctionDefinitionContext) Arglist() IArglistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArglistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *FunctionDefinitionContext) Typetuple() ITypetupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypetupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypetupleContext)
}

func (s *FunctionDefinitionContext) Genericparamslist() IGenericparamslistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericparamslistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericparamslistContext)
}

func (s *FunctionDefinitionContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *FunctionDefinitionContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitFunctionDefinition(s)
	}
}

func (p *DoParser) FunctionDefinition() (localctx IFunctionDefinitionContext) {
	localctx = NewFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DoParserRULE_functionDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(DoParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(DoParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DoParserT__14 {
		{
			p.SetState(85)
			p.Genericparamslist()
		}

	}
	{
		p.SetState(88)
		p.Match(DoParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.Arglist()
	}
	{
		p.SetState(90)
		p.Match(DoParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Typetuple()
	}
	{
		p.SetState(92)
		p.Match(DoParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&28573704) != 0 {
		{
			p.SetState(93)
			p.Statement()
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
		p.Match(DoParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeDefinitionContext is an interface to support dynamic dispatch.
type ITypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	Type_() ITypeContext
	Genericparamslist() IGenericparamslistContext

	// IsTypeDefinitionContext differentiates from other interfaces.
	IsTypeDefinitionContext()
}

type TypeDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefinitionContext() *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_typeDefinition
	return p
}

func InitEmptyTypeDefinitionContext(p *TypeDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_typeDefinition
}

func (*TypeDefinitionContext) IsTypeDefinitionContext() {}

func NewTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_typeDefinition

	return p
}

func (s *TypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefinitionContext) NAME() antlr.TerminalNode {
	return s.GetToken(DoParserNAME, 0)
}

func (s *TypeDefinitionContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeDefinitionContext) Genericparamslist() IGenericparamslistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericparamslistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericparamslistContext)
}

func (s *TypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterTypeDefinition(s)
	}
}

func (s *TypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitTypeDefinition(s)
	}
}

func (p *DoParser) TypeDefinition() (localctx ITypeDefinitionContext) {
	localctx = NewTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DoParserRULE_typeDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(DoParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(DoParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DoParserT__14 {
		{
			p.SetState(103)
			p.Genericparamslist()
		}

	}
	{
		p.SetState(106)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Typename() ITypenameContext
	Structdefinition() IStructdefinitionContext
	Behavourdefinition() IBehavourdefinitionContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeContext) Typename() ITypenameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypenameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypenameContext)
}

func (s *TypeContext) Structdefinition() IStructdefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructdefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructdefinitionContext)
}

func (s *TypeContext) Behavourdefinition() IBehavourdefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBehavourdefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBehavourdefinitionContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *DoParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DoParserRULE_type)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DoParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Match(DoParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Type_()
		}

	case DoParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.Match(DoParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Type_()
		}

	case DoParserNAME:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.Typename()
		}

	case DoParserT__9:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.Structdefinition()
		}

	case DoParserT__11:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(114)
			p.Behavourdefinition()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructdefinitionContext is an interface to support dynamic dispatch.
type IStructdefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBasetypefild() []IBasetypefildContext
	Basetypefild(i int) IBasetypefildContext
	AllVarfield() []IVarfieldContext
	Varfield(i int) IVarfieldContext
	AllGlobalvarfield() []IGlobalvarfieldContext
	Globalvarfield(i int) IGlobalvarfieldContext

	// IsStructdefinitionContext differentiates from other interfaces.
	IsStructdefinitionContext()
}

type StructdefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructdefinitionContext() *StructdefinitionContext {
	var p = new(StructdefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_structdefinition
	return p
}

func InitEmptyStructdefinitionContext(p *StructdefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_structdefinition
}

func (*StructdefinitionContext) IsStructdefinitionContext() {}

func NewStructdefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructdefinitionContext {
	var p = new(StructdefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_structdefinition

	return p
}

func (s *StructdefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *StructdefinitionContext) AllBasetypefild() []IBasetypefildContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBasetypefildContext); ok {
			len++
		}
	}

	tst := make([]IBasetypefildContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBasetypefildContext); ok {
			tst[i] = t.(IBasetypefildContext)
			i++
		}
	}

	return tst
}

func (s *StructdefinitionContext) Basetypefild(i int) IBasetypefildContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasetypefildContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasetypefildContext)
}

func (s *StructdefinitionContext) AllVarfield() []IVarfieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarfieldContext); ok {
			len++
		}
	}

	tst := make([]IVarfieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarfieldContext); ok {
			tst[i] = t.(IVarfieldContext)
			i++
		}
	}

	return tst
}

func (s *StructdefinitionContext) Varfield(i int) IVarfieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarfieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarfieldContext)
}

func (s *StructdefinitionContext) AllGlobalvarfield() []IGlobalvarfieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGlobalvarfieldContext); ok {
			len++
		}
	}

	tst := make([]IGlobalvarfieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGlobalvarfieldContext); ok {
			tst[i] = t.(IGlobalvarfieldContext)
			i++
		}
	}

	return tst
}

func (s *StructdefinitionContext) Globalvarfield(i int) IGlobalvarfieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalvarfieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobalvarfieldContext)
}

func (s *StructdefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructdefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructdefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterStructdefinition(s)
	}
}

func (s *StructdefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitStructdefinition(s)
	}
}

func (p *DoParser) Structdefinition() (localctx IStructdefinitionContext) {
	localctx = NewStructdefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DoParserRULE_structdefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(DoParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(DoParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16783104) != 0 {
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(119)
				p.Basetypefild()
			}

		case 2:
			{
				p.SetState(120)
				p.Varfield()
			}

		case 3:
			{
				p.SetState(121)
				p.Globalvarfield()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(124)
			p.Match(DoParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(131)
		p.Match(DoParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBehavourdefinitionContext is an interface to support dynamic dispatch.
type IBehavourdefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBehavourdefinitionContext differentiates from other interfaces.
	IsBehavourdefinitionContext()
}

type BehavourdefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBehavourdefinitionContext() *BehavourdefinitionContext {
	var p = new(BehavourdefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_behavourdefinition
	return p
}

func InitEmptyBehavourdefinitionContext(p *BehavourdefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_behavourdefinition
}

func (*BehavourdefinitionContext) IsBehavourdefinitionContext() {}

func NewBehavourdefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BehavourdefinitionContext {
	var p = new(BehavourdefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_behavourdefinition

	return p
}

func (s *BehavourdefinitionContext) GetParser() antlr.Parser { return s.parser }
func (s *BehavourdefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BehavourdefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BehavourdefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterBehavourdefinition(s)
	}
}

func (s *BehavourdefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitBehavourdefinition(s)
	}
}

func (p *DoParser) Behavourdefinition() (localctx IBehavourdefinitionContext) {
	localctx = NewBehavourdefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DoParserRULE_behavourdefinition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(DoParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Match(DoParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.Match(DoParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypetupleContext is an interface to support dynamic dispatch.
type ITypetupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllType_() []ITypeContext
	Type_(i int) ITypeContext

	// IsTypetupleContext differentiates from other interfaces.
	IsTypetupleContext()
}

type TypetupleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypetupleContext() *TypetupleContext {
	var p = new(TypetupleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_typetuple
	return p
}

func InitEmptyTypetupleContext(p *TypetupleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_typetuple
}

func (*TypetupleContext) IsTypetupleContext() {}

func NewTypetupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypetupleContext {
	var p = new(TypetupleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_typetuple

	return p
}

func (s *TypetupleContext) GetParser() antlr.Parser { return s.parser }

func (s *TypetupleContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *TypetupleContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypetupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypetupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypetupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterTypetuple(s)
	}
}

func (s *TypetupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitTypetuple(s)
	}
}

func (p *DoParser) Typetuple() (localctx ITypetupleContext) {
	localctx = NewTypetupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DoParserRULE_typetuple)
	var _la int

	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DoParserT__4, DoParserT__7, DoParserT__8, DoParserT__9, DoParserT__11, DoParserNAME:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16783104) != 0 {
			{
				p.SetState(137)
				p.Type_()
			}

		}

	case DoParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.Match(DoParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Type_()
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == DoParserT__12 {
			{
				p.SetState(142)
				p.Match(DoParserT__12)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(143)
				p.Type_()
			}

			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(149)
			p.Match(DoParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArglistContext is an interface to support dynamic dispatch.
type IArglistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArgsublist() []IArgsublistContext
	Argsublist(i int) IArgsublistContext

	// IsArglistContext differentiates from other interfaces.
	IsArglistContext()
}

type ArglistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArglistContext() *ArglistContext {
	var p = new(ArglistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_arglist
	return p
}

func InitEmptyArglistContext(p *ArglistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_arglist
}

func (*ArglistContext) IsArglistContext() {}

func NewArglistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArglistContext {
	var p = new(ArglistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_arglist

	return p
}

func (s *ArglistContext) GetParser() antlr.Parser { return s.parser }

func (s *ArglistContext) AllArgsublist() []IArgsublistContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgsublistContext); ok {
			len++
		}
	}

	tst := make([]IArgsublistContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgsublistContext); ok {
			tst[i] = t.(IArgsublistContext)
			i++
		}
	}

	return tst
}

func (s *ArglistContext) Argsublist(i int) IArgsublistContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsublistContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsublistContext)
}

func (s *ArglistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArglistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArglistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterArglist(s)
	}
}

func (s *ArglistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitArglist(s)
	}
}

func (p *DoParser) Arglist() (localctx IArglistContext) {
	localctx = NewArglistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DoParserRULE_arglist)
	var _la int

	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == DoParserNAME {
			{
				p.SetState(153)
				p.Argsublist()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Argsublist()
		}

		{
			p.SetState(157)
			p.Match(DoParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Argsublist()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgsublistContext is an interface to support dynamic dispatch.
type IArgsublistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArgname() []IArgnameContext
	Argname(i int) IArgnameContext
	Type_() ITypeContext

	// IsArgsublistContext differentiates from other interfaces.
	IsArgsublistContext()
}

type ArgsublistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsublistContext() *ArgsublistContext {
	var p = new(ArgsublistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_argsublist
	return p
}

func InitEmptyArgsublistContext(p *ArgsublistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_argsublist
}

func (*ArgsublistContext) IsArgsublistContext() {}

func NewArgsublistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsublistContext {
	var p = new(ArgsublistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_argsublist

	return p
}

func (s *ArgsublistContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsublistContext) AllArgname() []IArgnameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgnameContext); ok {
			len++
		}
	}

	tst := make([]IArgnameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgnameContext); ok {
			tst[i] = t.(IArgnameContext)
			i++
		}
	}

	return tst
}

func (s *ArgsublistContext) Argname(i int) IArgnameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgnameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgnameContext)
}

func (s *ArgsublistContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ArgsublistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsublistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsublistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterArgsublist(s)
	}
}

func (s *ArgsublistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitArgsublist(s)
	}
}

func (p *DoParser) Argsublist() (localctx IArgsublistContext) {
	localctx = NewArgsublistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DoParserRULE_argsublist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Argname()
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DoParserT__12 {
		{
			p.SetState(163)
			p.Match(DoParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.Argname()
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(170)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBasetypefildContext is an interface to support dynamic dispatch.
type IBasetypefildContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Genericarglist() IGenericarglistContext

	// IsBasetypefildContext differentiates from other interfaces.
	IsBasetypefildContext()
}

type BasetypefildContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasetypefildContext() *BasetypefildContext {
	var p = new(BasetypefildContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_basetypefild
	return p
}

func InitEmptyBasetypefildContext(p *BasetypefildContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_basetypefild
}

func (*BasetypefildContext) IsBasetypefildContext() {}

func NewBasetypefildContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasetypefildContext {
	var p = new(BasetypefildContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_basetypefild

	return p
}

func (s *BasetypefildContext) GetParser() antlr.Parser { return s.parser }

func (s *BasetypefildContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *BasetypefildContext) Genericarglist() IGenericarglistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericarglistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericarglistContext)
}

func (s *BasetypefildContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasetypefildContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasetypefildContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterBasetypefild(s)
	}
}

func (s *BasetypefildContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitBasetypefild(s)
	}
}

func (p *DoParser) Basetypefild() (localctx IBasetypefildContext) {
	localctx = NewBasetypefildContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DoParserRULE_basetypefild)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Type_()
	}
	{
		p.SetState(173)
		p.Genericarglist()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarfieldContext is an interface to support dynamic dispatch.
type IVarfieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Fieldname() IFieldnameContext
	Type_() ITypeContext

	// IsVarfieldContext differentiates from other interfaces.
	IsVarfieldContext()
}

type VarfieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarfieldContext() *VarfieldContext {
	var p = new(VarfieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_varfield
	return p
}

func InitEmptyVarfieldContext(p *VarfieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_varfield
}

func (*VarfieldContext) IsVarfieldContext() {}

func NewVarfieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarfieldContext {
	var p = new(VarfieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_varfield

	return p
}

func (s *VarfieldContext) GetParser() antlr.Parser { return s.parser }

func (s *VarfieldContext) Fieldname() IFieldnameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldnameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldnameContext)
}

func (s *VarfieldContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VarfieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarfieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarfieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterVarfield(s)
	}
}

func (s *VarfieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitVarfield(s)
	}
}

func (p *DoParser) Varfield() (localctx IVarfieldContext) {
	localctx = NewVarfieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DoParserRULE_varfield)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Fieldname()
	}
	{
		p.SetState(176)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGlobalvarfieldContext is an interface to support dynamic dispatch.
type IGlobalvarfieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Fieldname() IFieldnameContext
	Type_() ITypeContext

	// IsGlobalvarfieldContext differentiates from other interfaces.
	IsGlobalvarfieldContext()
}

type GlobalvarfieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalvarfieldContext() *GlobalvarfieldContext {
	var p = new(GlobalvarfieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_globalvarfield
	return p
}

func InitEmptyGlobalvarfieldContext(p *GlobalvarfieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_globalvarfield
}

func (*GlobalvarfieldContext) IsGlobalvarfieldContext() {}

func NewGlobalvarfieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalvarfieldContext {
	var p = new(GlobalvarfieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_globalvarfield

	return p
}

func (s *GlobalvarfieldContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalvarfieldContext) Fieldname() IFieldnameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldnameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldnameContext)
}

func (s *GlobalvarfieldContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *GlobalvarfieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalvarfieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlobalvarfieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterGlobalvarfield(s)
	}
}

func (s *GlobalvarfieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitGlobalvarfield(s)
	}
}

func (p *DoParser) Globalvarfield() (localctx IGlobalvarfieldContext) {
	localctx = NewGlobalvarfieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DoParserRULE_globalvarfield)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Fieldname()
	}
	{
		p.SetState(179)
		p.Match(DoParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldnameContext is an interface to support dynamic dispatch.
type IFieldnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode

	// IsFieldnameContext differentiates from other interfaces.
	IsFieldnameContext()
}

type FieldnameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldnameContext() *FieldnameContext {
	var p = new(FieldnameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_fieldname
	return p
}

func InitEmptyFieldnameContext(p *FieldnameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_fieldname
}

func (*FieldnameContext) IsFieldnameContext() {}

func NewFieldnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldnameContext {
	var p = new(FieldnameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_fieldname

	return p
}

func (s *FieldnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldnameContext) NAME() antlr.TerminalNode {
	return s.GetToken(DoParserNAME, 0)
}

func (s *FieldnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterFieldname(s)
	}
}

func (s *FieldnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitFieldname(s)
	}
}

func (p *DoParser) Fieldname() (localctx IFieldnameContext) {
	localctx = NewFieldnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DoParserRULE_fieldname)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(DoParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgnameContext is an interface to support dynamic dispatch.
type IArgnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode

	// IsArgnameContext differentiates from other interfaces.
	IsArgnameContext()
}

type ArgnameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgnameContext() *ArgnameContext {
	var p = new(ArgnameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_argname
	return p
}

func InitEmptyArgnameContext(p *ArgnameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_argname
}

func (*ArgnameContext) IsArgnameContext() {}

func NewArgnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgnameContext {
	var p = new(ArgnameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_argname

	return p
}

func (s *ArgnameContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgnameContext) NAME() antlr.TerminalNode {
	return s.GetToken(DoParserNAME, 0)
}

func (s *ArgnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterArgname(s)
	}
}

func (s *ArgnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitArgname(s)
	}
}

func (p *DoParser) Argname() (localctx IArgnameContext) {
	localctx = NewArgnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DoParserRULE_argname)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(DoParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypenameContext is an interface to support dynamic dispatch.
type ITypenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Dividedname() IDividednameContext
	Genericparamslist() IGenericparamslistContext

	// IsTypenameContext differentiates from other interfaces.
	IsTypenameContext()
}

type TypenameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypenameContext() *TypenameContext {
	var p = new(TypenameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_typename
	return p
}

func InitEmptyTypenameContext(p *TypenameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_typename
}

func (*TypenameContext) IsTypenameContext() {}

func NewTypenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypenameContext {
	var p = new(TypenameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_typename

	return p
}

func (s *TypenameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypenameContext) Dividedname() IDividednameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDividednameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDividednameContext)
}

func (s *TypenameContext) Genericparamslist() IGenericparamslistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericparamslistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericparamslistContext)
}

func (s *TypenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypenameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterTypename(s)
	}
}

func (s *TypenameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitTypename(s)
	}
}

func (p *DoParser) Typename() (localctx ITypenameContext) {
	localctx = NewTypenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DoParserRULE_typename)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Dividedname()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(187)
			p.Genericparamslist()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGenericparamslistContext is an interface to support dynamic dispatch.
type IGenericparamslistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNAME() []antlr.TerminalNode
	NAME(i int) antlr.TerminalNode

	// IsGenericparamslistContext differentiates from other interfaces.
	IsGenericparamslistContext()
}

type GenericparamslistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericparamslistContext() *GenericparamslistContext {
	var p = new(GenericparamslistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_genericparamslist
	return p
}

func InitEmptyGenericparamslistContext(p *GenericparamslistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_genericparamslist
}

func (*GenericparamslistContext) IsGenericparamslistContext() {}

func NewGenericparamslistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericparamslistContext {
	var p = new(GenericparamslistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_genericparamslist

	return p
}

func (s *GenericparamslistContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericparamslistContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(DoParserNAME)
}

func (s *GenericparamslistContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(DoParserNAME, i)
}

func (s *GenericparamslistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericparamslistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericparamslistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterGenericparamslist(s)
	}
}

func (s *GenericparamslistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitGenericparamslist(s)
	}
}

func (p *DoParser) Genericparamslist() (localctx IGenericparamslistContext) {
	localctx = NewGenericparamslistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DoParserRULE_genericparamslist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(DoParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DoParserNAME {
		{
			p.SetState(191)
			p.Match(DoParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == DoParserT__12 {
			{
				p.SetState(192)
				p.Match(DoParserT__12)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(193)
				p.Match(DoParserNAME)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(201)
		p.Match(DoParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGenericarglistContext is an interface to support dynamic dispatch.
type IGenericarglistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllType_() []ITypeContext
	Type_(i int) ITypeContext

	// IsGenericarglistContext differentiates from other interfaces.
	IsGenericarglistContext()
}

type GenericarglistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericarglistContext() *GenericarglistContext {
	var p = new(GenericarglistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_genericarglist
	return p
}

func InitEmptyGenericarglistContext(p *GenericarglistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_genericarglist
}

func (*GenericarglistContext) IsGenericarglistContext() {}

func NewGenericarglistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericarglistContext {
	var p = new(GenericarglistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_genericarglist

	return p
}

func (s *GenericarglistContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericarglistContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *GenericarglistContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *GenericarglistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericarglistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericarglistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterGenericarglist(s)
	}
}

func (s *GenericarglistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitGenericarglist(s)
	}
}

func (p *DoParser) Genericarglist() (localctx IGenericarglistContext) {
	localctx = NewGenericarglistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DoParserRULE_genericarglist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(DoParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16783104) != 0 {
		{
			p.SetState(204)
			p.Type_()
		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == DoParserT__12 {
			{
				p.SetState(205)
				p.Match(DoParserT__12)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(206)
				p.Type_()
			}

			p.SetState(211)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(214)
		p.Match(DoParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assign() IAssignContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *DoParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DoParserRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Assign()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expressiontuplelhv() IExpressiontuplelhvContext
	Expressiontuplerhv() IExpressiontuplerhvContext

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_assign
	return p
}

func InitEmptyAssignContext(p *AssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_assign
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) Expressiontuplelhv() IExpressiontuplelhvContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressiontuplelhvContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressiontuplelhvContext)
}

func (s *AssignContext) Expressiontuplerhv() IExpressiontuplerhvContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressiontuplerhvContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressiontuplerhvContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *DoParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DoParserRULE_assign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Expressiontuplelhv()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DoParserT__16 {
		{
			p.SetState(219)
			p.Match(DoParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)
			p.Expressiontuplerhv()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressiontupleContext is an interface to support dynamic dispatch.
type IExpressiontupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsExpressiontupleContext differentiates from other interfaces.
	IsExpressiontupleContext()
}

type ExpressiontupleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressiontupleContext() *ExpressiontupleContext {
	var p = new(ExpressiontupleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_expressiontuple
	return p
}

func InitEmptyExpressiontupleContext(p *ExpressiontupleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_expressiontuple
}

func (*ExpressiontupleContext) IsExpressiontupleContext() {}

func NewExpressiontupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressiontupleContext {
	var p = new(ExpressiontupleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_expressiontuple

	return p
}

func (s *ExpressiontupleContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressiontupleContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressiontupleContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressiontupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressiontupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressiontupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterExpressiontuple(s)
	}
}

func (s *ExpressiontupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitExpressiontuple(s)
	}
}

func (p *DoParser) Expressiontuple() (localctx IExpressiontupleContext) {
	localctx = NewExpressiontupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, DoParserRULE_expressiontuple)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Expression()
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DoParserT__12 {
		{
			p.SetState(224)
			p.Match(DoParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.Expression()
		}

		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Emptyexpression() IEmptyexpressionContext
	Variableuse() IVariableuseContext
	Expressiontuple() IExpressiontupleContext
	Constantuse() IConstantuseContext
	Functioncall() IFunctioncallContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Emptyexpression() IEmptyexpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyexpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyexpressionContext)
}

func (s *ExpressionContext) Variableuse() IVariableuseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableuseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableuseContext)
}

func (s *ExpressionContext) Expressiontuple() IExpressiontupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressiontupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressiontupleContext)
}

func (s *ExpressionContext) Constantuse() IConstantuseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantuseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantuseContext)
}

func (s *ExpressionContext) Functioncall() IFunctioncallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctioncallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctioncallContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DoParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, DoParserRULE_expression)
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(231)
			p.Emptyexpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(232)
			p.Variableuse()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(233)
			p.Match(DoParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Expressiontuple()
		}
		{
			p.SetState(235)
			p.Match(DoParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(237)
			p.Constantuse()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(238)
			p.Functioncall()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctioncallContext is an interface to support dynamic dispatch.
type IFunctioncallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Dividedname() IDividednameContext
	Expressiontuple() IExpressiontupleContext

	// IsFunctioncallContext differentiates from other interfaces.
	IsFunctioncallContext()
}

type FunctioncallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctioncallContext() *FunctioncallContext {
	var p = new(FunctioncallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_functioncall
	return p
}

func InitEmptyFunctioncallContext(p *FunctioncallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_functioncall
}

func (*FunctioncallContext) IsFunctioncallContext() {}

func NewFunctioncallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctioncallContext {
	var p = new(FunctioncallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_functioncall

	return p
}

func (s *FunctioncallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctioncallContext) Dividedname() IDividednameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDividednameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDividednameContext)
}

func (s *FunctioncallContext) Expressiontuple() IExpressiontupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressiontupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressiontupleContext)
}

func (s *FunctioncallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctioncallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctioncallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterFunctioncall(s)
	}
}

func (s *FunctioncallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitFunctioncall(s)
	}
}

func (p *DoParser) Functioncall() (localctx IFunctioncallContext) {
	localctx = NewFunctioncallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, DoParserRULE_functioncall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Dividedname()
	}
	{
		p.SetState(242)
		p.Match(DoParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&28573704) != 0 {
		{
			p.SetState(243)
			p.Expressiontuple()
		}

	}
	{
		p.SetState(246)
		p.Match(DoParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressiontuplelhvContext is an interface to support dynamic dispatch.
type IExpressiontuplelhvContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expressiontuple() IExpressiontupleContext

	// IsExpressiontuplelhvContext differentiates from other interfaces.
	IsExpressiontuplelhvContext()
}

type ExpressiontuplelhvContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressiontuplelhvContext() *ExpressiontuplelhvContext {
	var p = new(ExpressiontuplelhvContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_expressiontuplelhv
	return p
}

func InitEmptyExpressiontuplelhvContext(p *ExpressiontuplelhvContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_expressiontuplelhv
}

func (*ExpressiontuplelhvContext) IsExpressiontuplelhvContext() {}

func NewExpressiontuplelhvContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressiontuplelhvContext {
	var p = new(ExpressiontuplelhvContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_expressiontuplelhv

	return p
}

func (s *ExpressiontuplelhvContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressiontuplelhvContext) Expressiontuple() IExpressiontupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressiontupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressiontupleContext)
}

func (s *ExpressiontuplelhvContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressiontuplelhvContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressiontuplelhvContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterExpressiontuplelhv(s)
	}
}

func (s *ExpressiontuplelhvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitExpressiontuplelhv(s)
	}
}

func (p *DoParser) Expressiontuplelhv() (localctx IExpressiontuplelhvContext) {
	localctx = NewExpressiontuplelhvContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, DoParserRULE_expressiontuplelhv)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Expressiontuple()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressiontuplerhvContext is an interface to support dynamic dispatch.
type IExpressiontuplerhvContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expressiontuple() IExpressiontupleContext

	// IsExpressiontuplerhvContext differentiates from other interfaces.
	IsExpressiontuplerhvContext()
}

type ExpressiontuplerhvContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressiontuplerhvContext() *ExpressiontuplerhvContext {
	var p = new(ExpressiontuplerhvContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_expressiontuplerhv
	return p
}

func InitEmptyExpressiontuplerhvContext(p *ExpressiontuplerhvContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_expressiontuplerhv
}

func (*ExpressiontuplerhvContext) IsExpressiontuplerhvContext() {}

func NewExpressiontuplerhvContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressiontuplerhvContext {
	var p = new(ExpressiontuplerhvContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_expressiontuplerhv

	return p
}

func (s *ExpressiontuplerhvContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressiontuplerhvContext) Expressiontuple() IExpressiontupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressiontupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressiontupleContext)
}

func (s *ExpressiontuplerhvContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressiontuplerhvContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressiontuplerhvContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterExpressiontuplerhv(s)
	}
}

func (s *ExpressiontuplerhvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitExpressiontuplerhv(s)
	}
}

func (p *DoParser) Expressiontuplerhv() (localctx IExpressiontuplerhvContext) {
	localctx = NewExpressiontuplerhvContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, DoParserRULE_expressiontuplerhv)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.Expressiontuple()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableuseContext is an interface to support dynamic dispatch.
type IVariableuseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Dividedname() IDividednameContext

	// IsVariableuseContext differentiates from other interfaces.
	IsVariableuseContext()
}

type VariableuseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableuseContext() *VariableuseContext {
	var p = new(VariableuseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_variableuse
	return p
}

func InitEmptyVariableuseContext(p *VariableuseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_variableuse
}

func (*VariableuseContext) IsVariableuseContext() {}

func NewVariableuseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableuseContext {
	var p = new(VariableuseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_variableuse

	return p
}

func (s *VariableuseContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableuseContext) Dividedname() IDividednameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDividednameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDividednameContext)
}

func (s *VariableuseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableuseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableuseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterVariableuse(s)
	}
}

func (s *VariableuseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitVariableuse(s)
	}
}

func (p *DoParser) Variableuse() (localctx IVariableuseContext) {
	localctx = NewVariableuseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, DoParserRULE_variableuse)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Dividedname()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantuseContext is an interface to support dynamic dispatch.
type IConstantuseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Bool_() IBoolContext
	String_() IStringContext
	Number() INumberContext

	// IsConstantuseContext differentiates from other interfaces.
	IsConstantuseContext()
}

type ConstantuseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantuseContext() *ConstantuseContext {
	var p = new(ConstantuseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_constantuse
	return p
}

func InitEmptyConstantuseContext(p *ConstantuseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_constantuse
}

func (*ConstantuseContext) IsConstantuseContext() {}

func NewConstantuseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantuseContext {
	var p = new(ConstantuseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_constantuse

	return p
}

func (s *ConstantuseContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantuseContext) Bool_() IBoolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolContext)
}

func (s *ConstantuseContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *ConstantuseContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ConstantuseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantuseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantuseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterConstantuse(s)
	}
}

func (s *ConstantuseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitConstantuse(s)
	}
}

func (p *DoParser) Constantuse() (localctx IConstantuseContext) {
	localctx = NewConstantuseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, DoParserRULE_constantuse)
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DoParserBOOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(254)
			p.Bool_()
		}

	case DoParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(255)
			p.String_()
		}

	case DoParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(256)
			p.Number()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEmptyexpressionContext is an interface to support dynamic dispatch.
type IEmptyexpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEmptyexpressionContext differentiates from other interfaces.
	IsEmptyexpressionContext()
}

type EmptyexpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyexpressionContext() *EmptyexpressionContext {
	var p = new(EmptyexpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_emptyexpression
	return p
}

func InitEmptyEmptyexpressionContext(p *EmptyexpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_emptyexpression
}

func (*EmptyexpressionContext) IsEmptyexpressionContext() {}

func NewEmptyexpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyexpressionContext {
	var p = new(EmptyexpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_emptyexpression

	return p
}

func (s *EmptyexpressionContext) GetParser() antlr.Parser { return s.parser }
func (s *EmptyexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyexpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyexpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterEmptyexpression(s)
	}
}

func (s *EmptyexpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitEmptyexpression(s)
	}
}

func (p *DoParser) Emptyexpression() (localctx IEmptyexpressionContext) {
	localctx = NewEmptyexpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, DoParserRULE_emptyexpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(DoParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDividednameContext is an interface to support dynamic dispatch.
type IDividednameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNAME() []antlr.TerminalNode
	NAME(i int) antlr.TerminalNode

	// IsDividednameContext differentiates from other interfaces.
	IsDividednameContext()
}

type DividednameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDividednameContext() *DividednameContext {
	var p = new(DividednameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_dividedname
	return p
}

func InitEmptyDividednameContext(p *DividednameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_dividedname
}

func (*DividednameContext) IsDividednameContext() {}

func NewDividednameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DividednameContext {
	var p = new(DividednameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_dividedname

	return p
}

func (s *DividednameContext) GetParser() antlr.Parser { return s.parser }

func (s *DividednameContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(DoParserNAME)
}

func (s *DividednameContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(DoParserNAME, i)
}

func (s *DividednameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DividednameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DividednameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterDividedname(s)
	}
}

func (s *DividednameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitDividedname(s)
	}
}

func (p *DoParser) Dividedname() (localctx IDividednameContext) {
	localctx = NewDividednameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, DoParserRULE_dividedname)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(DoParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DoParserT__18 {
		{
			p.SetState(262)
			p.Match(DoParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)
			p.Match(DoParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolContext is an interface to support dynamic dispatch.
type IBoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOL() antlr.TerminalNode

	// IsBoolContext differentiates from other interfaces.
	IsBoolContext()
}

type BoolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolContext() *BoolContext {
	var p = new(BoolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_bool
	return p
}

func InitEmptyBoolContext(p *BoolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_bool
}

func (*BoolContext) IsBoolContext() {}

func NewBoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolContext {
	var p = new(BoolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_bool

	return p
}

func (s *BoolContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolContext) BOOL() antlr.TerminalNode {
	return s.GetToken(DoParserBOOL, 0)
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitBool(s)
	}
}

func (p *DoParser) Bool_() (localctx IBoolContext) {
	localctx = NewBoolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, DoParserRULE_bool)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.Match(DoParserBOOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(DoParserSTRING, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitString(s)
	}
}

func (p *DoParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, DoParserRULE_string)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Match(DoParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DoParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DoParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(DoParserNUMBER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DoListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *DoParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, DoParserRULE_number)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Match(DoParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
