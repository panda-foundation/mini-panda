namespace MicroPanda.AST.Declaration;

using Token;
using Type;
using Expression;

internal class Variable : Declaration
{
    private Type? _type;
    internal Type Type => _type!;
    private bool Const { get; set; }
    internal Token Token { get; set; }
    internal Expression? Value { get; set; }
    internal Struct? Parent { get; set; }

    internal override void ValidateType(Context context)
    {
        _type = context.Program.ValidateType(_type!);
    }

    internal override void Validate(Context context)
    {
        Value?.Validate(context, _type);
        if (Const)
        {
            if (Value == null)
            {
                context.Program.Error(Position, "const must be initialized when declared");
            }
            else if (!Value.Const)
            {
                context.Program.Error(Value.Position, "expect const expression");
            }
        }
        if (Value != null)
        {
            if (Value.Type == null)
            {
                context.Program.Error(Value.Position, "unknown type");
            }
            else if (!Value.Type.Equals(Type))
            {
                context.Program.Error(Value.Position, "init value type mismatch with define");
            }
        }
    }
}