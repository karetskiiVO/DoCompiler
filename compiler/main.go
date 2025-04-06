package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
	"text/template"

	"github.com/antlr4-go/antlr/v4"
	"github.com/jessevdk/go-flags"
	"github.com/karetskiiVO/DoCompiler/parser"
)

const typeDescriptorString = `type {{ .Name }}:
	.isfunc         = {{ .IsFunction }}
	.isbehavour     = {{ .IsBehavour }}
{{ if .IsBehavour }}
	.isselfbehavour = {{ .SelfBehavour }}
{{ end }}
`
const variablesDescriptorString = `var {{ .Name }}: .type {{ .VarType.Name }}
`

var typeDescriptor = template.Must(template.New("typedescriptor").Parse(typeDescriptorString))
var variableDescriptor = template.Must(template.New("variabledescriptor").Parse(variablesDescriptorString))

func main() {
	var options struct {
		Args struct {
			SourceFileNames []string
		} `positional-args:"yes" required:"1"`
	}

	flagsParser := flags.NewParser(&options, flags.Default&(^flags.PrintErrors))
	_, err := flagsParser.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	/********** Компиляция **********/
	Compile(options.Args.SourceFileNames...)
}

func Compile(srcFiles ...string) {
	listners := &Listners{
		listners: make([]antlr.ParseTreeListener, len(srcFiles)),
		roots:    make([]parser.IProgramContext, 0, len(srcFiles)),
	}
	for _, fileName := range srcFiles {
		input, err := antlr.NewFileStream(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		lexer := parser.NewDoLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		listners.roots = append(listners.roots, parser.NewDoParser(stream).Program())
	}

	program := NewProgram()

	listners.Set(func(int) antlr.ParseTreeListener {
		return NewDoTypeDeclarationListener(program)
	})
	listners.Exec()

	listners.Set(func(int) antlr.ParseTreeListener {
		return NewDoTypeDefinitionListener(program)
	})
	listners.Exec()

	listners.Set(func(int) antlr.ParseTreeListener {
		return NewDoVariableDeclarationListner(program)
	})
	listners.Exec()

	err := program.Validate()
	if err != nil {
		fmt.Println(err)
	}

	typeinfos := slices.Collect(maps.Values(program.types))
	slices.SortFunc(typeinfos, func(fst, snd *Type) int {
		return strings.Compare(string(fst.Name), string(snd.Name))
	})

	for _, typeinfo := range typeinfos {
		typeDescriptor.Execute(os.Stdout, typeinfo)
	}

	variables := slices.Collect(maps.Values(program.variables))
	slices.SortFunc(variables, func(fst, snd *Variable) int {
		return strings.Compare(string(fst.Name), string(snd.Name))
	})

	for _, varinfo := range variables {
		variableDescriptor.Execute(os.Stdout, varinfo)
	}
}
