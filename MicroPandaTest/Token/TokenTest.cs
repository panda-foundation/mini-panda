namespace MicroPandaTest.Token;

using MicroPanda.Token;

[TestClass]
public class TokenTest
{
    [TestMethod]
    public void TestTokenUtil()
    {
        Assert.IsTrue(TokenUtil.IsLiteral(Token.IDENT));
        Assert.IsTrue(TokenUtil.IsOperator(Token.Plus));
        Assert.IsTrue(TokenUtil.IsKeyword(Token.If));
        Assert.IsTrue(TokenUtil.IsScalar(Token.Bool));
        Assert.IsTrue(TokenUtil.IsInteger(Token.Int64));
        Assert.IsTrue(TokenUtil.IsFloat(Token.Float64));
        Assert.IsTrue(TokenUtil.IsNumber(Token.Uint64));
        Assert.IsTrue(TokenUtil.IsAssign(Token.Assign));

        Assert.AreEqual(Token.IDENT, TokenUtil.FromString("foo"));
        Assert.AreEqual(Token.BOOL, TokenUtil.FromString("true"));
        Assert.AreEqual(Token.BOOL, TokenUtil.FromString("false"));
        Assert.AreEqual(Token.NULL, TokenUtil.FromString("null"));

        Assert.AreEqual("identifier", TokenUtil.ToString(Token.IDENT));
        Assert.AreEqual("bool_literal", TokenUtil.ToString(Token.BOOL));
        Assert.AreEqual("null", TokenUtil.ToString(Token.NULL));

        Assert.AreEqual(TokenUtil._string2Token.Count, TokenUtil._token2String.Count);
    }
}
