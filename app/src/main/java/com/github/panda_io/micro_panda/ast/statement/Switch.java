package com.github.panda_io.micro_panda.ast.statement;

import java.util.List;

import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.ast.Node;
import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.scanner.Token;

public class Switch extends Statement {
	public static class Case extends Node {
		public Token token;
		public com.github.panda_io.micro_panda.ast.expression.Expression caseExpr;
		public Statement body;

		public void validate(Context context, Type operandType) {
			if (this.caseExpr == null) {
				context.addError(this.getOffset(), "expect case expression");
			} else {
				this.caseExpr.validate(context, operandType);
				if (!this.caseExpr.getType().equal(operandType)) {
					context.addError(this.getOffset(), "case operand type mismatch with switch operand");
				}
			}
			if (this.body != null) {
				this.body.validate(context);
			}
		}
	}

	public Statement initialization;
	public com.github.panda_io.micro_panda.ast.expression.Expression operand;
	public List<Case> cases;
	public Case defaultCase;

	public void validate(Context context) {
		Context ctx = context.newContext();
		if (this.initialization != null) {
			this.initialization.validate(context);
		}
		Type operandType;
		if (this.operand == null) {
			context.addError(this.getOffset(), "expect switch operand");
			return;
		} else {
			this.operand.validate(context, null);
			operandType = this.operand.getType();
			if (!operandType.isInteger()) {
				context.addError(this.operand.getOffset(), "expect integer operand");
			}
		}
		for (Case caseNode : this.cases) {
			Context caseCtx = ctx.newContext();
			caseNode.validate(caseCtx, operandType);
		}
		if (this.defaultCase != null) {
			Context defaultCtx = ctx.newContext();
			this.defaultCase.validate(defaultCtx, operandType);
		}
	}
}
