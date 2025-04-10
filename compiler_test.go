package main

import (
	"testing"

	example "github.com/karetskiiVO/DoCompiler/example"
)

func TestCompile(t *testing.T) {
	Compile(".\\doexamples\\test1.txt", ".\\doexamples\\test2.txt")
}

func TestExample(t *testing.T) {
	example.Run()
}