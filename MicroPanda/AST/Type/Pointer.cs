namespace MicroPanda.AST.Type;

internal class Pointer : Type
{
    internal Type ElementType { get; set; }

    internal Pointer(Type elementType)
    {
        ElementType = elementType;
    }

    override internal bool Equal(Type type)
    {
        if (type is Pointer typePointer)
        {
            return ElementType.Equal(typePointer.ElementType);
        }
        else if (type is Array typeArray)
        {
            if (typeArray.Dimension.Count == 1)
            {
                return this.ElementType.Equal(typeArray.ElementType);
            }
        }
        return false;
    }
}