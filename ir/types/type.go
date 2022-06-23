package types

import "github.com/panda-io/micro-panda/ir/core"

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

func IsVoid(t core.Type) bool {
	_, ok := t.(*VoidType)
	return ok
}

func IsFunc(t core.Type) bool {
	_, ok := t.(*FuncType)
	return ok
}

func IsInt(t core.Type) bool {
	if i, ok := t.(*IntType); ok {
		// bit size == 1, bool
		return i.BitSize > 1
	}
	return false
}

func IsBool(t core.Type) bool {
	if i, ok := t.(*IntType); ok {
		return i.BitSize == 1
	}
	return false
}

func IsFloat(t core.Type) bool {
	_, ok := t.(*FloatType)
	return ok
}

func IsNumber(t core.Type) bool {
	return IsInt(t) || IsFloat(t)
}

func IsPointer(t core.Type) bool {
	_, ok := t.(*PointerType)
	return ok
}

func IsLabel(t core.Type) bool {
	_, ok := t.(*LabelType)
	return ok
}

func IsArray(t core.Type) bool {
	_, ok := t.(*ArrayType)
	return ok
}

func IsStruct(t core.Type) bool {
	_, ok := t.(*StructType)
	return ok
}
