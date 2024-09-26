namespace MicroPanda.AST.Type;

using Node;
using Token;

internal abstract class Type : Node
{
    internal abstract bool Equal(Type type);
}

internal class TypeHelper
{
    internal static TypeBuiltin TypeBool = new(Token.Bool);
    internal static TypeBuiltin TypeU8 = new(Token.Uint8);
    internal static TypeBuiltin TypeU32 = new(Token.Uint32);
    internal static TypeBuiltin TypeI32 = new(Token.Int32);
    internal static TypeBuiltin TypeF16 = new(Token.Float16);
    internal static TypeBuiltin TypeF32 = new(Token.Float32);
    internal static Pointer TypePointer = new(TypeU8);

    internal static bool IsInteger(Type? type)
    {
        if (type is TypeBuiltin typeBuiltin)
        {
            return TokenHelper.IsInteger(typeBuiltin.Token);
        }
        return false;
    }

    internal static bool IsFloat(Type? type)
    {
        if (type is TypeBuiltin typeBuiltin)
        {
            return TokenHelper.IsFloat(typeBuiltin.Token);
        }
        return false;
    }

    internal static bool IsNumber(Type? type)
    {
        if (type is TypeBuiltin typeBuiltin)
        {
            return TokenHelper.IsNumber(typeBuiltin.Token);
        }
        return false;
    }

    internal static bool IsBool(Type? type)
    {
        if (type is TypeBuiltin typeBuiltin)
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
        return type is TypeFunction;
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

    internal static int TypeBuiltinBits(TypeBuiltin type)
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

    internal static int TypeBuiltinSize(TypeBuiltin type)
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
                var array = new Array(typeArray.ElementType, []);
                for (int i = 1; i < typeArray.Dimension.Count; i++)
                {
                    array.Dimension.Add(typeArray.Dimension[i]);
                }
                return array;
            }
        }
        return null;
    }
/*
    internal static Type ValidateType(Type v, Program p)
    {
        switch (v)
        {
            case TypeName t:
                var d = p.FindType(t);
                if (d == null)
                {
                    p.Error(v.GetPosition(), "type not defined");
                }
                else
                {
                    if (d is Function f)
                    {
                        return f.Type;
                    }
                    else if (d is Struct)
                    {
                        t.Qualified = d.QualifiedName();
                    }
                    else
                    {
                        p.Error(v.GetPosition(), "type not defined");
                    }
                }
                return t;

            case TypeArray t:
                t.ElementType = ValidateType(t.ElementType, p);
                if (t.Dimension[0] < 0)
                {
                    p.Error(v.GetPosition(), "invalid array index");
                }
                for (int i = 1; i < t.Dimension.Count; i++)
                {
                    if (t.Dimension[i] < 1)
                    {
                        p.Error(v.GetPosition(), "invalid array index");
                    }
                }
                return t;

            case TypePointer t:
                t.ElementType = ValidateType(t.ElementType, p);
                return t;

            case TypeFunction t:
                t.ReturnType = ValidateType(t.ReturnType, p);
                for (int i = 0; i < t.Parameters.Count; i++)
                {
                    t.Parameters[i] = ValidateType(t.Parameters[i], p);
                    if (TypeHelper.IsStruct(t.Parameters[i]))
                    {
                        p.Error(t.Parameters[i].GetPosition(), "struct is not allowed as parameter, use pointer instead");
                    }
                    if (TypeHelper.IsArray(t.Parameters[i]))
                    {
                        p.Error(t.Parameters[i].GetPosition(), "array is not allowed as parameter, use pointer instead");
                    }
                }
                return t;

            default:
                return v;
        }
    }*/
}