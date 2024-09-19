namespace MicroPanda.Scanner;

using Token;

internal partial class Scanner
{
    readonly File _file;
    readonly RuneReader _reader;

    internal Scanner(File file, byte[] source)
    {
        _file = file;
        _reader = new RuneReader(file, source);
        _reader.Consume();
    }

    internal (int offset, Token token, string literal) Scan()
    {
        while(_reader.Rune == ' ' || _reader.Rune == '\t' || _reader.Rune == '\n' || _reader.Rune == '\r')
        {
            _reader.Consume();
        }

        var offset = _reader.Offset;
        var literal = string.Empty;
        Token token;

        if (RuneHelper.IsLetter(_reader.Rune))
        {
            literal = ScanIdentifier();
            token = TokenHelper.FromString(literal);
        }
        else if (RuneHelper.IsDecimal(_reader.Rune) || (_reader.Rune == '.' && RuneHelper.IsDecimal(_reader.Peek())))
        {
            (token, literal) = ScanNumber();
        }
        else
        {
            var rune = _reader.Rune;
            _reader.Consume();
            switch (rune)
            {
                case RuneReader.EOF:
                    token = Token.EOF;
                    break;

                case '\'':
                    token = Token.CHAR;
                    literal = ScanChar();
                    break;

                case '"':
                    token = Token.STRING;
                    literal = ScanString();
                    break;

                case '`':
                    token = Token.STRING;
                    literal = ScanRawString();
                    break;

                case '/':
                    if (_reader.Rune == '/' || _reader.Rune == '*')
                    {
                        token = Token.COMMENT;
                        literal = ScanComment();
                        break;
                    }
                    (token, literal) = ScanOperators();
                    break;

                case '@':
                    token = Token.ANNOTATION;
                    literal = "@";
                    break;

                case ';':
                    token = Token.Semi;
                    literal = ";";
                    break;

                default:
                    (token, literal) = ScanOperators();
                    if (token == Token.ILLEGAL)
                    {
                        Error(offset, "invalid token");
                        _reader.Consume();
                    }
                    break;
            }
        }
        return (offset, token, literal);
    }

    private void Error(int offset, string message)
    {
        throw new Exception($"error: {_file.GetPosition(offset).ToString} {message}");
    }
}