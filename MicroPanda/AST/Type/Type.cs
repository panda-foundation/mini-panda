namespace MicroPanda.AST.Type;

using Node;
using Token;

internal abstract class Type : Node
{
    internal abstract bool Equal(Type type);
}

internal class TypeHelper
{
    internal static Builtin TypeBool = new(Token.Bool);
    internal static Builtin TypeU8 = new(Token.Uint8);
    internal static Builtin TypeU32 = new(Token.Uint32);
    internal static Builtin TypeI32 = new(Token.Int32);
    internal static Builtin TypeF16 = new(Token.Float16);
    internal static Builtin TypeF32 = new(Token.Float32);
    internal static Pointer TypePointer = new() { ElementType = TypeU8};

    internal static bool IsInteger(Type? type)
    {
        if (type is Builtin typeBuiltin)
        {
            return TokenHelper.IsInteger(typeBuiltin.Token);
        }
        return false;
    }

    internal static bool IsFloat(Type? type)
    {
        if (type is Builtin typeBuiltin)
        {
            return TokenHelper.IsFloat(typeBuiltin.Token);
        }
        return false;
    }

    internal static bool IsNumber(Type? type)
    {
        if (type is Builtin typeBuiltin)
        {
            return TokenHelper.IsNumber(typeBuiltin.Token);
        }
        return false;
    }

    internal static bool IsBool(Type? type)
    {
        if (type is Builtin typeBuiltin)
        {
            return typeBuiltin.Token == Token.Bool;
        }
        return false;
    }

    internal static bool IsStruct(Type? type)
    {
        if (type is TypeName typeName)
        {
            return !typeName.IsEnum;
        }
        return false;
    }

    internal static bool IsArray(Type? type)
    {
        if (type is Array typeArray)
        {
            return typeArray.Dimension.Count > 0 && typeArray.Dimension[0] != 0;
        }
        return false;
    }

    internal static bool IsFunction(Type? type)
    {
        return type is Function;
    }

    internal static bool IsPointer(Type? type)
    {
        if (type is Pointer)
        {
            return true;
        }
        if (type is Array typeArray)
        {
            return typeArray.Dimension.Count > 0 && typeArray.Dimension[0] == 0;
        }
        return false;
    }

    internal static int TypeBuiltinBits(Builtin type)
    {
        return type.Token switch
        {
            Token.Bool => 1,
            Token.Int8 or Token.Uint8 => 8,
            Token.Int16 or Token.Uint16 or Token.Float16 => 16,
            Token.Int32 or Token.Uint32 or Token.Float32 => 32,
            Token.Int64 or Token.Uint64 or Token.Float64 => 64,
            _ => 0,
        };
    }

    internal static int TypeBuiltinSize(Builtin type)
    {
        return type.Token switch
        {
            Token.Bool => 1,
            Token.Int8 or Token.Uint8 => 1,
            Token.Int16 or Token.Uint16 or Token.Float16 => 2,
            Token.Int32 or Token.Uint32 or Token.Float32 => 4,
            Token.Int64 or Token.Uint64 or Token.Float64 => 8,
            _ => 0,
        };
    }

    internal static Type? ElementType(Type? type)
    {
        if (type is Pointer typePointer)
        {
            return typePointer.ElementType;
        }
        if (type is Array typeArray && typeArray.Dimension.Count > 0 && typeArray.Dimension[0] == 0)
        {
            if (typeArray.Dimension.Count == 1)
            {
                return typeArray.ElementType;
            }
            else
            {
                var array = new Array(){ ElementType = typeArray.ElementType };
                for (int i = 1; i < typeArray.Dimension.Count; i++)
                {
                    array.Dimension.Add(typeArray.Dimension[i]);
                }
                return array;
            }
        }
        return null;
    }
}