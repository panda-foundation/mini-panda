package declaration

import (
	"fmt"

	"github.com/panda-io/micro-panda/ast/ast"
	"github.com/panda-io/micro-panda/ast/ast_types"
)

type Struct struct {
	DeclarationBase
	Functions []*Function
	Variables []*Variable
}

func (s *Struct) IsConstant() bool {
	return false
}

func (s *Struct) Kind() ast.DeclarationKind {
	return ast.DeclarationStruct
}

func (s *Struct) AddVariable(v *Variable) error {
	if s.IsRedeclared(v.Name.Name) {
		return fmt.Errorf("%s redeclared", v.Name.Name)
	}
	v.Parent = s
	s.Variables = append(s.Variables, v)
	return nil
}

func (s *Struct) AddFunction(f *Function) error {
	if s.IsRedeclared(f.Name.Name) {
		return fmt.Errorf("%s redeclared", f.Name.Name)
	}
	f.Parent = s
	s.Functions = append(s.Functions, f)
	return nil
}

func (s *Struct) IsRedeclared(name string) bool {
	for _, variable := range s.Variables {
		if name == variable.Name.Name {
			return true
		}
	}
	for _, function := range s.Functions {
		if name == function.Name.Name {
			return true
		}
	}
	return false
}

func (s *Struct) HasMember(member string) bool {
	return s.IsRedeclared(member)
}

func (s *Struct) MemberType(member string) ast.Type {
	for _, variable := range s.Variables {
		if member == variable.Name.Name {
			return variable.Typ
		}
	}
	for _, function := range s.Functions {
		if member == function.Name.Name {
			return function.Typ
		}
	}
	return nil
}

func (s *Struct) Type() ast.Type {
	return &ast_types.TypeName{
		Name:      s.Name.Name,
		Qualified: s.Qualified,
		IsEnum:    false,
	}
}

func (s *Struct) PointerType() *ast_types.TypePointer {
	return &ast_types.TypePointer{
		ElementType: s.Type(),
	}
}

func (s *Struct) ResolveType(c ast.Context) {
	if len(s.Variables) == 0 {
		c.Error(s.GetPosition(), "struct should contain at least 1 variable member.")
	}
	for _, v := range s.Variables {
		v.ResolveType(c)
	}
	for _, f := range s.Functions {
		f.ResolveType(c)
		f.Qualified = s.Qualified + "." + f.Name.Name
	}
}

func (s *Struct) Validate(c ast.Context) {
	for _, v := range s.Variables {
		v.Validate(c)
		if v.Value != nil {
			c.Error(v.GetPosition(), "struct member has no initialize value")
		}
	}
	for _, f := range s.Functions {
		f.Validate(c)
	}
}

func (s *Struct) ValidateInitializer(c ast.Context, expressions []ast.Expression) {
	if len(s.Variables) == len(expressions) {
		for idx, e := range expressions {
			e.Validate(c, s.Variables[idx].Typ)
			if !e.IsConstant() {
				c.Error(e.GetPosition(), "expect constant expression initializer")
			}
			if e.Type() != nil && !e.Type().Equal(s.Variables[idx].Typ) {
				c.Error(e.GetPosition(), "type mismatch")
			}
		}
	} else {
		c.Error(expressions[0].GetPosition(), "element number mismatch")
	}
}
