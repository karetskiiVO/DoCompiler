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
	values [][]value.Value
}

func NewDoSourceListener(program *compiler.Program) antlr.ParseTreeListener {
	if program.Error() != nil {
		return new(parser.BaseDoListener)
	}

	return &DoSourceListener{
		program: program,
	}
}

func (l *DoSourceListener) pushBlock(block *ir.Block) { l.blocks = append(l.blocks, block) }
func (l *DoSourceListener) popBlock()                 { l.blocks = l.blocks[:len(l.blocks)-1] }
func (l *DoSourceListener) topBlock() *ir.Block       { return l.blocks[len(l.blocks)-1] }
func (l *DoSourceListener) addValue(val ...value.Value) {
	l.values[len(l.values)-1] = append(l.values[len(l.values)-1], val...)
}
func (l *DoSourceListener) newBlock() *ir.Block {
	return l.currfunc.NewBlock("")
}

func (l *DoSourceListener) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	funcname := ctx.NAME().GetText() // Объвить нужные переменные

	function, _ := l.program.GetFunction(funcname)

	l.currfunc = function
	entry := l.newBlock()
	l.pushBlock(entry)
}

func (l *DoSourceListener) ExitFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	rettype := l.currfunc.Sig.RetType.(*types.StructType)
	if len(rettype.Fields) == 0 {
		retval := l.topBlock().NewAlloca(l.currfunc.Sig.RetType)
		retval.SetName("#ret")
		l.topBlock().NewRet(retval)
	} else {
		l.topBlock().NewUnreachable()
	}

	l.popBlock()
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

			l.program.AddErrorf("%v:%v:%v: %w", stream, line, start, err)
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
	values := l.values[len(l.values)-1]
	if lhvLen != len(values) {
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
			l.topBlock().NewStore(values[i], variable)
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

func (l *DoSourceListener) EnterStatement(ctx *parser.StatementContext) {
	l.values = append(l.values, []value.Value{})
}
func (l *DoSourceListener) ExitStatement(ctx *parser.StatementContext) {
	l.values = l.values[:len(l.values)-1]
}

func (l *DoSourceListener) EnterIfstatement(ctx *parser.IfstatementContext) {
	trueBlock := l.newBlock()
	falseBlock := l.newBlock()

	l.pushBlock(falseBlock)
	l.pushBlock(trueBlock)
}

func (l *DoSourceListener) ExitIfstatement(ctx *parser.IfstatementContext) {
	baseBlock := l.blocks[len(l.blocks)-3]
	falseBlock := l.blocks[len(l.blocks)-2]
	trueBlock := l.blocks[len(l.blocks)-1]

	values := l.values[len(l.values)-1]

	if len(values) != 1 {
		line := ctx.Expression().GetStart().GetLine()
		start := ctx.Expression().GetStart().GetColumn()
		stream := ctx.Expression().GetStart().GetInputStream().GetSourceName()

		l.program.AddErrorf("%v:%v:%v: if statement expected 1 bool, actual: %v", stream, line, start, len(values))
		return
	}

	// TODO: проверка на буль values[0]

	l.popBlock()
	l.popBlock()
	l.popBlock()

	resBlock := falseBlock
	if ctx.Elsestatement() != nil {
		resBlock = l.newBlock()
		falseBlock.NewBr(resBlock)
	}

	baseBlock.NewCondBr(values[0], trueBlock, falseBlock)
	trueBlock.NewBr(resBlock)

	l.pushBlock(resBlock)
}

func (l *DoSourceListener) EnterElsestatement(ctx *parser.ElsestatementContext) {
	idx1, idx2 := len(l.blocks)-1, len(l.blocks)-2
	l.blocks[idx1], l.blocks[idx2] = l.blocks[idx2], l.blocks[idx1]
	l.values = append(l.values, []value.Value{})
}

func (l *DoSourceListener) ExitElsestatement(ctx *parser.ElsestatementContext) {
	idx1, idx2 := len(l.blocks)-1, len(l.blocks)-2
	l.blocks[idx1], l.blocks[idx2] = l.blocks[idx2], l.blocks[idx1]
	l.values = l.values[:len(l.values)-1]
}

func (l *DoSourceListener) ExitReturnstatement(ctx *parser.ReturnstatementContext) {
	values := l.values[len(l.values)-1]

	rettype := l.currfunc.Sig.RetType.(*types.StructType)

	if len(rettype.Fields) != len(values) {
		line := ctx.GetStart().GetLine()
		start := ctx.GetStart().GetColumn()
		stream := ctx.GetStart().GetInputStream().GetSourceName()

		l.program.AddErrorf(
			"%v:%v:%v: incorrect number of expressions in the return expected: %v actual: %v",
			stream,
			line,
			start,
			len(rettype.Fields),
			len(values),
		)
		return
	}

	retval := l.topBlock().NewAlloca(l.currfunc.Sig.RetType)
	for i, field := range rettype.Fields {

		l.topBlock().NewStore(
			values[i],
			l.topBlock().NewGetElementPtr(field, retval, constant.NewIndex(constant.NewInt(types.I64, int64(i)))),
		)
	}
	retval.SetName("#ret")

	l.topBlock().NewRet(retval)
	l.popBlock()
	l.pushBlock(l.newBlock())
}
