package doListeners

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"github.com/karetskiiVO/DoCompiler/compiler"
)

type DoSourceListener struct {
	*parser.BaseDoListener

	currfunc *ir.Func
	program  *compiler.Program

	blocks []*ir.Block
	values []value.Value
}

func NewDoSourceListener(program *compiler.Program) antlr.ParseTreeListener {
	if program.Error() != nil {
		return new(parser.BaseDoListener)
	}

	return &DoSourceListener{
		program: program,
	}
}

func (l *DoSourceListener) pushBlock(block *ir.Block)   { l.blocks = append(l.blocks, block) }
func (l *DoSourceListener) popBlock()                   { l.blocks = l.blocks[:len(l.blocks)-1] }
func (l *DoSourceListener) topBlock() *ir.Block         { return l.blocks[len(l.blocks)-1] }
func (l *DoSourceListener) addValue(val ...value.Value) { l.values = append(l.values, val...) }

func (l *DoSourceListener) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	funcname := ctx.NAME().GetText() // Объвить нужные переменные

	function, _ := l.program.GetFunction(funcname)

	entry := function.NewBlock("")
	l.pushBlock(entry)
	l.currfunc = function
}

func (l *DoSourceListener) ExitFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	retval := l.topBlock().NewAlloca(l.currfunc.Sig.RetType)
	retval.SetName("#ret")

	l.topBlock().NewRet(retval)
}

func (l *DoSourceListener) ExitFunctioncall(ctx *parser.FunctioncallContext) {
	funcname := ctx.Dividedname().GetText()
	function, err := l.program.GetFunction(funcname)
	if err != nil {
		line := ctx.Dividedname().GetStart().GetLine()
		start := ctx.Dividedname().GetStart().GetColumn()
		stream := ctx.Dividedname().GetStart().GetInputStream().GetSourceName()

		l.program.AddErrorf("%v:%v:%v: %w", stream, line, start, err)
		return
	}

	call := l.topBlock().NewCall(function)

	// unpack return
	for i := range function.Sig.RetType.(*types.StructType).Fields {
		l.addValue(
			l.topBlock().NewExtractValue(call, uint64(i)),
		)
		// не работает
		//l.topBlock().NewGetElementPtr(types.NewPointer(typ), call, constant.NewInt(types.I64, int64(i)))
	}
}

func (l *DoSourceListener) ExitConstantuse(ctx *parser.ConstantuseContext) {
	switch {
	case ctx.Bool_() != nil:
		typ, _ := l.program.GetType("bool")
		val := int64(0)
		if ctx.Bool_().GetText() == "true" {
			val = 1
		}

		l.addValue(
			constant.NewInt(typ.(*types.IntType), val),
		)
	case ctx.Number() != nil:
		val, err := strconv.ParseInt(ctx.Number().GetText(), 0, 64)
		if err != nil { // unricheable
			line := ctx.GetStart().GetLine()
			start := ctx.GetStart().GetColumn()
			stream := ctx.GetStart().GetInputStream().GetSourceName()

			l.program.AddError(fmt.Errorf("%v:%v:%v: %w", stream, line, start, err))
			return
		}

		typ, _ := l.program.GetType("int")
		l.addValue(
			constant.NewInt(typ.(*types.IntType), int64(val)),
		)
	default:
		panic(fmt.Errorf("not implemented: %v", ctx.GetText()))
	}
}

func (l *DoSourceListener) ExitVariableuse(ctx *parser.VariableuseContext) {
}

func (l *DoSourceListener) ExitAssign(ctx *parser.AssignContext) {
	if ctx.Expressiontuplelhv() == nil {
		return
	}

	lhvLen := len(ctx.Expressiontuplelhv().Expressiontuple().AllExpression())
	if lhvLen != len(l.values) {
		line := ctx.Expressiontuplelhv().GetStop().GetLine()
		start := ctx.Expressiontuplelhv().GetStop().GetColumn()
		stream := ctx.Expressiontuplelhv().GetStop().GetInputStream().GetSourceName()

		l.program.AddErrorf(
			"%v:%v:%v: incorrect number of expressions in the assignment expected: %v actual: %v",
			stream,
			line,
			start,
			len(l.values),
			lhvLen,
		)
		return
	}

	for i, lhvi := range ctx.Expressiontuplelhv().Expressiontuple().AllExpression() {
		switch {
		case lhvi.Variableuse() != nil:
			varname := lhvi.GetText()

			// TODO: тут должна быть таблица переменных
			variable, err := l.program.GetVariable(varname)
			if err != nil {
				line := ctx.Expressiontuplelhv().Expressiontuple().AllExpression()[i].GetStart().GetLine()
				start := ctx.Expressiontuplelhv().Expressiontuple().AllExpression()[i].GetStart().GetColumn()
				stream := ctx.Expressiontuplelhv().Expressiontuple().AllExpression()[i].GetStart().GetInputStream().GetSourceName()

				l.program.AddErrorf("%v:%v:%v: %w", stream, line, start, err)
				continue
			}

			// TODO: более умная проверка
			l.topBlock().NewStore(l.values[i], variable)
		case lhvi.Emptyexpression() != nil:
		default:
			line := ctx.Expressiontuplelhv().Expressiontuple().AllExpression()[i].GetStart().GetLine()
			start := ctx.Expressiontuplelhv().Expressiontuple().AllExpression()[i].GetStart().GetColumn()
			stream := ctx.Expressiontuplelhv().Expressiontuple().AllExpression()[i].GetStart().GetInputStream().GetSourceName()

			l.program.AddErrorf(
				"%v:%v:%v: expression `%v` is not mutable",
				stream,
				line,
				start,
				ctx.Expressiontuplelhv().Expressiontuple().AllExpression()[i].GetText(),
			)
		}
	}
}

func (l *DoSourceListener) ExitStatement(ctx *parser.StatementContext) {
	l.values = l.values[:0]
}

func (l *DoSourceListener) ExitEmptyexpression(*parser.EmptyexpressionContext) {
}
