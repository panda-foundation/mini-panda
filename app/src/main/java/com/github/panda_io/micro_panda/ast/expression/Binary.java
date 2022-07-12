package com.github.panda_io.micro_panda.ast.expression;

import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.scanner.Token;

public class Binary extends Expression {
	public Expression left;
	public Expression right;
	public Token operator;

	public void validate(Context context, Type expected) {
		this.left.validate(context, expected);
		if (this.left.type == null) {
			return;
		}
		this.right.validate(context, this.left.type);
		if (this.right.type == null) {
			return;
		}
		if (!this.left.type.equal(this.right.type)) {
			context.addError(this.left.getOffset(), "mismatch type for binary expression");
		}

		switch (this.operator) {
			case LeftShift:
			case RightShift:
			case BitXor:
			case BitOr:
			case BitAnd:
				this.constant = this.left.isConstant() && this.right.isConstant();
				if (!this.left.type.isInteger()) {
					context.addError(this.left.getOffset(), "expect integer for bit operation");
				} else if (!this.right.type.isInteger()) {
					context.addError(this.right.getOffset(), "expect integer for bit operation");
				} else {
					this.type = this.left.type;
				}
				break;

			case Assign:
				this.constant = false;
				if (this.left.isConstant()) {
					context.addError(this.left.getOffset(), "expect variable");
				}
				if (this.left.type.isArrayWithSize()) {
					context.addError(this.left.getOffset(), "array type is not assignable");
				}
				break;

			case MulAssign:
			case DivAssign:
			case RemAssign:
			case PlusAssign:
			case MinusAssign:
				this.constant = false;
				if (this.left.isConstant()) {
					context.addError(this.left.getOffset(), "expect variable");
				}
				if (!this.left.type.isNumber()) {
					context.addError(this.left.getOffset(), "expect number for binary expression");
				}
				break;

			case LeftShiftAssign:
			case RightShiftAssign:
			case AndAssign:
			case OrAssign:
			case XorAssign:
				this.constant = false;
				if (this.left.isConstant()) {
					context.addError(this.left.getOffset(), "expect variable");
				}
				if (!this.left.type.isInteger()) {
					context.addError(this.left.getOffset(), "expect integer for binary expression");
				}
				break;

			case Or:
			case And:
				this.constant = this.left.isConstant() && this.right.isConstant();
				if (!this.left.type.isBool()) {
					context.addError(this.left.getOffset(), "expect bool binary expression");
				}
				break;

			case Less:
			case LessEqual:
			case Greater:
			case GreaterEqual:
				this.constant = this.left.isConstant() && this.right.isConstant();
				if (this.left.type.isNumber()) {
					this.type = Type.bool;
				} else {
					context.addError(this.left.getOffset(), "expect number for comparison");
				}
				break;

			case Equal:
			case NotEqual:
				this.constant = this.left.isConstant() && this.right.isConstant();
				if (this.left.type.isNumber() || this.left.type.isPointer()) {
					this.type = Type.bool;
				} else {
					context.addError(this.left.getOffset(), "expect number or pointer for comparison");
				}
				break;

			case Plus:
			case Minus:
			case Mul:
			case Div:
			case Rem:
				this.constant = this.left.isConstant() && this.right.isConstant();
				if (this.left.type.isNumber()) {
					this.type = this.left.type;
				} else {
					context.addError(this.left.getOffset(), "expect number for binary expression");
				}
				break;

			default:
				context.addError(this.left.getOffset(), "invalid operator for binary expression");
		}
	}
}
