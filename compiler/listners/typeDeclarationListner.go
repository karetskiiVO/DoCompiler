package dolistners

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"

	"github.com/karetskiiVO/DoCompiler/compiler"
)

type DoTypeDeclarationListener struct {
	*parser.BaseDoListener

	program  *compiler.Program
}

func NewDoTypeDeclarationListener(program *compiler.Program) antlr.ParseTreeListener {
	if program.Error() != nil {
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
	// TODO: парсить тип
	// TODO: generics
	_ = newtype
}
