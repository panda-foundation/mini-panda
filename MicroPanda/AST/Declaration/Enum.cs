namespace MicroPanda.AST.Declaration;

using Expression;

internal class Enum : Declaration
{
    internal List<Variable> Members = [];
    internal List<byte> Values = [];

    internal bool AddMember(Variable member)
    {
        if (HasMember(member.Name!.Name!))
        {
            throw new Exception($"{member.Name!.Name!} redeclared");
        }
        Members.Add(member);
        return true;
    }

    internal bool HasMember(string name)
    {
        foreach (var member in Members)
        {
            if (member.Name!.Name! == name)
            {
                return true;
            }
        }
        return false;
    }

    internal override void ValidateType(Context context)
    {
    }

    internal override void Validate(Context context)
    {
        int index = 0;
        foreach (var member in Members)
        {
            if (index >= 256)
            {
                context.Program.Error(member.Position, "enum value should be less than 256");
            }
            if (member.Value == null)
            {
                Values.Add((byte)index);
                index++;
            }
            else
            {
                if (member.Value is Literal literal)
                {
                    var i = literal.AsInt(context);
                    if (i >= index)
                    {
                        index = i;
                        Values.Add((byte)index);
                        index++;
                    }
                    else
                    {
                        context.Program.Error(member.Position, $"enum value here should be greater than {i - 1}.");
                    }
                }
                else
                {
                    context.Program.Error(member.Position, "enum value must be const integer.");
                }
            }
        }
    }
}