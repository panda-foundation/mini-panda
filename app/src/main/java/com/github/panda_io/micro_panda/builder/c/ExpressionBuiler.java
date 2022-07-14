package com.github.panda_io.micro_panda.builder.c;

import com.github.panda_io.micro_panda.ast.expression.*;

public class ExpressionBuiler {

    static void writeExpression(StringBuilder builder, Expression expression) {
        if (expression instanceof Binary) {

        } else if (expression instanceof Conversion) {

        } else if (expression instanceof Decrement) {

        } else if (expression instanceof Identifier) {

        } else if (expression instanceof Increment) {

        } else if (expression instanceof Initializer) {

        } else if (expression instanceof Invocation) {

        } else if (expression instanceof Literal) {

        } else if (expression instanceof MemberAccess) {

        } else if (expression instanceof Parentheses) {

        } else if (expression instanceof Subscripting) {

        } else if (expression instanceof This) {

        } else if (expression instanceof Unary) {

        }
    }
}
