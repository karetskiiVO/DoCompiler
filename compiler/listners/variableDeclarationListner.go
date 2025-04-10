package dolistners

import (
	"fmt"
	"slices"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"

	
	"github.com/karetskiiVO/DoCompiler/compiler"
)

type DoVariableDeclarationListner struct {
	*parser.BaseDoListener

	program *compiler.Program
}

func NewDoVariableDeclarationListner(program *compiler.Program) antlr.ParseTreeListener {
	if program.Error() != nil {
		return new(parser.BaseDoListener)
	}

	return &DoVariableDeclarationListner{
		program: program,
	}
}

func (l *DoVariableDeclarationListner) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	// TODO: генерировать массив типов вместо этого
	argtypes := make([]string, 0)
	rettypes := make([]string, 0)

	// TODO: роаспарсить тип 
	for _, argsublist := range ctx.Arglist().AllArgsublist() {
		argtypes = append(
			argtypes,
			slices.Repeat([]string{argsublist.Type_().GetText()}, len(argsublist.AllArgname()))...,
		)
	}
	for _, rettypename := range ctx.Typetuple().AllType_() {
		rettypes = append(rettypes, rettypename.GetText())
	}

	functype := fmt.Sprintf("act(%v)(%v)", strings.Join(argtypes, ","), strings.Join(rettypes, ","))

	newfunc, err := l.program.RegisterGlobalVariable(ctx.NAME().GetText(), functype)
	if err != nil {
		line := ctx.NAME().GetSymbol().GetLine()
		start := ctx.NAME().GetSymbol().GetColumn()
		stream := ctx.NAME().GetSymbol().GetInputStream().GetSourceName()

		l.program.AddError(fmt.Errorf("%v:%v:%v: %w", stream, line, start, err))
	}

	// TODO: generics
	_ = newfunc
}

func (l *DoVariableDeclarationListner) EnterGlobalVariableDefinition(ctx *parser.GlobalVariableDefinitionContext) {
	newvar, err := l.program.RegisterGlobalVariable(ctx.NAME().GetText(), ctx.Typename().GetText())
	if err != nil {
		line := ctx.NAME().GetSymbol().GetLine()
		start := ctx.NAME().GetSymbol().GetColumn()
		stream := ctx.NAME().GetSymbol().GetInputStream().GetSourceName()

		l.program.AddError(fmt.Errorf("%v:%v:%v: %w", stream, line, start, err))
	}

	// TODO: generics
	_ = newvar
}
