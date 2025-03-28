package main

// func main() {
// 	entries, err := os.ReadDir(os.Args[1])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, e := range entries {
// 		// Parse the source files.
// 		fset := token.NewFileSet()
// 		f, err := parser.ParseFile(fset, "hello.go", parser.ParseComments)
// 		if err != nil {
// 			fmt.Print(err) // parse error
// 			return
// 		}
// 		files := []*ast.File{f}
// 		// Create the type-checker's package.
// 		pkg := types.NewPackage("", "")
// 		// Type-check the package, load dependencies.
// 		// Create and build the SSA program.
// 		hello, _, err := ssautil.BuildPackage(
// 			&types.Config{Importer: importer.Default()}, fset, pkg, files, ssa.SanityCheckFunctions)
// 		if err != nil {
// 			fmt.Print(err) // type error in some package
// 			return
// 		}
// 		// Print out the package.
// 		hello.WriteTo(os.Stdout)
// 		// Print out the package-level functions.
// 		hello.Func("init").WriteTo(os.Stdout)
// 		hello.Func("main").WriteTo(os.Stdout)
// 		//  hello.Func("f").WriteTo(os.Stdout)
// 		// for idx, val := range hello.Members {
// 		// 	fmt.Printf("%v:\t%v\t%v:%v\n", idx, val, val.Type(), val.Object())
// 		// }
// 	}
// }
