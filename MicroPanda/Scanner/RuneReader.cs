namespace MicroPanda.Scanner;

using Token;

internal class RuneReader
{
    internal const int EOF = -1;

    private readonly File _file;
    private readonly byte[] _source;
    private int _offset = EOF;
    private int _readOffset;
    private int _cutIn;
    private int _rune = EOF;

    internal RuneReader(File file, byte[] source)
    {
        _file = file;
        _source = source;
    }

    internal int Consume()
    {
        if (_source.Length == 0)
        {
            _offset = 0;
            _rune = EOF;
            return _rune;
        }
        else if (_readOffset >= _source.Length)
        {
            _offset = _source.Length;
            if (_rune == '\n')
            {
                _file.AddLine(_offset);
            }
            _rune = EOF;
            return _rune;
        }

        _offset = _readOffset;
        if (_rune == '\n')
        {
             _file.AddLine(_offset);
        }
        System.Text.Rune.DecodeFromUtf8(new ReadOnlySpan<byte>(_source)[_readOffset..], out var rune, out var bytesConsumed);
        _readOffset += bytesConsumed;
        _rune = rune.Value;

        return _rune;
    }

    internal int Peek()
    {
        if (_source.Length == 0 || _readOffset >= _source.Length)
        {
            return EOF;
        }

        System.Text.Rune.DecodeFromUtf8(new ReadOnlySpan<byte>(_source)[_readOffset..], out var rune, out _);
        return rune.Value;
    }

    internal void CutIn(int offset)
    {
        _cutIn = offset;
    }

    internal string CutOut(int offset)
    {
        var cut = _source[_cutIn..offset];
        return System.Text.Encoding.UTF8.GetString(cut);
    }

    internal int Offset => _offset;
    internal int CutFrom => _cutIn;
    internal int Rune => _rune;
    internal byte[] Source => _source;
}

internal static class RuneHelper
{

    internal static bool IsLetter(int rune)
    {
        return rune == '_' || 'a' <= rune && rune <= 'z' || 'A' <= rune && rune <= 'Z';
    }

    internal static bool IsDecimal(int rune)
    {
        return '0' <= rune && rune <= '9';
    }

    internal static int Lower(int rune)
    {
        return ('a' - 'A') | rune;
    }

    internal static int DigitValue(int rune)
    {
        if ('0' <= rune && rune <= '9')
        {
            return rune - '0';
        }
        if ('a' <= Lower(rune) && Lower(rune) <= 'f')
        {
            return Lower(rune) - 'a' + 10;
        }
        return 16; // return invalid value for invalid digit
    }
}