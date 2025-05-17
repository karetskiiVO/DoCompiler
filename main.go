package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/jessevdk/go-flags"
	"github.com/karetskiiVO/DoCompiler/parser"

	"tinygo.org/x/go-llvm"

	"github.com/karetskiiVO/DoCompiler/compiler"
	doListeners "github.com/karetskiiVO/DoCompiler/compiler/listeners"
)

func main() {
	var options struct {
		Args struct {
			SourceFileNames []string
		} `positional-args:"yes" required:"1"`
		OutputFile string `short:"o" long:"output" description:"Output file" default:"a.bc"`
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
		os.Exit(1)
	}

	llvm.InitializeAllTargets()
	llvm.InitializeAllTargetMCs()
	llvm.InitializeAllAsmPrinters()
	llvm.InitializeAllAsmParsers()

	tempdir := os.TempDir()
	tempFile := filepath.Join(tempdir, "program.ll")
	err = os.WriteFile(tempFile, []byte(program.String()), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buf, err := llvm.NewMemoryBufferFromFile(tempFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//defer buf.Dispose()
	llvmctx := llvm.NewContext()
	//defer ctx.Dispose()
	mod, err := llvmctx.ParseIR(buf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out, err := os.CreateTemp(tempdir, "program*.bc")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer out.Close()
	llvm.WriteBitcodeToFile(mod, out)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	command := exec.CommandContext(
		ctx,
		"clang", "-o",
		options.OutputFile,
		out.Name(),
		"./std/tmpprint.c",
	)

	err = command.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Compile(srcFiles ...string) (*compiler.Program, error) {
	Listeners := &Listeners{
		Listeners: make([]antlr.ParseTreeListener, len(srcFiles)),
		roots:     make([]parser.IProgramContext, 0, len(srcFiles)),
	}
	for _, fileName := range srcFiles {
		input, err := antlr.NewFileStream(fileName)
		// file, err := os.Open(fileName)
		// input := antlr.NewIoStream(file)

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
