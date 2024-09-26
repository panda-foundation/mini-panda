namespace MicroPanda.AST.Expression;

using Type;

internal class Sizeof : Expression
{
    internal Type? Target { get; set; }

    internal override void Validate(Context context, Type? expected)
    {
        Target = context.Program.ValidateType(Target!);
        _type = TypeHelper.TypeU32;
        _const = true;
    }
}