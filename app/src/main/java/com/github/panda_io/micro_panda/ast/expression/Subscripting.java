package com.github.panda_io.micro_panda.ast.expression;

import java.util.*;

import com.github.panda_io.micro_panda.ast.type.TypeArray;
import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;

public class Subscripting extends Expression {
	public Expression parent;
	public List<Expression> indexes;

	public void validate(Context context, Type expected) {
		this.constant = false;
		this.parent.validate(context, null);
		if (this.parent.type instanceof TypeArray) {
			TypeArray array = (TypeArray) this.parent.type;
			for (Expression index : this.indexes) {
				index.validate(context, null);
				if (!index.type.isInteger()) {
					context.addError(index.getOffset(),
							String.format("expect integer index for array, got '%s'", index.type.string()));
				}
			}
			if (this.indexes.size() == array.dimensions.size()) {
				this.type = array.elementType;
			} else if (this.indexes.size() < array.dimensions.size()) {
				TypeArray elementType = new TypeArray();
				elementType.elementType = array.elementType;
				elementType.dimensions.add(0);
				for (int i = array.dimensions.size() - this.indexes.size() - 1; i > 0; i--) {
					elementType.dimensions.add(array.dimensions.get(array.dimensions.size() - i));
				}
				this.type = elementType;
			} else {
				context.addError(this.getOffset(), "mismatch array dimensions");
			}
		} else {
			context.addError(this.getOffset(), "expect array type");
		}
	}
}
