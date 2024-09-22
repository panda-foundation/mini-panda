namespace MicroPanda.Scanner;

using Token;

internal partial class Scanner
{
    //#if #else #elif #end

    // ||
    // &&
    // == !=
    // !

    private const string PREPROSSESOR_IF = "#if";
    private const string PREPROSSESOR_ELIF = "#elif";
    private const string PREPROSSESOR_ELSE = "#else";
    private const string PREPROSSESOR_END = "#end";

    private readonly Stack<Preprocessor> _preprocessors = new();
    private (int offset, Token token, string literal) _preprocessorToken = new(0, Token.ILLEGAL, "");
}

internal class Preprocessor(string keyword, bool evaluated)
{
    internal string Keyword { get; set; } = keyword;
    internal bool Evaluated { get; set; } = evaluated;
}

internal partial class Scanner
{
    private void Preprocess()
    {
        var offset = _reader.CutIn();
        _reader.Consume();

        if (!RuneHelper.IsLetter(_reader.Peek()))
        {
            Error(offset, "Unexpected preprocessor");
        }
        var keyword = ScanIdentifier();
        if (keyword == PREPROSSESOR_IF)
        {
            var evaluated = ParseExpression().Evaluate(_flags);
            _preprocessors.Push(new Preprocessor(PREPROSSESOR_IF, evaluated));
            if (!evaluated)
            {
                SkipPreprocessor();
            }
        }
        else if (keyword == PREPROSSESOR_ELIF)
        {
            if (_preprocessors.Count == 0 || _preprocessors.Peek().Keyword == PREPROSSESOR_ELSE)
            {
                Error(offset, "Unexpected #elif");
            }
            else if (_preprocessors.Peek().Evaluated)
            {
                SkipPreprocessor();
            }
            else 
            {
                var evaluated = ParseExpression().Evaluate(_flags);
                _preprocessors.Peek().Evaluated = evaluated;
                if (!evaluated)
                {
                    SkipPreprocessor();
                }
            }
            _preprocessors.Peek().Keyword = PREPROSSESOR_ELIF;
        }
        else if (keyword == PREPROSSESOR_ELSE)
        {
            if (_preprocessors.Count == 0 || _preprocessors.Peek().Keyword == PREPROSSESOR_ELSE)
            {
                Error(offset, "Unexpected #else");
            }
            else if (_preprocessors.Peek().Evaluated)
            {
                SkipPreprocessor();
            }
            _preprocessors.Peek().Keyword = PREPROSSESOR_ELSE;
        }
        else if (keyword == PREPROSSESOR_END)
        {
            if (_preprocessors.Count == 0)
            {
                Error(offset, "Unexpected #end");
            }
            _preprocessors.Pop();
        }
        else
        {
            Error(offset, "Unexpected preprocessor");
        }
    }

    private void SkipPreprocessor()
    {
        var count = _preprocessors.Count;

        while (true)
        {
            while (_reader.Peek() >= 0 && _reader.Peek() != '#')
            {
                _reader.Consume();
            }
            if (_reader.Peek() < 0)
            {
                Error(_reader.CutIn(), "Preprocessor not terminated, expecting #end");
            }

            var offset = _reader.CutIn();
            _reader.Consume();
            if (!RuneHelper.IsLetter(_reader.Peek()))
            {
                Error(offset, "Unexpected preprocessor");
            }

            var keyword = ScanIdentifier();
            if (keyword == PREPROSSESOR_IF)
            {
                _preprocessors.Push(new Preprocessor(PREPROSSESOR_IF, false));
            }
            else if (keyword == PREPROSSESOR_ELIF)
            {
                if (_preprocessors.Count == count)
                {
                    _reader.Back(5);
                    break;
                }
                if (_preprocessors.Peek().Keyword == PREPROSSESOR_ELSE)
                {
                    Error(offset, "Unexpected #elif");
                }
                _preprocessors.Peek().Keyword = PREPROSSESOR_ELIF;
            }
            else if (keyword == PREPROSSESOR_ELSE)
            {
                if (_preprocessors.Count == count)
                {
                    _reader.Back(5);
                    break;
                }
                if (_preprocessors.Peek().Keyword == PREPROSSESOR_ELSE)
                {
                    Error(offset, "Unexpected #else");
                }
                _preprocessors.Peek().Keyword = PREPROSSESOR_ELSE;
            }
            else if (keyword == PREPROSSESOR_END)
            {
                if (_preprocessors.Count == count)
                {
                    _reader.Back(4);
                    break;
                }
                _preprocessors.Pop();
            }
            else
            {
                Error(offset, "Unexpected preprocessor");
            }
        }
    }
    
    private Expression ParseExpression()
    {
        _preprocessorToken = Scan();
        var expression = ParseOr();
        return expression;
    }

    private Expression ParseOr()
    {
        var left = ParseAnd();
        if (_preprocessorToken.token == Token.Or)
        {
            var op = _preprocessorToken.token;
            _preprocessorToken = Scan();
            var right = ParseAnd();
            left = new Binary(left, right, op);
        }
        return left;
    }

    private Expression ParseAnd()
    {
        var left = ParseEquality();
        if (_preprocessorToken.token == Token.And)
        {
            var op = _preprocessorToken.token;
            _preprocessorToken = Scan();
            var right = ParseEquality();
            left = new Binary(left, right, op);
        }
        return left;
    }

    private Expression ParseEquality()
    {
        var left = ParseUnary();
        if (_preprocessorToken.token == Token.Equal || _preprocessorToken.token == Token.NotEqual)
        {
            var op = _preprocessorToken.token;
            _preprocessorToken = Scan();
            var right = ParseUnary();
            left = new Binary(left, right, op);
        }
        return left;
    }

    private Expression ParseUnary()
    {
        if (_preprocessorToken.token == Token.Not)
        {
            var op = _preprocessorToken.token;
            _preprocessorToken = Scan();
            var expression = ParsePrimary();
            return new Unary(expression, op);
        }
        return ParsePrimary();
    }

    private Expression ParsePrimary()
    {
        if (_preprocessorToken.token == Token.IDENT)
        {
            var name = _preprocessorToken.literal;
            _preprocessorToken = Scan();
            return new Identifier(name);
        }
        if (_preprocessorToken.token == Token.LeftParen)
        {
            var expression = ParseExpression();
            if (_preprocessorToken.token != Token.RightParen)
            {
                Error(_preprocessorToken.offset, "Expecting ')'");
                return new Identifier("Invalid expression");
            }
            _preprocessorToken = Scan();
            return new Parentheses(expression);
        }
        Error(_preprocessorToken.offset, "Invalid expression");
        return new Identifier("Invalid expression");
    }
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