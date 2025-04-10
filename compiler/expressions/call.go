package expr

import compilertypes "github.com/karetskiiVO/DoCompiler/compiler/types"

type Call struct {
	Function  compilertypes.Variable
	Arguments []compilertypes.Expression
}

var _ compilertypes.Expression = (*Call)(nil)

func (Call) IsLHV() bool {
	return false
}
