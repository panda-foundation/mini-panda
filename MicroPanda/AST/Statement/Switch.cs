namespace MicroPanda.AST.Statement;

using Expression;
using Node;
using Token;
using Type;

internal class Switch : Statement
{
    internal Statement? Initialization { get; set; }
    internal Expression? Operand { get; set; }
    internal List<Case> Cases = [];
    internal Case? Default { get; set; }

    internal override void Validate(Context context)
    {
        var ctx = context.NewContext();
        Initialization?.Validate(ctx);
        Type operandType;
        if (Operand == null)
        {
            context.Program.Error(Position, "expect switch operand");
            return;
        }
        else
        {
            Operand.Validate(ctx, null);
            operandType = Operand.Type!;
            if (!TypeHelper.IsInteger(operandType))
            {
                context.Program.Error(Operand.Position, "expect integer operand");
                return;
            }
        }
        foreach (var ca in Cases)
        {
            var caseCtx = ctx.NewContext();
            ca.Validate(caseCtx, operandType);
        }
        if (Default != null)
        {
            var defaultCtx = ctx.NewContext();
            Default.Validate(defaultCtx, operandType);
        }
    }
}

internal class Case : Node
{
    internal Token Token { get; set; }
    internal Expression? CaseExpression { get; set; }
    internal Statement? Body { get; set; }

    internal void Validate(Context context, Type operandType)
    {
        if (CaseExpression == null)
        {
            context.Program.Error(Position, "expect case expression");
        }
        else
        {
            CaseExpression.Validate(context, operandType);
            if (!CaseExpression.Type!.Equal(operandType))
            {
                context.Program.Error(Position, "case operand type mismatch with switch operand");
            }
        }
        Body?.Validate(context);
    }
}