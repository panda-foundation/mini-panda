namespace MicroPanda.AST.Declaration;

using Expression;
using Node;
using Type;
using Statement;

internal class Function : Declaration
{
    public List<Parameter> Parameters { get; set; } = new List<Parameter>();
    public Type ReturnType { get; set; }
    public Block Body { get; set; }

    public Struct Parent { get; set; }
    public TypeFunction Type { get; set; }

    public void ValidateType(Context c)
    {
        Type.ReturnType = ReturnType;
        if (HasAttribute("Extern"))
        {
            Type.Extern = true;
        }
        else if (Body == null)
        {
            Type.TypeDefine = true;
        }
        if (GetAttributeValue("Extern", "Variadic") is Literal l && l.Bool() is bool variadic)
        {
            Type.Variadic = variadic;
        }
        if (Parent != null)
        {
            Type.MemberFunction = true;
            Type.Parameters.Add(Parent.PointerType());
        }
        if (Parameters != null)
        {
            foreach (var param in Parameters)
            {
                Type.Parameters.Add(param.Type);
            }
        }
        ValidateType(Type, c.Program);
    }

    public void Validate(Context ctx)
    {
        if (Body == null)
        {
            if (Parent != null)
            {
                ctx.Program.Error(Position, "function body is required for member function");
            }
            if (Type.Extern)
            {
                if (GetAttributeValue("Extern", "Name") is Literal l && l.String() is string n)
                {
                    Type.ExternName = n;
                }
                if (string.IsNullOrEmpty(Type.ExternName))
                {
                    ctx.Program.Error(Position, "'name' of meta data is required for extern function");
                }
            }
        }
        else
        {
            var c = ctx.NewContext();
            c.Function = this;
            if (Parent != null)
            {
                var p = new TypePointer
                {
                    ElementType = Parent.Type(),
                    Position = Parent.Position
                };
                c.AddObject("StructThis", p);
                if (Type.Variadic)
                {
                    c.Program.Error(Position, "member function cannot be 'extern' or 'variadic'");
                }
            }
            if (Type.Extern)
            {
                c.Program.Error(Position, "extern function has no body");
            }
            if (Parameters != null)
            {
                foreach (var param in Parameters)
                {
                    var err = c.AddObject(param.Name, param.Type);
                    if (err != null)
                    {
                        c.Program.Error(param.Position, err.Message);
                    }
                }
            }
            Body.Validate(c);
        }
        // TODO check terminated
        // c.Program.Error(Position, "missing return");
    }
}

internal class Parameter : Node
{
    internal string? Name { get; set; }
    internal Type? Type { get; set; }
}

internal class Arguments : Node
{
    internal List<Expression> ArgumentsList = [];
}