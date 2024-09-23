namespace MicroPanda.AST.Type;

using Token;

internal class TypeBuiltin : Type
{
    internal Token Token { get; set; }

    internal TypeBuiltin(Token token)
    {
        Token = token;
    }
    
    override internal bool Equal(Type type)
    {
        if (type is TypeBuiltin typeBuiltin)
        {
            return typeBuiltin.Token == Token;
        }
        return false;
    }
}