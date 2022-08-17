package llvm

import (
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
	"github.com/panda-io/micro-panda/target/llvm/ir/ir"
)

func IdentifierIR(c *Context, i *expression.Identifier) ir.Value {
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

func IdentifierConstIR(p *Program, i *expression.Identifier) constant.Constant {
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
