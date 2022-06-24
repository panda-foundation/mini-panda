package expression

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/ast/types"
	"github.com/panda-io/micro-panda/token"
)

type Literal struct {
	ExpressionBase
	Token token.Token
	Value string
}

func (l *Literal) Validate(c ast.Context, expected core.Type) {
	l.Const = true
	switch l.Token {
	case token.STRING:
		length := 0
		if l.Value[0] == '"' {
			// string
			str, _ := strconv.Unquote(l.Value)
			length = len(str) + 1
		} else {
			// `` raw string
			length = len(l.Value) - 1
		}
		l.Typ = &types.TypeArray{
			ElementType: types.TypeU8,
			Dimension:   []int{length},
		}

	case token.CHAR:
		l.Typ = types.TypeU8

	case token.FLOAT:
		if expected != nil {
			if types.IsFloat(expected) {
				l.Typ = expected
			} else {
				c.Error(l.GetPosition(), fmt.Sprintf("type mismatch, expect '%s' got 'float'", expected.String()))
			}
		} else {
			l.Typ = types.TypeF32
		}

	case token.INT:
		if expected != nil {
			if types.IsNumber(expected) {
				l.Typ = expected
			} else {
				c.Error(l.GetPosition(), fmt.Sprintf("type mismatch, expect '%s' got 'int'", expected.String()))
			}
		} else {
			l.Typ = types.TypeI32
		}

	case token.BOOL:
		if expected != nil && !types.IsBool(expected) {
			c.Error(l.GetPosition(), fmt.Sprintf("type mismatch, expect '%s' got 'bool'", expected.String()))
		} else {
			l.Typ = types.TypeBool
		}

	case token.NULL:
		if expected == nil {
			c.Error(l.GetPosition(), "expect type for 'null'")
		} else {
			if types.IsPointer(expected) {
				l.Typ = expected
			} else {
				c.Error(l.GetPosition(), fmt.Sprintf("type mismatch, expect 'pointer' got '%s'", expected.String()))
			}
		}
	}
}

func (l *Literal) Bool() (bool, bool) {
	if l.Token == token.BOOL {
		return l.Value == "true", true
	}
	return false, false
}

func (l *Literal) Char() (int64, bool) {
	if l.Token == token.CHAR {
		r := []rune(l.Value[1 : len(l.Value)-1])
		return int64(r[0]), true
	}
	return 0, false
}

func (l *Literal) Int() (int64, bool) {
	if l.Token == token.INT {
		x, _ := (&big.Int{}).SetString(l.Value, 10)
		return x.Int64(), true
	}
	return 0, false
}

func (l *Literal) Float() (float64, bool) {
	if l.Token == token.FLOAT || l.Token == token.INT {
		x, _, _ := big.ParseFloat(l.Value, 10, 24, big.ToNearestEven)
		f, _ := x.Float64()
		return f, true
	}
	return 0, false
}

func (l *Literal) String() (string, bool) {
	if l.Token == token.STRING {
		return l.Value[1 : len(l.Value)-1], true
	}
	return "", false
}
