package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/jessevdk/go-flags"
	"github.com/karetskiiVO/DoCompiler/parser"

	"github.com/karetskiiVO/DoCompiler/compiler"
	doListeners "github.com/karetskiiVO/DoCompiler/compiler/listeners"
)

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
	program, err := Compile(options.Args.SourceFileNames...)
	if err != nil {
		fmt.Println(err)
	} else {
		os.WriteFile("program.ll", []byte(program.String()), 0644)
	}
}

func Compile(srcFiles ...string) (*compiler.Program, error) {
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

	// Listeners.Set(func(int) antlr.ParseTreeListener {
	// 	return doListeners.NewDoTypeDefinitionListener(program)
	// })
	// Listeners.Exec()

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
		return nil, err
	}

	return program, nil
}
