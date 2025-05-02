package compiler

import (
	"fmt"
	"strings"

	compilertypes "github.com/karetskiiVO/DoCompiler/compiler/types"
	"github.com/karetskiiVO/slices"
	"tinygo.org/x/go-llvm"
)

type Program struct {
	// TODO: таблица типов и видимостей
	types     map[string]*compilertypes.DoType
	variables map[string]*compilertypes.Variable
	ctx       llvm.Context
	err       error
}

func NewProgram() *Program {
	return (&Program{
		types:     make(map[string]*compilertypes.DoType),
		variables: make(map[string]*compilertypes.Variable),
	}).init()
}

func (prog *Program) Error() error {
	return prog.err
}

func (prog *Program) init() *Program {
	// prog.RegisterType("int")
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

func (prog *Program) RegisterType(typename string) (*compilertypes.DoType, error) {
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

	res := new(compilertypes.DoType)
	prog.types[typename] = res
	return res, nil
}

func (prog *Program) registerFuncType(typename string) (*compilertypes.DoType, error) {
	// Убераем лишнее
	preparedTypename, _ := strings.CutPrefix(typename, FunctionKeyword+"(")

	argTypenames, retTuple, _ := strings.Cut(preparedTypename, ")")

	var err error
	typeConverter := func(typename string) compilertypes.DoType {
		res, locerr := prog.GetType(typename)

		if locerr != nil {
			if err == nil {
				err = locerr
			} else {
				err = fmt.Errorf("%w, %w", err, locerr)
			}
		}

		return *res
	}
	retType, err := prog.GetType(retTuple)
	argTypes := slices.Map(strings.Split(argTypenames, ","), typeConverter)

	if err != nil {
		return nil, err
	}

	res := new(compilertypes.DoType)
	*res = llvm.FunctionType(*retType, argTypes, false)
	prog.types[typename] = res

	return res, nil
}

func (prog *Program) registerTypeTupple(typename string) (*compilertypes.DoType, error) {
	preparedTypename := (string)(([]byte)(typename)[1 : len(typename)-1])

	var err error
	typeConverter := func(typename string) compilertypes.DoType {
		res, locerr := prog.GetType(typename)

		if locerr != nil {
			if err == nil {
				err = locerr
			} else {
				err = fmt.Errorf("%w, %w", err, locerr)
			}
		}

		return *res
	}
	types := slices.Map(strings.Split(preparedTypename, ","), typeConverter)
	res := new(compilertypes.DoType)
	*res = llvm.StructType(types, false)

	return res, nil
}

func (prog *Program) GetType(typename string) (*compilertypes.DoType, error) {
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

func (prog Program) Types() map[string]*compilertypes.DoType {
	return prog.types
}

func (prog Program) Variables() map[string]*compilertypes.Variable {
	return prog.variables
}
