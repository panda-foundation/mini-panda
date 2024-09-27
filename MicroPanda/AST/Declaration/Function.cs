namespace MicroPanda.AST.Declaration;

using Expression;
using Node;
using Type;
using Statement;

internal class Function : Declaration
{
    internal List<Parameter> Parameters = [];
    internal Type? ReturnType { get; set; }
    internal Block? Body { get; set; }

    internal Struct? Parent { get; set; }
    internal TypeFunction? Type { get; set; }

    internal override void ValidateType(Context context)
    {
        Type!.ReturnType = ReturnType;
        if (HasAnnotation(Annotation.Extern))
        {
            Type.Extern = true;
        }
        else if (Body == null)
        {
            Type.TypeDefine = true;
        }
        if (Parent != null)
        {
            Type.MemberFunction = true;
            Type.Parameters.Add(Parent.PointerType());
        }
        foreach (var param in Parameters)
        {
            Type.Parameters.Add(param.Type!);
        }
        context.Program.ValidateType(Type);
    }

    internal override void Validate(Context context)
    {
        if (Body == null)
        {
            if (Parent != null)
            {
                context.Program.Error(Position, "function body is required for member function");
            }
            if (Type!.Extern)
            {
                Type.ExternName = GetAnnotationField(Annotation.Extern, Annotation.ExternName)!.AsString(context);
                if (string.IsNullOrEmpty(Type.ExternName))
                {
                    context.Program.Error(Position, "'name' of meta data is required for extern function");
                }
            }
        }
        else
        {
            var ctx = context.NewContext();
            ctx.Function = this;
            if (Parent != null)
            {
                var pointer = new Pointer
                {
                    ElementType = Parent.Type(),
                    Position = Parent.Position
                };
                ctx.AddObject(Program.This, pointer);
            }
            if (Type!.Extern)
            {
                context.Program.Error(Position, "extern function has no body");
            }
            foreach (var param in Parameters)
            {
                var added = ctx.AddObject(param.Name!, param.Type!);
                if (!added)
                {
                    context.Program.Error(param.Position, $"redeclared variable: {param.Name}");
                }
            }
            Body.Validate(ctx);
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