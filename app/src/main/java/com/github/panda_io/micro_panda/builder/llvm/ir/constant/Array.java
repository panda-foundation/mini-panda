package com.github.panda_io.micro_panda.builder.llvm.ir.constant;

import java.util.List;

import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Array extends Constant {
	com.github.panda_io.micro_panda.builder.llvm.ir.type.Array type;
	List<Constant> elements;

	public Array(com.github.panda_io.micro_panda.builder.llvm.ir.type.Array type, List<Constant> elements) {
		this.type = type;
		this.elements = elements;
	}

	public String string() {
		return String.format("%s %s", this.type.string(), this.identifier());
	}

	public Type getType() {
		return this.type;
	}

	public String identifier() {
		StringBuilder builder = new StringBuilder();
		builder.append("[");
		for (int i = 0; i < this.elements.size(); i++) {
			if (i != 0) {
				builder.append(", ");
			}
			builder.append(this.elements.get(i).string());
		}
		builder.append("]");
		return builder.toString();
	}
}
