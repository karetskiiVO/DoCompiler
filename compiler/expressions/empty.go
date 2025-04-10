package expr

import compilertypes "github.com/karetskiiVO/DoCompiler/compiler/types"

type Empty struct{}

var _ compilertypes.Expression = (*Empty)(nil)

func NewEmpty() *Empty {
	return &Empty{}
}

func (Empty) IsLHV() bool { return true }

func (e *Empty) ReturnTypes() []*compilertypes.Type {
	return []*compilertypes.Type{}
}
