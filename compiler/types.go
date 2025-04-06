package main

type TypeName string
type Type struct {
	// TODO
	genericArgs map[TypeName]struct {
		BehavourName string
		DepeendsOn   []TypeName
	}

	isBehavour bool
	/******** Особые свойства интерфейсов ********/
	selfBehavour bool
}

type Function struct{}
