package compiler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/karetskiiVO/slices"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Program struct {
	// TODO: таблица типов и видимостей
	types     map[string]types.Type
	variables []map[string]value.Value
	functions map[string]*ir.Func

	mod *ir.Module
	err error
}

func NewProgram() *Program {
	return (&Program{
		types:     make(map[string]types.Type),
		variables: []map[string]value.Value{make(map[string]value.Value)}, // init globalscope
		functions: make(map[string]*ir.Func),
	}).init()
}

func (prog *Program) Error() error {
	return prog.err
}

func (prog *Program) init() *Program {
	prog.mod = ir.NewModule()

	prog.types["int"] = types.I32
	prog.types["bool"] = types.I1

	// очень временное решение
	prog.RegisterFunction("tmpPrint", []string{}, []string{}, []string{})

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

func (prog *Program) AddErrorf(format string, a ...any) {
	prog.AddError(fmt.Errorf(format, a...))
}

func isFuncType(typename string) bool {
	return strings.HasPrefix(typename, FunctionKeyword+"(")
}

func isTypeTupple(typename string) bool {
	return strings.HasPrefix(typename, "(")
}

func (prog *Program) RegisterType(typename string) (types.Type, error) {
	if _, ok := prog.types[typename]; ok {
		return nil, fmt.Errorf("type `%v` is already exist", typename)
	}

	if isFuncType(typename) {
		return prog.registerFuncType(typename)
	}

	if isTypeTupple(typename) {
		return prog.registerTypeTupple(typename)
	}

	// TODO: struct, behovour and so on

	res := types.Type(nil)
	prog.types[typename] = res
	return res, nil
}

func (prog *Program) registerFuncType(typename string) (types.Type, error) {
	// Убераем лишнее
	preparedTypename, _ := strings.CutPrefix(typename, FunctionKeyword+"(")

	argTypenames, retTuple, _ := strings.Cut(preparedTypename, ")")

	var err error
	typeConverter := func(typename string) types.Type {
		res, locerr := prog.GetType(typename)
		err = errors.Join(err, locerr)

		return res
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

	res := types.NewFunc(retType, argTypes...)
	prog.types[typename] = res

	return res, nil
}

func (prog *Program) registerTypeTupple(typename string) (types.Type, error) {
	preparedTypename := (string)(([]byte)(typename)[1 : len(typename)-1])

	var err error
	typeConverter := func(typename string) types.Type {
		res, locerr := prog.GetType(typename)
		err = errors.Join(err, locerr)
		return res
	}

	sep := ","
	if len(preparedTypename) == 0 {
		sep = ""
	}

	res := types.NewStruct(
		slices.Map(strings.Split(preparedTypename, sep), typeConverter)...,
	)
	prog.types[typename] = res

	return res, nil
}

func (prog *Program) GetType(typename string) (types.Type, error) {
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

	return nil, fmt.Errorf("type `%v` is not declared in this scope", typename)
}

func (prog *Program) RegisterFunction(funcname string, argnames, argtypenames, rettypenames []string) (*ir.Func, error) {
	if _, ok := prog.types[funcname]; ok {
		return nil, fmt.Errorf("function `%v` is already exist as type", funcname)
	}

	if _, ok := prog.functions[funcname]; ok {
		return nil, fmt.Errorf("function `%v` is already exist as function", funcname)
	}

	if _, ok := prog.variables[0][funcname]; ok {
		return nil, fmt.Errorf("function `%v` is already exist as variable", funcname)
	}

	var err error
	argtypes := slices.Map(argtypenames, func(typename string) types.Type {
		argType, locerr := prog.GetType(typename)
		err = errors.Join(err, locerr)
		return argType
	})
	// rettypes := slices.Map(rettypenames, func(typename string) types.Type {
	// 	argType, locerr := prog.GetType(typename)
	// 	err = errors.Join(err, locerr)
	// 	return argType
	// })

	if err != nil {
		return nil, err
	}

	functypeRaw, _ := prog.GetType(
		fmt.Sprintf(
			"%v(%v)(%v)",
			FunctionKeyword,
			strings.Join(argtypenames, ","),
			strings.Join(rettypenames, ","),
		),
	)

	functype := functypeRaw.(*types.FuncType)
	res := prog.mod.NewFunc(
		funcname,
		functype.RetType,
		slices.Map(
			slices.Join(argnames, argtypes),
			func(arg slices.Pair[string, types.Type]) *ir.Param {
				return ir.NewParam(arg.First, arg.Second)
			},
		)...,
	)
	prog.functions[funcname] = res

	return res, nil
}

func (prog *Program) GetFunction(funcname string) (*ir.Func, error) {
	res, ok := prog.functions[funcname]
	if ok {
		return res, nil
	}

	return nil, fmt.Errorf("function `%v` is not declared in this scope", funcname)
}

func (prog *Program) GetVariable(varname string) (value.Value, error) {
	for i_ := range prog.variables {
		i := len(prog.variables) - i_ - 1

		res, ok := prog.variables[i][varname]
		if ok {
			return res, nil
		}
	}

	if funcres, ok := prog.functions[varname]; ok {
		return funcres, nil
	}

	return nil, fmt.Errorf("variable `%v` is not declared in this scope", varname)
}

func (prog Program) Types() map[string]types.Type {
	return prog.types
}

func (prog *Program) RegisterGlobalVariable(varname, vartype string) (*ir.Global, error) {
	if _, ok := prog.types[varname]; ok {
		return nil, fmt.Errorf("variable %v is already exist as type", varname)
	}

	if _, ok := prog.functions[varname]; ok {
		return nil, fmt.Errorf("variable %v is already exist as function", varname)
	}

	if _, ok := prog.variables[0][varname]; ok {
		return nil, fmt.Errorf("variable %v is already exist as variable", varname)
	}

	varType, err := prog.GetType(vartype)
	if err != nil {
		return nil, err
	}

	res := prog.mod.NewGlobal(varname, varType)
	res.Init = constant.NewZeroInitializer(varType)
	prog.variables[0][varname] = res

	return res, nil
}

func (prog *Program) RegisterVariable(block *ir.Block, varname, vartype string) (value.Value, error) {
	if len(prog.variables) == 1 {
		return prog.RegisterGlobalVariable(varname, vartype)
	}

	if _, ok := prog.types[varname]; ok {
		return nil, fmt.Errorf("variable %v is already exist as type", varname)
	}

	if _, ok := prog.functions[varname]; ok {
		return nil, fmt.Errorf("variable %v is already exist as function", varname)
	}

	varType, err := prog.GetType(vartype)
	if err != nil {
		return nil, err
	}

	res := block.NewAlloca(varType)
	res.SetName(varname)
	prog.variables[len(prog.variables)-1][varname] = res

	return res, nil
}

func (prog Program) GlobalVariables() map[string]value.Value {
	return prog.variables[0]
}

func (prog Program) Functions() map[string]*ir.Func {
	return prog.functions
}

func (prog Program) Module() *ir.Module {
	return prog.mod
}

func (prog *Program) NewScope() {
	prog.variables = append(prog.variables, make(map[string]value.Value))
}

func (prog *Program) TerminateScope() {
	if len(prog.variables) <= 1 {
		panic("can't terminate")
	}

	prog.variables = prog.variables[:len(prog.variables)-1]
}
