namespace MicroPanda.AST.Statement;

internal class Block : Statement
{
    internal List<Statement> Statements = [];

    internal override void Validate(Context context)
    {
        // TODO: warning: unreachable code //Start, End of block
        foreach (var statement in Statements)
        {
            var ctx = context;
            if (statement is Block)
            {
                ctx = context.NewContext();
            }
            statement.Validate(ctx);
        }
    }
}