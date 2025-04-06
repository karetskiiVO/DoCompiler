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
			SourceFileNames []string
		} `positional-args:"yes" required:"1"`
	}

	flagsParser := flags.NewParser(&options, flags.Default&(^flags.PrintErrors))
	_, err := flagsParser.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	listners := &Listners{
		listners: make([]antlr.ParseTreeListener, len(options.Args.SourceFileNames)),
		parsers:  make([]*parser.DoParser, 0, len(options.Args.SourceFileNames)),
	}
	for _, fileName := range options.Args.SourceFileNames {
		input, err := antlr.NewFileStream(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		lexer := parser.NewDoLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		listners.parsers = append(listners.parsers, parser.NewDoParser(stream))
	}

	program := NewProgram()

	listners.Set(func() antlr.ParseTreeListener { return NewGoDeclarationListener(program) })
	listners.Exec()

	listners.Set(func() antlr.ParseTreeListener { return NewDoDefinitionListener(program) })
	listners.Exec()

	err = program.Validate()
	if err != nil {
		fmt.Println(err)
	}
}

type Listners struct {
	listners []antlr.ParseTreeListener
	parsers  []*parser.DoParser
}

func (l *Listners) Set(newListner func() antlr.ParseTreeListener) {
	for i := range l.listners {
		l.listners[i] = newListner()
	}
}

func (l *Listners) Exec() {
	for i := range l.listners {
		antlr.ParseTreeWalkerDefault.Walk(l.listners[i], l.parsers[i].Program())
	}
}
