namespace MicroPanda.AST;

using Declaration;

internal class Context
{
    internal Program Program { get; set; }
    internal Function? Function { get; set; }
    internal Context? Parent;
    private readonly Dictionary<string, Type.Type> objects = [];

    internal Context(Program program)
    {
        Program = program;
    }

    internal Context NewContext()
    {
        var context = new Context(Program)
        {
            Function = Function,
            Parent = this
        };
        return context;
    }

    internal bool AddObject(string name, Type.Type type)
    {
        if (objects.ContainsKey(name))
        {
            return false;
        }
        objects[name] = type;
        return true;
    }

    internal Type.Type? FindObject(string name)
    {
        if (objects.TryGetValue(name, out var type))
        {
            return type;
        }
        if (Parent != null)
        {
            type = Parent.FindObject(name);
            if (type != null)
            {
                return type;
            }
        }
        if (Function != null && Function.Parent != null)
        {
            return Function.Parent.MemberType(name);
        }
        return null;
    }

    public Type.Type? FindSelector(string selector, string member)
    {
        var parent = FindObject(selector);
        if (parent == null)
        {
            var declaration = Program.FindSelector(selector, member);
            if (declaration == null)
            {
                // could be an enum
                var e = Program.FindMember(selector);
                if (e is Enum ee && ee.HasMember(member))
                {
                    return Type.TypeHelper.TypeU8;
                }
                return null;
            }
            switch (declaration)
            {
                case Variable variable:
                    return variable.Type;

                case Function function:
                    return function.Type;

                case Enum _:
                    // enum itself is not a type, its member has u8 type
                    return null;
            }
        }
        else if (parent is Type.TypeName typeName)
        {
            var declaration = Program.FindByTypeName(typeName);
            if (declaration is Struct s)
            {
                return s.MemberType(member);
            }
        }
        return null;
    }
}
