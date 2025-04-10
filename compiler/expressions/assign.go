package expr

import compilertypes "github.com/karetskiiVO/DoCompiler/compiler/types"

type Assign struct {
	Lhv []compilertypes.Expression
	Rhv []compilertypes.Expression
}

var _ compilertypes.Expression = (*Assign)(nil)

func (Assign) IsLHV() bool {
	return false
}
