package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

const hello = `
package main

import "fmt"

const message = "Hello, World!"

func f () (int, int) {
	ch := make(chan int)
	_ = ch
	return 1, 2
}

func g () {}

func main() {
	g()
	a, b := f()
	fmt.Println(a, b)
}
`

func main() {
	// Parse the source files.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", hello, parser.ParseComments)
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

	// Print out the package.
	// hello.WriteTo(os.Stdout)

	// Print out the package-level functions.
	//  hello.Func("init").WriteTo(os.Stdout)
	//  hello.Func("main").WriteTo(os.Stdout)
	//  hello.Func("f").WriteTo(os.Stdout)

	for idx, val := range hello.Members {
		fmt.Printf("%v:\t%v\t%v:%v\n", idx, val, val.Type(), val.Object())
	}
}
