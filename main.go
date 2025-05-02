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

	"github.com/karetskiiVO/DoCompiler/compiler"
	dolistners "github.com/karetskiiVO/DoCompiler/compiler/listners"
	compilertypes "github.com/karetskiiVO/DoCompiler/compiler/types"
)

const typeDescriptorString = `type {{ .Name }}:
	.isfunc         = {{ .IsFunction }}
	.isbehavour     = {{ .IsBehavour }}
{{ if .IsBehavour -}}
	.isselfbehavour = {{ .SelfBehavour }}
{{- end }}
`
const variablesDescriptorString = `var {{ .Name }}: .type {{ .VarType }}
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

	program := compiler.NewProgram()

	listners.Set(func(int) antlr.ParseTreeListener {
		return dolistners.NewDoTypeDeclarationListener(program)
	})
	listners.Exec()

	listners.Set(func(int) antlr.ParseTreeListener {
		return dolistners.NewDoTypeDefinitionListener(program)
	})
	listners.Exec()

	listners.Set(func(int) antlr.ParseTreeListener {
		return dolistners.NewDoVariableDeclarationListner(program)
	})
	listners.Exec()

	err := program.Validate()
	if err != nil {
		fmt.Println(err)
	}

	typeinfos := slices.Collect(maps.Values(program.Types()))
	slices.SortFunc(typeinfos, func(fst, snd *compilertypes.DoType) int {
		return strings.Compare(string(fst.String()), string(fst.String()))
	})

	for _, typeinfo := range typeinfos {
		typeDescriptor.Execute(os.Stdout, typeinfo)
	}

	variables := slices.Collect(maps.Values(program.Variables()))
	slices.SortFunc(variables, func(fst, snd *compilertypes.Variable) int {
		return strings.Compare(string(fst.Name), string(snd.Name))
	})

	for _, varinfo := range variables {
		variableDescriptor.Execute(os.Stdout, varinfo)
	}
}
