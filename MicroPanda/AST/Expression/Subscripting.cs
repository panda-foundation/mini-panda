namespace MicroPanda.AST.Expression;

using Type;

internal class Subscripting : Expression
{
    internal Expression? Parent { get; set; }
    internal List<Expression> Indexes = [];

    internal override void Validate(Context context, Type? expected)
    {
        _const = false;
        Parent!.Validate(context, null);
        if (Parent.Type is Array typeArray)
        {
            if (Indexes.Count == typeArray.Dimension.Count)
            {
                _type = typeArray.ElementType;
                foreach (var index in Indexes)
                {
                    index.Validate(context, null);
                    if (!TypeHelper.IsInteger(index.Type))
                    {
                        context.Program.Error(index.Position, "expect integer index for array");
                    }
                }
            }
            else if (Indexes.Count < typeArray.Dimension.Count)
            {
                var array = new Array
                {
                    ElementType = typeArray.ElementType,
                };
                array.Dimension.Add(0);
                foreach (var index in Indexes)
                {
                    index.Validate(context, null);
                    if (!TypeHelper.IsInteger(index.Type))
                    {
                        context.Program.Error(index.Position, "expect integer index for array");
                    }
                }
                for (int i = typeArray.Dimension.Count - Indexes.Count - 1; i > 0; i--)
                {
                    array.Dimension.Add(typeArray.Dimension[^i]);
                }
                _type = array;
            }
            else
            {
                context.Program.Error(Position, "mismatch array dimension");
            }
        }
        else
        {
            context.Program.Error(Position, "expect array type");
        }
    }
}