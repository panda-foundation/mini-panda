namespace MicroPanda.AST;

using System;
using System.Collections.Generic;

internal class Program
{
    internal const string Global = "global";
    internal const string This = "this";
    internal const string FunctionEntry = "entry";
    internal const string FunctionBody = "body";
    internal const string FunctionExit = "exit";
    internal const string ProgramEntry = "global.main";

    private Dictionary<string, Module> _modules;
    private Module _module = null!;

    private Dictionary<string, Declaration.Declaration> _declarations;
    private List<Error> _errors;

    internal Program()
    {
        _modules = [];
        _declarations = [];
        _errors = [];
    }

    internal void Reset()
    {
        _modules = [];
        _declarations = [];
        _errors = [];
    }
/*
    public (string, Declaration) FindSelector(string selector, string member)
    {
        if (string.IsNullOrEmpty(selector))
        {
            return FindMember(member);
        }
        foreach (var i in Module.Usings)
        {
            if (i.Alias == selector)
            {
                var qualified = i.Namespace + "." + member;
                return (qualified, Declarations.ContainsKey(qualified) ? Declarations[qualified] : null);
            }
        }
        return (null, null);
    }

    public (string, Declaration) FindMember(string member)
    {
        var qualified = Module.Namespace + "." + member;
        if (Declarations.ContainsKey(qualified))
        {
            return (qualified, Declarations[qualified]);
        }
        qualified = Constants.Global + "." + member;
        if (Declarations.ContainsKey(qualified))
        {
            return (qualified, Declarations[qualified]);
        }
        return (null, null);
    }

    public Declaration FindType(TypeName t)
    {
        var (q, d) = FindSelector(t.Selector, t.Name);
        if (d is Enum)
        {
            t.IsEnum = true;
        }
        t.Qualified = q;
        return d;
    }

    internal Declaration FindQualified(string qualified)
    {
        return Declarations.ContainsKey(qualified) ? Declarations[qualified] : null;
    }*/

    internal bool IsNamespace(string name)
    {
        /*
        foreach (var i in Module.Usings)
        {
            if (i.Alias == name)
            {
                return true;
            }
        }*/
        return false;
    }

    internal void Validate()
    {
        foreach (var module in _modules.Values)
        {
            // TO-DO check if Using is valid // must be valid, cannot Using self, cannot duplicated
            //module.ValidateType(this);
        }
        foreach (var module in _modules.Values)
        {
            //module.Validate(this);
        }
    }

    internal Type.Type? ValidateType(Type.Type v)
    {
        // TODO
        return null;
    }

    internal void Error(int offset, string message)
    {
        //_errors.Add(new Error(Module.File.Position(offset), message));
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
