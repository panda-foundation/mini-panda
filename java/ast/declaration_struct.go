package ast

import (
	"fmt"
)

type Struct struct {
	DeclarationBase
	Functions []*Function
	Variables []*Variable
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

func (s *Struct) MemberType(member string) Type {
	for _, variable := range s.Variables {
		if member == variable.Name.Name {
			return variable.Type
		}
	}
	for _, function := range s.Functions {
		if member == function.Name.Name {
			return function.Type
		}
	}
	return nil
}

func (s *Struct) Type() *TypeName {
	return &TypeName{
		Name:      s.Name.Name,
		Qualified: s.Qualified,
		IsEnum:    false,
	}
}

func (s *Struct) PointerType() *TypePointer {
	return &TypePointer{
		ElementType: s.Type(),
	}
}

func (s *Struct) ValidateType(c *Context) {
	if len(s.Variables) == 0 {
		c.Program.Error(s.Position, "struct should contain at least 1 variable member.")
	}
	for _, v := range s.Variables {
		v.ValidateType(c)
	}
	for _, f := range s.Functions {
		f.ValidateType(c)
		f.Qualified = s.Qualified + "." + f.Name.Name
	}
}

func (s *Struct) Validate(c *Context) {
	for _, v := range s.Variables {
		v.Validate(c)
		if v.Value != nil {
			c.Program.Error(v.Position, "struct member has no initialize value")
		}
	}
	for _, f := range s.Functions {
		f.Validate(c)
	}
}
