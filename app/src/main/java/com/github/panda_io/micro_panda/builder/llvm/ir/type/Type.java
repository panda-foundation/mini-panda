package com.github.panda_io.micro_panda.builder.llvm.ir.type;

public abstract class Type {
    public abstract boolean equal(Type type);
    public abstract String string();
    public abstract void writeIR(StringBuilder builder);

	// Basic types.
	public static Void Void = new Void();
	public static Label Label = new Label();

	// Integer types.
	public static Int I1   = new Int(1, false);
	public static Int I8   = new Int(8, false);
	public static Int I16  = new Int(16, false);
	public static Int I32  = new Int(32, false);
	public static Int I64  = new Int(64, false);
	public static Int UI8  = new Int(8, true);
	public static Int UI16 = new Int(16, true);
	public static Int UI32 = new Int(32, true);
	public static Int UI64 = new Int(64, true);

	// Floating-point types.
	public static Float Float16 = new Float(Float.Kind.Half);
	public static Float Float32 = new Float(Float.Kind.Float);
	public static Float Float64 = new Float(Float.Kind.Double);

    public static boolean isVoid(Type type){
        return type instanceof Void;
    }

    public static boolean isFunction(Type type){
        return type instanceof Function;
    }

    public static boolean isInt(Type type){
        if (type instanceof Int) {
            return ((Int)type).bitSize > 1;
        }
        return false;
    }

    public static boolean isBool(Type type){
        if (type instanceof Int) {
            return ((Int)type).bitSize == 1;
        }
        return false;
    }

    public static boolean isFloat(Type type){
        return type instanceof Float;
    }

    public static boolean isNumber(Type type){
        return isInt(type) || isFloat(type);
    }

    public static boolean isPointer(Type type){
        return type instanceof Pointer;
    }

    public static boolean isLabel(Type type){
        return type instanceof Label;
    }

    public static boolean isArray(Type type){
        return type instanceof Array;
    }

    public static boolean isStruct(Type type){
        return type instanceof Struct;
    }
}
