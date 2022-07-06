package com.github.panda_io.micro_panda.ast.expression;

import com.github.panda_io.micro_panda.ast.type.TypeArray;
import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.scanner.Token;

public class Literal extends Expression {
	public Token token;
	public String value;
	
    public void validate(Context context, Type expected) {
		this.constant = true;
		switch (this.token) {
		case STRING:
			//TO-DO check & test unquote
			TypeArray array = new TypeArray(Type.u8);
			array.dimensions.add(this.value.length() - 1);
			this.type = array;
			break;

		case CHAR:
			if (expected == null || expected.equal(Type.u8)) {
				this.type = Type.u8;
			} else {
				context.addError(this.getOffset(), String.format("type mismatch, expect '%s' got 'u8'", expected.string()));
			}
			break;

		case FLOAT:
			if (expected == null) {
				this.type = Type.f32;
			} else if (expected.isFloat()) {
				this.type = expected;
			} else {
				context.addError(this.getOffset(), String.format("type mismatch, expect '%s' got 'float'", expected.string()));
			}
			break;

		case INT:
			if (expected == null) {
				this.type = Type.i32;
			} else if (expected.isNumber()) {
				this.type = expected;
			} else {
				context.addError(this.getOffset(), String.format("type mismatch, expect '%s' got 'int'", expected.string()));
			}
			break;

		case BOOL:
			if (expected == null || expected.isBool()) {
				this.type = Type.bool;
			} else {
				context.addError(this.getOffset(), String.format("type mismatch, expect '%s' got 'bool'", expected.string()));
			}
			break;

		case NULL:
			if (expected == null) {
				context.addError(this.getOffset(), "expect type for 'null'");
			} else if (expected.isPointer()) {
				this.type = expected;
			} else {
				context.addError(this.getOffset(), String.format("type mismatch, expect 'pointer' got '%s'", expected.string()));
			}
			break;

		default:
			context.addError(this.getOffset(), "invalid token");
		}
    } 
}
