package com.github.panda_io.micro_panda.ast.expression;

import com.github.panda_io.micro_panda.ast.type.*;
import com.github.panda_io.micro_panda.ast.declaration.*;
import com.github.panda_io.micro_panda.ast.Context;

public class MemberAccess extends Expression {
    public Expression parent;
	public Identifier member;

	public String qualified;
	public boolean isNamespace;

	/*
	parent expression could be: identifier$, member_access$, parentheses, subscripting, 'this', invocation
	possible incomplete parent expression. it need to combine with member access
	*/
	public void validate(Context context, Type expected) {
		this.parent.validate(context, null);
		if (this.parent.type == null) {
			if (this.parent instanceof Identifier) {
				Identifier identifier = (Identifier) this.parent;
				if (identifier.isNamespace) {
					String qualified = String.format("%s.%s", identifier.name, this.member.name);
					Declaration declaration = context.findQualifiedDeclaration(qualified);
					// struct has no static members
					if (declaration != null && declaration instanceof Struct) {
						this.type = declaration.getType();
						this.constant = declaration.isConstant();
						this.qualified = declaration.qualified;
					} else if (context.isNamespace(qualified)) {
						this.isNamespace = true;
						this.qualified = qualified;
					}
				} else {
					Declaration declaration = context.findQualifiedDeclaration(identifier.qualified);
					if (declaration != null && declaration instanceof Enumeration) {
						Enumeration enumeration = (Enumeration)declaration;
						if (enumeration.hasMember(this.member.name)) {
							this.type = Type.u8;
							this.constant = true;
						}
					}
				}
			} else if (this.parent instanceof MemberAccess) {
				MemberAccess memberAccess = (MemberAccess)this.parent;
				if (memberAccess.isNamespace) {
					String qualified = String.format("%s.%s", memberAccess.qualified, this.member.name);
					Declaration declaration = context.findQualifiedDeclaration(qualified);
					// struct has no static members
					if (declaration != null && declaration instanceof Struct) {
						this.type = declaration.getType();
						this.constant = declaration.isConstant();
						this.qualified = declaration.qualified;
					} else if (context.isNamespace(qualified) ){
						this.isNamespace = true;
						this.qualified = qualified;
					}
				} else {
					Declaration declaration = context.findQualifiedDeclaration(memberAccess.qualified);
					if (declaration != null && declaration instanceof Enumeration) {
						Enumeration enumeration = (Enumeration)declaration;
						if (enumeration.hasMember(this.member.name)) {
							this.type = Type.u8;
							this.constant = true;
						}
					}
				}
			}
		} else {
			Type parentType = this.parent.type;
			if (parentType.isPointer()) {
				parentType = ((TypePointer)parentType).elementType;
			}
			if (parentType instanceof TypeName) {
				Declaration declaration = context.findDeclaration(parentType);
				if (declaration != null) {
					if (declaration instanceof Struct) {
						Struct struct = (Struct)declaration;
						this.type = struct.memberType(this.member.name);
						this.constant = false;
					}
				}
			}
		}
		// * type would be nil for enum (its member has type u8)
		if (this.type == null && this.qualified == null) {
			context.addError(this.getOffset(), String.format("undefined: %s", this.member.name));
		}
	}
}
