package compilertypes

import "go/types"

type Expression interface {
	IsLHV() bool
	ReturnType() types.Type
}
