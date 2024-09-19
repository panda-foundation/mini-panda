namespace MicroPanda.Scanner;

using Token;

internal partial class Scanner
{
    private string ScanComment()
    {
        _reader.CutIn(_reader.Offset - 1);
        if (_reader.Rune == '/')
        {
            // Single-line comment
            _reader.Read();
            while (_reader.Rune != '\n' && _reader.Rune >= 0)
            {
                _reader.Read();
            }
        }
        else
        {
            // Multi-line comment
            bool terminated = false;
            _reader.Read();
            while (_reader.Rune >= 0)
            {
                int charBefore = _reader.Rune;
                _reader.Read();
                if (charBefore == '*' && _reader.Rune == '/')
                {
                    _reader.Read();
                    terminated = true;
                    break;
                }
            }
            if (!terminated)
            {
                Error(_reader.CutFrom, "comment not terminated");
            }
        }
        return _reader.CutOut(_reader.Offset);
    }

    private string ScanIdentifier()
    {
        _reader.CutIn(_reader.Offset);
        while (RuneHelper.IsLetter(_reader.Rune) || RuneHelper.IsDecimal(_reader.Rune))
        {
            _reader.Read();
        }
        return _reader.CutOut(_reader.Offset);
    }

    private (Token, string) ScanNumber()
    {
        _reader.CutIn(_reader.Offset);
        Token token = Token.INT;

        if (_reader.Rune != '.')
        {
            if (_reader.Rune == '0')
            {
                _reader.Read();
                if (_reader.Rune != '.')
                {
                    int numberBase = 10;
                    switch (RuneHelper.Lower(_reader.Rune))
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
                            if (RuneHelper.IsDecimal(_reader.Rune))
                            {
                                Error(_reader.CutFrom, "invalid integer");
                                token = Token.ILLEGAL;
                            }
                            else
                            {
                                return (token, "0");
                            }
                            break;
                    }
                    if (token != Token.ILLEGAL)
                    {
                        _reader.Read();
                        BypassDigits(numberBase);
                        if (_reader.Offset - _reader.CutFrom <= 2)
                        {
                            token = Token.ILLEGAL;
                            Error(_reader.CutFrom, "illegal number");
                        }
                        if (_reader.Rune == '.')
                        {
                            token = Token.ILLEGAL;
                            Error(_reader.CutFrom, "invalid radix point");
                        }
                    }
                }
            }
            else
            {
                BypassDigits(10);
            }
        }

        if (_reader.Rune == '.')
        {
            int offsetFraction = _reader.Offset;
            token = Token.FLOAT;
            _reader.Read();
            BypassDigits(10);
            if (offsetFraction == _reader.Offset - 1)
            {
                token = Token.ILLEGAL;
                Error(_reader.CutFrom, "float has no digits after .");
            }
        }

        return (token, _reader.CutOut(_reader.Offset));
    }

    private void BypassDigits(int numberBase)
    {
        while (RuneHelper.DigitValue(_reader.Rune) < numberBase)
        {
            _reader.Read();
        }
    }

    private string ScanChar()
    {
        _reader.CutIn(_reader.Offset - 1);
        var rune = _reader.Rune;
        if (rune == '\n' || rune < 0)
        {
            Error(_reader.CutFrom, "char not terminated");
        }
        _reader.Read();
        if (rune == '\\')
        {
            BypassEscape();
        }
        if (_reader.Rune != '\'')
        {
            Error(_reader.CutFrom, "illegal char literal");
        }
        _reader.Read();
        return _reader.CutOut(_reader.Offset);
    }

    private string ScanString()
    {
        _reader.CutIn(_reader.Offset - 1);

        while (true)
        {
            var rune = _reader.Rune;
            if (rune == '\n' || rune < 0)
            {
                Error(_reader.CutFrom, "string not terminated");
            }
            _reader.Read();
            if (rune == '"')
            {
                break;
            }
            if (rune == '\\')
            {
                BypassEscape();
            }
        }

        return _reader.CutOut(_reader.Offset);
    }

    private void BypassEscape()
    {
        switch (_reader.Rune)
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
                _reader.Read();
                return;

            default:
                string message = "unknown escape sequence";
                if (_reader.Rune < 0)
                {
                    message = "escape sequence not terminated";
                }
                Error(_reader.Offset, message);
                return;
        }
    }

    private string ScanRawString()
    {
        _reader.CutIn(_reader.Offset - 1);

        while (true)
        {
            var rune = _reader.Rune;
            if (rune < 0)
            {
                Error(_reader.CutFrom, "string not terminated");
            }
            _reader.Read();
            if (rune == '`')
            {
                break;
            }
        }

        return _reader.CutOut(_reader.Offset);
    }

    private (Token, string) ScanOperators()
    {
        _reader.CutIn(_reader.Offset - 1);
        (Token t, int length) = TokenHelper.ReadOperator(_reader.Source, _reader.Offset - 1);
        string literal = string.Empty;

        if (length > 0)
        {
            for (int i = 1; i < length; i++)
            {
                _reader.Read();
            }
            literal = _reader.CutOut(_reader.Offset);
        }

        return (t, literal);
    }
}