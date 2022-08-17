package com.github.panda_io.micro_panda.builder.llvm.ir.instruction;

import java.util.ArrayList;
import java.util.List;

import com.github.panda_io.micro_panda.builder.llvm.ir.Identifier;
import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.constant.Constant;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class GetElementPtr extends Instruction {
    public class Index {
        boolean hasValue;
        int value;
    
        public Index(int value) {
            if (value < 0) {
                this.hasValue = false;
            } else {
                this.hasValue = true;
                this.value = value;
            }
        }
    }

    Type elementType;
    Value source;
    List<Value> indexes;
    Type type;

    public GetElementPtr(Type elementType, Value source, List<Value> indexes) {
        this.elementType = elementType;
        this.source = source;
        this.indexes = indexes;
        this.identifier = new Identifier(false);
        this.getType();
    }

    public String string() {
        return String.format("%s %s", this.type.string(), this.identifier.identifier());
    }

    public Type getType() {
        if (this.type == null) {
            List<Index> gepIndexes = new ArrayList<>();
            for (Value index : this.indexes) {
                if (index instanceof Constant) {
                    Index gepIndex = getGepIndex((Constant) index);
                    gepIndexes.add(gepIndex);
                }
            }
            this.type = getGepType(this.elementType, gepIndexes);
        }
        return this.type;
    }

    public void writeIR(StringBuilder builder) {
        builder.append(String.format("%s = getelementptr %s, %s", this.identifier.identifier(),
                this.elementType.string(), this.source.string()));
        for (Value index : this.indexes) {
            builder.append(String.format(", %s", index.string()));
        }
    }

    public static Type getGepType(Type elementType, List<Index> indexes) {
        //TO-DO
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

    public static Index getGepIndex(Constant index) {
        //TO-DO
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