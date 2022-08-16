package com.github.panda_io.micro_panda.builder.llvm.ir.constant;

import java.util.List;

import com.github.panda_io.micro_panda.builder.llvm.ir.type.Type;

public class Struct extends Constant {
	com.github.panda_io.micro_panda.builder.llvm.ir.type.Struct type;
	List<Constant> fields;

	public Struct(com.github.panda_io.micro_panda.builder.llvm.ir.type.Struct type, List<Constant> fields) {
		this.type = type;
		this.fields = fields;
	}

	public String string() {
		return String.format("%s %s", this.type.string(), this.identifier());
	}

	public Type getType() {
		return this.type;
	}

	public String identifier() {
		if (this.fields == null || this.fields.size() == 0) {
			return "{}";
		}
		StringBuilder builder = new StringBuilder();
		builder.append("{ ");
		for (int i = 0; i < this.fields.size(); i++) {
			if (i != 0) {
				builder.append(", ");
			}
			builder.append(this.fields.get(i).string());
		}
		builder.append(" }");
		return builder.toString();
	}
}
