package compilertypes

type Expression interface {
	IsLHV() bool
	ReturnTypes() []*Type
}
