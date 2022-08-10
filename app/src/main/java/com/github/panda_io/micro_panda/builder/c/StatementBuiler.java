package com.github.panda_io.micro_panda.builder.c;

import com.github.panda_io.micro_panda.ast.expression.Expression;
import com.github.panda_io.micro_panda.ast.statement.*;
import com.github.panda_io.micro_panda.ast.type.TypeArray;

public class StatementBuiler {

    static void writeStatement(StringBuilder builder, Statement statement, int indent) {
        if (statement instanceof BlockStatement) {
            writeBlockStatement(builder, (BlockStatement) statement, indent);

        } else if (statement instanceof BreakStatement) {
            writeIndent(builder, indent);
            builder.append("break;\n");

        } else if (statement instanceof ContinueStatement) {
            writeIndent(builder, indent);
            builder.append("continue;\n");

        } else if (statement instanceof DeclarationStatement) {
            writeIndent(builder, indent);
            writeSimpleStatement(builder, statement);
            builder.append(";\n");

        } else if (statement instanceof EmptyStatement) {
            writeIndent(builder, indent);
            builder.append(";\n");

        } else if (statement instanceof ExpressionStatement) {
            writeIndent(builder, indent);
            writeSimpleStatement(builder, statement);
            builder.append(";\n");

        } else if (statement instanceof ForStatement) {
            writeForStatement(builder, (ForStatement) statement, indent);

        } else if (statement instanceof IfStatement) {
            writeIndent(builder, indent);
            writeIfStatement(builder, (IfStatement) statement, indent);

        } else if (statement instanceof ReturnStatement) {
            writeReturnStatement(builder, (ReturnStatement) statement, indent);

        } else if (statement instanceof SwitchStatement) {
            writeSwitchStatement(builder, (SwitchStatement) statement, indent);
        }
    }

    static void writeIfStatement(StringBuilder builder, IfStatement statement, int indent) {
        builder.append("if (");
        ExpressionBuiler.writeExpression(builder, statement.condition);
        builder.append(")\n");
        writeStatement(builder, statement.body, indent);
        if (statement.elseStatement != null) {
            writeIndent(builder, indent);
            builder.append("else ");
            if (statement.elseStatement instanceof IfStatement) {
                writeIfStatement(builder, (IfStatement) statement.elseStatement, indent);
            } else if (statement.elseStatement instanceof BlockStatement) {
                builder.append("\n");
                writeBlockStatement(builder, (BlockStatement) statement.elseStatement, indent);
            }
        }
    }

    static void writeForStatement(StringBuilder builder, ForStatement statement, int indent) {
        writeIndent(builder, indent);
        builder.append("for (");
        if (statement.initialization != null) {
            writeSimpleStatement(builder, statement.initialization);
        }
        builder.append(";");
        if (statement.condition != null) {
            builder.append(" ");
            ExpressionBuiler.writeExpression(builder, statement.condition);
        }
        builder.append(";");
        if (statement.post != null) {
            builder.append(" ");
            writeSimpleStatement(builder, statement.post);
        }
        builder.append(")\n");
        writeBlockStatement(builder, statement.body, indent);
    }

    static void writeSwitchStatement(StringBuilder builder, SwitchStatement statement, int indent) {
        writeIndent(builder, indent);
        builder.append("switch (");
        ExpressionBuiler.writeExpression(builder, statement.operand);
        builder.append(")\n");
        writeIndent(builder, indent);
        builder.append("{\n");
        indent++;
        for (SwitchStatement.Case caseStmt : statement.cases) {
            for (Expression expr : caseStmt.casesExpr) {
                writeIndent(builder, indent);
                builder.append("case ");
                ExpressionBuiler.writeExpression(builder, expr);
                builder.append(":\n");
            }
            indent++;
            if (caseStmt.body != null) {
                for (Statement stmt : caseStmt.body.statements) {
                    writeStatement(builder, stmt, indent);
                }
            }
            writeIndent(builder, indent);
            builder.append("break;\n");
            indent--;
        }
        if (statement.defaultCase != null) {
            writeIndent(builder, indent);
            builder.append("default:\n");
            if (statement.defaultCase.body != null) {
                indent++;
                for (Statement stmt : statement.defaultCase.body.statements) {
                    writeStatement(builder, stmt, indent);
                }
                indent--;
            }
        }
        indent--;
        writeIndent(builder, indent);
        builder.append("}\n");
    }

    static void writeReturnStatement(StringBuilder builder, ReturnStatement statement, int indent) {
        writeIndent(builder, indent);
        builder.append("return");
        if (statement.expression != null) {
            builder.append(" ");
            ExpressionBuiler.writeExpression(builder, statement.expression);
        }
        builder.append(";\n");
    }

    static void writeSimpleStatement(StringBuilder builder, Statement statement) {
        if (statement instanceof ExpressionStatement) {
            ExpressionBuiler.writeExpression(builder, ((ExpressionStatement) statement).expression);

        } else if (statement instanceof DeclarationStatement) {
            DeclarationStatement declaration = (DeclarationStatement) statement;
            TypeBuiler.writeStructPrefix(builder, declaration.type);
            TypeBuiler.writeType(builder, declaration.type);
            builder.append(String.format(" %s", declaration.name.name));
            if (declaration.type instanceof TypeArray) {
                TypeBuiler.writeArrayIndex(builder, (TypeArray) declaration.type);
            }
            if (declaration.value == null) {
                if (!declaration.type.isArrayWithSize() && !declaration.type.isStruct()) {
                    builder.append(" = 0");
                }
            } else {
                builder.append(" = ");
                ExpressionBuiler.writeExpression(builder, declaration.value);
            }
        }
    }

    static void writeBlockStatement(StringBuilder builder, BlockStatement block, int indent) {
        writeIndent(builder, indent);
        builder.append("{\n");
        indent++;
        for (Statement statement : block.statements) {
            writeStatement(builder, statement, indent);
        }
        indent--;
        writeIndent(builder, indent);
        builder.append("}\n");
    }

    static void writeIndent(StringBuilder builder, int indent) {
        for (int i = 0; i < indent; i++) {
            builder.append("    ");
        }
    }
}
