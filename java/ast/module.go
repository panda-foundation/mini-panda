package ast

import "github.com/panda-io/micro-panda/token"

type Import struct {
	NodeBase
	Alias     string
	Namespace string
}

type Module struct {
	File *token.File

	Namespace string
	Imports   []*Import

	Attributes []*Attribute
	Variables  []*Variable
	Functions  []*Function
	Enums      []*Enum
	Structs    []*Struct
}

func (m *Module) ValidateType(p *Program) {
	p.Module = m
	c := NewContext(p)
	for _, v := range m.Variables {
		v.ValidateType(c)
	}
	for _, f := range m.Functions {
		f.ValidateType(c)
	}
	for _, e := range m.Enums {
		e.ValidateType(c)
	}
	for _, s := range m.Structs {
		s.ValidateType(c)
	}
}

func (m *Module) Validate(p *Program) {
	p.Module = m
	c := NewContext(p)
	for _, v := range m.Variables {
		v.Validate(c)
	}
	for _, f := range m.Functions {
		f.Validate(c)
	}
	for _, e := range m.Enums {
		e.Validate(c)
	}
	for _, s := range m.Structs {
		s.Validate(c)
	}
}
