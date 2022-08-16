package com.github.panda_io.micro_panda.builder.llvm.ir.constant.expression;

import java.util.List;

import com.github.panda_io.micro_panda.builder.llvm.ir.constant.Constant;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class GepIndex {
    boolean hasValue;
    int value;

    public GepIndex(int value) {
        if (value < 0) {
            this.hasValue = false;
        } else {
            this.hasValue = true;
            this.value = value;
        }
    }
    
    public static Type getGepType(Type elementType, List<GepIndex> indexes) {
        /*
        e := elemType
        for i, index := range indices {
            if i == 0 {
                continue
            }
            switch elm := e.(type) {
            case *ir_types.PointerType:
                panic(fmt.Errorf("cannot index into pointer type at %d:th gep index, only valid at 0:th gep index; see https://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep", i))
            case *ir_types.ArrayType:
                e = elm.ElemType
            case *ir_types.StructType:
                if !index.HasVal {
                    panic(fmt.Errorf("unable to index into struct type `%v` using gep with non-constant index", e))
                }
                e = elm.Fields[index.Val]
            default:
                panic(fmt.Errorf("cannot index into type %T using gep", e))
            }
        }
        return ir_types.NewPointerType(e)*/
        return null;
    }

    public static GepIndex getGepIndex(Constant index) {
        /*
        if idx, ok := index.(*Index); ok {
            index = idx.Index
        }

        switch index := index.(type) {
        case *constant.Int:
            val := index.X.Int64()
            return NewGepIndex(val)

        case *constant.ZeroInitializer:
            return NewGepIndex(0)

        case Expression:
            return &GepIndex{HasVal: false}

        default:
            panic(fmt.Errorf("support for gep index type %T not yet implemented", index))
        }*/
        return null;
    }
}
