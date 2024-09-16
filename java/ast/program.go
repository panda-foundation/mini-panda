package ast

import (
	"fmt"

	"github.com/panda-io/micro-panda/token"
)

const (
	Global        = "global"
	FunctionEntry = "entry"
	FunctionBody  = "body"
	FunctionExit  = "exit"
	StructThis    = "this"
	ProgramEntry  = "global.main"

	// meta define
	Extern   = "extern"
	Variadic = "variadic"
	Name     = "name"
)

type Error struct {
	Position *token.Position
	Message  string
}

type Program struct {
	Modules map[string]*Module
	Module  *Module

	Declarations map[string]Declaration

	Errors []*Error
}

func NewProgram() *Program {
	p := &Program{}
	p.Reset()
	return p
}

func (p *Program) Reset() {
	p.Modules = make(map[string]*Module)
	p.Declarations = make(map[string]Declaration)
	p.Errors = p.Errors[:0]
}

func (p *Program) FindSelector(selector, member string) (string, Declaration) {
	if selector == "" {
		return p.FindMember(member)
	}
	for _, i := range p.Module.Imports {
		if i.Alias == selector {
			qualified := i.Namespace + "." + member
			return qualified, p.Declarations[qualified]
		}
	}
	return "", nil
}

func (p *Program) FindMember(member string) (string, Declaration) {
	qualified := p.Module.Namespace + "." + member
	if d, ok := p.Declarations[qualified]; ok {
		return qualified, d
	}
	qualified = Global + "." + member
	if d, ok := p.Declarations[qualified]; ok {
		return qualified, d
	}
	return "", nil
}

func (p *Program) FindType(t *TypeName) Declaration {
	q, d := p.FindSelector(t.Selector, t.Name)
	if _, ok := d.(*Enum); ok {
		t.IsEnum = true
	}
	t.Qualified = q
	return d
}

func (p *Program) FindQualified(qualified string) Declaration {
	return p.Declarations[qualified]
}

func (p *Program) IsNamespace(name string) bool {
	for _, i := range p.Module.Imports {
		if i.Alias == name {
			return true
		}
	}
	return false
}

func (p *Program) Validate() {
	for _, module := range p.Modules {
		// TO-DO check if import is valid // must be valid, cannot import self, cannot duplicated
		module.ValidateType(p)
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
