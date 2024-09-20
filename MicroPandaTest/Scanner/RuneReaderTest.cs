namespace MicroPandaTest.Scanner;

using MicroPanda.Scanner;
using MicroPanda.Token;

[TestClass]
public class RuneReaderTest
{
    [TestMethod]
    public void TestRuneReader()
    {
        var file = new File("Test", 11);
        var reader = new RuneReader(file, "abcde你好"u8.ToArray());

        Assert.AreEqual('a', reader.Peek());
        Assert.AreEqual('a', reader.Consume());
        Assert.AreEqual(1, reader.CutIn());
        Assert.AreEqual('b', reader.Peek());
        Assert.AreEqual('b', reader.Consume());
        Assert.AreEqual('c', reader.Consume());
        Assert.AreEqual('d', reader.Peek());
        reader.Back();
        Assert.AreEqual('c', reader.Peek());
        Assert.AreEqual('c', reader.Consume());
        Assert.AreEqual('d', reader.Consume());
        Assert.AreEqual("bcd", reader.CutOut());
        Assert.AreEqual('e', reader.Peek());
        Assert.AreEqual('e', reader.Consume());
        Assert.AreEqual('你', reader.Peek());
        Assert.AreEqual("bcde", reader.CutOut());
        Assert.AreEqual('你', reader.Consume());
        Assert.AreEqual('好', reader.Peek());
        Assert.AreEqual("bcde你", reader.CutOut());
        Assert.AreEqual('好', reader.Consume());
        Assert.AreEqual("bcde你好", reader.CutOut());
        Assert.AreEqual(RuneReader.EOF, reader.Peek());

        Assert.AreEqual(RuneReader.EOF, reader.Consume());
        Assert.AreEqual("bcde你好", reader.CutOut());
        Assert.AreEqual(RuneReader.EOF, reader.Peek());

        Assert.AreEqual(RuneReader.EOF, reader.Consume());
        Assert.AreEqual("bcde你好", reader.CutOut());
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
        Assert.IsTrue(RuneHelper.DigitValue('+') == 16);
    }
}