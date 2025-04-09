package dolistners

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"
	
	. "github.com/karetskiiVO/DoCompiler/compiler"
	"github.com/karetskiiVO/DoCompiler/compiler/expressions"
)

type DoSourceListner struct {
	*parser.BaseDoListener
	
	program *Program
	expression []expr.Expression
}

func NewDoSourceListner(program *Program) antlr.ParseTreeListener {
	if program.Error() != nil {
		return new(parser.BaseDoListener)
	}

	return &DoSourceListner{
		program:    program,
		expression: make([]expr.Expression, 0),
	}
}

func (l *DoSourceListner) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {

}

func (l *DoSourceListner) ExitFunctionDefinition(ctx *parser.GlobalVariableDefinitionContext) {

}
