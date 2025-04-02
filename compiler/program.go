package main

import "fmt"

type Program struct {
	types     map[string]*Type
	functions map[string]*Function
	err       error
}

func NewProgram() *Program {
	return &Program{
		types:     make(map[string]*Type),
		functions: make(map[string]*Function),
	}
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

	newType := new(Type)
	prog.types[typename] = newType
	
	return newType, nil
}

func (prog *Program) RegisterFunc(funcname string) (*Function, error) {
	if _, ok := prog.functions[funcname]; ok {
		return nil, fmt.Errorf("function %v is already exist", funcname)
	}

	newFunc := new(Function)
	prog.functions[funcname] = newFunc
	
	return newFunc, nil
}

func (prog *Program) GetType(typename string) (*Type, error) {
	val, ok := prog.types[typename]
	if !ok {
		return nil, fmt.Errorf("can't find %v", typename)
	}

	return val, nil
}

func (prog *Program) GetFunc(funcname string) (*Function, error) {
	val, ok := prog.functions[funcname]
	if !ok {
		return nil, fmt.Errorf("can't find %v", funcname)
	}

	return val, nil
}