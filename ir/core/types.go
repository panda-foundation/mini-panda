package core

import (
	"fmt"
	"io"
)

type irWriter interface {
	writeIR(io.Writer) error
}

// === [ Types ] ===
// Type is an LLVM IR type.
//
// A Type has one of the following underlying types.
//
//    *types.VoidType       // https://godoc.org/github.com/llir/llvm/ir/types#VoidType
//    *types.FuncType       // https://godoc.org/github.com/llir/llvm/ir/types#FuncType
//    *types.IntType        // https://godoc.org/github.com/llir/llvm/ir/types#IntType
//    *types.FloatType      // https://godoc.org/github.com/llir/llvm/ir/types#FloatType
//    *types.PointerType    // https://godoc.org/github.com/llir/llvm/ir/types#PointerType
//    *types.LabelType      // https://godoc.org/github.com/llir/llvm/ir/types#LabelType
//    *types.ArrayType      // https://godoc.org/github.com/llir/llvm/ir/types#ArrayType
//    *types.StructType     // https://godoc.org/github.com/llir/llvm/ir/types#StructType

type Type interface {
	irWriter
	fmt.Stringer
	Equal(u Type) bool
}

var (
	// Basic types.
	Void  = &VoidType{}  // void
	Label = &LabelType{} // label

	// Integer types.
	I1   = &IntType{BitSize: 1}                  // i1
	I8   = &IntType{BitSize: 8}                  // i8
	I16  = &IntType{BitSize: 16}                 // i16
	I32  = &IntType{BitSize: 32}                 // i32
	I64  = &IntType{BitSize: 64}                 // i64
	UI8  = &IntType{BitSize: 8, Unsigned: true}  // i8
	UI16 = &IntType{BitSize: 16, Unsigned: true} // i16
	UI32 = &IntType{BitSize: 32, Unsigned: true} // i32
	UI64 = &IntType{BitSize: 64, Unsigned: true} // i64

	// Floating-point types.
	Float16 = &FloatType{Kind: FloatKindHalf}   // half
	Float32 = &FloatType{Kind: FloatKindFloat}  // float
	Float64 = &FloatType{Kind: FloatKindDouble} // double

	// Integer pointer types.
	I8Ptr  = &PointerType{ElemType: I8}  // i8*
	I16Ptr = &PointerType{ElemType: I16} // i16*
	I32Ptr = &PointerType{ElemType: I32} // i32*
	I64Ptr = &PointerType{ElemType: I64} // i64*
)

func IsVoid(t Type) bool {
	_, ok := t.(*VoidType)
	return ok
}

func IsFunc(t Type) bool {
	_, ok := t.(*FuncType)
	return ok
}

func IsInt(t Type) bool {
	if i, ok := t.(*IntType); ok {
		// bit size == 1, bool
		return i.BitSize > 1
	}
	return false
}

func IsBool(t Type) bool {
	if i, ok := t.(*IntType); ok {
		return i.BitSize == 1
	}
	return false
}

func IsFloat(t Type) bool {
	_, ok := t.(*FloatType)
	return ok
}

func IsNumber(t Type) bool {
	return IsInt(t) || IsFloat(t)
}

func IsPointer(t Type) bool {
	_, ok := t.(*PointerType)
	return ok
}

func IsLabel(t Type) bool {
	_, ok := t.(*LabelType)
	return ok
}

func IsArray(t Type) bool {
	_, ok := t.(*ArrayType)
	return ok
}

func IsStruct(t Type) bool {
	_, ok := t.(*StructType)
	return ok
}

type VoidType struct {
}

func (t *VoidType) Equal(u Type) bool {
	if _, ok := u.(*VoidType); ok {
		return true
	}
	return false
}

func (t *VoidType) String() string {
	return "void"
}

func (t *VoidType) writeIR(w io.Writer) error {
	_, err := w.Write([]byte("void"))
	return err
}

type FuncType struct {
	Qualified string
	RetType   Type
	Params    []Type
}

func NewFuncType(qualified string, retType Type, params ...Type) *FuncType {
	return &FuncType{
		Qualified: qualified,
		RetType:   retType,
		Params:    params,
	}
}

func (t *FuncType) Equal(u Type) bool {
	if u, ok := u.(*FuncType); ok {
		if !t.RetType.Equal(u.RetType) {
			return false
		}
		if len(t.Params) != len(u.Params) {
			return false
		}
		for i := range t.Params {
			if !t.Params[i].Equal(u.Params[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (t *FuncType) String() string {
	return t.Qualified
}

func (t *FuncType) writeIR(w io.Writer) error {
	err := t.RetType.writeIR(w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(" ("))
	if err != nil {
		return err
	}
	for i, param := range t.Params {
		if i != 0 {
			_, err = w.Write([]byte(", "))
			if err != nil {
				return err
			}
		}
		err = param.writeIR(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte(")"))
	return err
}

type IntType struct {
	BitSize  int
	Unsigned bool
}

func NewIntType(bitSize int) *IntType {
	return &IntType{
		BitSize: bitSize,
	}
}

func (t *IntType) Equal(u Type) bool {
	if u, ok := u.(*IntType); ok {
		return t.BitSize == u.BitSize
	}
	return false
}

func (t *IntType) String() string {
	return fmt.Sprintf("i%d", t.BitSize)
}

func (t *IntType) writeIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "i%d", t.BitSize)
	return err
}

type FloatType struct {
	Kind FloatKind
}

func (t *FloatType) Equal(u Type) bool {
	if u, ok := u.(*FloatType); ok {
		return t.Kind == u.Kind
	}
	return false
}

func (t *FloatType) String() string {
	return string(t.Kind)
}

func (t *FloatType) writeIR(w io.Writer) error {
	_, err := w.Write([]byte(t.Kind))
	return err
}

type FloatKind string

const (
	FloatKindHalf   FloatKind = "half"
	FloatKindFloat  FloatKind = "float"
	FloatKindDouble FloatKind = "double"
)

type PointerType struct {
	ElemType Type
}

func NewPointerType(elemType Type) *PointerType {
	p := &PointerType{
		ElemType: elemType,
	}
	return p
}

func (t *PointerType) Equal(u Type) bool {
	if u, ok := u.(*PointerType); ok {
		return t.ElemType.Equal(u.ElemType)
	}
	return false
}

func (t *PointerType) String() string {
	return fmt.Sprintf("%s*", t.ElemType.String())
}

func (t *PointerType) writeIR(w io.Writer) error {
	err := t.ElemType.writeIR(w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("*"))
	return err
}

type LabelType struct {
}

func (t *LabelType) Equal(u Type) bool {
	if _, ok := u.(*LabelType); ok {
		return true
	}
	return false
}

func (t *LabelType) String() string {
	return "label"
}

func (t *LabelType) writeIR(w io.Writer) error {
	_, err := w.Write([]byte("label"))
	return err
}

type ArrayType struct {
	Len      uint64
	ElemType Type
}

func NewArrayType(len uint64, elemType Type) *ArrayType {
	return &ArrayType{
		Len:      len,
		ElemType: elemType,
	}
}

func (t *ArrayType) Equal(u Type) bool {
	if u, ok := u.(*ArrayType); ok {
		if t.Len != u.Len {
			return false
		}
		return t.ElemType.Equal(u.ElemType)
	}
	return false
}

func (t *ArrayType) String() string {
	return fmt.Sprintf("[%d x %s]", t.Len, t.ElemType.String())
}

func (t *ArrayType) writeIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "[%d x ", t.Len)
	if err != nil {
		return nil
	}
	err = t.ElemType.writeIR(w)
	if err != nil {
		return nil
	}
	_, err = w.Write([]byte("]"))
	return err
}

type StructType struct {
	Qualified string
	Fields    []Type
}

func NewStructType(qualified string, fields ...Type) *StructType {
	return &StructType{
		Qualified: qualified,
		Fields:    fields,
	}
}

func (t *StructType) Equal(u Type) bool {
	if u, ok := u.(*StructType); ok {
		if len(t.Fields) != len(u.Fields) {
			return false
		}
		for i := range t.Fields {
			if !t.Fields[i].Equal(u.Fields[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (t *StructType) String() string {
	return t.Qualified
}

func (t *StructType) writeIR(w io.Writer) error {
	if len(t.Fields) == 0 {
		_, err := w.Write([]byte("{}"))
		return err
	}
	_, err := w.Write([]byte("{ "))
	if err != nil {
		return err
	}
	for i, field := range t.Fields {
		if i != 0 {
			_, err := w.Write([]byte(", "))
			if err != nil {
				return err
			}
		}
		err = field.writeIR(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte(" }"))
	return err
}
