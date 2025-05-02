package compiler

import (
	"errors"
	"fmt"
	"strings"

	compilertypes "github.com/karetskiiVO/DoCompiler/compiler/types"
	"github.com/karetskiiVO/slices"
	"tinygo.org/x/go-llvm"
)

type Program struct {
	// TODO: таблица типов и видимостей
	types     map[string]*compilertypes.Type
	variables map[string]*compilertypes.Variable
	functions map[string]*compilertypes.Function
	ctx       llvm.Context
	mod       llvm.Module
	builder   llvm.Builder
	err       error
}

func NewProgram() *Program {
	return (&Program{
		types:     make(map[string]*compilertypes.Type),
		variables: make(map[string]*compilertypes.Variable),
		functions: make(map[string]*compilertypes.Function),
	}).init()
}

func (prog *Program) Error() error {
	return prog.err
}

func (prog *Program) init() *Program {
	prog.ctx = llvm.NewContext()
	prog.mod = prog.ctx.NewModule("main")
	prog.builder = prog.ctx.NewBuilder()

	intType, _ := prog.RegisterType("int")
	*intType = prog.ctx.IntType(32)

	// очень временное решение
	prog.RegisterFunction("tmpPrint", []string{}, []string{})
	prog.RegisterGlobalVariable("tmpOut", "int")
	// prog.RegisterType("bool")
	// prog.RegisterType("string")

	// /* временное решение */
	// prog.RegisterGlobalVariable("tmpOut", "int")
	// prog.RegisterType("act()()")
	// prog.RegisterGlobalVariable("tmpPrint", "act()()")
	/*********************/

	return prog
}

func (prog *Program) Validate() error {
	if prog.err != nil {
		return prog.err
	}

	return nil
}

func (prog *Program) AddError(err error) {
	if prog.err == nil {
		prog.err = err
	} else {
		prog.err = fmt.Errorf("%w\n%w", prog.err, err)
	}
}

func isFuncType(typename string) bool {
	return strings.HasPrefix(typename, FunctionKeyword+"(")
}

func isTypeTupple(typename string) bool {
	return strings.HasPrefix(typename, "(")
}

func (prog *Program) RegisterType(typename string) (*compilertypes.Type, error) {
	if _, ok := prog.types[typename]; ok {
		return nil, fmt.Errorf("type %v is already exist", typename)
	}

	if isFuncType(typename) {
		return prog.registerFuncType(typename)
	}

	if isTypeTupple(typename) {
		return prog.registerTypeTupple(typename)
	}

	// TODO: struct, behovour and so on

	res := new(compilertypes.Type)
	prog.types[typename] = res
	return res, nil
}

func (prog *Program) registerFuncType(typename string) (*compilertypes.Type, error) {
	// Убераем лишнее
	preparedTypename, _ := strings.CutPrefix(typename, FunctionKeyword+"(")

	argTypenames, retTuple, _ := strings.Cut(preparedTypename, ")")

	var err error
	typeConverter := func(typename string) compilertypes.Type {
		res, locerr := prog.GetType(typename)
		err = errors.Join(err, locerr)

		return *res
	}
	retType, err := prog.GetType(retTuple)

	// Пустая страка должна возвращать слайс длины 0
	sep := ","
	if len(argTypenames) == 0 {
		sep = ""
	}

	argTypes := slices.Map(strings.Split(argTypenames, sep), typeConverter)

	if err != nil {
		return nil, err
	}

	res := new(compilertypes.Type)
	*res = llvm.FunctionType(*retType, argTypes, false)
	prog.types[typename] = res

	return res, nil
}

func (prog *Program) registerTypeTupple(typename string) (*compilertypes.Type, error) {
	preparedTypename := (string)(([]byte)(typename)[1 : len(typename)-1])

	var err error
	typeConverter := func(typename string) compilertypes.Type {
		res, locerr := prog.GetType(typename)
		err = errors.Join(err, locerr)
		return *res
	}

	sep := ","
	if len(preparedTypename) == 0 {
		sep = ""
	}

	types := slices.Map(strings.Split(preparedTypename, sep), typeConverter)
	res := new(compilertypes.Type)
	prog.types[typename] = res
	*res = llvm.StructType(types, false)

	return res, nil
}

func (prog *Program) GetType(typename string) (*compilertypes.Type, error) {
	res, ok := prog.types[typename]
	if ok {
		return res, nil
	}

	if !ok && isFuncType(typename) {
		return prog.registerFuncType(typename)
	}

	if !ok && isTypeTupple(typename) {
		return prog.registerTypeTupple(typename)
	}

	return nil, fmt.Errorf("type %v is not declared in this scope", typename)
}

func (prog *Program) RegisterFunction(funcname string, argtypenames, rettypenames []string) (*compilertypes.Function, error) {
	if _, ok := prog.types[funcname]; ok {
		return nil, fmt.Errorf("function %v is already exist as type", funcname)
	}

	if _, ok := prog.functions[funcname]; ok {
		return nil, fmt.Errorf("function %v is already exist as function", funcname)
	}

	if _, ok := prog.variables[funcname]; ok {
		return nil, fmt.Errorf("function %v is already exist as variable", funcname)
	}

	var err error
	argtypes := slices.Map(argtypenames, func(typename string) *compilertypes.Type {
		argType, locerr := prog.GetType(typename)
		err = errors.Join(err, locerr)
		return argType
	})
	rettypes := slices.Map(rettypenames, func(typename string) *compilertypes.Type {
		argType, locerr := prog.GetType(typename)
		err = errors.Join(err, locerr)
		return argType
	})

	if err != nil {
		return nil, err
	}

	functype, _ := prog.GetType(
		fmt.Sprintf("%v(%v)(%v)", FunctionKeyword, strings.Join(argtypenames, ","), strings.Join(rettypenames, ",")),
	)

	val := new(llvm.Value)
	*val = llvm.AddFunction(prog.mod, funcname, *functype)

	res := compilertypes.NewFunction(
		funcname,
		argtypes,
		rettypes,
		functype,
		val,
	)

	prog.functions[funcname] = res

	return res, nil
}

func (prog *Program) GetFunction(funcname string) (*compilertypes.Function, error) {
	res, ok := prog.functions[funcname]
	if ok {
		return res, nil
	}

	return nil, fmt.Errorf("function %v is not declared in this scope", funcname)
}

func (prog Program) Types() map[string]*compilertypes.Type {
	return prog.types
}

func (prog * Program) RegisterGlobalVariable(varname string, vartype string) (*compilertypes.Variable, error) {
	if _, ok := prog.types[varname]; ok {
		return nil, fmt.Errorf("variable %v is already exist as type", varname)
	}

	if _, ok := prog.functions[varname]; ok {
		return nil, fmt.Errorf("variable %v is already exist as function", varname)
	}

	if _, ok := prog.variables[varname]; ok {
		return nil, fmt.Errorf("variable %v is already exist as variable", varname)
	}

	varType, err := prog.GetType(vartype)
	if err != nil {
		return nil, err
	}

	val := new(llvm.Value)
	*val = llvm.AddGlobal(prog.mod, *varType, varname)

	res := compilertypes.NewVariable(varname, varType, val)
	prog.variables[varname] = res

	return res, nil
}

func (prog Program) Variables() map[string]*compilertypes.Variable {
	return prog.variables
}

func (prog Program) Functions() map[string]*compilertypes.Function {
	return prog.functions
}

func (prog Program) Builder() llvm.Builder {
	return prog.builder
}

func (prog Program) Module() llvm.Module {
	return prog.mod
}