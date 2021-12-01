; ModuleID = '../mini-panda/main.ll'
source_filename = "../mini-panda/main.ll"

@global.Color.r = local_unnamed_addr global i8 0
@global.Color.g = local_unnamed_addr global i8 1
@global.Color.b = local_unnamed_addr global i8 2
@string.cb091131e20d7842e7627e8736856b45 = constant [12 x i8] c"hello world\00"

; Function Attrs: nofree nounwind
declare i32 @puts(i8* nocapture readonly) local_unnamed_addr #0

; Function Attrs: nofree nounwind
define void @main() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([12 x i8], [12 x i8]* @string.cb091131e20d7842e7627e8736856b45, i64 0, i64 0))
  ret void
}

attributes #0 = { nofree nounwind }
