namespace MicroPanda.AST.Type;

using Token;

internal class Builtin : Type
{
    internal Token Token { get; set; }

    internal Builtin(Token token)
    {
        Token = token;
    }
    
    internal override bool Equal(Type type)
    {
        if (type is Builtin typeBuiltin)
        {
            return typeBuiltin.Token == Token;
        }
        return false;
    }
}