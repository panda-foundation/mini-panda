package com.github.panda_io.micro_panda.builder.llvm.ir;

import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public interface Value {
	String string();

	Type getType();

	String identifier();
}
