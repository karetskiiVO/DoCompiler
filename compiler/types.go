package compiler

type TypeName string
type Type struct {
	// TODO
	// genericArgs map[TypeName]struct {
	// 	BehavourName string
	// 	DepeendsOn   []TypeName
	// }

	Name TypeName

	IsFunction bool
	IsBehavour bool
	/******** Особые свойства интерфейсов ********/
	SelfBehavour bool
}

func NewType(typename TypeName) *Type {
	return &Type{
		Name: typename,
	}
}

type VarName string
type Variable struct {
	Name    VarName
	VarType *Type
}

func NewVariable(varname VarName, vartype *Type) *Variable {
	return &Variable{
		Name:    varname,
		VarType: vartype,
	}
}
