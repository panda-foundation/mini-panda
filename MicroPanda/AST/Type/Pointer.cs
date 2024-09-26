namespace MicroPanda.AST.Type;

internal class Pointer : Type
{
    internal Type? ElementType { get; set; }

    internal override bool Equal(Type type)
    {
        if (type is Pointer typePointer)
        {
            return ElementType!.Equal(typePointer.ElementType!);
        }
        else if (type is Array typeArray)
        {
            if (typeArray.Dimension.Count == 1)
            {
                return ElementType!.Equal(typeArray.ElementType!);
            }
        }
        return false;
    }
}