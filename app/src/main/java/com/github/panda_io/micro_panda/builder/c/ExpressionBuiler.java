package com.github.panda_io.micro_panda.builder.c;

import com.github.panda_io.micro_panda.ast.expression.*;

public class ExpressionBuiler {

    static void writeExpression(StringBuilder builder, Expression expression) {
        if (expression instanceof Binary) {
            Binary binary = (Binary) expression;
            writeExpression(builder, binary.left);
            builder.append(String.format(" %s ", binary.operator.toString()));
            writeExpression(builder, binary.right);
        } else if (expression instanceof Conversion) {
            Conversion conversion = (Conversion) expression;
            builder.append("((");
            TypeBuiler.writeType(builder, conversion.getType());
            builder.append(")(");
            writeExpression(builder, conversion.value);
            builder.append("))");
        } else if (expression instanceof Decrement) {
            writeExpression(builder, ((Decrement) expression).expression);
            builder.append("--");
        } else if (expression instanceof Identifier) {
            Identifier identifier = (Identifier) expression;
            if (identifier.qualified != null) {
                builder.append(identifier.qualified.replaceAll("\\.", "_"));
            } else {
                builder.append(((Identifier) expression).name);
            }
        } else if (expression instanceof Increment) {
            writeExpression(builder, ((Increment) expression).expression);
            builder.append("++");
        } else if (expression instanceof Initializer) {
            // TO-DO
        } else if (expression instanceof Invocation) {
            Invocation invocation = (Invocation) expression;
            if (invocation.define.isExtern) {
                builder.append(invocation.define.externName);
            } else {
                writeExpression(builder, invocation.function);
            }
            builder.append("(");
            for (int i = 0; i < invocation.arguments.size(); i++) {
                if (i != 0) {
                    builder.append(", ");
                }
                writeExpression(builder, invocation.arguments.get(i));
            }
            builder.append(")");
        } else if (expression instanceof Literal) {
            Literal literal = (Literal) expression;
            switch (literal.token) {
                case BOOL:
                    builder.append(literal.value.equals("true") ? "1" : "0");
                    break;

                case NULL:
                    builder.append("0");
                    break;

                default:
                    builder.append(literal.value);
            }
        } else if (expression instanceof MemberAccess) {
            MemberAccess memberAccess = (MemberAccess) expression;
            if (memberAccess.parent.getType() == null) {
                builder.append(memberAccess.qualified.replaceAll("\\.", "_"));
            } else {
                //TO-DO struct or pointer of struct
            }
        } else if (expression instanceof Parentheses) {
            builder.append("(");
            writeExpression(builder, ((Parentheses) expression).expression);
            builder.append(")");
        } else if (expression instanceof Subscripting) {
            Subscripting subscripting = (Subscripting) expression;
            writeExpression(builder, subscripting.parent);
            for (Expression index : subscripting.indexes) {
                builder.append("[");
                writeExpression(builder, index);
                builder.append("]");
            }
        } else if (expression instanceof This) {
            builder.append("this");
        } else if (expression instanceof Unary) {
            Unary unary = (Unary) expression;
            builder.append(unary.operator.toString());
            writeExpression(builder, unary.expression);
        }
    }
}
