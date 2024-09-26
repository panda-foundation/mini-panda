namespace MicroPanda.AST;

using Token;

internal class Using
{
    internal Using(string ns)
    {
        Namespace = ns;
    }
    internal string Namespace { get; set; }
}

// Module is a collection of variables, functions, enums, and structs
// It corresponds to a single source file
internal class Module
{
    public File File { get; set; }
/*
    public string Namespace { get; set; }

    public List<Using> Usings { get; set; }

    public List<Attribute> Attributes { get; set; }
    public List<Variable> Variables { get; set; }
    public List<Function> Functions { get; set; }
    public List<Enum> Enums { get; set; }
    public List<Struct> Structs { get; set; }

    internal Module(File file)
    {
        File = file;
        Usings = new List<Using>();
        Attributes = new List<Attribute>();
        Variables = new List<Variable>();
        Functions = new List<Function>();
        Enums = new List<Enum>();
        Structs = new List<Struct>();
    }*/

    internal void ValidateType(Program program)
    {
        /*
        p.Module = this;
        var c = new Context(p);
        foreach (var v in Variables)
        {
            v.ValidateType(c);
        }
        foreach (var f in Functions)
        {
            f.ValidateType(c);
        }
        foreach (var e in Enums)
        {
            e.ValidateType(c);
        }
        foreach (var s in Structs)
        {
            s.ValidateType(c);
        }*/
    }

    internal void Validate(Program program)
    {
        /*
        p.Module = this;
        var c = new Context(p);
        foreach (var v in Variables)
        {
            v.Validate(c);
        }
        foreach (var f in Functions)
        {
            f.Validate(c);
        }
        foreach (var e in Enums)
        {
            e.Validate(c);
        }
        foreach (var s in Structs)
        {
            s.Validate(c);
        }*/
    }
}