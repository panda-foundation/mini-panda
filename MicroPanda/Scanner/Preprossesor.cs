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
        if (keyword == KEYWORD_IF)
        {
            var evaluated = ParseExpression().Evaluate(_flags);
            _preprocessors.Push(new Preprocessor(KEYWORD_IF, evaluated));
            if (!evaluated)
            {
                SkipPreprocessor();
            }
        }
        else if (keyword == KEYWORD_ELIF)
        {
            if (_preprocessors.Count == 0 || _preprocessors.Peek().Keyword == KEYWORD_ELSE)
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
            _preprocessors.Peek().Keyword = KEYWORD_ELIF;
        }
        else if (keyword == KEYWORD_ELSE)
        {
            if (_preprocessors.Count == 0 || _preprocessors.Peek().Keyword == KEYWORD_ELSE)
            {
                Error(offset, "Unexpected #else");
            }
            else if (_preprocessors.Peek().Evaluated)
            {
                SkipPreprocessor();
            }
            _preprocessors.Peek().Keyword = KEYWORD_ELSE;
        }
        else if (keyword == KEYWORD_END)
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
            if (keyword == KEYWORD_IF)
            {
                _preprocessors.Push(new Preprocessor(KEYWORD_IF, false));
            }
            else if (keyword == KEYWORD_ELIF)
            {
                if (_preprocessors.Count == count)
                {
                    _reader.Back(5);
                    break;
                }
                if (_preprocessors.Peek().Keyword == KEYWORD_ELSE)
                {
                    Error(offset, "Unexpected #elif");
                }
                _preprocessors.Peek().Keyword = KEYWORD_ELIF;
            }
            else if (keyword == KEYWORD_ELSE)
            {
                if (_preprocessors.Count == count)
                {
                    _reader.Back(5);
                    break;
                }
                if (_preprocessors.Peek().Keyword == KEYWORD_ELSE)
                {
                    Error(offset, "Unexpected #else");
                }
                _preprocessors.Peek().Keyword = KEYWORD_ELSE;
            }
            else if (keyword == KEYWORD_END)
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
    
    /*
    func (s *Scanner) ParseExpression() bool {
        for s.char == ' ' || s.char == '\t' {
            s.next()
        }

        if s.isLetter(s.char) {
            flag := s.scanIdentifier()

            for s.char == ' ' || s.char == '\t' || s.char == '\r' {
                s.next()
            }
            if s.char != '\n' {
                s.error(s.offset, "unexpected: "+string(s.char))
            }

            result := false
            if _, ok := s.flags[flag]; ok {
                result = true
            }
            return result
        }

        s.error(s.offset, "unexpected: "+string(s.char))
        return false
    }*/
    private Expression ParseExpression()
    {
        return new Identifier("dummy");
        //var expression = ParseOr();
        //return expression;
    }
/*
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