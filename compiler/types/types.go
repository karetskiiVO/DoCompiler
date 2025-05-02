package compilertypes

import "tinygo.org/x/go-llvm"

type DoType = llvm.Type

func NewType() *llvm.Type {
	return &llvm.Type{}
}

type Variable struct {
	Name string
	Type *DoType
}

func NewVariable(varname string, vartype *llvm.Type) *Variable {
	return &Variable{
		Name: varname,
		Type: vartype,
	}
}
