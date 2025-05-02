package dolistners

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"

	"github.com/karetskiiVO/DoCompiler/compiler"
	compilertypes "github.com/karetskiiVO/DoCompiler/compiler/types"
)

type DoSourceListner struct {
	*parser.BaseDoListener

	program     *compiler.Program
	expressions []compilertypes.Expression
}

func NewDoSourceListner(program *compiler.Program) antlr.ParseTreeListener {
	if program.Error() != nil {
		return new(parser.BaseDoListener)
	}

	return &DoSourceListner{
		program:     program,
		expressions: nil,
	}
}

func (l *DoSourceListner) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	l.expressions = make([]compilertypes.Expression, 0, len(ctx.AllStatement()))
}

func (l *DoSourceListner) ExitFunctionDefinition(ctx *parser.GlobalVariableDefinitionContext) {

}

func (l *DoSourceListner) ExitConstantuse(ctx *parser.ConstantuseContext) {
}

func (l *DoSourceListner) ExitVariableuse(ctx *parser.VariableuseContext) {
}

func (l *DoSourceListner) ExitAssign(ctx *parser.AssignContext) {
}

func (l *DoSourceListner) ExitEmptyexpression(*parser.EmptyexpressionContext) {
}
