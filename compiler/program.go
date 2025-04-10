package compiler

import (
	"fmt"
	"strings"

	compilertypes "github.com/karetskiiVO/DoCompiler/compiler/types"
)

type Program struct {
	types     map[string]*compilertypes.Type
	variables map[string]*compilertypes.Variable
	err       error
}

func NewProgram() *Program {
	return (&Program{
		types:     make(map[string]*compilertypes.Type),
		variables: make(map[string]*compilertypes.Variable),
	}).init()
}

func (prog *Program) Types() map[string]*compilertypes.Type {
	return prog.types
}

func (prog *Program) Variables() map[string]*compilertypes.Variable {
	return prog.variables
}

func (prog *Program) Error() error {
	return prog.err
}

func (prog *Program) init() *Program {
	prog.RegisterType("int")

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

func (prog *Program) RegisterType(typename string) (*compilertypes.Type, error) {
	if _, ok := prog.types[typename]; ok {
		return nil, fmt.Errorf("type %v is already exist", typename)
	}
	
	newType := compilertypes.NewType(compilertypes.TypeName(typename))
	prog.types[typename] = newType

	return newType, nil
}

func (prog *Program) GetType(typename string) (*compilertypes.Type, error) {
	val, ok := prog.types[typename]
	if !ok {
		if strings.HasPrefix(typename, FunctionKeyword) {
			val = compilertypes.NewType(compilertypes.TypeName(typename))
			prog.types[typename] = val

			val.IsFunction = true
			typeTupples, _ := strings.CutPrefix(typename, FunctionKeyword)
			inputTupple, outputTupple, _ := strings.Cut(typeTupples[1:len(typeTupples)-1], ")(")

			inputs := strings.Split(inputTupple, ",")
			outputs := strings.Split(outputTupple, ",")

			for _, input_ := range inputs {
				_, err := prog.GetType(input_)
				if err != nil {
					return nil, err
				}
			}
			for _, output := range outputs {
				_, err := prog.GetType(output)
				if err != nil {
					return nil, err
				}
			}
		} else {
			return nil, fmt.Errorf("can't find %v", typename)
		}
	}

	return val, nil
}

func (prog *Program) RegisterGlobalVariable(varname, typename string) (*compilertypes.Variable, error) {
	var found bool
	_, found = prog.types[varname]
	if found {
		return nil, fmt.Errorf("%v is a type name", varname)
	}
	_, found = prog.variables[varname]
	if found {
		return nil, fmt.Errorf("%v is already declared in this scope", varname)
	}

	vartype, err := prog.GetType(typename)
	if err != nil {
		return nil, err
	}

	variable := compilertypes.NewVariable(compilertypes.VarName(varname), vartype)
	variable.IsConstant = vartype.IsFunction
	prog.variables[varname] = variable

	return variable, nil
}

func (prog *Program) GetVariable(varname string) (*compilertypes.Variable, error) {
	variable, ok := prog.variables[varname]
	if !ok {
		return nil, fmt.Errorf("%v is not declarated in this scope", varname)
	}

	return variable, nil
}
