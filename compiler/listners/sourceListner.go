package dolistners

import (
	"fmt"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	compilertypes "github.com/karetskiiVO/DoCompiler/compiler/types"
	"github.com/karetskiiVO/DoCompiler/parser"
	"tinygo.org/x/go-llvm"

	"github.com/karetskiiVO/DoCompiler/compiler"
)

type DoSourceListner struct {
	*parser.BaseDoListener

	function *compilertypes.Function
	program  *compiler.Program
}

func NewDoSourceListner(program *compiler.Program) antlr.ParseTreeListener {
	if program.Error() != nil {
		return new(parser.BaseDoListener)
	}

	return &DoSourceListner{
		program: program,
	}
}

func (l *DoSourceListner) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	funcname := ctx.NAME().GetText() // Объвить нужные переменные

	function, err := l.program.GetFunction(funcname)
	if err != nil {
		line := ctx.NAME().GetSymbol().GetLine()
		start := ctx.NAME().GetSymbol().GetColumn()
		stream := ctx.NAME().GetSymbol().GetInputStream().GetSourceName()

		l.program.AddError(fmt.Errorf("%v:%v:%v: %w", stream, line, start, err))
	}

	entry := llvm.AddBasicBlock(*function.LLVMFunction, "entry")
	l.program.Builder().SetInsertPoint(entry, *function.LLVMFunction)
}

func (l *DoSourceListner) ExitFunctionDefinition(ctx *parser.GlobalVariableDefinitionContext) {

	l.function = nil
}

func (l *DoSourceListner) ExitFunctioncall(ctx *parser.FunctioncallContext) {
	funcname := ctx.Dividedname().GetText()
	function, err := l.program.GetFunction(funcname)
	if err != nil {
		line := ctx.Dividedname().GetStart().GetLine()
		start := ctx.Dividedname().GetStart().GetColumn()
		stream := ctx.Dividedname().GetStart().GetInputStream().GetSourceName()

		l.program.AddError(fmt.Errorf("%v:%v:%v: %w", stream, line, start, err))
	}

	fmt.Println(funcname)

	l.program.Builder().CreateCall(*function.FunctionType, *function.LLVMFunction, []llvm.Value{}, funcname)
}

func (l *DoSourceListner) ExitConstantuse(ctx *parser.ConstantuseContext) {
}

func (l *DoSourceListner) ExitVariableuse(ctx *parser.VariableuseContext) {
}

func (l *DoSourceListner) ExitAssign(ctx *parser.AssignContext) {
}

func (l *DoSourceListner) ExitEmptyexpression(*parser.EmptyexpressionContext) {
}

func (l *DoSourceListner) ExitEveryRule (ctx antlr.ParserRuleContext) {
	fmt.Println(reflect.TypeOf(ctx))
}
