namespace MicroPanda.AST.Statement;

using Expression;
using Type;

internal class For : Statement
{
    internal Statement? Initialization { get; set; }
    internal Expression? Condition { get; set; }
    internal Statement? Post { get; set; }
    internal Statement? Body { get; set; }

    internal override void Validate(Context context)
    {
        var ctx = context.NewContext();
        Initialization?.Validate(ctx);
        if (Condition != null)
        {
            var conditionCtx = ctx.NewContext();
            Condition.Validate(conditionCtx, TypeHelper.TypeBool);
            if (Condition.Type != null && !Condition.Type.Equal(TypeHelper.TypeBool))
            {
                context.Program.Error(Condition.Position, "expect bool type condition");
            }
        }
        if (Post != null)
        {
            var postCtx = ctx.NewContext();
            Post.Validate(postCtx);
        }
        if (Body != null)
        {
            var bodyCtx = ctx.NewContext();
            Body.Validate(bodyCtx);
        }
    }
}