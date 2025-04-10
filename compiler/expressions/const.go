package expr

import compilertypes "github.com/karetskiiVO/DoCompiler/compiler/types"

type Const struct {
	Content any
}

var _ compilertypes.Expression = (*Const)(nil)

func NewConst(content any) *Const {
	return &Const{
		Content: content,
	}
}

func (Const) IsLHV() bool {
	return false
}
