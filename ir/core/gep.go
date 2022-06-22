package core

import (
	"fmt"
)

type GepIndex struct {
	HasVal bool
	Val    int64
}

func NewGepIndex(val int64) GepIndex {
	return GepIndex{
		HasVal: true,
		Val:    val,
	}
}

func GepResultType(elemType Type, indices []GepIndex) Type {
	e := elemType
	for i, index := range indices {
		if i == 0 {
			continue
		}
		switch elm := e.(type) {
		case *PointerType:
			panic(fmt.Errorf("cannot index into pointer type at %d:th gep index, only valid at 0:th gep index; see https://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep", i))
		case *ArrayType:
			e = elm.ElemType
		case *StructType:
			if !index.HasVal {
				panic(fmt.Errorf("unable to index into struct type `%v` using gep with non-constant index", e))
			}
			e = elm.Fields[index.Val]
		default:
			panic(fmt.Errorf("cannot index into type %T using gep", e))
		}
	}
	return NewPointerType(e)
}
