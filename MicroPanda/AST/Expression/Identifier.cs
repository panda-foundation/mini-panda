namespace MicroPanda.AST.Expression;

internal abstract class Identifier : Expression
{
    internal string? Name { get; set; }
    internal string? Qualified { get; set; }
    internal bool IsNamespace { get; set; }
}