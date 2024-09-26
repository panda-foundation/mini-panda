namespace MicroPanda.AST.Expression;

using System.Text;
using Token;
using Type;

internal class Literal : Expression
{
    internal Token Token { get; set; }
    internal string? Value { get; set; }

    internal override void Validate(Context context, Type? expected)
    {
        _const = true;
        switch (Token)
        {
            case Token.STRING:
                var length = Encoding.UTF8.GetByteCount(Value!) - 1;
                // char[] represents a string in c or llvm
                // end with a null character '\0' 
                // -2 for the quotes, then +1 for the null character
                var array = new Array
                {
                    ElementType = TypeHelper.TypeU8
                };
                array.Dimension.Add(length + 1);
                _type = array;
                break;

            case Token.CHAR:
                _type = TypeHelper.TypeU8;
                break;

            case Token.FLOAT:
                if (expected != null)
                {
                    if (TypeHelper.IsFloat(expected))
                    {
                        _type = expected;
                    }
                    else
                    {
                        context.Program.Error(Position, "type mismatch");
                    }
                }
                else
                {
                    _type = TypeHelper.TypeF32;
                }
                break;

            case Token.INT:
                if (expected != null)
                {
                    if (TypeHelper.IsNumber(expected))
                    {
                        _type = expected;
                    }
                    else
                    {
                        context.Program.Error(this.Position, "type mismatch");
                    }
                }
                else
                {
                    _type = TypeHelper.TypeI32;
                }
                break;

            case Token.BOOL:
                if (expected != null && !TypeHelper.IsBool(expected))
                {
                    context.Program.Error(Position, "type mismatch");
                }
                else
                {
                    _type = TypeHelper.TypeBool;
                }
                break;

            case Token.NULL:
                if (expected == null)
                {
                    context.Program.Error(Position, "expect type for 'null'");
                }
                else
                {
                    if (TypeHelper.IsPointer(expected))
                    {
                        _type = expected;
                    }
                    else
                    {
                        context.Program.Error(Position, "type mismatch");
                    }
                }
                break;
        }
    }

    internal int AsInt(Context context)
    {
        if (TokenHelper.IsInteger(Token))
        {
            return int.Parse(Value!);
        }
        else
        {
            context.Program.Error(Position, "type mismatch, expected int");
        }
        return 0;
    }

    internal float AsFloat(Context context)
    {
        if (TokenHelper.IsNumber(Token))
        {
            return float.Parse(Value!);
        }
        else
        {
            context.Program.Error(Position, "type mismatch, expected float");
        }
        return 0;
    }

    internal bool AsBool(Context context)
    {
        if (Token == Token.BOOL)
        {
            return bool.Parse(Value!);
        }
        else
        {
            context.Program.Error(Position, "type mismatch, expected bool");
        }
        return false;
    }

    internal string AsString(Context context)
    {
        if (Token == Token.STRING)
        {
            return Value!;
        }
        else
        {
            context.Program.Error(Position, "type mismatch, expected string");
        }
        return "";
    }
}