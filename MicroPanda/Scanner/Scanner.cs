namespace MicroPanda.Scanner;

using Token;

internal partial class Scanner
{
    readonly File _file;
    readonly RuneReader _reader;
    readonly HashSet<string> _flags = [];

    internal Scanner(File file, byte[] source, HashSet<string> flags)
    {
        _file = file;
        _reader = new RuneReader(file, source);
        _flags = flags;
    }

    internal (int offset, Token token, string literal) Scan()
    {
        var rune = _reader.Peek();
        while (rune == ' ' || rune == '\t' || rune == '\r')
        {
            _reader.Consume();
            rune = _reader.Peek();
        }

        var offset = _reader.CutIn();
        var literal = string.Empty;
        Token token;

        if (RuneHelper.IsLetter(rune))
        {
            literal = ScanIdentifier();
            token = TokenHelper.FromString(literal);
        }
        else if (RuneHelper.IsDecimal(rune))
        {
            (token, literal) = ScanNumber(offset);
        }
        else
        {
            switch (rune)
            {
                case RuneReader.EOF:
                    token = Token.EOF;
                    _reader.Consume();
                    break;

                case '\n':
                    token = Token.NEWLINE;
                    _reader.Consume();
                    literal = "\n";
                    break;

                case '\'':
                    token = Token.CHAR;
                    _reader.Consume();
                    literal = ScanChar(offset);
                    break;

                case '"':
                    token = Token.STRING;
                    _reader.Consume();
                    literal = ScanString(offset);
                    break;

                case '`':
                    token = Token.STRING;
                    _reader.Consume();
                    literal = ScanRawString(offset);
                    break;

                case '/':
                    _reader.Consume();
                    rune = _reader.Peek();
                    if (rune == '/' || rune == '*')
                    {
                        token = Token.COMMENT;
                        literal = ScanComment(offset);
                        break;
                    }
                    _reader.Back();
                    (token, literal) = ScanOperators();
                    break;

                case '@':
                    token = Token.ANNOTATION;
                    literal = "@";
                    _reader.Consume();
                    break;

                case '.':
                    _reader.Consume();
                    if (RuneHelper.IsDecimal(_reader.Peek()))
                    {
                        _reader.Back();
                        (token, literal) = ScanNumber(offset);
                        break;
                    }
                    token = Token.Dot;
                    literal = ".";
                    break;

                case '#':
                    Preprocess();
                    return Scan();

                default:
                    (token, literal) = ScanOperators();
                    if (token == Token.ILLEGAL)
                    {
                        Error(offset, "Invalid token");
                        return (offset, token, literal);
                    }
                    break;
            }
        }
        return (offset, token, literal);
    }

    private void Error(int offset, string message)
    {
        throw new Exception($"Error: {_file.GetPosition(offset).ToString} {message}");
    }
}