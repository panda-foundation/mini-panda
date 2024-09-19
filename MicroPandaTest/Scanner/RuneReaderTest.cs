namespace MicroPandaTest.Scanner;

using MicroPanda.Scanner;
using MicroPanda.Token;

[TestClass]
public class RuneReaderTest
{
    [TestMethod]
    public void TestRuneReader()
    {
        var file = new File("Test", 5);
        var reader = new RuneReader(file, "abcde你好"u8.ToArray());

        Assert.AreEqual(0, reader.Offset);
        Assert.AreEqual(RuneReader.EOF, reader.Rune);
        Assert.AreEqual('a', reader.Peek());

        reader.Read();
        reader.CutIn(reader.Offset);
        Assert.AreEqual(1, reader.Offset);
        Assert.AreEqual('a', reader.Rune);

        reader.Read();
        reader.Read();
        reader.Read();
        reader.Read();
        Assert.AreEqual(5, reader.Offset);
        Assert.AreEqual('e', reader.Rune);
        Assert.AreEqual('你', reader.Peek());
        Assert.AreEqual("bcde", reader.CutOut());

        reader.Read();
        Assert.AreEqual(8, reader.Offset);
        Assert.AreEqual('你', reader.Rune);
        Assert.AreEqual('好', reader.Peek());
        Assert.AreEqual("bcde你", reader.CutOut());

        reader.Read();
        Assert.AreEqual(11, reader.Offset);
        Assert.AreEqual('好', reader.Rune);
        Assert.AreEqual("bcde你好", reader.CutOut());
        Assert.AreEqual(RuneReader.EOF, reader.Peek());

        reader.Read();
        Assert.AreEqual(11, reader.Offset);
        Assert.AreEqual(RuneReader.EOF, reader.Rune);
        Assert.AreEqual(RuneReader.EOF, reader.Peek());

        reader.Read();
        Assert.AreEqual(11, reader.Offset);
        Assert.AreEqual(RuneReader.EOF, reader.Rune);
        Assert.AreEqual(RuneReader.EOF, reader.Peek());
    }

    [TestMethod]
    public void TestRuneHelper()
    {
        Assert.IsTrue(RuneHelper.IsLetter('a'));
        Assert.IsTrue(RuneHelper.IsLetter('z'));
        Assert.IsTrue(RuneHelper.IsLetter('A'));
        Assert.IsTrue(RuneHelper.IsLetter('Z'));
        Assert.IsFalse(RuneHelper.IsLetter('你'));
        Assert.IsFalse(RuneHelper.IsLetter('好'));
        Assert.IsFalse(RuneHelper.IsLetter('0'));
        Assert.IsFalse(RuneHelper.IsLetter('9'));
        Assert.IsFalse(RuneHelper.IsLetter(' '));
        Assert.IsFalse(RuneHelper.IsLetter('\t'));
        Assert.IsFalse(RuneHelper.IsLetter('\n'));
        Assert.IsFalse(RuneHelper.IsLetter('\r'));

        Assert.IsTrue(RuneHelper.IsDecimal('0'));
        Assert.IsTrue(RuneHelper.IsDecimal('9'));
        Assert.IsFalse(RuneHelper.IsDecimal('a'));
        Assert.IsFalse(RuneHelper.IsDecimal('z'));
        Assert.IsFalse(RuneHelper.IsDecimal('A'));
        Assert.IsFalse(RuneHelper.IsDecimal('Z'));
        Assert.IsFalse(RuneHelper.IsDecimal('你'));
        Assert.IsFalse(RuneHelper.IsDecimal('好'));
        Assert.IsFalse(RuneHelper.IsDecimal(' '));
        Assert.IsFalse(RuneHelper.IsDecimal('\t'));
        Assert.IsFalse(RuneHelper.IsDecimal('\n'));
        Assert.IsFalse(RuneHelper.IsDecimal('\r'));

        Assert.IsTrue(RuneHelper.Lower('A') == 'a');
        Assert.IsTrue(RuneHelper.Lower('Z') == 'z');
        Assert.IsTrue(RuneHelper.Lower('a') == 'a');
        Assert.IsTrue(RuneHelper.Lower('z') == 'z');

        Assert.IsTrue(RuneHelper.DigitValue('0') == 0);
        Assert.IsTrue(RuneHelper.DigitValue('9') == 9);
        Assert.IsTrue(RuneHelper.DigitValue('a') == 10);
        Assert.IsTrue(RuneHelper.DigitValue('f') == 15);
        Assert.ThrowsException<Exception>(() => RuneHelper.DigitValue('g'));
    }
}