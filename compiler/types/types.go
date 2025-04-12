package compilertypes

import "go/types"

type VarName string

// TODO: Разбить по итерфейсам
type Variable struct {
	Name        VarName
	VarType     types.Type
	IsConstant  bool
	Expressions []Expression
}

func NewVariable(varname VarName, vartype types.Type) *Variable {
	return &Variable{
		Name:    varname,
		VarType: vartype,
	}
}
