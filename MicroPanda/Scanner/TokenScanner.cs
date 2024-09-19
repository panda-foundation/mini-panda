namespace MicroPanda.Scanner;

using Token;

internal partial class Scanner
{
    internal string ScanComment()
    {
        _reader.SetCutPosition(_reader.Offset - 1);
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
                Error(_reader.CutPosition, "comment not terminated");
            }
        }
        return _reader.Cut();
    }

    internal string ScanIdentifier()
    {
        _reader.SetCutPosition(_reader.Offset);
        while (RuneHelper.IsLetter(_reader.Rune) || RuneHelper.IsDecimal(_reader.Rune))
        {
            _reader.Read();
        }
        return _reader.Cut();
    }

    public (Token, string) ScanNumber()
    {
        _reader.SetCutPosition(_reader.Offset);
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
                                Error(_reader.CutPosition, "invalid integer");
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
                        if (_reader.Offset - _reader.CutPosition <= 2)
                        {
                            token = Token.ILLEGAL;
                            Error(_reader.CutPosition, "illegal number");
                        }
                        if (_reader.Rune == '.')
                        {
                            token = Token.ILLEGAL;
                            Error(_reader.CutPosition, "invalid radix point");
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
                Error(_reader.CutPosition, "float has no digits after .");
            }
        }

        return (token, _reader.Cut());
    }

    private void BypassDigits(int numberBase)
    {
        while (RuneHelper.DigitValue(_reader.Rune) < numberBase)
        {
            _reader.Read();
        }
    }

    internal string ScanChar()
    {
        _reader.SetCutPosition(_reader.Offset - 1);
        var rune = _reader.Rune;
        if (rune == '\n' || rune < 0)
        {
            Error(_reader.CutPosition, "char literal not terminated");
        }
        _reader.Read();
        if (rune == '\\')
        {
            BypassEscape();
        }
        if (_reader.Rune != '\'')
        {
            Error(_reader.CutPosition, "illegal rune literal");
        }
        _reader.Read();
        return _reader.Cut();
    }

    internal string ScanString()
    {
        _reader.SetCutPosition(_reader.Offset - 1);

        while (true)
        {
            var rune = _reader.Rune;
            if (rune == '\n' || rune < 0)
            {
                Error(_reader.CutPosition, "string literal not terminated");
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

        return _reader.Cut();
    }

    private void BypassEscape()
    {
        int offset = _reader.Offset;
        int numberBase;
        int numberMax;
        int numberBytes;
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

            case 'u':
                _reader.Read();
                numberBytes = 4;
                numberBase = 16;
                numberMax = 0xFFFF;
                break;

            case 'U':
                _reader.Read();
                numberBytes = 8;
                numberBase = 16;
                numberMax = 0x10FFFF; // Unicode.MaxRune
                break;

            default:
                string message = "unknown escape sequence";
                if (_reader.Rune < 0)
                {
                    message = "escape sequence not terminated";
                }
                Error(offset, message);
                return;
        }

        int value = 0;
        while (numberBytes > 0)
        {
            int d = RuneHelper.DigitValue(_reader.Rune);
            if (d < 0 || d >= numberBase)
            {
                string message = $"illegal character {(char) _reader.Rune} in escape sequence";
                Error(offset, message);
                return;
            }
            value = value * numberBase + d;
            _reader.Read();
            numberBytes--;
        }

        if (value > numberMax)
        {
            Error(offset, "escape sequence is invalid Unicode code point");
        }
    }

    internal string ScanRawString()
    {
        _reader.SetCutPosition(_reader.Offset - 1);

        while (true)
        {
            var rune = _reader.Rune;
            if (rune < 0)
            {
                Error(_reader.CutPosition, "raw string literal not terminated");
            }
            _reader.Read();
            if (rune == '`')
            {
                break;
            }
        }

        return _reader.Cut();
    }

    internal (Token, string) ScanOperators()
    {
        _reader.SetCutPosition(_reader.Offset - 1);
        (Token t, int length) = TokenHelper.ReadOperator(_reader.Source, _reader.Offset);
        string literal = string.Empty;

        if (length > 0)
        {
            for (int i = 1; i < length; i++)
            {
                _reader.Read();
            }
            literal = _reader.Cut();
        }

        return (t, literal);
    }
}