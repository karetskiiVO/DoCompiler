package main

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/karetskiiVO/DoCompiler/parser"
)

type Listeners struct {
	Listeners []antlr.ParseTreeListener
	roots     []parser.IProgramContext
}

func (l *Listeners) Set(newListener func(int) antlr.ParseTreeListener) {
	for i := range l.Listeners {
		l.Listeners[i] = newListener(i)
	}
}

func (l *Listeners) Exec() {
	for i := range l.Listeners {
		antlr.ParseTreeWalkerDefault.Walk(l.Listeners[i], l.roots[i])
	}
}
