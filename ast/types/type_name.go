package types

import (
	"github.com/panda-io/micro-panda/ast/core"
)

type TypeName struct {
	TypeBase
	Name      string
	Qualified string

	IsEnum bool
}

func (n *TypeName) Equal(t core.Type) bool {
	if t, ok := t.(*TypeName); ok {
		return n.Qualified != "" && t.Qualified == n.Qualified
	}
	return false
}

func (n *TypeName) String() string {
	if n.Qualified != "" {
		return n.Qualified
	}
	return n.Name
}
