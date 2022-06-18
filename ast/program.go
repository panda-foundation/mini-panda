package ast

import (
	"fmt"

	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/token"
)

type Error struct {
	Position *token.Position
	Message  string
}

type Program struct {
	Modules map[string]*Module
	Module  *Module

	Declarations map[string]core.Declaration

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
	p.Errors = p.Errors[:0]
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
