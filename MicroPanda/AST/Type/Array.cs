namespace MicroPanda.AST.Type;

internal class Array : Type
{
    internal Type? ElementType { get; set; }

    // The elements are the size of that dimension
    // if the array is int[2][3] then Dimension = [2, 3]
    // of the size is 0, means pointer instead
    // Dimension[0] pointer to elemnt
    // Dimension[0][10] pointer to array of 10 elements
    internal List<int> Dimension = [];

    internal override bool Equal(Type type)
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
                return ElementType!.Equal(typePointer.ElementType);
            }
        }
        return false;
    }
}