package com.github.panda_io.micro_panda.ast.expression;

import com.github.panda_io.micro_panda.ast.type.Pointer;
import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.scanner.Token;

public class Unary extends Expression {
    public Token operator;
    public Expression expression;

    public void validate(Context context, Type expected) {
        this.expression.validate(context, expected);
        this.constant = this.expression.constant;
        this.type = this.expression.type;

        switch (this.operator) {
            case Plus:
            case Minus:
                if (!this.type.isNumber()) {
                    context.addError(this.getOffset(), "expect number expression");
                }
                break;

            case Not:
                if (!this.type.isBool()) {
                    context.addError(this.getOffset(), "expect boolean expression");
                }
                break;

            case Complement:
                if (!this.type.isInteger()) {
                    context.addError(this.getOffset(), "expect integer expression");
                }
                break;

            case BitAnd:
                if (this.type.isPointer() || this.type.isFunction() || this.type.isArray()) {
                    context.addError(this.getOffset(),
                            "pointer, function and array are not allowed to use '&' operator");
                    return;
                }
                this.type = new Pointer(this.type);
                if (this.expression.isConstant()) {
                    context.addError(this.getOffset(),
                            "expect variable, constant expression is not allowed to use '&' operator");
                }
                break;

            case Mul:
                if (this.type.isPointer()) {
                    this.type = ((Pointer) this.type).elementType();
                } else {
                    context.addError(this.getOffset(), "only pointer type is allowed to use '*' operator");
                }
                break;

            default:
                context.addError(this.getOffset(), "invalid operator for unary expression");
        }
    }
}
