package expr

import (
	"go/types"

	compilertypes "github.com/karetskiiVO/DoCompiler/compiler/types"
)

type Const struct {
	Content     any
	ContentType types.Type
}

var _ compilertypes.Expression = (*Const)(nil)

func NewConst(content any, contentType types.Type) *Const {
	return &Const{
		Content:     content,
		ContentType: contentType,
	}
}

func (Const) IsLHV() bool {
	return false
}

func (c *Const) ReturnType() types.Type {
	panic("implement")
}
