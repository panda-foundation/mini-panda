namespace MicroPandaTest.Token;

using MicroPanda.Token;

[TestClass]
public class TokenTest
{
    [TestMethod]
    public void TestTokenUtil()
    {
        Assert.IsTrue(true);
        Assert.AreEqual((int)Token.ILLEGAL, 0);
    }
}
