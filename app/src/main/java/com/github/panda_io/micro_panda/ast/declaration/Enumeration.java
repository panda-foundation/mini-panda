package com.github.panda_io.micro_panda.ast.declaration;

import java.util.*;

import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.ast.expression.Literal;
import com.github.panda_io.micro_panda.scanner.Token;

public class Enumeration extends Declaration {
	public List<Variable> members;

	public Enumeration() {
		this.members = new ArrayList<>();
	}

	public boolean isConstant() {
		return false;
	}

	public Type getType() {
		return null;
	}

	public void resolveType(Context context) {
		for (Variable member : this.members) {
			member.type = Type.u8;
			member.qualified = String.format("%s.%s", this.qualified, member.name.name);
		}

		int index = 0;
		for (Variable member : this.members) {
			if (index > 255) {
				context.addError(member.getOffset(), "enum value shoud be less than 256");
			}
			Literal literal = new Literal();
			literal.token = Token.INT;
			if (member.value == null) {
				literal.value = Integer.toString(index);
				member.value = literal;
				index++;
			} else {
				if ((member.value instanceof Literal) && ((Literal) member.value).token == Token.INT) {
					int i = Integer.parseInt(((Literal) member.value).value);
					if (i >= index) {
						index = i;
						literal.value = Integer.toString(index);
						member.value = literal;
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

	public void validate(Context context) {
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
			if (variable.name.name.equals(memberName)) {
				return true;
			}
		}
		return false;
	}

	public String getValue(String memberName) {
		for (Variable variable : this.members) {
			if (variable.name.name.equals(memberName)) {
				return ((Literal) (variable.value)).value;
			}
		}
		return null;
	}
}
