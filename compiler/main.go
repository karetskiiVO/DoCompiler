package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/jessevdk/go-flags"
	"github.com/karetskiiVO/DoCompiler/parser"
)

func main() {
	var options struct {
		Args struct {
			SourceFileName string
		} `positional-args:"yes" required:"1"`
	}

	flagsParser := flags.NewParser(&options, flags.Default&(^flags.PrintErrors))
	_, err := flagsParser.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input, err := antlr.NewFileStream(options.Args.SourceFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lexer := parser.NewDoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewDoParser(stream)

	tree := parser.Program() // Начинаем с корневого узла
	program := NewProgram()

	antlr.ParseTreeWalkerDefault.Walk(NewGoDeclarationListener(program), tree)



	if program.err != nil {
		fmt.Println(err)
	}
}
