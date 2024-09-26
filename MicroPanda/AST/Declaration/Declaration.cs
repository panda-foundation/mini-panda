namespace MicroPanda.AST.Declaration;

using Node;
using Expression;

internal abstract class Declaration : Node
{
    internal List<Attribute> Attributes = [];
    internal bool IsPublic { get; set; }
    internal Identifier? Name { get; set; }
    internal string? QualifiedName { get; set; }
    internal abstract void ValidateType(Context context);
    internal abstract void Validate(Context context);

    internal bool HasAnnotation(string annotation)
    {
        foreach (var attribute in Attributes)
        {
            if (attribute.Name == annotation)
            {
                return true;
            }
        }
        return false;
    }

    internal bool HasAnnotationField(string annotation, string field)
    {
        foreach (var attribute in Attributes)
        {
            if (attribute.Name == annotation)
            {
                return attribute.Values.ContainsKey(field);
            }
        }
        return false;
    }

    internal Literal? GetAnnotationField(string annotation, string field)
    {
        foreach (var attribute in Attributes)
        {
            if (attribute.Name == annotation)
            {
                return attribute.Values[field];
            }
        }
        return null;
    }
}

// Annotation member types: token.INT, token.FLOAT, token.STRING, token.BOOL
internal class Attribute
{
    internal int Position { get; set; }
    internal string? Name { get; set; }
    internal Dictionary<string, Literal> Values = [];
}