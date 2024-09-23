namespace MicroPanda.AST.Expression;

public interface IExpression : INode
{
    bool IsConstant();
    Type Type();
    void Validate(Context c, Type expected);
}

public abstract class ExpressionBase : NodeBase, IExpression
{
    public bool Const { get; set; }
    public Type Typ { get; set; }

    public bool IsConstant()
    {
        return Const;
    }

    public Type Type()
    {
        return Typ;
    }

    public abstract void Validate(Context c, Type expected);
}