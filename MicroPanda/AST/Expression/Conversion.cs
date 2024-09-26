namespace MicroPanda.AST.Expression;

using Type;

internal class Conversion : Expression
{
    internal Expression? Value { get; set; }

    internal override void Validate(Context context, Type? expected)
    {
        _type = context.Program.ValidateType(_type!);
        Value!.Validate(context, _type!);
        _const = Value.Const;
        if (!(TypeHelper.IsNumber(_type!) && TypeHelper.IsNumber(Value.Type!)
        || TypeHelper.IsPointer(_type!) && TypeHelper.IsPointer(Value.Type!)))
        {
            context.Program.Error(Position, "invalid type conversion");
        }
    }
}