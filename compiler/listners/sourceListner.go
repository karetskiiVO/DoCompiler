package dolistners

import (
	"fmt"
	"slices"
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
	contentTypename := ""

	switch {
	case ctx.Bool_() != nil:
		content = (ctx.Bool_().BOOL().GetText() == "true")
		contentTypename = "bool"
	case ctx.String_() != nil:
		content = ctx.String_().STRING().GetText() // TODO: отрезать лишнее
		contentTypename = "string"
	case ctx.Number() != nil:
		content, _ = strconv.Atoi(ctx.Number().GetText()) // TODO: хранить не int а numbertype
		contentTypename = "int"
	}

	Type, err := l.program.GetType(contentTypename)
	if err != nil { // TODO: не возможно в проде
		panic(err)
	}
	l.expressions = append(l.expressions, expr.NewConst(content, Type))
}

func (l *DoSourceListner) ExitVariableuse(ctx *parser.VariableuseContext) {
	variable, err := l.program.GetVariable(ctx.Dividedname().GetText())

	if err != nil {
		line := ctx.Dividedname().GetStart().GetLine()
		start := ctx.Dividedname().GetStart().GetColumn()
		stream := ctx.Dividedname().GetStart().GetInputStream().GetSourceName()

		l.program.AddError(fmt.Errorf("%v:%v:%v: %w", stream, line, start, err))
	}

	l.expressions = append(l.expressions, expr.NewVariable(variable))
}

func (l *DoSourceListner) ExitAssign(ctx *parser.AssignContext) {
	lhvLen := len(ctx.Expressiontuplelhv().Expressiontuple().AllExpression())
	rhvLen := len(ctx.Expressiontuplelhv().Expressiontuple().AllExpression())

	lhvStart := len(l.expressions) - rhvLen - lhvLen
	rhvStart := len(l.expressions) - rhvLen

	lhvExpressions := slices.Clone(l.expressions[lhvStart:rhvStart])
	rhvExpressions := slices.Clone(l.expressions[rhvStart:])
	l.expressions = l.expressions[:lhvStart]

	for i, lhvExpression := range lhvExpressions {
		if !lhvExpression.IsLHV() {
			astexpr := ctx.Expressiontuplelhv().Expressiontuple().AllExpression()[i]

			line := astexpr.GetStart().GetLine()
			start := astexpr.GetStart().GetColumn()
			stream := astexpr.GetStart().GetInputStream().GetSourceName()

			err := fmt.Errorf("`%v` is not rhv expresion", astexpr.GetText())
			l.program.AddError(fmt.Errorf("%v:%v:%v: %w", stream, line, start, err))
		}
	}

	for i, rhvExpression := range rhvExpressions {
		if _, ok := rhvExpression.(*expr.Empty); ok {
			astexpr := ctx.Expressiontuplelhv().Expressiontuple().AllExpression()[i]

			line := astexpr.GetStart().GetLine()
			start := astexpr.GetStart().GetColumn()
			stream := astexpr.GetStart().GetInputStream().GetSourceName()

			err := fmt.Errorf("`%v` empty expression is not allowed here", astexpr.GetText())
			l.program.AddError(fmt.Errorf("%v:%v:%v: %w", stream, line, start, err))
		}
	}

	// TODO: type assertion & comparetypes

	l.expressions = append(l.expressions, expr.NewAssign(lhvExpressions, rhvExpressions))
}

func (l *DoSourceListner) ExitEmptyexpression(*parser.EmptyexpressionContext) {
	l.expressions = append(l.expressions, expr.NewEmpty())
}
