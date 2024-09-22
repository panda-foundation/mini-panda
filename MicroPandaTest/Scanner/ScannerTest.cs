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

    public const string Source = """
        0
        #if a
        1
        #elif b
        2
        #elif c
        3
        #else
        4
        #end

        #if a && b
        5
        #elif b || c
        6
        #else
        7
        #end

        #if a && b || c
        8
        #elif a || b && c
        9
        #else
        10
        #end

        #if a && b || c && d
        11
        #elif a || b && c || d
        12
        #else
        13
        #end

        #if a && (b || c)
        14
        #elif (a || b) && c
        15
        #else
        16
        #end

        #if a
        17
        #if b
        18
        #if c
        19
        #end
        20
        #end
        21
        #end
    """;
    [TestMethod]
    [DataRow(new string[] { "a" }, new int[] { 0, 1, 7, 9, 12, 16, 17, 21 })]
    [DataRow(new string[] { "a", "b", "c" }, new int[] { 0, 1, 5, 8, 11, 14, 17, 18, 19, 20, 21 })]
    [DataRow(new string[] { "b" }, new int[] { 0, 2, 6, 10, 13, 16 })]
    [DataRow(new string[] { "b", "c" }, new int[] { 0, 2, 6, 8, 12, 15 })]
    [DataRow(new string[] { "a", "b" }, new int[] { 0, 1, 5, 8, 11, 14, 17, 18, 20, 21 })]
    public void TestPreprossesor(string[] flags, int[] expected)
    {
        var bytes = System.Text.Encoding.UTF8.GetBytes(Source);
        var scanner = new Scanner(new File("air-compile-source.mpd", bytes.Length), bytes, new HashSet<string>(flags));

        var results = new List<int>();
        while (true)
        {
            var (_, token, literal) = scanner.Scan();
            if (token == Token.EOF)
            {
                break;
            }
            if (token == Token.INT)
            {
                results.Add(int.Parse(literal));
            }
        }

        CollectionAssert.AreEqual(expected, results);
    }

    //Invalid preprossor test
}