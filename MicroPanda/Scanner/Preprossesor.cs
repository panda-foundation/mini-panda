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

    private const string KEYWORD_IF = "if";
    private const string KEYWORD_ELSE = "else";
    private const string KEYWORD_ELIF = "elif";
    private const string KEYWORD_END = "end";

    private readonly Stack<Preprocessor> _preprocessors = new();
}

internal class Preprocessor(string keyword, string expression)
{
    internal string Keyword { get; } = keyword;
    internal string Expression { get; } = expression;
}

/*
internal class Preprocessor
{
    private readonly Scanner _scanner;
    private readonly HashSet<string> _flags = new();

    public Preprocessor(Scanner scanner)
    {
        _scanner = scanner;
    }

    public void Process()
    {
        while (true)
        {
            var (offset, token, literal) = _scanner.Scan();
            if (token == Token.EOF)
            {
                break;
            }

            if (token == Token.IDENTIFIER && literal == KEYWORD_IF)
            {
                var expression = ParseExpression();
                if (expression.Evaluate(_flags))
                {
                    Process();
                }
                else
                {
                    Skip();
                }
            }
            else if (token == Token.IDENTIFIER && literal == KEYWORD_ELIF)
            {
                Skip();
                var expression = ParseExpression();
                if (expression.Evaluate(_flags))
                {
                    Process();
                }
                else
                {
                    Skip();
                }
            }
            else if (token == Token.IDENTIFIER && literal == KEYWORD_ELSE)
            {
                Skip();
                Process();
            }
            else if (token == Token.IDENTIFIER && literal == KEYWORD_END)
            {
                break;
            }
            else
            {
                _scanner.SkipLine();
            }
        }
    }

    private void Skip()
    {
        while (true)
        {
            var (offset, token, literal) = _scanner.Scan();
            if (token == Token.EOF)
            {
                break;
            }

            if (token == Token.IDENTIFIER && literal == KEYWORD_END)
            {
                break;
            }
        }
    }*/

internal partial class Scanner
{
    /*
    private Expression ParseExpression()
    {
        var expression = ParseOr();
        return expression;
    }

    private Expression ParseOr()
    {
        var left = ParseAnd();
        while (_scanner.Match(Token.OR))
        {
            var op = _scanner.Token;
            var right = ParseAnd();
            left = new Binary(left, right, op);
        }
        return left;
    }

    private Expression ParseAnd()
    {
        var left = ParseEquality();
        while (_scanner.Match(Token.AND))
        {
            var op = _scanner.Token;
            var right = ParseEquality();
            left = new Binary(left, right, op);
        }
        return left;
    }

    private Expression ParseEquality()
    {
        var left = ParseUnary();
        while (_scanner.Match(Token.EQUAL, Token.NOT_EQUAL))
        {
            var op = _scanner.Token;
            var right = ParseUnary();
            left = new Binary(left, right, op);
        }
        return left;
    }

    private Expression ParseUnary()
    {
        if (_reader.Peek() == '!')
        {
            var op = _scanner.Token;
            var expression = ParseUnary();
            return new Unary(expression, op);
        }
        return ParsePrimary();
    }

    private Expression ParsePrimary()
    {
        if (_scanner.Match(Token.IDENTIFIER))
        {
            var name = _scanner.Literal;
            return new Identifier(name);
        }
        if (_scanner.Match(Token.LPAREN))
        {
            var expression = ParseExpression();
            _scanner.Consume(Token.RPAREN);
            return new Parentheses(expression);
        }
        throw new Exception("Invalid expression");
    }*/
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