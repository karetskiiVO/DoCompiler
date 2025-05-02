package compilertypes

import (
	"github.com/karetskiiVO/slices"
	"tinygo.org/x/go-llvm"
)

type Type = llvm.Type

func NewType() *llvm.Type {
	return &llvm.Type{}
}

type Variable struct {
	Name string
	Type *Type
	Val  *llvm.Value
}

func NewVariable(varname string, vartype *llvm.Type, val *llvm.Value) *Variable {
	return &Variable{
		Name: varname,
		Type: vartype,
		Val:  val,
	}
}

type Function struct {
	Name    string
	Args    []*Type
	Returns []*Type

	FunctionType *Type
	LLVMFunction *llvm.Value
}

func NewFunction(funcname string, args, returns []*Type, FunctionType *Type, LLVMFunction *llvm.Value) *Function {
	return &Function{
		Name:         funcname,
		Args:         args,
		Returns:      returns,
		FunctionType: FunctionType,
		LLVMFunction: LLVMFunction,
	}
}

func (f *Function) InpitTypes() []*Type {
	return slices.Clone(f.Args)
}

func (f *Function) OutputTypes() []*Type {
	return slices.Clone(f.Returns)
}
