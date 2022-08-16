package com.github.panda_io.micro_panda.builder.llvm.ir;

public class Identifier {
	String name;
	int id;
	boolean isGlobal;

	public Identifier(boolean isGlobal) {
		this.isGlobal = isGlobal;
	}

	public String getName() {
		if (this.name == null || this.name.isEmpty()) {
			return Long.toString(this.id);
		}
		return this.name;
	}

	public void setName(String name) {
		this.name = name;
		this.id = 0;
	}

	public int getId() {
		return this.id;
	}

	public void setId(int id) {
		this.id = id;
	}

	public String identifier() {
		if (this.name == null || this.name.isEmpty()) {
			if (this.isGlobal) {
				return String.format("@%s", Integer.toString(this.id));
			}
			return String.format("%%s", Integer.toString(this.id));
		}
		if (this.isGlobal) {
			return String.format("@%s", Encode.escapeIdentifier(this.name));
		}
		return String.format("%%s", Encode.escapeIdentifier(this.name));
	}
}
