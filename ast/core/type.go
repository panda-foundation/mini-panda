package core

import (
	"github.com/panda-io/micro-panda/token"
)

var (
	TypeBool = &TypeBuiltin{
		Token: token.Bool,
	}

	TypeU8 = &TypeBuiltin{
		Token: token.Uint8,
	}

	TypeU32 = &TypeBuiltin{
		Token: token.Uint32,
	}

	TypeI32 = &TypeBuiltin{
		Token: token.Int32,
	}

	TypeF16 = &TypeBuiltin{
		Token: token.Float16,
	}

	TypeF32 = &TypeBuiltin{
		Token: token.Float32,
	}

	TypeF64 = &TypeBuiltin{
		Token: token.Float64,
	}

	TypePointerRaw = &TypePointer{
		ElementType: TypeU8,
	}
)

type Type interface {
	Node
	Equal(t Type) bool
}

type TypeBase struct {
	NodeBase
}

type TypeBuiltin struct {
	TypeBase
	Token token.Token
}

func (b *TypeBuiltin) Equal(t Type) bool {
	if tt, ok := t.(*TypeBuiltin); ok {
		return tt.Token == b.Token
	}
	return false
}

type TypeName struct {
	TypeBase
	Name     string
	Selector string

	Qualified string //qualified name
	IsEnum    bool
}

func (n *TypeName) Equal(t Type) bool {
	if tt, ok := t.(*TypeName); ok {
		return n.Qualified != "" && tt.Qualified == n.Qualified
	}
	return false
}

type TypeFunction struct {
	TypeBase
	ReturnType Type
	Parameters []Type

	MemberFunction bool
	Extern         bool
	ExternName     string
	TypeDefine     bool
}

func (f *TypeFunction) Equal(t Type) bool {
	if tt, ok := t.(*TypeFunction); ok {
		if f.ReturnType != nil && tt.ReturnType != nil {
			if !f.ReturnType.Equal(tt.ReturnType) {
				return false
			}
		} else if f.ReturnType != tt.ReturnType {
			return false
		}
		if len(f.Parameters) != len(tt.Parameters) {
			return false
		}
		for i := 0; i < len(f.Parameters); i++ {
			if !f.Parameters[i].Equal(tt.Parameters[i]) {
				return false
			}
		}
		return true
	}
	return false
}

type TypeArray struct {
	TypeBase
	ElementType Type
	Dimension   []int
}

func (a *TypeArray) Equal(t Type) bool {
	if tt, ok := t.(*TypeArray); ok {
		if len(a.Dimension) == len(tt.Dimension) {
			for i := 1; i < len(a.Dimension); i++ {
				if a.Dimension[i] != tt.Dimension[i] {
					return false
				}
			}
			return true
		}
	} else if tt, ok := t.(*TypePointer); ok {
		if len(a.Dimension) == 1 {
			return a.ElementType.Equal(tt.ElementType)
		}
	}
	return false
}

type TypePointer struct {
	TypeBase
	ElementType Type
}

func (p *TypePointer) Equal(t Type) bool {
	if tt, ok := t.(*TypePointer); ok {
		return p.ElementType.Equal(tt.ElementType)
	} else if tt, ok := t.(*TypeArray); ok {
		if len(tt.Dimension) == 1 {
			return p.ElementType.Equal(tt.ElementType)
		}
	}
	return false
}

func IsInteger(t Type) bool {
	if b, ok := t.(*TypeBuiltin); ok {
		return b.Token.IsInteger()
	}
	return false
}

func IsFloat(t Type) bool {
	if b, ok := t.(*TypeBuiltin); ok {
		return b.Token.IsFloat()
	}
	return false
}

func IsNumber(t Type) bool {
	if b, ok := t.(*TypeBuiltin); ok {
		return b.Token.IsNumber()
	}
	return false
}

func IsBool(t Type) bool {
	if b, ok := t.(*TypeBuiltin); ok {
		return b.Token == token.Bool
	}
	return false
}

func IsStruct(t Type) bool {
	tt, ok := t.(*TypeName)
	return ok && !tt.IsEnum
}

func IsArray(t Type) bool {
	tt, ok := t.(*TypeArray)
	return ok && tt.Dimension[0] != 0
}

func IsFunction(t Type) bool {
	_, ok := t.(*TypeFunction)
	return ok
}

func IsPointer(t Type) bool {
	_, ok := t.(*TypePointer)
	if ok {
		return true
	}
	array, ok := t.(*TypeArray)
	return ok && array.Dimension[0] == 0
}

func TypeBuilinBits(t *TypeBuiltin) int {
	switch t.Token {
	case token.Bool:
		return 1

	case token.Int8, token.Uint8:
		return 8

	case token.Int16, token.Uint16, token.Float16:
		return 16

	case token.Int32, token.Uint32, token.Float32:
		return 32

	case token.Int64, token.Uint64, token.Float64:
		return 64

	default:
		return 0
	}
}

func TypeBuilinSize(t *TypeBuiltin) int {
	switch t.Token {
	case token.Bool:
		return 1

	case token.Int8, token.Uint8:
		return 1

	case token.Int16, token.Uint16, token.Float16:
		return 2

	case token.Int32, token.Uint32, token.Float32:
		return 4

	case token.Int64, token.Uint64, token.Float64:
		return 8

	default:
		return 0
	}
}

func ElementType(t Type) Type {
	if tt, ok := t.(*TypePointer); ok {
		return tt.ElementType
	}
	if tt, ok := t.(*TypeArray); ok && tt.Dimension[0] == 0 {
		if len(tt.Dimension) == 1 {
			return tt.ElementType
		} else {
			array := &TypeArray{
				ElementType: tt.ElementType,
			}
			for i := 1; i < len(tt.Dimension); i++ {
				array.Dimension = append(array.Dimension, tt.Dimension[i])
			}
			return array
		}
	}
	return nil
}
