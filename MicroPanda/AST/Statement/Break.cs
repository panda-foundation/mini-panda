namespace MicroPanda.AST.Statement;

internal class Break : Statement
{
    internal override void Validate(Context context)
    {
        // TO-DO add check
        // if (context.LeaveBlock == null)
        // {
        //     context.Program.Error(Position, "invalid break");
        // }
    }
}
