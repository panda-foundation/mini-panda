package com.github.panda_io.micro_panda.builder.llvm;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import com.github.panda_io.micro_panda.ast.type.TypeArray;
import com.github.panda_io.micro_panda.ast.type.TypeBuiltin;
import com.github.panda_io.micro_panda.ast.type.TypeFunction;
import com.github.panda_io.micro_panda.ast.type.TypeName;
import com.github.panda_io.micro_panda.ast.type.TypePointer;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Array;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Function;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Pointer;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Struct;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;
import com.github.panda_io.micro_panda.scanner.Token;

public class TypeBuilder {
    static Map<Token, Type> builtinTypes;
    static {
        builtinTypes = new HashMap<>();
        builtinTypes.put(Token.Bool, Type.I1);
        builtinTypes.put(Token.Int8, Type.I8);
        builtinTypes.put(Token.Uint8, Type.UI8);
        builtinTypes.put(Token.Int16, Type.I16);
        builtinTypes.put(Token.Uint16, Type.UI16);
        builtinTypes.put(Token.Int32, Type.I32);
        builtinTypes.put(Token.Uint32, Type.UI32);
        builtinTypes.put(Token.Int64, Type.I64);
        builtinTypes.put(Token.Uint64, Type.UI64);
        builtinTypes.put(Token.Float16, Type.Float16);
        builtinTypes.put(Token.Float32, Type.Float32);
        builtinTypes.put(Token.Float64, Type.Float64);
        builtinTypes.put(Token.Void, Type.Void);
    }

    public static Type buildType(com.github.panda_io.micro_panda.ast.type.Type type) {
        if (type instanceof TypeBuiltin) {
            return builtinTypes.get(((TypeBuiltin) type).token);

        } else if (type instanceof TypeName) {
            return new Struct(((TypeName) type).qualified, null);
            
        } else if (type instanceof TypePointer) {
            return new Pointer(buildType(((TypePointer) type).elementType));

        } else if (type instanceof TypeArray) {
            TypeArray arrayType = (TypeArray) type;
            Type elementType = buildType(arrayType.elementType);
            if (arrayType.dimensions.get(0) == 0) {
                if (arrayType.dimensions.size() == 1) {
                    return new Pointer(elementType);
                } else {
                    Array array = new Array(arrayType.dimensions.get(arrayType.dimensions.size() - 1), elementType);
                    for (int i = arrayType.dimensions.size() - 3; i >= 0; i--) {
                        array = new Array(arrayType.dimensions.get(i), array);
                    }
                    return new Pointer(array);
                }
            } else {
                Array array = new Array(arrayType.dimensions.get(arrayType.dimensions.size() - 1), elementType);
                for (int i = arrayType.dimensions.size() - 2; i >= 0; i--) {
                    array = new Array(arrayType.dimensions.get(i), array);
                }
                return array;
            }
            
        } else if (type instanceof TypeFunction) {
            TypeFunction functionType = (TypeFunction) type;
            List<Type> parameterTypes = new ArrayList<>();
            for (com.github.panda_io.micro_panda.ast.type.Type parameter : functionType.parameters) {
                parameterTypes.add(buildType(parameter));
            }
            Type returnType = Type.Void;
            if (functionType.returnType != null) {
                returnType = buildType(functionType.returnType);
            }
            return new Pointer(new Function(functionType.qualified, returnType, parameterTypes));
        }
        return null;
    }

    /*
    func ParamIR(parameter *declaration.Parameter) *ir.Param {
	var param *ir.Param
	var paramType ir_core.Type
	switch t := parameter.Typ.(type) {
	case *ast_types.TypeBuiltin:
		paramType = TypeBuiltinIR(t)

	case *ast_types.TypeName:
		paramType = TypeNameIR(t)

	case *ast_types.TypePointer:
		paramType = TypePointerIR(t)

	case *ast_types.TypeArray:
		paramType = TypeArrayIR(t)

	case *ast_types.TypeFunction:
		paramType = TypeFunctionIR(t)
	}
	param = ir.NewParam(paramType)
	param.LocalName = parameter.Name
	return param
}

func StructIR(qualified string) *ir_types.StructType {
	return ir_types.NewStructType(qualified)
}

func StructPointerIR(qualified string) *ir_types.PointerType {
	t := ir_types.NewStructType(qualified)
	return ir_types.NewPointerType(t)
}*/
}
