namespace MicroPanda.AST.Expression;

using Type;

internal class Parentheses : Expression
{
    internal Expression? Expression { get; set; }

    internal override void Validate(Context context, Type? expected)
    {
        Expression!.Validate(context, expected);
        _const = Expression.Const;
        _type = Expression.Type;
    }
}