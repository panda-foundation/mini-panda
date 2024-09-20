namespace MicroPandaTest.Token;

using System.Text;
using MicroPanda.Token;

[TestClass]
public class TokenTest
{
    [TestMethod]
    public void TestTokenHelper()
    {
        Assert.IsTrue(TokenHelper.IsLiteral(Token.IDENT));
        Assert.IsTrue(TokenHelper.IsOperator(Token.Plus));
        Assert.IsTrue(TokenHelper.IsKeyword(Token.If));
        Assert.IsTrue(TokenHelper.IsScalar(Token.Bool));
        Assert.IsTrue(TokenHelper.IsInteger(Token.Int64));
        Assert.IsTrue(TokenHelper.IsFloat(Token.Float64));
        Assert.IsTrue(TokenHelper.IsNumber(Token.Uint64));
        Assert.IsTrue(TokenHelper.IsAssign(Token.Assign));

        Assert.AreEqual(Token.IDENT, TokenHelper.FromString("foo"));
        Assert.AreEqual(Token.BOOL, TokenHelper.FromString("true"));
        Assert.AreEqual(Token.BOOL, TokenHelper.FromString("false"));
        Assert.AreEqual(Token.NULL, TokenHelper.FromString("null"));

        Assert.AreEqual("identifier", TokenHelper.ToString(Token.IDENT));
        Assert.AreEqual("bool_literal", TokenHelper.ToString(Token.BOOL));
        Assert.AreEqual("null", TokenHelper.ToString(Token.NULL));

        Assert.AreEqual(1, TokenHelper.Precedence(Token.Assign));
        Assert.AreEqual(2, TokenHelper.Precedence(Token.Or));
    }
}
