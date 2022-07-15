package com.github.panda_io.micro_panda.ast.expression;

import com.github.panda_io.micro_panda.ast.type.TypePointer;
import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.scanner.Token;

public class Unary extends Expression {
    public Token operator;
    public Expression expression;

    public boolean isLvalue() {
        return false;
    }

    public void validate(Context context, Type expected) {
        switch (this.operator) {
            case Plus:
            case Minus:
            case Not:
            case Complement:
                this.expression.validate(context, expected);
                this.type = this.expression.type;
                break;

            case BitAnd:
                if (expected == null) {
                    this.expression.validate(context, null);
                    if (this.expression.type == null) {
                        return;
                    }
                    if (this.expression.type.isPointer() || this.expression.type.isFunction()) {
                        context.addError(this.getOffset(),
                                "pointer, function and array are not allowed to use '&' operator");
                        return;
                    }
                    TypePointer pointer = new TypePointer();
                    pointer.elementType = this.expression.type;
                    this.type = pointer;
                } else {
                    if (expected.isPointer()) {
                        this.expression.validate(context, expected.elementType());
                        this.type = expected;
                    } else {
                        context.addError(this.getOffset(),
                                String.format("type mismatch, expect '%s' got pointer", expected.string()));
                        return;
                    }
                }
                if (!this.expression.isLvalue()) {
                    context.addError(this.getOffset(),
                            "expect variable operand, rvalues is not allowed to use '&' operator");
                    return;
                }
                break;

            case Mul:
                if (expected == null) {
                    this.expression.validate(context, null);
                    if (this.expression.type == null) {
                        return;
                    }
                    if (this.expression.type.isPointer()) {
                        this.type = this.expression.type.elementType();
                    } else {
                        context.addError(this.getOffset(), "only pointer type is allowed to use '*' operator");
                    }
                } else {
                    if (!expected.isPointer()) {
                        TypePointer pointer = new TypePointer();
                        pointer.elementType = expected;
                        this.expression.validate(context, pointer);
                        this.type = expected;
                    } else {
                        context.addError(this.getOffset(), "type mismatch, expect non-pointer type");
                        return;
                    }
                }
                if (!this.expression.isLvalue()) {
                    context.addError(this.getOffset(),
                            "expect variable operand, rvalues is not allowed to use '*' operator");
                    return;
                }
                break;

            default:
                context.addError(this.getOffset(), "invalid operator for unary expression");
                return;
        }

        this.constant = this.expression.constant;
        if (this.type == null) {
            return;
        }

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

            default:
        }
    }
}
