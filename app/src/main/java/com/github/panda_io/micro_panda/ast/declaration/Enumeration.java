package com.github.panda_io.micro_panda.ast.declaration;

import java.util.*;

import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.ast.expression.Literal;
import com.github.panda_io.micro_panda.scanner.Token;

public class Enumeration extends Declaration {
	public List<Variable> members;
	public List<Integer> values;

	public boolean isConstant() {
		return false;
	}

	public Type getType() {
		return null;
	}

	public void resolveType(Context context) {
	}

	public void validate(Context context) {
		int index = 0;
		for (Variable member : this.members) {
			if (index > 255) {
				context.addError(member.getOffset(), "enum value shoud be less than 256");
			}
			if (member.value == null) {
				this.values.add(index);
				index++;
			} else {
				if ((member.value instanceof Literal) && ((Literal) member.value).token == Token.INT) {
					int i = Integer.parseInt(((Literal) member.value).value);
					if (i >= index) {
						index = i;
						this.values.add(index);
						index++;
					} else {
						context.addError(member.getOffset(),
								String.format("enum value here should be greater than %d.", i - 1));
					}
				} else {
					context.addError(member.getOffset(), "enum value must be constant integer.");
				}
			}
		}
	}

	public boolean addMember(Variable member) {
		if (this.hasMember(member.name.name)) {
			return false;
		}
		this.members.add(member);
		return true;
	}

	public boolean hasMember(String memberName) {
		for (Variable variable : this.members) {
			if (variable.name.name == memberName) {
				return true;
			}
		}
		return false;
	}
}
