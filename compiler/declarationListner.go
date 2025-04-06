package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"
)

type DoDeclarationListener struct {
	*parser.BaseDoListener

	program *Program
}

func NewGoDeclarationListener(program *Program) antlr.ParseTreeListener {
	if program.err != nil {
		return new(parser.BaseDoListener)
	}

	return &DoDeclarationListener{
		program: program,
	}
}

func (l *DoDeclarationListener) EnterTypeDefinition(ctx *parser.TypeDefinitionContext) {
	newtype, err := l.program.RegisterType(ctx.NAME().GetText())
	if err != nil {
		line := ctx.GetStart().GetLine()
		start := ctx.NAME().GetSourceInterval().Start

		l.program.AddError(fmt.Errorf("%v:%v: %w", line, start, err))
	}

	// TODO: generics
	_ = newtype
}

// func (l *DoDeclarationListener) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
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
