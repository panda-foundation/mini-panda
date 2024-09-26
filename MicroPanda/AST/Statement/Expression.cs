namespace MicroPanda.AST.Statement;

using Expression;

internal class ExpressionStatement : Statement
{
    internal Expression? Expression { get; set; }

    internal override void Validate(Context c)
    {
        Expression!.Validate(c, null);
    }
}