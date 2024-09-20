namespace MicroPanda.Scanner;

using Microsoft.VisualBasic;
using Token;

internal class RuneReader
{
    internal const int EOF = -1;

    private readonly File _file;
    private readonly byte[] _source;
    private int _offset;
    private int _cutIn;

    internal RuneReader(File file, byte[] source)
    {
        _file = file;
        _source = source;
    }

    internal int Consume()
    {
        var rune = Peek(out var bytesConsumed);
        if (rune == '\n')
        {
             _file.AddLine(_offset);
        }
        _offset += bytesConsumed;
        return rune;
    }

    internal void Back()
    {
        Back(1);
    }

    internal void Back(int step)
    {
        // won't back an UTF-8 character, only ASCII
        if (_offset >= step)
        {
            _offset -= step;
        }
    }

    internal int Peek() => Peek(out _);

    internal int Peek(out int bytesConsumed)
    {
        if (_source.Length == 0 || _offset >= _source.Length)
        {
            bytesConsumed = 0;
            return EOF;
        }
        if (_source[_offset] < 0x80)
        {
            bytesConsumed = 1;
            return _source[_offset];
        }

        System.Text.Rune.DecodeFromUtf8(new ReadOnlySpan<byte>(_source)[_offset..], out var rune, out bytesConsumed);
        return rune.Value;
    }

    internal int CutIn()
    {
        _cutIn = _offset;
        return _cutIn;
    }

    internal string CutOut()
    {
        var cut = _source[_cutIn.._offset];
        return System.Text.Encoding.UTF8.GetString(cut);
    }
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