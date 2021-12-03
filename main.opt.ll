; ModuleID = '../mini-panda/main.ll'
source_filename = "../mini-panda/main.ll"

%global.Data = type { i8, float, [8 x i8] }

@global.Color.r = local_unnamed_addr global i8 0
@global.Color.g = local_unnamed_addr global i8 1
@global.Color.b = local_unnamed_addr global i8 2
@global.values = local_unnamed_addr global [5 x i8] c"\01\02\03\04\05"
@string.cb091131e20d7842e7627e8736856b45 = constant [12 x i8] c"hello world\00"
@string.17dc8feeedfe47c12c0d109e5e0da235 = constant [12 x i8] c"value: %d \0A\00"
@string.328de3303ca25a967f81f9c8c805e8a1 = constant [13 x i8] c"1 + 2 = %d \0A\00"
@string.f9b6d891c5ca674309c459ad55eb01c8 = constant [9 x i8] c"v1: %d \0A\00"
@string.f52c63a936a31e2b2d03c5c746b8d5b9 = constant [9 x i8] c"v2: %d \0A\00"
@string.da88a2e8a843d3d238dc43a4378c6887 = constant [9 x i8] c"v3: %u \0A\00"
@string.c3c4ff0f83dad5387535d315826c22f8 = constant [9 x i8] c"b1: %d \0A\00"
@string.e40a403ffbdf9c2c70921d6bb7739cd8 = constant [9 x i8] c"b2: %d \0A\00"
@string.2dcc97a590ca083991ffe9b43c08dd02 = constant [9 x i8] c"v4: %u \0A\00"
@string.fc631314303f7db146188786b60902e8 = constant [8 x i8] c"x: %d \0A\00"
@string.a753fba743a9e6b08cb7a2627f69b75d = constant [11 x i8] c"bool: %d \0A\00"
@string.9155af3e03234ca6017e6a626fa48d60 = constant [18 x i8] c"parentheses: %d \0A\00"
@string.6815af516458351e77683ead5f501317 = constant [8 x i8] c"c: %d \0A\00"
@string.d58ddb72e75f1acfc4203e33bddc08a1 = constant [8 x i8] c"b: %d \0A\00"
@string.5b8b2fafadbddfa000cd0e716725d4a4 = constant [8 x i8] c"i: %d \0A\00"
@string.dd42ef93dc06a72b063baa72848d660c = constant [10 x i8] c"f16: %f \0A\00"
@string.dc24ff6a55a1c588a346f9dff66c25a0 = constant [10 x i8] c"f32: %f \0A\00"
@string.7a828f7c003ac662930a932d14c84f48 = constant [10 x i8] c"f64: %f \0A\00"
@string.3aff445dea2b63e4d3b135c5219ba7dc = constant [12 x i8] c"some string\00"
@string.ccbd06f65fb69a974bb7bbe132352fd5 = constant [15 x i8] c"array[0]: %d \0A\00"
@string.502edb90c5d63a7982b92c4846005a12 = constant [15 x i8] c"array[3]: %d \0A\00"
@string.5b8a1afb98c4b2718e7e1f29b27539e6 = constant [19 x i8] c"data.integer: %d \0A\00"
@string.4703b4d82797dc9d0990618793a935a5 = constant [17 x i8] c"data.float: %f \0A\00"
@string.e839e54fd3fe1d952dd8a33030d97634 = constant [20 x i8] c"data.array[3]: %d \0A\00"
@string.7499c41e0f87337e5f3f93200f97701e = constant [21 x i8] c"new_data.value: %d \0A\00"
@string.5c65bb89388b87cc845b7ed6cc4e0933 = constant [28 x i8] c"new_data.data.integer: %d \0A\00"
@string.18fc68733fbf6df1ade57d0706714eec = constant [29 x i8] c"new_data.data.array[3]: %d \0A\00"
@string.07af74d61c4bcfd65e300c22c36df6a3 = constant [14 x i8] c"a(%d) >= 10 \0A\00"
@string.12625b519c0ef75b350a9963cafc3f42 = local_unnamed_addr constant [17 x i8] c"shouldn't happen\00"
@string.7c13f0ed550e89d5fe0dab15a8790a6b = constant [9 x i8] c"I'm else\00"
@string.e509c213bf338f03d246b720ec617c01 = constant [11 x i8] c"loop: %d \0A\00"
@string.ba86886fe05268c3936c4741a0d07a6e = local_unnamed_addr constant [14 x i8] c"switch case 0\00"
@string.162d9796d41e74535694f9688ea21a49 = constant [14 x i8] c"switch case 3\00"
@string.ba4ed99596c7e9aa2595a8f23577c2a9 = constant [16 x i8] c"values[0]: %d \0A\00"
@string.291bd270faa3b66dd92c4af584f01044 = constant [16 x i8] c"values[4]: %d \0A\00"
@string.bcfa829c5c86235c99443fb88b9d9699 = constant [15 x i8] c"array[2]: %d \0A\00"
@string.47b89087c0546b3ff5a4ec613cfd034c = constant [19 x i8] c"this.integer: %d \0A\00"

; Function Attrs: nofree nounwind
declare i32 @puts(i8* nocapture readonly) local_unnamed_addr #0

; Function Attrs: nofree nounwind
declare i32 @printf(i8* nocapture readonly, ...) local_unnamed_addr #0

; Function Attrs: nofree nounwind
define void @main() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([12 x i8], [12 x i8]* @string.cb091131e20d7842e7627e8736856b45, i64 0, i64 0)) #2
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([12 x i8], [12 x i8]* @string.17dc8feeedfe47c12c0d109e5e0da235, i64 0, i64 0), i32 1) #2
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([12 x i8], [12 x i8]* @string.17dc8feeedfe47c12c0d109e5e0da235, i64 0, i64 0), i32 2) #2
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([13 x i8], [13 x i8]* @string.328de3303ca25a967f81f9c8c805e8a1, i64 0, i64 0), i32 3) #2
  tail call void @global.expression()
  tail call void @global.statement()
  %4 = load i8, i8* getelementptr inbounds ([5 x i8], [5 x i8]* @global.values, i64 0, i64 0), align 1
  %5 = sext i8 %4 to i32
  %6 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.ba4ed99596c7e9aa2595a8f23577c2a9, i64 0, i64 0), i32 %5) #2
  %7 = load i8, i8* getelementptr inbounds ([5 x i8], [5 x i8]* @global.values, i64 0, i64 4), align 1
  %8 = sext i8 %7 to i32
  %9 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.291bd270faa3b66dd92c4af584f01044, i64 0, i64 0), i32 %8) #2
  %10 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.5b8a1afb98c4b2718e7e1f29b27539e6, i64 0, i64 0), i32 1) #2
  %11 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.4703b4d82797dc9d0990618793a935a5, i64 0, i64 0), double 0x40091EB860000000) #2
  %12 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([20 x i8], [20 x i8]* @string.e839e54fd3fe1d952dd8a33030d97634, i64 0, i64 0), i32 4) #2
  %13 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.47b89087c0546b3ff5a4ec613cfd034c, i64 0, i64 0), i32 1) #2
  %14 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.47b89087c0546b3ff5a4ec613cfd034c, i64 0, i64 0), i32 3) #2
  %15 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.bcfa829c5c86235c99443fb88b9d9699, i64 0, i64 0), i32 2) #2
  ret void
}

; Function Attrs: nofree nounwind
define void @global.extern() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([12 x i8], [12 x i8]* @string.cb091131e20d7842e7627e8736856b45, i64 0, i64 0))
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([12 x i8], [12 x i8]* @string.17dc8feeedfe47c12c0d109e5e0da235, i64 0, i64 0), i32 1)
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([12 x i8], [12 x i8]* @string.17dc8feeedfe47c12c0d109e5e0da235, i64 0, i64 0), i32 2)
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([13 x i8], [13 x i8]* @string.328de3303ca25a967f81f9c8c805e8a1, i64 0, i64 0), i32 3)
  ret void
}

; Function Attrs: norecurse nounwind readnone
define i32 @global.add(i32 %a, i32 %b) local_unnamed_addr #1 {
entry:
  %0 = add i32 %b, %a
  ret i32 %0
}

; Function Attrs: nofree nounwind
define void @global.expression() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.f9b6d891c5ca674309c459ad55eb01c8, i64 0, i64 0), i32 1) #2
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.f52c63a936a31e2b2d03c5c746b8d5b9, i64 0, i64 0), i32 -1) #2
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.da88a2e8a843d3d238dc43a4378c6887, i64 0, i64 0), i32 -2) #2
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.c3c4ff0f83dad5387535d315826c22f8, i64 0, i64 0), i32 -1) #2
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.e40a403ffbdf9c2c70921d6bb7739cd8, i64 0, i64 0), i32 0) #2
  %5 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.2dcc97a590ca083991ffe9b43c08dd02, i64 0, i64 0), i32 11) #2
  %6 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.2dcc97a590ca083991ffe9b43c08dd02, i64 0, i64 0), i32 19) #2
  tail call void @global.binary()
  %7 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([18 x i8], [18 x i8]* @string.9155af3e03234ca6017e6a626fa48d60, i64 0, i64 0), i32 33) #2
  %8 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.6815af516458351e77683ead5f501317, i64 0, i64 0), i32 97) #2
  %9 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.d58ddb72e75f1acfc4203e33bddc08a1, i64 0, i64 0), i32 -1) #2
  %10 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.5b8b2fafadbddfa000cd0e716725d4a4, i64 0, i64 0), i32 123) #2
  %11 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.dd42ef93dc06a72b063baa72848d660c, i64 0, i64 0), double 0x4009200000000000) #2
  %12 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.dc24ff6a55a1c588a346f9dff66c25a0, i64 0, i64 0), double 0x40091EB860000000) #2
  %13 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.7a828f7c003ac662930a932d14c84f48, i64 0, i64 0), double 3.140000e+00) #2
  %14 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([12 x i8], [12 x i8]* @string.3aff445dea2b63e4d3b135c5219ba7dc, i64 0, i64 0)) #2
  %15 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.ccbd06f65fb69a974bb7bbe132352fd5, i64 0, i64 0), i32 0) #2
  %16 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.502edb90c5d63a7982b92c4846005a12, i64 0, i64 0), i32 3) #2
  %17 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.5b8a1afb98c4b2718e7e1f29b27539e6, i64 0, i64 0), i32 0) #2
  %18 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.5b8a1afb98c4b2718e7e1f29b27539e6, i64 0, i64 0), i32 5) #2
  %19 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.4703b4d82797dc9d0990618793a935a5, i64 0, i64 0), double 0x40091EB860000000) #2
  %20 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([20 x i8], [20 x i8]* @string.e839e54fd3fe1d952dd8a33030d97634, i64 0, i64 0), i32 3) #2
  %21 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([21 x i8], [21 x i8]* @string.7499c41e0f87337e5f3f93200f97701e, i64 0, i64 0), i32 5) #2
  %22 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([28 x i8], [28 x i8]* @string.5c65bb89388b87cc845b7ed6cc4e0933, i64 0, i64 0), i32 8) #2
  %23 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([29 x i8], [29 x i8]* @string.18fc68733fbf6df1ade57d0706714eec, i64 0, i64 0), i32 9) #2
  ret void
}

; Function Attrs: nofree nounwind
define void @global.unary() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.f9b6d891c5ca674309c459ad55eb01c8, i64 0, i64 0), i32 1)
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.f52c63a936a31e2b2d03c5c746b8d5b9, i64 0, i64 0), i32 -1)
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.da88a2e8a843d3d238dc43a4378c6887, i64 0, i64 0), i32 -2)
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.c3c4ff0f83dad5387535d315826c22f8, i64 0, i64 0), i32 -1)
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.e40a403ffbdf9c2c70921d6bb7739cd8, i64 0, i64 0), i32 0)
  ret void
}

; Function Attrs: nofree nounwind
define void @global.increment_decrement() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.2dcc97a590ca083991ffe9b43c08dd02, i64 0, i64 0), i32 11)
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.2dcc97a590ca083991ffe9b43c08dd02, i64 0, i64 0), i32 19)
  ret void
}

; Function Attrs: nofree nounwind
define void @global.binary() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 5)
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 10)
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 9)
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 18)
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 2)
  %5 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 3)
  %6 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 12)
  %7 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 6)
  %8 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 15)
  %9 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 7)
  %10 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 6)
  %11 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 15)
  %12 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 7)
  %13 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 8)
  %14 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 -1)
  %15 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 0)
  %16 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 0)
  %17 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 -1)
  %18 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 0)
  %19 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 -1)
  %20 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 80)
  %21 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 5)
  %22 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 8)
  %23 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 2)
  %24 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 15)
  %25 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 1)
  %26 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i64 0, i64 0), i32 2)
  %27 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.a753fba743a9e6b08cb7a2627f69b75d, i64 0, i64 0), i32 -1)
  %28 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.a753fba743a9e6b08cb7a2627f69b75d, i64 0, i64 0), i32 0)
  ret void
}

; Function Attrs: nofree nounwind
define void @global.parentheses() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([18 x i8], [18 x i8]* @string.9155af3e03234ca6017e6a626fa48d60, i64 0, i64 0), i32 33)
  ret void
}

; Function Attrs: nofree nounwind
define void @global.literal() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.6815af516458351e77683ead5f501317, i64 0, i64 0), i32 97)
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.d58ddb72e75f1acfc4203e33bddc08a1, i64 0, i64 0), i32 -1)
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.5b8b2fafadbddfa000cd0e716725d4a4, i64 0, i64 0), i32 123)
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.dd42ef93dc06a72b063baa72848d660c, i64 0, i64 0), double 0x4009200000000000)
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.dc24ff6a55a1c588a346f9dff66c25a0, i64 0, i64 0), double 0x40091EB860000000)
  %5 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.7a828f7c003ac662930a932d14c84f48, i64 0, i64 0), double 3.140000e+00)
  %6 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([12 x i8], [12 x i8]* @string.3aff445dea2b63e4d3b135c5219ba7dc, i64 0, i64 0))
  ret void
}

; Function Attrs: nofree nounwind
define void @global.subscripting() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.ccbd06f65fb69a974bb7bbe132352fd5, i64 0, i64 0), i32 0)
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.502edb90c5d63a7982b92c4846005a12, i64 0, i64 0), i32 3)
  ret void
}

; Function Attrs: nofree nounwind
define void @global.member_access() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.5b8a1afb98c4b2718e7e1f29b27539e6, i64 0, i64 0), i32 0)
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.5b8a1afb98c4b2718e7e1f29b27539e6, i64 0, i64 0), i32 5)
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.4703b4d82797dc9d0990618793a935a5, i64 0, i64 0), double 0x40091EB860000000)
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([20 x i8], [20 x i8]* @string.e839e54fd3fe1d952dd8a33030d97634, i64 0, i64 0), i32 3)
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([21 x i8], [21 x i8]* @string.7499c41e0f87337e5f3f93200f97701e, i64 0, i64 0), i32 5)
  %5 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([28 x i8], [28 x i8]* @string.5c65bb89388b87cc845b7ed6cc4e0933, i64 0, i64 0), i32 8)
  %6 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([29 x i8], [29 x i8]* @string.18fc68733fbf6df1ade57d0706714eec, i64 0, i64 0), i32 9)
  ret void
}

; Function Attrs: nofree nounwind
define void @global.statement() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([14 x i8], [14 x i8]* @string.07af74d61c4bcfd65e300c22c36df6a3, i64 0, i64 0), i32 10)
  %1 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.7c13f0ed550e89d5fe0dab15a8790a6b, i64 0, i64 0))
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 0)
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 1)
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 2)
  %5 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 3)
  %6 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 4)
  %7 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 5)
  %8 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 6)
  %9 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 7)
  %10 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 8)
  %11 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 9)
  %12 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([14 x i8], [14 x i8]* @string.162d9796d41e74535694f9688ea21a49, i64 0, i64 0))
  ret void
}

; Function Attrs: norecurse nounwind readnone
define void @global.pointers() local_unnamed_addr #1 {
entry:
  ret void
}

; Function Attrs: norecurse nounwind readnone
define void @global.conversions() local_unnamed_addr #1 {
entry:
  ret void
}

; Function Attrs: nofree nounwind
define void @global.structs() local_unnamed_addr #0 {
entry:
  %0 = load i8, i8* getelementptr inbounds ([5 x i8], [5 x i8]* @global.values, i64 0, i64 0), align 1
  %1 = sext i8 %0 to i32
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.ba4ed99596c7e9aa2595a8f23577c2a9, i64 0, i64 0), i32 %1)
  %3 = load i8, i8* getelementptr inbounds ([5 x i8], [5 x i8]* @global.values, i64 0, i64 4), align 1
  %4 = sext i8 %3 to i32
  %5 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.291bd270faa3b66dd92c4af584f01044, i64 0, i64 0), i32 %4)
  %6 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.5b8a1afb98c4b2718e7e1f29b27539e6, i64 0, i64 0), i32 1)
  %7 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.4703b4d82797dc9d0990618793a935a5, i64 0, i64 0), double 0x40091EB860000000)
  %8 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([20 x i8], [20 x i8]* @string.e839e54fd3fe1d952dd8a33030d97634, i64 0, i64 0), i32 4)
  %9 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.47b89087c0546b3ff5a4ec613cfd034c, i64 0, i64 0), i32 1) #2
  %10 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.47b89087c0546b3ff5a4ec613cfd034c, i64 0, i64 0), i32 3) #2
  %11 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.bcfa829c5c86235c99443fb88b9d9699, i64 0, i64 0), i32 2) #2
  ret void
}

; Function Attrs: nofree nounwind
define void @global.call_print(%global.Data* nocapture %data) local_unnamed_addr #0 {
entry:
  %0 = getelementptr %global.Data, %global.Data* %data, i64 0, i32 0
  store i8 3, i8* %0, align 1
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.47b89087c0546b3ff5a4ec613cfd034c, i64 0, i64 0), i32 3) #2
  ret void
}

; Function Attrs: nofree nounwind
define void @global.call_array(i8* nocapture %data) local_unnamed_addr #0 {
entry:
  %0 = getelementptr i8, i8* %data, i64 2
  store i8 2, i8* %0, align 1
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.bcfa829c5c86235c99443fb88b9d9699, i64 0, i64 0), i32 2)
  ret void
}

; Function Attrs: norecurse nounwind readnone
define void @global.functions() local_unnamed_addr #1 {
entry:
  ret void
}

; Function Attrs: nofree nounwind
define void @global.Data.print_integer(%global.Data* nocapture readonly %this) local_unnamed_addr #0 {
entry:
  %0 = getelementptr %global.Data, %global.Data* %this, i64 0, i32 0
  %1 = load i8, i8* %0, align 1
  %2 = sext i8 %1 to i32
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.47b89087c0546b3ff5a4ec613cfd034c, i64 0, i64 0), i32 %2)
  ret void
}

attributes #0 = { nofree nounwind }
attributes #1 = { norecurse nounwind readnone }
attributes #2 = { nounwind }
