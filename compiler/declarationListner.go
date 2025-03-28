package main

import (
	"github.com/karetskiiVO/DoCompiler/parser"
)

type DoDefinitionListener struct {
	*parser.BaseDoListener

	functions map[string]Function
}

func NewGoDeclarationListener() *DoDefinitionListener {
	return &DoDefinitionListener{
		functions: make(map[string]Function),
	}
}

func (l *DoDefinitionListener) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	funcname := ctx.NAME()
	_ = funcname
}
