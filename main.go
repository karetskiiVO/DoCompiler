package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/antlr4-go/antlr/v4"
	"github.com/jessevdk/go-flags"
	"github.com/karetskiiVO/DoCompiler/parser"
	"github.com/karetskiiVO/slices"

	"github.com/karetskiiVO/DoCompiler/compiler"
	doListeners "github.com/karetskiiVO/DoCompiler/compiler/Listeners"
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
	Listeners := &Listeners{
		Listeners: make([]antlr.ParseTreeListener, len(srcFiles)),
		roots:     make([]parser.IProgramContext, 0, len(srcFiles)),
	}
	for _, fileName := range srcFiles {
		input, err := antlr.NewFileStream(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		lexer := parser.NewDoLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		Listeners.roots = append(Listeners.roots, parser.NewDoParser(stream).Program())
	}

	program := compiler.NewProgram()

	Listeners.Set(func(int) antlr.ParseTreeListener {
		return doListeners.NewDoTypeDeclarationListener(program)
	})
	Listeners.Exec()

	Listeners.Set(func(int) antlr.ParseTreeListener {
		return doListeners.NewDoTypeDefinitionListener(program)
	})
	Listeners.Exec()

	Listeners.Set(func(int) antlr.ParseTreeListener {
		return doListeners.NewDoVariableDeclarationListener(program)
	})
	Listeners.Exec()

	Listeners.Set(func(int) antlr.ParseTreeListener {
		return doListeners.NewDoSourceListener(program)
	})
	Listeners.Exec()

	err := program.Validate()
	if err != nil {
		fmt.Println(err)
	}

	typeinfos := slices.CollectMap(program.Types())
	slices.Map(typeinfos, func(pair slices.Pair[string, *compilertypes.Type]) struct{} {
		fmt.Printf("%v:\n", pair.First)
		fmt.Printf("\tllvm: %p\n", pair.Second)

		return struct{}{}
	})

	os.WriteFile("program.ll", []byte(program.Module().String()), 0644)

	/*
		slices.SortFunc(typeinfos, func(fst, snd *compilertypes.Type) int {
			return strings.Compare(string(fst.String()), string(snd.String()))
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
	*/
}
