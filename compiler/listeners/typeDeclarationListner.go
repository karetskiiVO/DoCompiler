package doListeners

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"
	"github.com/llir/llvm/ir/types"

	"github.com/karetskiiVO/DoCompiler/compiler"
)

type DoTypeDeclarationListener struct {
	*parser.BaseDoListener

	program *compiler.Program
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
	typename := ctx.NAME().GetText()

	var definition []compiler.TypeDefiner
	var fieldnames []string
	if structdefCtx := ctx.Type_().Structdefinition(); structdefCtx != nil {
		var fieldtypes []types.Type
		for _, fieldToken := range structdefCtx.AllVarfield() {
			fieldtypename := fieldToken.Type_().GetText()
			fieldname := fieldToken.Fieldname().GetText()
			fieldtype, err := l.program.GetType(fieldtypename)

			if err != nil {
				line := fieldToken.Type_().GetStart().GetLine()
				start := fieldToken.Type_().GetStart().GetColumn()
				stream := fieldToken.Type_().GetStart().GetInputStream().GetSourceName()

				l.program.AddError(fmt.Errorf("%v:%v:%v: %w", stream, line, start, err))
				return
			}

			fieldtypes = append(fieldtypes, fieldtype)
			fieldnames = append(fieldnames, fieldname)
		}

		definition = append(definition,
			compiler.TypeDefiner{
				Typ: types.NewStruct(fieldtypes...),
				Fieldnames: fieldnames,
			},
		)

	}

	_, err := l.program.RegisterType(typename, definition...)
	if err != nil {
		line := ctx.NAME().GetSymbol().GetLine()
		start := ctx.NAME().GetSymbol().GetColumn()
		stream := ctx.NAME().GetSymbol().GetInputStream().GetSourceName()

		l.program.AddError(fmt.Errorf("%v:%v:%v: %w", stream, line, start, err))
		return
	}

}
