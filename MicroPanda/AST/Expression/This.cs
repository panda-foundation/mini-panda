namespace MicroPanda.AST.Expression;

using Type;

internal class This : Expression
{
    internal override void Validate(Context context, Type? expected)
    {
        _const = false;
        _type = context.FindObject(Program.This);
        if (_type == null)
        {
            context.Program.Error(Position, "undefined 'this'");
        }
    }
}