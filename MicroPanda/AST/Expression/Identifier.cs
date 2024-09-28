namespace MicroPanda.AST.Expression;

using Type;

internal abstract class Identifier : Expression
{
    internal string? Name { get; set; }
    internal string? Qualified { get; set; }
    internal bool IsNamespace { get; set; }

    internal override void Validate(Context context, Type? expected)
    {
        var type = context.FindObject(Name!);
        if (t == null)
        {
            var (found, d) = c.Program.FindMember(Name);
            if (d is Variable v)
            {
                Const = v.Const;
                Typ = v.Type;
                Qualified = d.QualifiedName();
            }
            else if (d is Function f)
            {
                Const = true;
                Typ = f.Type;
                Qualified = d.QualifiedName();
            }
            else if (d is Enum)
            {
                Qualified = d.QualifiedName();
            }
            else if (d == null)
            {
                IsNamespace = c.Program.IsNamespace(Name);
            }
        }
        else
        {
            Const = false;
            Typ = t;
        }

        // * type would be null for enum (its member has type u8)
        // * type is null when identifier is namespace
        if (Typ == null && string.IsNullOrEmpty(Qualified) && !IsNamespace)
        {
            c.Program.Error(Position, $"undefined {Name}");
        }
    }
}