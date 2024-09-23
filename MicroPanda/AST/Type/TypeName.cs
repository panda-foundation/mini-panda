namespace MicroPanda.AST.Type;

internal class TypeName : Type
{
    internal string Name { get; set; }
    internal string? Selector { get; set; }
    internal string Qualified { get; set; }
    internal bool IsEnum { get; set; }

    internal TypeName(string name, string qualified, bool isEnum)
    {
        Name = name;
        Qualified = qualified;
        IsEnum = isEnum;
    }

    override internal bool Equal(Type type)
    {
        if (type is TypeName typeName)
        {
            return !string.IsNullOrEmpty(Qualified) && typeName.Qualified == Qualified;
        }
        return false;
    }
}