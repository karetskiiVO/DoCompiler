package dolistners

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"
	"github.com/karetskiiVO/slices"

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
	funcname := ctx.NAME().GetText()

	argTypenames := make([]string, 0)
	slices.Map(ctx.Arglist().AllArgsublist(), func(arglistctx parser.IArgsublistContext) struct{} {
		// TODO: сделать не только typename - сейчас очень коряво
		typename := arglistctx.Type_().Typename().Dividedname().GetText()

		argTypenames = append(argTypenames,
			slices.Repeat(
				[]string{typename},
				len(arglistctx.AllArgname()),
			)...,
		)

		return struct{}{}
	})
	retTypenames := slices.Map(ctx.Typetuple().AllType_(), func(typectx parser.ITypeContext) string {
		return typectx.GetText()
	})

	_, err := l.program.RegisterFunction(funcname, argTypenames, retTypenames)
	if err != nil {
		line := ctx.NAME().GetSymbol().GetLine()
		start := ctx.NAME().GetSymbol().GetColumn()
		stream := ctx.NAME().GetSymbol().GetInputStream().GetSourceName()

		l.program.AddError(fmt.Errorf("%v:%v:%v: %w", stream, line, start, err))
	}
}

func (l *DoVariableDeclarationListner) EnterGlobalVariableDefinition(ctx *parser.GlobalVariableDefinitionContext) {

}
