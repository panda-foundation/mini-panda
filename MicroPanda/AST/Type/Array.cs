namespace MicroPanda.AST.Type;

internal class Array : Type
{
    internal Type ElementType { get; set; }
    internal List<int> Dimension { get; set; }

    internal Array(Type elementType, List<int> dimension)
    {
        ElementType = elementType;
        Dimension = dimension;
    }

    override internal bool Equal(Type type)
    {
        if (type is Array typeArray)
        {
            if (Dimension.Count == typeArray.Dimension.Count)
            {
                for (int i = 1; i < Dimension.Count; i++)
                {
                    if (Dimension[i] != typeArray.Dimension[i])
                    {
                        return false;
                    }
                }
                return true;
            }
        }
        else if (type is Pointer typePointer)
        {
            if (Dimension.Count == 1)
            {
                return ElementType.Equal(typePointer.ElementType);
            }
        }
        return false;
    }
}