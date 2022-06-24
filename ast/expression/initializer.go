package expression

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/ast/types"
)

type Initializer struct {
	ExpressionBase
	Expressions []ast.Expression
}

func (i *Initializer) Validate(c ast.Context, expected core.Type) {
	if array, ok := expected.(*types.TypeArray); ok {
		i.Typ = array
		i.Const = true
		i.ValidateTypeArray(c, array, i.Expressions)
	} else if t, ok := expected.(*types.TypeName); ok {
		d := c.FindDeclaration(t)
		i.Typ = t
		i.Const = true
		if d.Kind() == core.DeclarationStruct {
			d.(core.Struct).ValidateInitializer(c, i.Expressions)
		} else {
			c.Error(i.GetPosition(), "enum has no initializer {} expression")
		}
	} else {
		c.Error(i.GetPosition(), "unexpected initializer")
	}
}

func (i *Initializer) ValidateTypeArray(c ast.Context, t *types.TypeArray, exprs []ast.Expression) {
	if t.Dimension[0] == 0 {
		c.Error(i.GetPosition(), "initializer is not allowed to pointer type variable")
	}
	if len(t.Dimension) == 1 {
		if len(exprs) == t.Dimension[0] {
			for _, expr := range exprs {
				expr.Validate(c, t.ElementType)
				if !expr.IsConstant() {
					c.Error(expr.GetPosition(), "expect constant expression initializer")
				}
				if !expr.Type().Equal(t.ElementType) {
					c.Error(expr.GetPosition(), "type mismatch")
				}
			}
		} else {
			c.Error(i.GetPosition(), "array length mismatch")
		}
	} else {
		if len(exprs) == t.Dimension[0] {
			sub := &types.TypeArray{
				ElementType: t.ElementType,
				Dimension:   t.Dimension[1:],
			}
			for _, expr := range exprs {
				if subExprs, ok := expr.(*Initializer); ok {
					i.ValidateTypeArray(c, sub, subExprs.Expressions)
				} else {
					c.Error(expr.GetPosition(), "expect array initializer")
				}
			}
		} else {
			c.Error(i.GetPosition(), "array length mismatch")
		}
	}
}
