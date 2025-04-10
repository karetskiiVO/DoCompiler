package dolistners

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"

	"github.com/karetskiiVO/DoCompiler/compiler"
	expr "github.com/karetskiiVO/DoCompiler/compiler/expressions"
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
	function, _ := l.program.GetVariable(ctx.NAME().GetText())
	function.Expressions = l.expressions
	l.expressions = nil
}

func (l *DoSourceListner) ExitConstantuse(ctx *parser.ConstantuseContext) {
	var content any

	switch {
	case ctx.Bool_() != nil:
		content = (ctx.Bool_().BOOL().GetText() == "true")
	case ctx.String_() != nil:
		content = ctx.String_().STRING().GetText() // TODO: отрезать лишнее
	case ctx.Number() != nil:
		content, _ = strconv.Atoi(ctx.Number().GetText()) // TODO: хранить не int а numbertype
	}

	l.expressions = append(l.expressions, expr.NewConst(content))
}

func (l *DoSourceListner) ExitVariableuse(ctx *parser.VariableuseContext) {
	variable, err := l.program.GetVariable(ctx.Dividedname().GetText())

	if err != nil {
		line := ctx.Dividedname().AllNAME()[0].GetSymbol().GetLine()
		start := ctx.Dividedname().AllNAME()[0].GetSymbol().GetColumn()
		stream := ctx.Dividedname().AllNAME()[0].GetSymbol().GetInputStream().GetSourceName()

		l.program.AddError(fmt.Errorf("%v:%v:%v: %w", stream, line, start, err))
	}

	l.expressions = append(l.expressions, expr.NewVariable(variable))
}

func (l *DoSourceListner) ExitAssign(ctx *parser.AssignContext) {
	lhvlen, rhvlen := len(ctx.Expressiontuplelhv().Expressiontuple().AllExpression()), len(ctx.Expressiontuplelhv().Expressiontuple().AllExpression())
	
}
