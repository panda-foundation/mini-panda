namespace MicroPanda.AST.Statement;

using Node;

internal abstract class Statement : Node
{
    internal abstract void Validate(Context context);
}