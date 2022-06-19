package core

import (
	"fmt"
	"strings"

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

type TypeBase struct {
	NodeBase
}

type TypeBuiltin struct {
	TypeBase
	Token token.Token
}

func (b *TypeBuiltin) Equal(t Type) bool {
	if t, ok := t.(*TypeBuiltin); ok {
		return t.Token == b.Token
	}
	return false
}

func (b *TypeBuiltin) String() string {
	return b.Token.String()
}

type TypeName struct {
	TypeBase
	Name      string
	Qualified string

	IsEnum bool
}

func (n *TypeName) Equal(t Type) bool {
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
	if t, ok := t.(*TypeFunction); ok {
		if f.ReturnType != nil && t.ReturnType != nil {
			if !f.ReturnType.Equal(t.ReturnType) {
				return false
			}
		} else if f.ReturnType != t.ReturnType {
			return false
		}
		if len(f.Parameters) != len(t.Parameters) {
			return false
		}
		for i := 0; i < len(f.Parameters); i++ {
			if !f.Parameters[i].Equal(t.Parameters[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (f *TypeFunction) String() string {
	var b strings.Builder
	b.WriteString("function(")
	for i, t := range f.Parameters {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(t.String())
	}
	b.WriteString(")->")
	if f.ReturnType == nil {
		b.WriteString("null")
	} else {
		b.WriteString(f.ReturnType.String())
	}
	return b.String()
}

type TypeArray struct {
	TypeBase
	ElementType Type
	Dimension   []int
}

func (a *TypeArray) Equal(t Type) bool {
	if array, ok := t.(*TypeArray); ok {
		if len(a.Dimension) == len(array.Dimension) {
			for i := 1; i < len(a.Dimension); i++ {
				if a.Dimension[i] != array.Dimension[i] {
					return false
				}
			}
			return true
		}
	} else if pointer, ok := t.(*TypePointer); ok {
		if len(a.Dimension) == 1 {
			return a.ElementType.Equal(pointer.ElementType)
		}
	}
	return false
}

func (a *TypeArray) String() string {
	return fmt.Sprintf("array[%s]", a.ElementType.String())
}

type TypePointer struct {
	TypeBase
	ElementType Type
}

func (p *TypePointer) Equal(t Type) bool {
	if pointer, ok := t.(*TypePointer); ok {
		return p.ElementType.Equal(pointer.ElementType)
	} else if array, ok := t.(*TypeArray); ok {
		if len(array.Dimension) == 1 {
			return p.ElementType.Equal(array.ElementType)
		}
	}
	return false
}

func (p *TypePointer) String() string {
	return fmt.Sprintf("pointer<%s>", p.ElementType.String())
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
	n, ok := t.(*TypeName)
	return ok && !n.IsEnum
}

func IsArray(t Type) bool {
	array, ok := t.(*TypeArray)
	return ok && array.Dimension[0] != 0
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

func TypeBuiltinBits(t *TypeBuiltin) int {
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

func TypeBuiltinSize(t *TypeBuiltin) int {
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

func GetElementType(t Type) Type {
	if t, ok := t.(*TypePointer); ok {
		return t.ElementType
	}
	if t, ok := t.(*TypeArray); ok && t.Dimension[0] == 0 {
		if len(t.Dimension) == 1 {
			return t.ElementType
		} else {
			array := &TypeArray{
				ElementType: t.ElementType,
			}
			for i := 1; i < len(t.Dimension); i++ {
				array.Dimension = append(array.Dimension, t.Dimension[i])
			}
			return array
		}
	}
	return nil
}
