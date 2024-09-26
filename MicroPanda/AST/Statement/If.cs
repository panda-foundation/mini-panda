namespace MicroPanda.AST.Statement;

using Expression;
using Type;

internal class If : Statement
{
    internal Statement? Initialization { get; set; }
    internal Expression? Condition { get; set; }
    internal Statement? Body { get; set; }
    internal Statement? Else { get; set; }

    internal override void Validate(Context context)
    {
        var ctx = context.NewContext();
        Initialization?.Validate(ctx);
        if (Condition == null)
        {
            context.Program.Error(Position, "expect condition expression");
        }
        else
        {
            Condition.Validate(ctx, TypeHelper.TypeBool);
            if (Condition.Type != null && !Condition.Type.Equal(TypeHelper.TypeBool))
            {
                context.Program.Error(Condition.Position, "expect bool type condition");
            }
        }
        if (Body != null)
        {
            var bodyCtx = ctx.NewContext();
            Body.Validate(bodyCtx);
        }
        if (Else != null)
        {
            var elseCtx = ctx.NewContext();
            Else.Validate(elseCtx);
        }
    }
}