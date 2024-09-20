namespace MicroPandaTest.Scanner;

using MicroPanda.Scanner;
using MicroPanda.Token;

[TestClass]
public class ScannerTest
{
    [TestMethod]
    [DataRow("", "EOF", "")]
    [DataRow("@", "ANNOTATION", "@")]
    [DataRow("// single line\ncomment", "COMMENT", "// single line")]
    [DataRow("/* multiline comment\n hello micro panda */", "COMMENT", "/* multiline comment\n hello micro panda */")]
    [DataRow("abc", "IDENT", "abc")]
    [DataRow("\n", "NEWLINE", "\n")]
    [DataRow("_abc123", "IDENT", "_abc123")]
    [DataRow("123", "INT", "123")]
    [DataRow("0x123", "INT", "0x123")]
    [DataRow("0b1001", "INT", "0b1001")]
    [DataRow("0o567", "INT", "0o567")]
    [DataRow(".123", "FLOAT", ".123")]
    [DataRow("123.456", "FLOAT", "123.456")]
    [DataRow("""'a'""", "CHAR", """'a'""")]
    [DataRow("""'\n'""", "CHAR", """'\n'""")]
    [DataRow("""'\a'""", "CHAR", """'\a'""")]
    [DataRow("""'\''""", "CHAR", """'\''""")]
    [DataRow("""'\"'""", "CHAR", """'\"'""")]
    [DataRow("""'\\'""", "CHAR", """'\\'""")]
    [DataRow("\"你\\n好\"", "STRING", "\"你\\n好\"")]
    [DataRow("`hello\nworld`", "STRING", "`hello\nworld`")]
    [DataRow("break", "Break", "break")]
    [DataRow("*", "Mul", "*")]
    [DataRow("--", "MinusMinus", "--")]
    [DataRow("+123", "Plus", "+")]
    [DataRow("/123", "Div", "/")]
    [DataRow("++123", "PlusPlus", "++")]
    [DataRow(">>123", "RightShift", ">>")]
    [DataRow(">>=123", "RightShiftAssign", ">>=")]
    public void TestScanner(string source, string expectedToken, string expectedLiteral)
    {
        var bytes = System.Text.Encoding.UTF8.GetBytes(source);
        var scanner = new Scanner(new File("air-compile-source.mpd", bytes.Length), bytes);

        var (_, token, literal) = scanner.Scan();
        Assert.AreEqual(expectedToken, token.ToString());
        Assert.AreEqual(expectedLiteral, literal);
    }

    [TestMethod]
    [DataRow("'aa'", "illegal char")]
    [DataRow("'\na'", "char not terminated")]
    [DataRow("'a", "illegal char")]
    [DataRow("\"string not terminated", "string not terminated")]
    [DataRow("`raw string not terminated", "string not terminated")]
    [DataRow("123.", "illegal fraction")]
    [DataRow("0x", "illegal integer")]
    [DataRow("0x123.456", "illegal radix point")]
    [DataRow("00", "illegal integer")]
    [DataRow("/* comment", "comment not terminated")]
    [DataRow("你好\n", "invalid token")]
    public void TestScannerWithInvalidToken(string source, string expectedException)
    {
        var bytes = System.Text.Encoding.UTF8.GetBytes(source);
        var scanner = new Scanner(new File("air-compile-source.mpd", bytes.Length), bytes);

        var exception = Assert.ThrowsException<Exception>(() => scanner.Scan());
        Assert.IsTrue(exception.Message.Contains(expectedException));
    }
}