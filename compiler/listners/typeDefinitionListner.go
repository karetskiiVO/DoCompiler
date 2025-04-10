package dolistners

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"

	"github.com/karetskiiVO/DoCompiler/compiler"
)

type DoDefenitionListener struct {
	*parser.BaseDoListener

	program  *compiler.Program
}

func NewDoTypeDefinitionListener(program *compiler.Program) antlr.ParseTreeListener {
	if program.Error() != nil {
		return new(parser.BaseDoListener)
	}

	return &DoDefenitionListener{
		program:  program,
	}
}
