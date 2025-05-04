package doListeners

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"
	"github.com/llir/llvm/ir"

	"github.com/karetskiiVO/DoCompiler/compiler"
)

type DoSourceListener struct {
	*parser.BaseDoListener

	function *ir.Func
	program  *compiler.Program
}

func NewDoSourceListener(program *compiler.Program) antlr.ParseTreeListener {
	if program.Error() != nil {
		return new(parser.BaseDoListener)
	}

	return &DoSourceListener{
		program: program,
	}
}

func (l *DoSourceListener) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	funcname := ctx.NAME().GetText() // Объвить нужные переменные

	function, err := l.program.GetFunction(funcname)
	if err != nil {
		line := ctx.NAME().GetSymbol().GetLine()
		start := ctx.NAME().GetSymbol().GetColumn()
		stream := ctx.NAME().GetSymbol().GetInputStream().GetSourceName()

		l.program.AddError(fmt.Errorf("%v:%v:%v: %w", stream, line, start, err))
		return
	}
	_ = function

	// entry := llvm.AddBasicBlock(*function.LLVMFunction, funcname+"#entry")
	// l.program.Builder().SetInsertPoint(entry, *function.LLVMFunction)
}

func (l *DoSourceListener) ExitFunctionDefinition(ctx *parser.FunctionDefinitionContext) {

	l.function = nil
}

func (l *DoSourceListener) ExitFunctioncall(ctx *parser.FunctioncallContext) {
	funcname := ctx.Dividedname().GetText()
	function, err := l.program.GetFunction(funcname)
	if err != nil {
		line := ctx.Dividedname().GetStart().GetLine()
		start := ctx.Dividedname().GetStart().GetColumn()
		stream := ctx.Dividedname().GetStart().GetInputStream().GetSourceName()

		l.program.AddError(fmt.Errorf("%v:%v:%v: %w", stream, line, start, err))
		return
	}
}

func (l *DoSourceListener) ExitConstantuse(ctx *parser.ConstantuseContext) {
}

func (l *DoSourceListener) ExitVariableuse(ctx *parser.VariableuseContext) {
}

func (l *DoSourceListener) ExitAssign(ctx *parser.AssignContext) {
}

func (l *DoSourceListener) ExitEmptyexpression(*parser.EmptyexpressionContext) {
}
