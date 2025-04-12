package compiler

import (
	"fmt"
	"go/types"
	"strings"

	compilertypes "github.com/karetskiiVO/DoCompiler/compiler/types"
)

type Program struct {
	types     map[string]*types.Type
	variables map[string]*compilertypes.Variable
	err       error
}

func NewProgram() *Program {
	return (&Program{
		types:     make(map[string]*types.Type),
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

func (prog *Program) RegisterType(typename string) (*types.Type, error) {
	if _, ok := prog.types[typename]; ok {
		return nil, fmt.Errorf("type %v is already exist", typename)
	}

	if strings.HasPrefix("act(", ) {}


	// newType := compilertypes.NewType(compilertypes.TypeName(typename))
	prog.types[typename] = newType

	return newType, nil
}
