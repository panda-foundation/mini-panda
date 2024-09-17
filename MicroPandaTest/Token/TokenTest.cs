namespace MicroPandaTest.Token;

using System.Text;
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
        foreach (var (key, value) in TokenUtil._string2Token)
        {
            Assert.IsTrue(TokenUtil._token2String.ContainsKey(value));
            Assert.AreEqual(key, TokenUtil._token2String[value]);
        }

        var result = TokenUtil.ReadOperator(GetSpan("++ 123"));
        Assert.AreEqual(Token.PlusPlus, result.Token);
        Assert.AreEqual(2, result.Length);

        result = TokenUtil.ReadOperator(GetSpan("<<= 123"));
        Assert.AreEqual(Token.LeftShiftAssign, result.Token);
        Assert.AreEqual(3, result.Length);

        result = TokenUtil.ReadOperator(GetSpan("> 123"));
        Assert.AreEqual(Token.Greater, result.Token);
        Assert.AreEqual(1, result.Length);

        result = TokenUtil.ReadOperator(GetSpan("123"));
        Assert.AreEqual(Token.ILLEGAL, result.Token);
        Assert.AreEqual(0, result.Length);

        Assert.AreEqual(1, TokenUtil.Precedence(Token.Assign));
        Assert.AreEqual(2, TokenUtil.Precedence(Token.Or));
    }

    private static Span<byte> GetSpan(string str)
    {
        byte[] byteArray = Encoding.UTF8.GetBytes(str);
        return new Span<byte>(byteArray);
    }
}
