namespace MicroPandaTest.Scanner;

using MicroPanda.Scanner;
using MicroPanda.Token;

[TestClass]
public class ScannerTest
{
    [TestMethod]
    [DataRow("// single line comment\n", "COMMENT", "// single line comment")]
    [DataRow("/* multiline comment\n hello micro panda */\n", "COMMENT", "/* multiline comment\n hello micro panda */")]
    [DataRow("abc\n", "IDENT", "abc")]
    [DataRow("_abc123\n", "IDENT", "_abc123")]
    [DataRow("123\n", "INT", "123")]
    [DataRow("0x123\n", "INT", "0x123")]
    [DataRow("0b1001\n", "INT", "0b1001")]
    [DataRow("0o567\n", "INT", "0o567")]
    [DataRow("123.456\n", "FLOAT", "123.456")]
    [DataRow(".123\n", "FLOAT", ".123")]
    [DataRow("""'a'\n""", "CHAR", """'a'""")]
    [DataRow("""'\n'\n""", "CHAR", """'\n'""")]
    [DataRow("""'\a'\n""", "CHAR", """'\a'""")]
    [DataRow("""'\''\n""", "CHAR", """'\''""")]
    [DataRow("""'\"'\n""", "CHAR", """'\"'""")]
    [DataRow("""'\\'\n""", "CHAR", """'\\'""")]
    [DataRow(""" "你好"\n """, "STRING", "\"你好\"")]
    [DataRow("`hello`\n", "STRING", "`hello`")]
    [DataRow("break\n", "Break", "break")]
    [DataRow("+\n", "Plus", "+")]
    [DataRow(">>=\n", "RightShiftAssign", ">>=")]
    public void TestScanner(string source, string expectedToken, string expectedLiteral)
    {
        var bytes = System.Text.Encoding.UTF8.GetBytes(source);
        var scanner = new Scanner(new File("air-compile-source.mpd", bytes.Length), bytes);

        var (_, token, literal) = scanner.Scan();
        Assert.AreEqual(expectedToken, token.ToString());
        Assert.AreEqual(expectedLiteral, literal);
    }

    [TestMethod]
    [DataRow("'aa'", "illegal char literal")]
    [DataRow("'\na'", "char not terminated")]
    [DataRow("'a", "illegal char literal")]
    [DataRow("\"string not terminated", "string not terminated")]
    [DataRow("`raw string not terminated", "string not terminated")]
    [DataRow("123.", "float has no digits after .")]
    [DataRow("0x", "illegal number")]
    [DataRow("0x123.456", "invalid radix point")]
    [DataRow("00", "invalid integer")]
    [DataRow("/* comment", "comment not terminated")]
    public void TestScannerWithInvalidToken(string source, string expectedException)
    {
        var bytes = System.Text.Encoding.UTF8.GetBytes(source);
        var scanner = new Scanner(new File("air-compile-source.mpd", bytes.Length), bytes);

        var exception = Assert.ThrowsException<Exception>(() => scanner.Scan());
        Assert.IsTrue(exception.Message.Contains(expectedException));
    }
}