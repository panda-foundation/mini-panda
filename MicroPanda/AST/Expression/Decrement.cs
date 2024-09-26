namespace MicroPanda.AST.Expression;

using Type;

internal class Decrement : Expression
{
    internal Expression? Expression { get; set; }

    internal override void Validate(Context context, Type? expected)
    {
        _const = false;
        Expression!.Validate(context, expected);
        if (Expression.Const)
        {
            context.Program.Error(Position, "expect variable");
        }
        if (!TypeHelper.IsInteger(Expression.Type!))
        {
            context.Program.Error(Position, "expect integer expression");
        }
    }
}