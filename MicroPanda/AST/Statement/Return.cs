namespace MicroPanda.AST.Statement;

using Expression;

internal class Return : Statement
{
    internal Expression? Expression { get; set; }

    internal override void Validate(Context context)
    {
        if (Expression == null)
        {
            if (context.Function.ReturnType != null)
            {
                context.Program.Error(Position, "mismatch return type, expect 'null'");
            }
        }
        else
        {
            Expression!.Validate(context, context.Function.ReturnType);
            if (Expression.Type != null && !Expression.Type.Equal(context.Function.ReturnType))
            {
                context.Program.Error(Position, "mismatch return type");
            }
        }
    }
}