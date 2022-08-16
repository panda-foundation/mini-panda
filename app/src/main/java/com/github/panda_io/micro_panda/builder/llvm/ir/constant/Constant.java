package com.github.panda_io.micro_panda.builder.llvm.ir.constant;

import com.github.panda_io.micro_panda.builder.llvm.ir.Value;
import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public abstract class Constant implements Value {
	public static Int True = new Int(Type.I1, 1);
	public static Int False = new Int(Type.I1, 0);
}