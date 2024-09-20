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
        var scanner = new Scanner(new File("air-compile-source.mpd", bytes.Length), bytes, []);

        var (_, token, literal) = scanner.Scan();
        Assert.AreEqual(expectedToken, token.ToString());
        Assert.AreEqual(expectedLiteral, literal);
    }

    [TestMethod]
    [DataRow("'aa'", "Illegal char")]
    [DataRow("'\na'", "Char not terminated")]
    [DataRow("'a", "Illegal char")]
    [DataRow("\"string not terminated", "String not terminated")]
    [DataRow("`raw string not terminated", "String not terminated")]
    [DataRow("123.", "Illegal fraction")]
    [DataRow("0x", "Illegal integer")]
    [DataRow("0x123.456", "Illegal radix point")]
    [DataRow("00", "Illegal integer")]
    [DataRow("/* comment", "Comment not terminated")]
    [DataRow("你好\n", "Invalid token")]
    public void TestScannerWithInvalidToken(string source, string expectedException)
    {
        var bytes = System.Text.Encoding.UTF8.GetBytes(source);
        var scanner = new Scanner(new File("air-compile-source.mpd", bytes.Length), bytes, []);

        var exception = Assert.ThrowsException<Exception>(() => scanner.Scan());
        Assert.IsTrue(exception.Message.Contains(expectedException));
    }

    [TestMethod]
    [DataRow("#hello", "Unexpected preprocessor")]
    public void TestScannerWithInvalidPreprossesor(string source, string expectedException)
    {
        var bytes = System.Text.Encoding.UTF8.GetBytes(source);
        var scanner = new Scanner(new File("air-compile-source.mpd", bytes.Length), bytes, []);

        var exception = Assert.ThrowsException<Exception>(() => scanner.Scan());
        Assert.IsTrue(exception.Message.Contains(expectedException));
    }
}