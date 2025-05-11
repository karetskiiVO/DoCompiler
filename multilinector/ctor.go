package multilinector

import (
	"fmt"
	"strings"
)

type MultilineCtor struct {
	builder *strings.Builder
}

func New() *MultilineCtor {
	return &MultilineCtor{
		builder: &strings.Builder{},
	}
}

func (m *MultilineCtor) AddLine(line string) *MultilineCtor {
	if m.builder.Len() != 0 {
		m.builder.WriteRune('\n')
	}

	fmt.Fprint(m.builder, line)
	return m
}

func (m *MultilineCtor) String() string {
	return m.builder.String()
}
