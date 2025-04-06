package main

import (
	"fmt"
	"strings"
)

type Program struct {
	types     map[string]*Type
	variables map[string]*Variable
	err       error
}

func NewProgram() *Program {
	return (&Program{
		types:     make(map[string]*Type),
		variables: make(map[string]*Variable),
	}).init()
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

func (prog *Program) RegisterType(typename string) (*Type, error) {
	if _, ok := prog.types[typename]; ok {
		return nil, fmt.Errorf("type %v is already exist", typename)
	}

	newType := NewType(TypeName(typename))
	prog.types[typename] = newType

	return newType, nil
}

// func (prog *Program) RegisterFunc(funcname string) (*Function, error) {
// 	if _, ok := prog.functions[funcname]; ok {
// 		return nil, fmt.Errorf("function %v is already exist", funcname)
// 	}
// 	newFunc := new(Function)
// 	prog.functions[funcname] = newFunc
// 	return newFunc, nil
// }

func (prog *Program) GetType(typename string) (*Type, error) {
	val, ok := prog.types[typename]
	if !ok {
		if strings.HasPrefix(typename, FunctionKeyword) {
			val = NewType(TypeName(typename))
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

func (prog *Program) RegisterGlobalVariable(varname, typename string) (*Variable, error) {
	var found bool
	_, found = prog.types[varname]
	if found {
		return nil, fmt.Errorf("%v is a type name", varname)
	}
	_, found = prog.variables[varname]
	if found {
		return nil, fmt.Errorf("%v is akready declared in this scope", varname)
	}

	vartype, err := prog.GetType(typename)
	if err != nil {
		return nil, err
	}

	variable := NewVariable(VarName(varname), vartype)
	prog.variables[varname] = variable

	return variable, nil
}

// func (prog *Program) GetFunc(funcname string) (*Function, error) {
// 	val, ok := prog.functions[funcname]
// 	if !ok {
// 		return nil, fmt.Errorf("can't find %v", funcname)
// 	}
// 	return val, nil
// }
