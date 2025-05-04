package expr

import (
	"go/types"

	compilertypes "github.com/karetskiiVO/DoCompiler/compiler/types"
)

type Variable struct {
	Variable *compilertypes.Variable
}

var _ compilertypes.Expression = (*Variable)(nil)

func NewVariable(variable *compilertypes.Variable) *Variable {
	return &Variable{
		Variable: variable,
	}
}

func (v Variable) IsLHV() bool {
	return !v.Variable.IsConstant
}

func (v *Variable) ReturnType() types.Type {
	panic("implement")
}
