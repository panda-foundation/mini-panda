namespace MicroPanda.AST.Declaration;

using Type;

internal class Struct : Declaration
{
    internal List<Function> Functions = [];
    internal List<Variable> Variables = [];

    internal bool AddVariable(Variable variable)
    {
        if (IsRedeclared(variable.Name!.Name!))
        {
            return false;
        }
        variable.Parent = this;
        Variables.Add(variable);
        return true;
    }

    internal bool AddFunction(Function function)
    {
        if (IsRedeclared(function.Name!.Name!))
        {
            return false;
        }
        function.Parent = this;
        Functions.Add(function);
        return true;
    }

    internal bool IsRedeclared(string name)
    {
        foreach (var variable in Variables)
        {
            if (name == variable.Name!.Name!)
            {
                return true;
            }
        }
        foreach (var function in Functions)
        {
            if (name == function.Name!.Name!)
            {
                return true;
            }
        }
        return false;
    }

    internal bool HasMember(string member)
    {
        return IsRedeclared(member);
    }

    internal Type? MemberType(string member)
    {
        foreach (var variable in Variables)
        {
            if (member == variable.Name!.Name!)
            {
                return variable.Type;
            }
        }
        foreach (var function in Functions)
        {
            if (member == function.Name!.Name!)
            {
                return function.Type;
            }
        }
        return null;
    }

    internal TypeName Type()
    {
        return new TypeName
        {
            Name = Name!.Name!,
            QualifiedName = QualifiedName,
            IsEnum = false
        };
    }

    internal Pointer PointerType()
    {
        return new Pointer
        {
            ElementType = Type()
        };
    }

    internal override void ValidateType(Context context)
    {
        if (Variables.Count == 0)
        {
            context.Program.Error(Position, "struct should contain at least 1 variable member.");
        }
        foreach (var variable in Variables)
        {
            variable.ValidateType(context);
        }
        foreach (var function in Functions)
        {
            function.ValidateType(context);
            function.QualifiedName = $"{QualifiedName}.{function.Name!.Name!}";
        }
    }

    internal override void Validate(Context context)
    {
        foreach (var variable in Variables)
        {
            variable.Validate(context);
            if (variable.Value != null)
            {
                context.Program.Error(variable.Position, "struct member has no initialize value");
            }
        }
        foreach (var function in Functions)
        {
            function.Validate(context);
        }
    }
}