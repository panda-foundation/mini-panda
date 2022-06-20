package ast

import (
	"fmt"
	"strings"

	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/token"
)

type Error struct {
	Position *token.Position
	Message  string
}

type Namespace struct {
	Name         string
	Qualified    string
	Declarations map[string]core.Declaration // by local name
	Children     map[string]*Namespace       // by sub namespace
}

func (n *Namespace) AddDeclaration(d core.Declaration) error {
	qualified := d.QualifiedName()
	names := strings.Split(qualified, ".")
	namespace := n
	for i, name := range names {
		if i == len(names)-1 {
			if namespace.Children[name] != nil {
				return fmt.Errorf("declaration qualified name conflict with existing namespace '%s'", namespace.Children[name].Qualified)
			}
			namespace.Declarations[name] = d
		} else {
			if namespace.Declarations[name] != nil {
				return fmt.Errorf("namespace conflict with existing declaration's qualified name '%s'", namespace.Declarations[name].QualifiedName())
			}
			if namespace.Children[name] == nil {
				namespace.Children[name] = &Namespace{
					Name:         name,
					Qualified:    fmt.Sprintf("%s.%s", namespace.Qualified, name),
					Declarations: make(map[string]core.Declaration),
					Children:     make(map[string]*Namespace),
				}
			}
			namespace = namespace.Children[name]
		}
	}
	return nil
}

func (n *Namespace) IsNamespace(name string) bool {
	names := strings.Split(name, ".")
	namespace := n
	for i, name := range names {
		if i == len(names)-1 {
			return namespace.Children[name] != nil
		} else {
			if namespace.Children[name] == nil {
				return false
			}
			namespace = namespace.Children[name]
		}
	}
	return false
}

type Program struct {
	Modules map[string]*Module
	Module  *Module

	Declarations map[string]core.Declaration // by qualified name
	Namespace    *Namespace

	Errors []*Error
}

func NewProgram() *Program {
	p := &Program{}
	p.Reset()
	return p
}

func (p *Program) Reset() {
	p.Modules = make(map[string]*Module)
	p.Declarations = make(map[string]core.Declaration)
	p.Namespace = &Namespace{}
	p.Errors = p.Errors[:0]
}

func (p *Program) AddDeclaration(d core.Declaration) error {
	qualified := d.QualifiedName()
	if p.Declarations[qualified] != nil {
		return fmt.Errorf("duplicated declaration with qualified name: %s", qualified)
	}
	p.Declarations[qualified] = d
	return p.Namespace.AddDeclaration(d)
}

func (p *Program) IsNamespace(name string) bool {
	return p.Namespace.IsNamespace(name)
}

func (p *Program) Validate() {
	for _, module := range p.Modules {
		// TO-DO check if import is valid // must be valid, cannot import self, cannot duplicated
		module.ResolveType(p)
	}
	for _, module := range p.Modules {
		module.Validate(p)
	}
}

func (p *Program) Error(offset int, message string) {
	p.Errors = append(p.Errors, &Error{
		Position: p.Module.File.Position(offset),
		Message:  message,
	})
}

func (p *Program) PrintErrors() bool {
	for _, e := range p.Errors {
		fmt.Printf("%s : %s \n", e.Position.String(), e.Message)
	}
	return len(p.Errors) > 0
}
