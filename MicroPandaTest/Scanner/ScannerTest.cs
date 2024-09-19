namespace MicroPandaTest.Scanner;

using MicroPanda.Scanner;
using MicroPanda.Token;

[TestClass]
public class ScannerTest
{
    [TestMethod]
    public void TestScanner()
    {
        var source = @"
// single line comment
/* multiline comment
 * hello micro panda
*/
a ab _abc _abc123 123abc"u8;
        var bytes = source.ToArray();

        var scanner = new Scanner(new File("air-compile-source.mpd", bytes.Length), bytes);
        while(true)
        {
            var (_, token, literal) = scanner.Scan();
            Console.WriteLine($"{token}:{literal}");
            if (token == Token.EOF)
            {
                break;
            }
        }
    }
}