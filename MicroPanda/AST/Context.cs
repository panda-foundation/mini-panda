namespace MicroPanda.AST;

using Declaration;

internal class Context
{
    public Program Program { get; set; }
    public Function Function { get; set; }
    //public bool Returned { get; set; }

    private Context parent;
    //private Dictionary<string, Type> objects;

    internal Context(Program program)
    {
        Program = program;
        //objects = new Dictionary<string, Type>();
    }

    internal Context NewContext()
    {
        return new Context(Program);
        /*
        {
            Function = this.Function,
            parent = this,
            objects = new Dictionary<string, Type>()
        };*/
    }

    internal bool AddObject(string name, Type.Type t)
    {
        return false;
        /*
        if (objects.ContainsKey(name))
        {
            return new Exception($"redeclared variable: {name}");
        }
        objects[name] = t;*/
    }

    internal Type.Type? FindObject(string name)
    {
        /*
        if (objects.TryGetValue(name, out var v))
        {
            return v;
        }
        if (parent != null)
        {
            v = parent.FindObject(name);
            if (v != null)
            {
                return v;
            }
        }
        if (Function != null && Function.Parent != null)
        {
            return Function.Parent.MemberType(name);
        }*/
        return null;
    }
/*
    public Type FindSelector(string selector, string member)
    {
        var parent = FindObject(selector);
        if (parent == null)
        {
            var (qualified, d) = Program.FindSelector(selector, member);
            if (d == null)
            {
                // could be an enum
                var (_, e) = Program.FindMember(selector);
                if (e is Enum ee && ee.HasMember(member))
                {
                    return TypeU8;
                }
                return null;
            }
            switch (d)
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
        else if (parent is TypeName t)
        {
            var d = Program.FindType(t);
            if (d is Struct s)
            {
                return s.MemberType(member);
            }
        }
        return null;
    }*/
}
