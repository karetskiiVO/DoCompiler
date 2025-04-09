package expr

type Assign struct {
	Lhv []Expression
	Rhv []Expression
}

func NewAssign() (*Assign, error) {
	return nil, nil
}
