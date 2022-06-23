package types

import (
	"github.com/panda-io/micro-panda/ast/core"
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
	core.NodeBase
}

func IsInteger(t core.Type) bool {
	if b, ok := t.(*TypeBuiltin); ok {
		return b.Token.IsInteger()
	}
	return false
}

func IsFloat(t core.Type) bool {
	if b, ok := t.(*TypeBuiltin); ok {
		return b.Token.IsFloat()
	}
	return false
}

func IsNumber(t core.Type) bool {
	if b, ok := t.(*TypeBuiltin); ok {
		return b.Token.IsNumber()
	}
	return false
}

func IsBool(t core.Type) bool {
	if b, ok := t.(*TypeBuiltin); ok {
		return b.Token == token.Bool
	}
	return false
}

func IsStruct(t core.Type) bool {
	n, ok := t.(*TypeName)
	return ok && !n.IsEnum
}

func IsArray(t core.Type) bool {
	array, ok := t.(*TypeArray)
	return ok && array.Dimension[0] != 0
}

func IsFunction(t core.Type) bool {
	_, ok := t.(*TypeFunction)
	return ok
}

func IsPointer(t core.Type) bool {
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

func GetElementType(t core.Type) core.Type {
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
