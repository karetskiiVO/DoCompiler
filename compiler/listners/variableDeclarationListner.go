package dolistners

import (
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

}

func (l *DoVariableDeclarationListner) EnterGlobalVariableDefinition(ctx *parser.GlobalVariableDefinitionContext) {

}
