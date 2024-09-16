package llvm

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ir"
)

func IdentifierIR(c *Context, i *ast.Identifier) ir.Value {
	v := c.FindObject(i.Name)
	if v == nil {
		if i.Qualified == "" {
			return nil
		}
		d := c.Program.FindDeclaration(i.Qualified)
		switch t := d.(type) {
		case *Variable:
			return t.Variable
		case *Function:
			return t.Function
		}
	}
	return v
}

func IdentifierConstIR(p *Program, i *ast.Identifier) ir.Constant {
	d := p.FindDeclaration(i.Qualified)
	switch t := d.(type) {
	case *Variable:
		if t.Const {
			return t.Variable
		}
	case *Function:
		if t.Parent == nil {
			return t.Function
		}
	}
	return nil
}
