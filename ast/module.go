package ast

import (
	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/declaration"
	"github.com/panda-io/micro-panda/token"
)

type Import struct {
	ast.NodeBase
	Namespace string
}

type Module struct {
	File *token.File

	Namespace string
	Imports   []*Import

	Attributes []*declaration.Attribute
	Variables  []*declaration.Variable
	Functions  []*declaration.Function
	Enums      []*declaration.Enum
	Structs    []*declaration.Struct
}

func (m *Module) ResolveType(p *Program) {
	p.Module = m
	c := NewContext(p)
	for _, v := range m.Variables {
		v.ResolveType(c)
	}
	for _, f := range m.Functions {
		f.ResolveType(c)
	}
	for _, e := range m.Enums {
		e.ResolveType(c)
	}
	for _, s := range m.Structs {
		s.ResolveType(c)
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
