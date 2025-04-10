package example

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"strings"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func Run() {
	// input, _ := os.Open(os.Args[1])
	// defer input.Close()
	// buff, _ := io.ReadAll(input)

	// Parse the source files.	
	file := ".\\example\\test\\test1.txt"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		fmt.Print(err) // parse error
		return
	}
	files := []*ast.File{f}

	// Create the type-checker's package.
	pkg := types.NewPackage("hello", "")

	// Type-check the package, load dependencies.
	// Create and build the SSA program.
	hello, _, err := ssautil.BuildPackage(
		&types.Config{Importer: importer.Default()}, fset, pkg, files, ssa.SanityCheckFunctions)
	if err != nil {
		fmt.Print(err) // type error in some package
		return
	}

	output, _ := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY, 0644)
	defer output.Close()

	// Print out the package.
	hello.WriteTo(output)

	// Print out the package-level functions.
	for _, member := range hello.Members {
		if strings.HasPrefix(member.Type().String(), "func") {
			hello.Func(member.Name()).WriteTo(output)
		}
	}
}
