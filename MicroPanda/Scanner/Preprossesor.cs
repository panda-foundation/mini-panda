namespace MicroPanda.Scanner;

using Token;

internal partial class Scanner
{
    //#if #else #elif #end
	// TO-DO add preprocessor expressions
	// () == != ! && ||

    // ||
    // &&
    // == !=
    // !  Unary > Binary

    private const string IF = "if";
    private const string ELSE = "else";
    private const string ELIF = "elif";
    private const string END = "end";

}

internal abstract class Expression
{
    public abstract bool Evaluate(HashSet<string> flags);
}

internal class Binary(Expression left, Expression right, Token op) : Expression
{
    public Expression Left { get; } = left;
    public Expression Right { get; } = right;
    public Token Operator { get; } = op;

    public override bool Evaluate(HashSet<string> flags)
    {
        return Operator switch
        {
            Token.Or => Left.Evaluate(flags) || Right.Evaluate(flags),
            Token.And => Left.Evaluate(flags) && Right.Evaluate(flags),
            Token.Equal => Left.Evaluate(flags) == Right.Evaluate(flags),
            Token.NotEqual => Left.Evaluate(flags) != Right.Evaluate(flags),
            _ => throw new NotImplementedException(),
        };
    }
}

internal class Unary(Expression expression, Token op) : Expression
{
    public Expression Expression { get; } = expression;
    public Token Operator { get; } = op;

    public override bool Evaluate(HashSet<string> flags)
    {
        return Operator switch
        {
            Token.Not => !Expression.Evaluate(flags),
            _ => throw new NotImplementedException(),
        };
    }
}

internal class Parentheses(Expression expression) : Expression
{
    public Expression Expression { get; } = expression;

    public override bool Evaluate(HashSet<string> flags)
    {
        return Expression.Evaluate(flags);
    }
}

internal class Identifier(string name) : Expression
{
    public string Name { get; } = name;

    public override bool Evaluate(HashSet<string> flags)
    {
        return flags.Contains(Name);
    }
}