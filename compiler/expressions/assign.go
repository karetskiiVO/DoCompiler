package expr

import compilertypes "github.com/karetskiiVO/DoCompiler/compiler/types"

type Assign struct {
	Lhv []compilertypes.Expression
	Rhv []compilertypes.Expression
}

var _ compilertypes.Expression = (*Assign)(nil)

func NewAssign(lhv, rhv []compilertypes.Expression) *Assign {
	return &Assign{
		Lhv: lhv,
		Rhv: rhv,
	}
}

func (Assign) IsLHV() bool {
	return false
}

func (a *Assign) ReturnTypes() []*compilertypes.Type {
	return []*compilertypes.Type{}
}
