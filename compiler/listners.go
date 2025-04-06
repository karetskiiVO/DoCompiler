package main

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"
)

type Listners struct {
	listners []antlr.ParseTreeListener
	roots    []parser.IProgramContext
}

func (l *Listners) Set(newListner func(int) antlr.ParseTreeListener) {
	for i := range l.listners {
		l.listners[i] = newListner(i)
	}
}

func (l *Listners) Exec() {
	for i := range l.listners {
		antlr.ParseTreeWalkerDefault.Walk(l.listners[i], l.roots[i])
	}
}
