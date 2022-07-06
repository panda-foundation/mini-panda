package com.github.panda_io.micro_panda.ast.expression;

import java.util.*;

import com.github.panda_io.micro_panda.ast.type.TypeArray;
import com.github.panda_io.micro_panda.ast.type.TypeName;
import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.declaration.Declaration;
import com.github.panda_io.micro_panda.ast.declaration.Struct;
import com.github.panda_io.micro_panda.ast.Context;

public class Initializer extends Expression {
	public List<Expression> expressions;

	public void validate(Context context, Type expected) {
		if (expected != null && expected instanceof TypeArray) {
			TypeArray array = (TypeArray) expected;
			this.type = array;
			this.constant = true;
			this.validateArray(context, array, this.expressions);
		} else if (expected != null && expected instanceof TypeName) {
			Declaration declaration = context.findDeclaration(expected);
			this.type = expected;
			this.constant = true;
			if (declaration instanceof Struct) {
				this.validateStruct(context, (Struct) declaration, this.expressions);
			} else {
				context.addError(this.getOffset(), "enum has no initializer {} expression");
			}
		} else {
			context.addError(this.getOffset(), "unexpected initializer");
		}
	}

	public void validateArray(Context context, TypeArray array, List<Expression> expressions) {
		if (array.dimensions.get(0) == 0) {
			context.addError(this.getOffset(), "initializer is not allowed for pointer");
		}
		if (array.dimensions.size() == 1) {
			if (expressions.size() == array.dimensions.get(0)) {
				for (Expression expression : expressions) {
					expression.validate(context, array.elementType);
					if (!expression.isConstant()) {
						context.addError(expression.getOffset(), "expect constant expression initializer");
					} else if (!expression.type.equal(array.elementType)) {
						context.addError(expression.getOffset(), "type mismatch");
					}
				}
			} else {
				context.addError(this.getOffset(), "array length mismatch");
			}
		} else if (expressions.size() == array.dimensions.get(0)) {
			TypeArray sub = new TypeArray();
			sub.elementType = array.elementType;
			sub.dimensions = array.dimensions.subList(1, array.dimensions.size());
			for (Expression expression : expressions) {
				if (expression instanceof Initializer) {
					this.validateArray(context, sub, ((Initializer) expression).expressions);
				} else {
					context.addError(expression.getOffset(), "expect array initializer");
				}
			}
		} else {
			context.addError(this.getOffset(), "array length mismatch");
		}
	}

	public void validateStruct(Context context, Struct struct, List<Expression> expressions) {
		if (struct.variables.size() == expressions.size()) {
			for (int i = 0; i < expressions.size(); i++) {
				Expression expression = expressions.get(i);
				expression.validate(context, struct.variables.get(i).type);
				if (!expression.isConstant()) {
					context.addError(expression.getOffset(), "expect constant expression initializer");
				}
				if (expression.type == null || !expression.type.equal(struct.variables.get(i).type)) {
					context.addError(expression.getOffset(), "type mismatch");
				}
			}
		} else {
			context.addError(expressions.get(0).getOffset(), "element number mismatch");
		}
	}
}
