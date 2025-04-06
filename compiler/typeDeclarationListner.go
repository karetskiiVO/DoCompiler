package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"
)

type DoTypeDeclarationListener struct {
	*parser.BaseDoListener

	program  *Program
}

func NewDoTypeDeclarationListener(program *Program) antlr.ParseTreeListener {
	if program.err != nil {
		return new(parser.BaseDoListener)
	}

	return &DoTypeDeclarationListener{
		program: program,
	}
}

func (l *DoTypeDeclarationListener) EnterTypeDefinition(ctx *parser.TypeDefinitionContext) {
	newtype, err := l.program.RegisterType(ctx.NAME().GetText())
	if err != nil {
		line := ctx.NAME().GetSymbol().GetLine()
		start := ctx.NAME().GetSymbol().GetColumn()
		stream := ctx.NAME().GetSymbol().GetInputStream().GetSourceName()

		l.program.AddError(fmt.Errorf("%v:%v:%v: %w", stream, line, start, err))
	}

	// TODO: generics
	_ = newtype
}

// func (l *DoTypeDeclarationListener) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
// 	newfunc, err := l.program.RegisterFunc(ctx.NAME().GetText())
// 	if err != nil {
// 		line := ctx.GetStart().GetLine()
// 		start := ctx.NAME().GetSourceInterval().Start
// 		l.program.AddError(fmt.Errorf("%v:%v: %w", line, start, err))
// 	}
// 	// TODO: generics
// 	_ = newfunc
// }

// TODO: глобальные переменные и константы
