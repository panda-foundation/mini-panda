namespace MicroPanda.AST.Expression;

using Node;
using Type;

internal abstract class Expression : Node
{
    protected bool _const;
    protected Type? _type;
    internal abstract void Validate(Context context, Type? expected);
    internal bool Const => _const;
    internal Type? Type => _type;
}