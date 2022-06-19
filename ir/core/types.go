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
	Equal(u Type) bool
}

// Convenience types.
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

// IsFunc reports whether the given type is a function type.
func IsFunc(t Type) bool {
	_, ok := t.(*FuncType)
	return ok
}

// IsInt reports whether the given type is an integer type.
func IsInt(t Type) bool {
	if i, ok := t.(*IntType); ok {
		// bit size == 1, bool
		return i.BitSize > 1
	}
	return false
}

// IsBool reports whether the given type is an integer type.
func IsBool(t Type) bool {
	if i, ok := t.(*IntType); ok {
		return i.BitSize == 1
	}
	return false
}

// IsFloat reports whether the given type is a floating-point type.
func IsFloat(t Type) bool {
	_, ok := t.(*FloatType)
	return ok
}

func IsNumber(t Type) bool {
	return IsInt(t) || IsFloat(t)
}

// IsPointer reports whether the given type is a pointer type.
func IsPointer(t Type) bool {
	_, ok := t.(*PointerType)
	return ok
}

// IsLabel reports whether the given type is a label type.
func IsLabel(t Type) bool {
	_, ok := t.(*LabelType)
	return ok
}

// IsArray reports whether the given type is an array type.
func IsArray(t Type) bool {
	_, ok := t.(*ArrayType)
	return ok
}

// IsStruct reports whether the given type is a struct type.
func IsStruct(t Type) bool {
	_, ok := t.(*StructType)
	return ok
}

// --- [ Void types ] ----------------------------------------------------------

// VoidType is an LLVM IR void type.
type VoidType struct {
}

// Equal reports whether t and u are of equal type.
func (t *VoidType) Equal(u Type) bool {
	if _, ok := u.(*VoidType); ok {
		return true
	}
	return false
}

func (t *VoidType) writeIR(w io.Writer) error {
	_, err := w.Write([]byte("void"))
	return err
}

// --- [ Function types ] ------------------------------------------------------

// FuncType is an LLVM IR function type.
type FuncType struct {
	// Return type.
	RetType Type
	// Function parameters.
	Params []Type
}

// NewFunc returns a new function type based on the given return type and
// function parameter types.
func NewFuncType(retType Type, params ...Type) *FuncType {
	return &FuncType{
		RetType: retType,
		Params:  params,
	}
}

// Equal reports whether t and u are of equal type.
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

// --- [ Integer types ] -------------------------------------------------------

// IntType is an LLVM IR integer type.
type IntType struct {
	// Integer size in number of bits.
	BitSize int
	// If int is unsigned
	Unsigned bool
}

// NewIntType returns a new integer type based on the given integer bit size.
func NewIntType(bitSize int) *IntType {
	return &IntType{
		BitSize: bitSize,
	}
}

// Equal reports whether t and u are of equal type.
func (t *IntType) Equal(u Type) bool {
	if u, ok := u.(*IntType); ok {
		return t.BitSize == u.BitSize
	}
	return false
}

func (t *IntType) writeIR(w io.Writer) error {
	_, err := fmt.Fprintf(w, "i%d", t.BitSize)
	return err
}

// --- [ Floating-point types ] ------------------------------------------------

// FloatType is an LLVM IR floating-point type.
type FloatType struct {
	// Floating-point kind.
	Kind FloatKind
}

// Equal reports whether t and u are of equal type.
func (t *FloatType) Equal(u Type) bool {
	if u, ok := u.(*FloatType); ok {
		return t.Kind == u.Kind
	}
	return false
}

func (t *FloatType) writeIR(w io.Writer) error {
	_, err := w.Write([]byte(t.Kind))
	return err
}

// FloatKind represents the set of floating-point kinds.
type FloatKind string

// Floating-point kinds.
const (
	// 16-bit floating-point type (IEEE 754 half precision).
	FloatKindHalf FloatKind = "half"
	// 32-bit floating-point type (IEEE 754 single precision).
	FloatKindFloat FloatKind = "float"
	// 64-bit floating-point type (IEEE 754 double precision).
	FloatKindDouble FloatKind = "double"
)

// --- [ Pointer types ] -------------------------------------------------------

// PointerType is an LLVM IR pointer type.
type PointerType struct {
	// Element type.
	ElemType Type
}

// NewPointerType returns a new pointer type based on the given element type.
func NewPointerType(elemType Type) *PointerType {
	p := &PointerType{
		ElemType: elemType,
	}
	return p
}

// Equal reports whether t and u are of equal type.
func (t *PointerType) Equal(u Type) bool {
	if u, ok := u.(*PointerType); ok {
		return t.ElemType.Equal(u.ElemType)
	}
	return false
}

func (t *PointerType) writeIR(w io.Writer) error {
	err := t.ElemType.writeIR(w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("*"))
	return err
}

// --- [ Label types ] ---------------------------------------------------------

// LabelType is an LLVM IR label type, which is used for basic block values.
type LabelType struct {
	// Type name; or empty if not present.
	TypeName string
}

// Equal reports whether t and u are of equal type.
func (t *LabelType) Equal(u Type) bool {
	if _, ok := u.(*LabelType); ok {
		return true
	}
	return false
}

func (t *LabelType) writeIR(w io.Writer) error {
	_, err := w.Write([]byte("label"))
	return err
}

// --- [ Array types ] ---------------------------------------------------------

// ArrayType is an LLVM IR array type.
type ArrayType struct {
	// Array length.
	Len uint64
	// Element type.
	ElemType Type
}

// NewArrayType returns a new array type based on the given array length and element
// type.
func NewArrayType(len uint64, elemType Type) *ArrayType {
	return &ArrayType{
		Len:      len,
		ElemType: elemType,
	}
}

// Equal reports whether t and u are of equal type.
func (t *ArrayType) Equal(u Type) bool {
	if u, ok := u.(*ArrayType); ok {
		if t.Len != u.Len {
			return false
		}
		return t.ElemType.Equal(u.ElemType)
	}
	return false
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
	_, err = w.Write([]byte(" ]"))
	return err
}

// --- [ Structure types ] -----------------------------------------------------

// StructType is an LLVM IR structure type. Identified (named) struct types are
// uniqued by type names, not by structural identity.
type StructType struct {
	// Struct fields.
	Fields []Type
}

// NewStructType returns a new struct type based on the given field types.
func NewStructType(fields ...Type) *StructType {
	return &StructType{
		Fields: fields,
	}
}

// Equal reports whether t and u are of equal type.
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
