namespace MicroPandaTest.Token;

using MicroPanda.Token;
using System.Collections.Generic;

[TestClass]
public class PositionTest
{
    [TestMethod]
    public void TestPosition()
    {
        var file = new File("test.mpd", 10);
        var position = new Position(file, 5);
        Assert.AreEqual("test.mpd:1:6", position.ToString());
        Assert.AreEqual(5, position.GlobalOffset());
    }

    [TestMethod]
    public void TestFile()
    {
        var file = new File("test.mpd", 20);
        file.AddLine(10);
        Assert.AreEqual("test.mpd", file.Name);
        Assert.AreEqual(20, file.Size);
        Assert.AreEqual(0, file.FileOffset);
        Assert.AreEqual((1, 1), file.GetLocation(0));
        Assert.AreEqual((1, 2), file.GetLocation(1));
        Assert.AreEqual((1, 10), file.GetLocation(9));
        Assert.AreEqual((2, 1), file.GetLocation(10));
        Assert.AreEqual((2, 2), file.GetLocation(11));
        Assert.AreEqual((2, 10), file.GetLocation(19));
    }

    [TestMethod]
    public void TestFileSet()
    {
        var fileSet = new FileSet();
        var file = fileSet.AddFile("test.mpd", 20);
        file.AddLine(10);
        var file2 = fileSet.AddFile("test2.mpd", 30);
        file2.AddLine(15);
        Assert.AreEqual(0, file.FileOffset);
        Assert.AreEqual(21, file2.FileOffset);
        Assert.AreEqual(file.Name, fileSet.GetFile(20)!.Name);
        Assert.AreEqual(file2.Name, fileSet.GetFile(21)!.Name);
        Assert.AreEqual("test.mpd:1:1", fileSet.GetPosition(0)!.ToString());
        Assert.AreEqual("test.mpd:2:10", fileSet.GetPosition(19)!.ToString());
        Assert.AreEqual("test2.mpd:1:1", fileSet.GetPosition(21)!.ToString());
        Assert.AreEqual("test2.mpd:2:16", fileSet.GetPosition(51)!.ToString());

        fileSet.UpdateFileSize("test.mpd", 30);
        Assert.AreEqual(30, file.Size);
        Assert.AreEqual("test2.mpd:1:2", fileSet.GetPosition(32)!.ToString());
    }
}