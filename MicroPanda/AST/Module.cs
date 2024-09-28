namespace MicroPanda.AST;

using Token;
using Declaration;

internal class Import
{
    internal string? Package { get; set; }
    internal string? Alias { get; set; }
}

// Module is a collection of variables, functions, enums, and structs
// It corresponds to a single source file
internal class Module
{
    public File File { get; set; }

    public string? Package { get; set; }

    public List<Import> Imports = [];
    public List<Attribute> Attributes = [];
    public List<Variable> Variables = [];
    public List<Function> Functions = [];
    public List<Enum> Enums = [];
    public List<Struct> Structs = [];

    internal Module(File file)
    {
        File = file;
    }

    internal void ValidateType(Program program)
    {
        var context = new Context(program);
        foreach (var variable in Variables)
        {
            variable.ValidateType(context);
        }
        foreach (var function in Functions)
        {
            function.ValidateType(context);
        }
        foreach (var e in Enums)
        {
            e.ValidateType(context);
        }
        foreach (var s in Structs)
        {
            s.ValidateType(context);
        }
    }

    internal void Validate(Program program)
    {
        var context = new Context(program);
        foreach (var variable in Variables)
        {
            variable.Validate(context);
        }
        foreach (var function in Functions)
        {
            function.Validate(context);
        }
        foreach (var e in Enums)
        {
            e.Validate(context);
        }
        foreach (var s in Structs)
        {
            s.Validate(context);
        }
    }
}