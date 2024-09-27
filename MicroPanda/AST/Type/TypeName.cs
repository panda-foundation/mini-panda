namespace MicroPanda.AST.Type;

internal class TypeName : Type
{
    internal string? Name { get; set; }
    internal string? QualifiedName { get; set; }
    internal bool IsEnum { get; set; }
    internal string? Selector { get; set; }

    internal override bool Equal(Type type)
    {
        if (type is TypeName typeName)
        {
            return !string.IsNullOrEmpty(QualifiedName) && typeName.QualifiedName == QualifiedName;
        }
        return false;
    }
}