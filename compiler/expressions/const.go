package expr

import compilertypes "github.com/karetskiiVO/DoCompiler/compiler/types"

type Const struct {
	Content     any
	ContentType *compilertypes.Type
}

var _ compilertypes.Expression = (*Const)(nil)

func NewConst(content any, contentType *compilertypes.Type) *Const {
	return &Const{
		Content:     content,
		ContentType: contentType,
	}
}

func (Const) IsLHV() bool {
	return false
}

func (c *Const) ReturnTypes() []*compilertypes.Type {
	return []*compilertypes.Type{c.ContentType}
}
