namespace MicroPanda.AST.Statement;

using Expression;
using Type;

internal class DeclarationStatement : Statement
{
    internal Identifier? Name { get; set; }
    internal Type? Type { get; set; }
    internal Expression? Value { get; set; }

    internal override void Validate(Context context)
    {
        Type = context.Program.ValidateType(Type!);
        if (Value != null)
        {
            Value.Validate(context, Type);
            if (Value.Type != null && Type != null && !Value.Type.Equal(Type))
            {
                context.Program.Error(Value.Position, "init value type mismatch with define");
            }
        }
        if (Type != null)
        {
            if(!context.AddObject(Name!.Name!, Type))
            {
                context.Program.Error(Name!.Position, $"redeclared variable: {Name.Name}");
            }
        }
    }
}