namespace MicroPanda.AST;

using Declaration;

internal class Program
{
    internal const string Global = "global";
    internal const string This = "this";
    internal const string FunctionEntry = "entry";
    internal const string FunctionBody = "body";
    internal const string FunctionExit = "exit";
    internal const string ProgramEntry = "global.main";

    private Dictionary<string, Module> _modules = [];
    private Dictionary<string, Declaration.Declaration> _declarations = [];
    private List<Error> _errors = [];
    private Module? _module;

    internal void Reset()
    {
        _modules = [];
        _declarations = [];
        _errors = [];
    }

    internal Declaration.Declaration? FindSelector(string? selector, string member)
    {
        if (string.IsNullOrEmpty(selector))
        {
            return FindMember(member);
        }
        foreach (var import in _module!.Imports)
        {
            if (import.Alias == selector)
            {
                var qualifiedName = $"{import.Package}.{member}";
                return _declarations[qualifiedName];
            }
        }
        return null;
    }

    internal Declaration.Declaration? FindMember(string member)
    {
        var qualifiedName = $"{_module!.Package}.{member}";
        if (_declarations.TryGetValue(qualifiedName, out Declaration.Declaration? value))
        {
            return value;
        }
        qualifiedName = $"{Program.Global}.{member}";
        return _declarations[qualifiedName];
    }

    internal Declaration.Declaration? FindByTypeName(Type.TypeName typeName)
    {
        var declaration = FindSelector(typeName.Selector, typeName.Name!);
        if (declaration is Enum)
        {
            typeName.IsEnum = true;
        }
        if (declaration != null)
        {
            typeName.QualifiedName = declaration!.QualifiedName;
        }
        return declaration;
    }

    internal Declaration.Declaration? FindByQualifiedName(string qualifiedName)
    {
        return _declarations[qualifiedName];
    }

    internal bool IsPackage(string package)
    {
        foreach (var import in _module!.Imports)
        {
            if (import.Alias == package)
            {
                return true;
            }
        }
        return false;
    }

    internal void Validate()
    {
        foreach (var module in _modules.Values)
        {
            // TO-DO check if Import is valid, cannot Import self, cannot duplicated
            _module = module;
            module.ValidateType(this);
        }
        foreach (var module in _modules.Values)
        {
            _module = module;
            module.Validate(this);
        }
    }

    internal Type.Type? ValidateType(Type.Type type)
    {
        switch (type)
        {
            case Type.TypeName typeName:
                var declaration = FindByTypeName(typeName);
                if (declaration == null)
                {
                    Error(type.Position, "type not defined");
                }
                else
                {
                    if (declaration is Function function)
                    {
                        return function.Type;
                    }
                    else if (declaration is Struct)
                    {
                        typeName.QualifiedName = declaration.QualifiedName;
                    }
                    else
                    {
                        Error(type.Position, "type not defined");
                    }
                }
                return typeName;

            case Type.Array array:
                array.ElementType = ValidateType(array.ElementType!);
                if (array.Dimension[0] < 0)
                {
                    Error(type.Position, "invalid array index");
                }
                for (int i = 1; i < array.Dimension.Count; i++)
                {
                    if (array.Dimension[i] < 1)
                    {
                        Error(type.Position, "invalid array index");
                    }
                }
                return array;

            case Type.Pointer pointer:
                pointer.ElementType = ValidateType(pointer.ElementType!);
                return pointer;

            case Type.Function function:
                function.ReturnType = ValidateType(function.ReturnType!);
                for (int i = 0; i < function.Parameters.Count; i++)
                {
                    function.Parameters[i] = ValidateType(function.Parameters[i])!;
                    if (Type.TypeHelper.IsStruct(function.Parameters[i]))
                    {
                        Error(function.Parameters[i].Position, "struct is not allowed as parameter, use pointer instead");
                    }
                    if (Type.TypeHelper.IsArray(function.Parameters[i]))
                    {
                        Error(function.Parameters[i].Position, "array is not allowed as parameter, use pointer instead");
                    }
                }
                return function;

            default:
                return type;
        }
    }

    internal void Error(int offset, string message)
    {
        _errors.Add(new Error(_module!.File.GetPosition(offset), message));
    }

    internal bool PrintErrors()
    {
        foreach (var error in _errors)
        {
            Console.WriteLine($"{error.Position} : {error.Message}");
        }
        return _errors.Count > 0;
    }
}

internal class Error
{
    public Token.Position Position { get; set; }
    public string Message { get; set; }

    internal Error(Token.Position position, string message)
    {
        Position = position;
        Message = message;
    }
}
