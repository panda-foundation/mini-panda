package com.github.panda_io.micro_panda.ast.statement;

import java.util.List;

import com.github.panda_io.micro_panda.ast.Context;
import com.github.panda_io.micro_panda.ast.Node;
import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.scanner.Token;
import com.github.panda_io.micro_panda.ast.expression.Expression;

public class SwitchStatement extends Statement {
	public static class Case extends Node {
		public Token token;
		public List<Expression> casesExpr;
		public Statement body;

		public void validate(Context context, Type operandType) {
			for (Expression expr : this.casesExpr) {
				expr.validate(context, operandType);
				if (!expr.getType().equal(operandType)) {
					context.addError(expr.getOffset(), "case operand type mismatch with switch operand");
				}
			}
			if (this.body != null) {
				this.body.validate(context);
			}
		}
	}

	public Expression operand;
	public List<Case> cases;
	public Case defaultCase;

	public void validate(Context context) {
		Context ctx = context.newContext();
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
