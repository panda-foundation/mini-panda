namespace MicroPanda.Scanner;

using Token;

internal partial class Scanner
{
    private string ScanComment(int offset)
    {
        var rune = _reader.Peek();
        if (rune == '/')
        {
            // Single-line comment
            _reader.Consume();
            rune = _reader.Peek();
            while (rune != '\n' && rune >= 0)
            {
                _reader.Consume();
                rune = _reader.Peek();
            }
        }
        else
        {
            // Multi-line comment
            bool terminated = false;
            _reader.Consume();
            rune = _reader.Peek();
            while (rune >= 0)
            {
                _reader.Consume();
                if (rune == '*' && _reader.Peek() == '/')
                {
                    _reader.Consume();
                    terminated = true;
                    break;
                }
                rune = _reader.Peek();
            }
            if (!terminated)
            {
                Error(offset, "Comment not terminated");
                return _reader.CutOut();
            }
        }
        return _reader.CutOut();
    }

    private string ScanIdentifier()
    {
        _reader.Consume();
        var rune = _reader.Peek();
        while (RuneHelper.IsLetter(rune) || RuneHelper.IsDecimal(rune))
        {
            _reader.Consume();
            rune = _reader.Peek();
        }
        return _reader.CutOut();
    }

    private (Token, string) ScanNumber(int offset)
    {
        Token token = Token.INT;
        var rune = _reader.Consume();
        if (rune != '.')
        {
            if (rune == '0')
            {
                rune = _reader.Consume();
                if (rune != '.')
                {
                    int numberBase;
                    switch (RuneHelper.Lower(rune))
                    {
                        case 'x':
                            numberBase = 16;
                            break;
                        case 'b':
                            numberBase = 2;
                            break;
                        case 'o':
                            //TODO omit o or convert to decimal if the compiler does not support it
                            numberBase = 8;
                            break;
                        default:
                            if (RuneHelper.IsDecimal(rune))
                            {
                                token = Token.ILLEGAL;
                                Error(offset, "Illegal integer");
                                return (token, _reader.CutOut());
                            }
                            return (token, "0");
                    }
                    
                    var digitLength = BypassDigits(numberBase);
                    if (digitLength == 0)
                    {
                        token = Token.ILLEGAL;
                        Error(offset, "Illegal integer");
                        return (token, _reader.CutOut());
                    }
                    rune = _reader.Peek();
                    if (rune == '.')
                    {
                        token = Token.ILLEGAL;
                        Error(offset, "Illegal radix point");
                        return (token, _reader.CutOut());
                    }
                }
            }
            else
            {
                BypassDigits(10);
                rune = _reader.Peek();
                if (rune == '.')
                {
                    _reader.Consume();
                }
            }
        }

        if (rune == '.')
        {
            token = Token.FLOAT;
            var fractionLength = BypassDigits(10);
            if (fractionLength == 0)
            {
                token = Token.ILLEGAL;
                Error(offset, "Illegal fraction");
                return (token, _reader.CutOut());
            }
        }

        return (token, _reader.CutOut());
    }

    private int BypassDigits(int numberBase)
    {
        var length = 0;
        var rune = _reader.Peek();
        while (RuneHelper.DigitValue(rune) < numberBase)
        {
            _reader.Consume();
            rune = _reader.Peek();
            length++;
        }
        return length;
    }

    private string ScanChar(int offset)
    {
        var rune = _reader.Peek();
        if (rune == '\n' || rune < 0)
        {
            Error(offset, "Char not terminated");
            return _reader.CutOut();
        }
        _reader.Consume();
        if (rune == '\\')
        {
            BypassEscape(offset);
        }
        rune = _reader.Peek();
        if (rune != '\'')
        {
            Error(offset, "Illegal char");
            return _reader.CutOut();
        }
        _reader.Consume();
        return _reader.CutOut();
    }

    private string ScanString(int offset)
    {
        while (true)
        {
            var rune = _reader.Peek();
            if (rune == '\n' || rune < 0)
            {
                Error(offset, "String not terminated");
                return _reader.CutOut();
            }
            _reader.Consume();
            rune = _reader.Peek();
            if (rune == '"')
            {
                _reader.Consume();
                break;
            }
            if (rune == '\\')
            {
                BypassEscape(offset);
            }
        }
        return _reader.CutOut();
    }

    private void BypassEscape(int offset)
    {
        var rune = _reader.Peek();
        switch (rune)
        {
            case '\'':
            case '\"':
            case '\\':
            case '0':
            case 'a':
            case 'b':
            case 'e':
            case 'f':
            case 'n':
            case 'r':
            case 't':
            case 'v':
                _reader.Consume();
                return;

            default:
                string message = "Unknown escape sequence";
                if (rune < 0)
                {
                    message = "Escape sequence not terminated";
                }
                Error(offset, message);
                return;
        }
    }

    private string ScanRawString(int offset)
    {
        var rune = _reader.Peek();
        while (true)
        {
            if (rune < 0)
            {
                Error(offset, "String not terminated");
                return _reader.CutOut();
            }
            if (rune == '`')
            {
                _reader.Consume();
                break;
            }
            _reader.Consume();
            rune = _reader.Peek();
        }
        return _reader.CutOut();
    }

    private (Token, string) ScanOperators()
    {
        var token = Token.ILLEGAL;
        if (_reader.Peek() < 0)
        {
            return (token, _reader.CutOut());
        }
        while(_reader.Peek() >= 0)
        {
            _reader.Consume();
            var literal = _reader.CutOut();

			if (TokenHelper.IsOperator(literal))
            {
                token = TokenHelper.FromString(literal);
            }
            else
            {
                _reader.Back();
                literal = _reader.CutOut();
                return (token, literal);
            }
        }

		return (token, _reader.CutOut());
    }
}