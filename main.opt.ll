; ModuleID = '../mini-panda/main.ll'
source_filename = "../mini-panda/main.ll"

@test.Color.r = local_unnamed_addr global i8 0
@test.Color.g = local_unnamed_addr global i8 1
@test.Color.b = local_unnamed_addr global i8 2
@test.global_array = local_unnamed_addr constant [5 x i8] c"\01\02\03\04\05"
@string.4576fbff7ad2d9fa622f16573db7b286 = constant [42 x i8] c"============ test expression ============\00"
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
@string.ccbd06f65fb69a974bb7bbe132352fd5 = constant [15 x i8] c"array[0]: %d \0A\00"
@string.502edb90c5d63a7982b92c4846005a12 = constant [15 x i8] c"array[3]: %d \0A\00"
@string.d18cb31ff3a37014a9ed64a2687344d4 = constant [14 x i8] c"Color.g: %d \0A\00"
@string.9fcfb18ceb0d348e69c2e13fa41b241d = constant [22 x i8] c"global_array[2]: %d \0A\00"
@string.5f0f1578abd44713c746ded55bf898ea = constant [41 x i8] c"============ test statement ============\00"
@string.07af74d61c4bcfd65e300c22c36df6a3 = constant [14 x i8] c"a(%d) >= 10 \0A\00"
@string.12625b519c0ef75b350a9963cafc3f42 = local_unnamed_addr constant [17 x i8] c"shouldn't happen\00"
@string.7c13f0ed550e89d5fe0dab15a8790a6b = constant [9 x i8] c"I'm else\00"
@string.e509c213bf338f03d246b720ec617c01 = constant [11 x i8] c"loop: %d \0A\00"
@string.ba86886fe05268c3936c4741a0d07a6e = local_unnamed_addr constant [14 x i8] c"switch case 0\00"
@string.162d9796d41e74535694f9688ea21a49 = constant [14 x i8] c"switch case 3\00"

; Function Attrs: nofree nounwind
declare i32 @puts(i8* nocapture readonly) local_unnamed_addr #0

; Function Attrs: nofree nounwind
declare i32 @printf(i8* nocapture readonly, ...) local_unnamed_addr #0

; Function Attrs: nofree nounwind
define void @main() local_unnamed_addr #0 {
entry:
  tail call void @test.expression() #2
  tail call void @test.statement() #2
  ret void
}

; Function Attrs: norecurse nounwind readnone
define void @test.conversions() local_unnamed_addr #1 {
entry:
  ret void
}

; Function Attrs: nofree nounwind
define void @test.expression() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([42 x i8], [42 x i8]* @string.4576fbff7ad2d9fa622f16573db7b286, i64 0, i64 0))
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.f9b6d891c5ca674309c459ad55eb01c8, i64 0, i64 0), i32 1) #2
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.f52c63a936a31e2b2d03c5c746b8d5b9, i64 0, i64 0), i32 -1) #2
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.da88a2e8a843d3d238dc43a4378c6887, i64 0, i64 0), i32 -2) #2
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.c3c4ff0f83dad5387535d315826c22f8, i64 0, i64 0), i32 -1) #2
  %5 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.e40a403ffbdf9c2c70921d6bb7739cd8, i64 0, i64 0), i32 0) #2
  %6 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.2dcc97a590ca083991ffe9b43c08dd02, i64 0, i64 0), i32 11) #2
  %7 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.2dcc97a590ca083991ffe9b43c08dd02, i64 0, i64 0), i32 19) #2
  tail call void @test.binary()
  %8 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([18 x i8], [18 x i8]* @string.9155af3e03234ca6017e6a626fa48d60, i64 0, i64 0), i32 33) #2
  %9 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.6815af516458351e77683ead5f501317, i64 0, i64 0), i32 97) #2
  %10 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.d58ddb72e75f1acfc4203e33bddc08a1, i64 0, i64 0), i32 -1) #2
  %11 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.5b8b2fafadbddfa000cd0e716725d4a4, i64 0, i64 0), i32 123) #2
  %12 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.dd42ef93dc06a72b063baa72848d660c, i64 0, i64 0), double 0x4009200000000000) #2
  %13 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.dc24ff6a55a1c588a346f9dff66c25a0, i64 0, i64 0), double 0x40091EB860000000) #2
  %14 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.7a828f7c003ac662930a932d14c84f48, i64 0, i64 0), double 3.140000e+00) #2
  %15 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.ccbd06f65fb69a974bb7bbe132352fd5, i64 0, i64 0), i32 0) #2
  %16 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.502edb90c5d63a7982b92c4846005a12, i64 0, i64 0), i32 3) #2
  %17 = load i8, i8* @test.Color.g, align 1
  %18 = sext i8 %17 to i32
  %19 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([14 x i8], [14 x i8]* @string.d18cb31ff3a37014a9ed64a2687344d4, i64 0, i64 0), i32 %18) #2
  %20 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([22 x i8], [22 x i8]* @string.9fcfb18ceb0d348e69c2e13fa41b241d, i64 0, i64 0), i32 3) #2
  ret void
}

; Function Attrs: nofree nounwind
define void @test.unary() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.f9b6d891c5ca674309c459ad55eb01c8, i64 0, i64 0), i32 1)
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.f52c63a936a31e2b2d03c5c746b8d5b9, i64 0, i64 0), i32 -1)
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.da88a2e8a843d3d238dc43a4378c6887, i64 0, i64 0), i32 -2)
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.c3c4ff0f83dad5387535d315826c22f8, i64 0, i64 0), i32 -1)
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.e40a403ffbdf9c2c70921d6bb7739cd8, i64 0, i64 0), i32 0)
  ret void
}

; Function Attrs: nofree nounwind
define void @test.increment_decrement() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.2dcc97a590ca083991ffe9b43c08dd02, i64 0, i64 0), i32 11)
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.2dcc97a590ca083991ffe9b43c08dd02, i64 0, i64 0), i32 19)
  ret void
}

; Function Attrs: nofree nounwind
define void @test.binary() local_unnamed_addr #0 {
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
define void @test.parentheses() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([18 x i8], [18 x i8]* @string.9155af3e03234ca6017e6a626fa48d60, i64 0, i64 0), i32 33)
  ret void
}

; Function Attrs: nofree nounwind
define void @test.literal() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.6815af516458351e77683ead5f501317, i64 0, i64 0), i32 97)
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.d58ddb72e75f1acfc4203e33bddc08a1, i64 0, i64 0), i32 -1)
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([8 x i8], [8 x i8]* @string.5b8b2fafadbddfa000cd0e716725d4a4, i64 0, i64 0), i32 123)
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.dd42ef93dc06a72b063baa72848d660c, i64 0, i64 0), double 0x4009200000000000)
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.dc24ff6a55a1c588a346f9dff66c25a0, i64 0, i64 0), double 0x40091EB860000000)
  %5 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.7a828f7c003ac662930a932d14c84f48, i64 0, i64 0), double 3.140000e+00)
  ret void
}

; Function Attrs: nofree nounwind
define void @test.subscripting() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.ccbd06f65fb69a974bb7bbe132352fd5, i64 0, i64 0), i32 0)
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.502edb90c5d63a7982b92c4846005a12, i64 0, i64 0), i32 3)
  ret void
}

; Function Attrs: nofree nounwind
define void @test.member_access() local_unnamed_addr #0 {
entry:
  %0 = load i8, i8* @test.Color.g, align 1
  %1 = sext i8 %0 to i32
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([14 x i8], [14 x i8]* @string.d18cb31ff3a37014a9ed64a2687344d4, i64 0, i64 0), i32 %1)
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([22 x i8], [22 x i8]* @string.9fcfb18ceb0d348e69c2e13fa41b241d, i64 0, i64 0), i32 3)
  ret void
}

; Function Attrs: norecurse nounwind readnone
define void @test.functions() local_unnamed_addr #1 {
entry:
  ret void
}

; Function Attrs: norecurse nounwind readnone
define void @test.pointers() local_unnamed_addr #1 {
entry:
  ret void
}

; Function Attrs: nofree nounwind
define void @test.statement() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([41 x i8], [41 x i8]* @string.5f0f1578abd44713c746ded55bf898ea, i64 0, i64 0))
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([14 x i8], [14 x i8]* @string.07af74d61c4bcfd65e300c22c36df6a3, i64 0, i64 0), i32 10)
  %2 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.7c13f0ed550e89d5fe0dab15a8790a6b, i64 0, i64 0))
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 0)
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 1)
  %5 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 2)
  %6 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 3)
  %7 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 4)
  %8 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 5)
  %9 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 6)
  %10 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 7)
  %11 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 8)
  %12 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i64 0, i64 0), i32 9)
  %13 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([14 x i8], [14 x i8]* @string.162d9796d41e74535694f9688ea21a49, i64 0, i64 0))
  ret void
}

; Function Attrs: norecurse nounwind readnone
define void @test.structs() local_unnamed_addr #1 {
entry:
  ret void
}

; Function Attrs: nofree nounwind
define void @test.test() local_unnamed_addr #0 {
entry:
  tail call void @test.expression()
  tail call void @test.statement()
  ret void
}

attributes #0 = { nofree nounwind }
attributes #1 = { norecurse nounwind readnone }
attributes #2 = { nounwind }
