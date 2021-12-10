; ModuleID = '../mini-panda/main.ll'
source_filename = "../mini-panda/main.ll"

%test.Printer = type { i32, [8 x i8], %test.Driver }
%test.Driver = type { i8 }

@test.Color.r = local_unnamed_addr constant i8 0
@test.Color.g = local_unnamed_addr constant i8 1
@test.Color.b = local_unnamed_addr constant i8 2
@test.global_array = local_unnamed_addr constant [5 x i8] c"\01\02\03\04\05"
@string.4576fbff7ad2d9fa622f16573db7b286 = constant [42 x i8] c"============ test expression ============\00"
@string.f9b6d891c5ca674309c459ad55eb01c8 = constant [9 x i8] c"v1: %d \0A\00"
@string.f52c63a936a31e2b2d03c5c746b8d5b9 = constant [9 x i8] c"v2: %d \0A\00"
@string.67c5acbbce37db2c12b92a427bc08a84 = constant [9 x i8] c"v3: %d \0A\00"
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
@string.f078cc2571d60d58e6a551d92df567c4 = constant [16 x i8] c"array2[0]: %d \0A\00"
@string.966d3dfa4a8527741a06bd9fbaa21f93 = constant [16 x i8] c"array2[3]: %d \0A\00"
@string.f9fe529bc21937b1ee14af38e842590b = constant [16 x i8] c"array3[0]: %d \0A\00"
@string.85ac8a2eb899f708bbb12d753db07868 = constant [16 x i8] c"array3[3]: %d \0A\00"
@string.8824f4dbad52bbb3684995480092775e = constant [16 x i8] c"array4[0]: %d \0A\00"
@string.5f0dc37317410eea89a43776ec4ac6e1 = constant [16 x i8] c"array4[3]: %d \0A\00"
@string.db08573c403e33d25bb325d1df98c844 = constant [19 x i8] c"new_array[0]: %d \0A\00"
@string.b1faf43818ca0e7e4b4b1e24b441f795 = constant [19 x i8] c"new_array[3]: %d \0A\00"
@string.d18cb31ff3a37014a9ed64a2687344d4 = constant [14 x i8] c"Color.g: %d \0A\00"
@string.9fcfb18ceb0d348e69c2e13fa41b241d = constant [22 x i8] c"global_array[2]: %d \0A\00"
@string.1daf4144552c4db57e99d55450ed346e = constant [18 x i8] c"sub.integer: %d \0A\00"
@string.560e3347d8fe3fd15f15ce5db418664f = constant [16 x i8] c"sub.float: %f \0A\00"
@string.b585a7adc3e8d68bbf60cb859044df1e = constant [19 x i8] c"sub.array[3]: %d \0A\00"
@string.84ad90c9c520f1a4e80779cfa15248b6 = constant [17 x i8] c"data.value: %d \0A\00"
@string.07ce14d972194d598243322dc9f50250 = constant [23 x i8] c"data.sub.integer: %d \0A\00"
@string.6db0fbcde59d77fa7fc3126dc45321f0 = constant [24 x i8] c"data.sub.array[3]: %d \0A\00"
@test.ff = local_unnamed_addr global void (i8*)* @test.do_something
@string.80c523c134f2b89c9ec7f6652a2dbdd7 = constant [40 x i8] c"============ test function ============\00"
@string.44083ed8ce984d51a6ecfdba2a6c2105 = constant [15 x i8] c"do something 1\00"
@string.b5b7eec21a3c4ab41dc70340c8ae1d93 = constant [15 x i8] c"do something 2\00"
@string.5f0f1578abd44713c746ded55bf898ea = constant [41 x i8] c"============ test statement ============\00"
@string.07af74d61c4bcfd65e300c22c36df6a3 = constant [14 x i8] c"a(%d) >= 10 \0A\00"
@string.12625b519c0ef75b350a9963cafc3f42 = local_unnamed_addr constant [17 x i8] c"shouldn't happen\00"
@string.7c13f0ed550e89d5fe0dab15a8790a6b = constant [9 x i8] c"I'm else\00"
@string.e509c213bf338f03d246b720ec617c01 = constant [11 x i8] c"loop: %d \0A\00"
@string.ba86886fe05268c3936c4741a0d07a6e = local_unnamed_addr constant [14 x i8] c"switch case 0\00"
@string.162d9796d41e74535694f9688ea21a49 = constant [14 x i8] c"switch case 3\00"
@string.ec374cb30dabe78ccd41f1bcfddac7db = constant [9 x i8] c"a1: %d \0A\00"
@test.global_printer = local_unnamed_addr constant %test.Printer { i32 80, [8 x i8] c"\01\02\03\04\05\06\07\08", %test.Driver { i8 88 } }
@string.91a35f7e30ee87849a8fb990c35dabf1 = constant [38 x i8] c"============ test struct ============\00"
@string.8c16759f16bae00294081efad1d55ec3 = constant [19 x i8] c"printer.line: %d \0A\00"
@string.c316f30584ee0ac304e8eed7e3af175f = constant [24 x i8] c"printer.buffer[7]: %d \0A\00"
@string.09e58fc876babc8908c9040bd77d8624 = constant [26 x i8] c"printer.driver.type: %d \0A\00"
@string.263c2d145bd0257bade41874fd5a73ec = constant [15 x i8] c"hello printer!\00"
@string.8c85cb3ae23186673c0ee88126a99c83 = constant [15 x i8] c"hello pointer!\00"
@string.6e04f1d448592af0a363c48cd79347e3 = constant [26 x i8] c"global_printer.line: %d \0A\00"
@string.569e8d7da8dcd242b4520ca536accffb = constant [31 x i8] c"global_printer.buffer[7]: %d \0A\00"
@string.e1297fae8db86112c4fd38cff8aca961 = constant [33 x i8] c"global_printer.driver.type: %d \0A\00"
@string.b5abd14716ff1d42a2c76d0bae14c3cf = constant [16 x i8] c"buffer[2]: %d \0A\00"
@string.f229d6156f4a2e6f6e5c4ee96406192b = constant [10 x i8] c"type:%d \0A\00"

; Function Attrs: nofree nounwind
declare i32 @puts(i8* nocapture readonly) local_unnamed_addr #0

; Function Attrs: nofree nounwind
declare i32 @printf(i8* nocapture readonly, ...) local_unnamed_addr #0

define void @main() local_unnamed_addr {
entry:
  tail call void @test.expression()
  tail call void @test.statement()
  tail call void @test.structs()
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([40 x i8], [40 x i8]* @string.80c523c134f2b89c9ec7f6652a2dbdd7, i64 0, i64 0))
  %1 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.44083ed8ce984d51a6ecfdba2a6c2105, i64 0, i64 0)) #2
  %2 = load void (i8*)*, void (i8*)** @test.ff, align 8
  tail call void %2(i8* getelementptr inbounds ([15 x i8], [15 x i8]* @string.b5b7eec21a3c4ab41dc70340c8ae1d93, i64 0, i64 0))
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
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.67c5acbbce37db2c12b92a427bc08a84, i64 0, i64 0), i32 -2) #2
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
  %15 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([12 x i8], [12 x i8]* @string.3aff445dea2b63e4d3b135c5219ba7dc, i64 0, i64 0)) #2
  tail call void @test.subscripting()
  %16 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([14 x i8], [14 x i8]* @string.d18cb31ff3a37014a9ed64a2687344d4, i64 0, i64 0), i32 1) #2
  %17 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([22 x i8], [22 x i8]* @string.9fcfb18ceb0d348e69c2e13fa41b241d, i64 0, i64 0), i32 3) #2
  %18 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([18 x i8], [18 x i8]* @string.1daf4144552c4db57e99d55450ed346e, i64 0, i64 0), i32 0) #2
  %19 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([18 x i8], [18 x i8]* @string.1daf4144552c4db57e99d55450ed346e, i64 0, i64 0), i32 5) #2
  %20 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.560e3347d8fe3fd15f15ce5db418664f, i64 0, i64 0), double 0x40091EB860000000) #2
  %21 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.b585a7adc3e8d68bbf60cb859044df1e, i64 0, i64 0), i32 3) #2
  %22 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.84ad90c9c520f1a4e80779cfa15248b6, i64 0, i64 0), i32 5) #2
  %23 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([23 x i8], [23 x i8]* @string.07ce14d972194d598243322dc9f50250, i64 0, i64 0), i32 8) #2
  %24 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([24 x i8], [24 x i8]* @string.6db0fbcde59d77fa7fc3126dc45321f0, i64 0, i64 0), i32 9) #2
  ret void
}

; Function Attrs: nofree nounwind
define void @test.unary() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.f9b6d891c5ca674309c459ad55eb01c8, i64 0, i64 0), i32 1)
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.f52c63a936a31e2b2d03c5c746b8d5b9, i64 0, i64 0), i32 -1)
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.67c5acbbce37db2c12b92a427bc08a84, i64 0, i64 0), i32 -2)
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
  %6 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([12 x i8], [12 x i8]* @string.3aff445dea2b63e4d3b135c5219ba7dc, i64 0, i64 0))
  ret void
}

; Function Attrs: nofree nounwind
define void @test.subscripting() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.ccbd06f65fb69a974bb7bbe132352fd5, i64 0, i64 0), i32 0)
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.502edb90c5d63a7982b92c4846005a12, i64 0, i64 0), i32 3)
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.f078cc2571d60d58e6a551d92df567c4, i64 0, i64 0), i32 0)
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.966d3dfa4a8527741a06bd9fbaa21f93, i64 0, i64 0), i32 3)
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.f9fe529bc21937b1ee14af38e842590b, i64 0, i64 0), i32 0)
  %5 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.85ac8a2eb899f708bbb12d753db07868, i64 0, i64 0), i32 3)
  %6 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.8824f4dbad52bbb3684995480092775e, i64 0, i64 0), i32 0)
  %7 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.5f0dc37317410eea89a43776ec4ac6e1, i64 0, i64 0), i32 3)
  %8 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.db08573c403e33d25bb325d1df98c844, i64 0, i64 0), i32 1)
  %9 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.b1faf43818ca0e7e4b4b1e24b441f795, i64 0, i64 0), i32 4)
  ret void
}

; Function Attrs: nofree nounwind
define void @test.member_access() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([14 x i8], [14 x i8]* @string.d18cb31ff3a37014a9ed64a2687344d4, i64 0, i64 0), i32 1)
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([22 x i8], [22 x i8]* @string.9fcfb18ceb0d348e69c2e13fa41b241d, i64 0, i64 0), i32 3)
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([18 x i8], [18 x i8]* @string.1daf4144552c4db57e99d55450ed346e, i64 0, i64 0), i32 0)
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([18 x i8], [18 x i8]* @string.1daf4144552c4db57e99d55450ed346e, i64 0, i64 0), i32 5)
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.560e3347d8fe3fd15f15ce5db418664f, i64 0, i64 0), double 0x40091EB860000000)
  %5 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.b585a7adc3e8d68bbf60cb859044df1e, i64 0, i64 0), i32 3)
  %6 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.84ad90c9c520f1a4e80779cfa15248b6, i64 0, i64 0), i32 5)
  %7 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([23 x i8], [23 x i8]* @string.07ce14d972194d598243322dc9f50250, i64 0, i64 0), i32 8)
  %8 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([24 x i8], [24 x i8]* @string.6db0fbcde59d77fa7fc3126dc45321f0, i64 0, i64 0), i32 9)
  ret void
}

define void @test.functions() local_unnamed_addr {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([40 x i8], [40 x i8]* @string.80c523c134f2b89c9ec7f6652a2dbdd7, i64 0, i64 0))
  %1 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.44083ed8ce984d51a6ecfdba2a6c2105, i64 0, i64 0)) #2
  %2 = load void (i8*)*, void (i8*)** @test.ff, align 8
  tail call void %2(i8* getelementptr inbounds ([15 x i8], [15 x i8]* @string.b5b7eec21a3c4ab41dc70340c8ae1d93, i64 0, i64 0))
  ret void
}

; Function Attrs: nofree nounwind
define void @test.do_something(i8* nocapture readonly %msg) #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) %msg)
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
  %14 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.ec374cb30dabe78ccd41f1bcfddac7db, i64 0, i64 0), i32 123)
  ret void
}

; Function Attrs: nofree nounwind
define void @test.structs() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([38 x i8], [38 x i8]* @string.91a35f7e30ee87849a8fb990c35dabf1, i64 0, i64 0))
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.8c16759f16bae00294081efad1d55ec3, i64 0, i64 0), i32 100)
  %2 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([24 x i8], [24 x i8]* @string.c316f30584ee0ac304e8eed7e3af175f, i64 0, i64 0), i32 8)
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([26 x i8], [26 x i8]* @string.09e58fc876babc8908c9040bd77d8624, i64 0, i64 0), i32 99)
  %4 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.263c2d145bd0257bade41874fd5a73ec, i64 0, i64 0)) #2
  %5 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.f229d6156f4a2e6f6e5c4ee96406192b, i64 0, i64 0), i32 99) #2
  %6 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.263c2d145bd0257bade41874fd5a73ec, i64 0, i64 0)) #2
  %7 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.f229d6156f4a2e6f6e5c4ee96406192b, i64 0, i64 0), i32 99) #2
  %8 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.8c85cb3ae23186673c0ee88126a99c83, i64 0, i64 0)) #2
  %9 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.f229d6156f4a2e6f6e5c4ee96406192b, i64 0, i64 0), i32 99) #2
  %10 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.263c2d145bd0257bade41874fd5a73ec, i64 0, i64 0)) #2
  %11 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.f229d6156f4a2e6f6e5c4ee96406192b, i64 0, i64 0), i32 3) #2
  %12 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.b5abd14716ff1d42a2c76d0bae14c3cf, i64 0, i64 0), i32 2) #2
  %13 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([26 x i8], [26 x i8]* @string.6e04f1d448592af0a363c48cd79347e3, i64 0, i64 0), i32 80)
  %14 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([31 x i8], [31 x i8]* @string.569e8d7da8dcd242b4520ca536accffb, i64 0, i64 0), i32 8)
  %15 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([33 x i8], [33 x i8]* @string.e1297fae8db86112c4fd38cff8aca961, i64 0, i64 0), i32 88)
  ret void
}

; Function Attrs: nofree nounwind
define void @test.pass_struct_pointer(%test.Printer* nocapture %printer) local_unnamed_addr #0 {
entry:
  %0 = getelementptr %test.Printer, %test.Printer* %printer, i64 0, i32 2, i32 0
  store i8 3, i8* %0, align 1
  %1 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.263c2d145bd0257bade41874fd5a73ec, i64 0, i64 0)) #2
  %2 = load i8, i8* %0, align 1
  %3 = sext i8 %2 to i32
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.f229d6156f4a2e6f6e5c4ee96406192b, i64 0, i64 0), i32 %3) #2
  ret void
}

; Function Attrs: nofree nounwind
define void @test.pass_array(i8* nocapture %buffer) local_unnamed_addr #0 {
entry:
  %0 = getelementptr i8, i8* %buffer, i64 2
  store i8 2, i8* %0, align 1
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.b5abd14716ff1d42a2c76d0bae14c3cf, i64 0, i64 0), i32 2)
  ret void
}

; Function Attrs: nofree nounwind
define void @test.Driver.print(%test.Driver* nocapture readonly %this, i8* nocapture readonly %message) local_unnamed_addr #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) %message)
  %1 = getelementptr %test.Driver, %test.Driver* %this, i64 0, i32 0
  %2 = load i8, i8* %1, align 1
  %3 = sext i8 %2 to i32
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.f229d6156f4a2e6f6e5c4ee96406192b, i64 0, i64 0), i32 %3)
  ret void
}

; Function Attrs: nofree nounwind
define void @test.Printer.print(%test.Printer* nocapture readonly %this, i8* nocapture readonly %message) local_unnamed_addr #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) %message) #2
  %1 = getelementptr %test.Printer, %test.Printer* %this, i64 0, i32 2, i32 0
  %2 = load i8, i8* %1, align 1
  %3 = sext i8 %2 to i32
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([10 x i8], [10 x i8]* @string.f229d6156f4a2e6f6e5c4ee96406192b, i64 0, i64 0), i32 %3) #2
  ret void
}

define void @test.test() local_unnamed_addr {
entry:
  tail call void @test.expression()
  tail call void @test.statement()
  tail call void @test.structs()
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([40 x i8], [40 x i8]* @string.80c523c134f2b89c9ec7f6652a2dbdd7, i64 0, i64 0))
  %1 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.44083ed8ce984d51a6ecfdba2a6c2105, i64 0, i64 0)) #2
  %2 = load void (i8*)*, void (i8*)** @test.ff, align 8
  tail call void %2(i8* getelementptr inbounds ([15 x i8], [15 x i8]* @string.b5b7eec21a3c4ab41dc70340c8ae1d93, i64 0, i64 0))
  ret void
}

attributes #0 = { nofree nounwind }
attributes #1 = { norecurse nounwind readnone }
attributes #2 = { nounwind }
