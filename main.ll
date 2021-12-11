%test.Data = type { %test.SubData, i8 }
%test.SubData = type { i8, float, [5 x i8] }
%test.Driver = type { i8 }
%test.Printer = type { i32, [8 x i8], %test.Driver }

@test.Color.r = constant i8 0
@test.Color.g = constant i8 1
@test.Color.b = constant i8 2
@test.global_array = constant [5 x i8] [i8 1, i8 2, i8 3, i8 4, i8 5]
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
@string.3ef4b443add330c8a95ca5f96c0ff649 = constant [25 x i8] c"convert i32 to i16: %d \0A\00"
@string.4ed512d2145b6a7ca1372858fd18ec55 = constant [24 x i8] c"convert i16 to i8: %d \0A\00"
@string.22f4c765769b2c3e259cdec04f7ab9fc = constant [26 x i8] c"convert float to i8: %d \0A\00"
@string.b8daa6164c5a20e98d469e71a9da3125 = constant [26 x i8] c"convert float to u8: %d \0A\00"
@string.c3c73595ffd209f6a04a209eb631fdd3 = constant [25 x i8] c"convert f32 to f16: %f \0A\00"
@string.68c37f72a2c1fdcc1cdcc2f848d60eae = constant [25 x i8] c"convert f32 to f64: %f \0A\00"
@string.ef448b1e1b7d3a83d28a1ce23787d883 = constant [25 x i8] c"convert i32 to f32: %f \0A\00"
@string.b0cbe247e2f4af55fd9dcce28109e925 = constant [21 x i8] c"address to f32: %d \0A\00"
@string.b25abb667c59353b4d5e4cda2e7cc58e = constant [21 x i8] c"address to i32: %d \0A\00"
@string.825793f8ecd43c3dc72996595a4131e4 = constant [18 x i8] c"bits of f32: %d \0A\00"
@test.ff = global void (i8*)* @test.do_something
@string.80c523c134f2b89c9ec7f6652a2dbdd7 = constant [40 x i8] c"============ test function ============\00"
@string.44083ed8ce984d51a6ecfdba2a6c2105 = constant [15 x i8] c"do something 1\00"
@string.b5b7eec21a3c4ab41dc70340c8ae1d93 = constant [15 x i8] c"do something 2\00"
@string.5f0f1578abd44713c746ded55bf898ea = constant [41 x i8] c"============ test statement ============\00"
@string.07af74d61c4bcfd65e300c22c36df6a3 = constant [14 x i8] c"a(%d) >= 10 \0A\00"
@string.12625b519c0ef75b350a9963cafc3f42 = constant [17 x i8] c"shouldn't happen\00"
@string.7c13f0ed550e89d5fe0dab15a8790a6b = constant [9 x i8] c"I'm else\00"
@string.e509c213bf338f03d246b720ec617c01 = constant [11 x i8] c"loop: %d \0A\00"
@string.ba86886fe05268c3936c4741a0d07a6e = constant [14 x i8] c"switch case 0\00"
@string.162d9796d41e74535694f9688ea21a49 = constant [14 x i8] c"switch case 3\00"
@string.ec374cb30dabe78ccd41f1bcfddac7db = constant [9 x i8] c"a1: %d \0A\00"
@test.global_printer = constant %test.Printer { i32 80, [8 x i8] [i8 1, i8 2, i8 3, i8 4, i8 5, i8 6, i8 7, i8 8], %test.Driver { i8 88 } }
@string.91a35f7e30ee87849a8fb990c35dabf1 = constant [38 x i8] c"============ test struct ============\00"
@string.8c16759f16bae00294081efad1d55ec3 = constant [19 x i8] c"printer.line: %d \0A\00"
@string.c316f30584ee0ac304e8eed7e3af175f = constant [24 x i8] c"printer.buffer[7]: %d \0A\00"
@string.09e58fc876babc8908c9040bd77d8624 = constant [26 x i8] c"printer.driver.type: %d \0A\00"
@string.263c2d145bd0257bade41874fd5a73ec = constant [15 x i8] c"hello printer!\00"
@string.8c85cb3ae23186673c0ee88126a99c83 = constant [15 x i8] c"hello pointer!\00"
@string.6e04f1d448592af0a363c48cd79347e3 = constant [26 x i8] c"global_printer.line: %d \0A\00"
@string.569e8d7da8dcd242b4520ca536accffb = constant [31 x i8] c"global_printer.buffer[7]: %d \0A\00"
@string.e1297fae8db86112c4fd38cff8aca961 = constant [33 x i8] c"global_printer.driver.type: %d \0A\00"
@string.7ffa0c893047b73021f29c1b48c9892b = constant [17 x i8] c"sizeof i32: %d \0A\00"
@string.451c6e0c15d560a4cfc5a46a02d53bfa = constant [17 x i8] c"sizeof f64: %d \0A\00"
@string.adf1b889a9023b93577417ca23d21793 = constant [21 x i8] c"sizeof pointer: %d \0A\00"
@string.846c6e4a2aac8de8fa1cce667d7cd481 = constant [21 x i8] c"sizeof printer: %d \0A\00"
@string.fbd4ad59a9864656f8875863ac3b7dab = constant [22 x i8] c"sizeof array[5]: %d \0A\00"
@string.b5abd14716ff1d42a2c76d0bae14c3cf = constant [16 x i8] c"buffer[2]: %d \0A\00"
@string.f229d6156f4a2e6f6e5c4ee96406192b = constant [10 x i8] c"type:%d \0A\00"

declare i32 @puts(i8* %text)

declare i32 @printf(i8* %format, ...)

declare i8* @malloc(i32 %size)

declare i8* @calloc(i32 %number, i32 %size)

declare i8* @realloc(i8* %address, i32 %size)

declare void @free(i8* %address)

declare i32 @memcmp(i8* %dest, i8* %source, i32 %size)

declare void @memcpy(i8* %dest, i8* %source, i32 %size)

declare void @memset(i8* %source, i32 %value, i32 %size)

define void @main() {
entry:
	br label %body


body:
	call void @test.test()
	br label %exit


exit:
	ret void

}

define void @test.expression() {
entry:
	br label %body


body:
	%0 = getelementptr [42 x i8], [42 x i8]* @string.4576fbff7ad2d9fa622f16573db7b286, i32 0, i32 0
	%1 = call i32 @puts(i8* %0)
	call void @test.unary()
	call void @test.increment_decrement()
	call void @test.binary()
	call void @test.parentheses()
	call void @test.literal()
	call void @test.subscripting()
	call void @test.member_access()
	call void @test.conversion()
	br label %exit


exit:
	ret void

}

define void @test.unary() {
entry:
	%0 = alloca i32
	%1 = alloca i32
	%2 = alloca i32
	%3 = alloca i1
	%4 = alloca i1
	br label %body


body:
	store i32 1, i32* %0
	%5 = getelementptr [9 x i8], [9 x i8]* @string.f9b6d891c5ca674309c459ad55eb01c8, i32 0, i32 0
	%6 = load i32, i32* %0
	%7 = call i32 (i8*, ...) @printf(i8* %5, i32 %6)
	%8 = load i32, i32* %0
	%9 = sub i32 0, %8
	store i32 %9, i32* %1
	%10 = getelementptr [9 x i8], [9 x i8]* @string.f52c63a936a31e2b2d03c5c746b8d5b9, i32 0, i32 0
	%11 = load i32, i32* %1
	%12 = call i32 (i8*, ...) @printf(i8* %10, i32 %11)
	%13 = xor i32 1, -1
	store i32 %13, i32* %2
	%14 = getelementptr [9 x i8], [9 x i8]* @string.67c5acbbce37db2c12b92a427bc08a84, i32 0, i32 0
	%15 = load i32, i32* %2
	%16 = call i32 (i8*, ...) @printf(i8* %14, i32 %15)
	store i1 true, i1* %3
	%17 = getelementptr [9 x i8], [9 x i8]* @string.c3c4ff0f83dad5387535d315826c22f8, i32 0, i32 0
	%18 = load i1, i1* %3
	%19 = sext i1 %18 to i32
	%20 = call i32 (i8*, ...) @printf(i8* %17, i32 %19)
	%21 = load i1, i1* %3
	%22 = xor i1 %21, true
	store i1 %22, i1* %4
	%23 = getelementptr [9 x i8], [9 x i8]* @string.e40a403ffbdf9c2c70921d6bb7739cd8, i32 0, i32 0
	%24 = load i1, i1* %4
	%25 = sext i1 %24 to i32
	%26 = call i32 (i8*, ...) @printf(i8* %23, i32 %25)
	br label %exit


exit:
	ret void

}

define void @test.increment_decrement() {
entry:
	%0 = alloca i32
	br label %body


body:
	store i32 10, i32* %0
	%1 = load i32, i32* %0
	%2 = add i32 %1, 1
	store i32 %2, i32* %0
	%3 = getelementptr [9 x i8], [9 x i8]* @string.2dcc97a590ca083991ffe9b43c08dd02, i32 0, i32 0
	%4 = load i32, i32* %0
	%5 = call i32 (i8*, ...) @printf(i8* %3, i32 %4)
	store i32 20, i32* %0
	%6 = load i32, i32* %0
	%7 = sub i32 %6, 1
	store i32 %7, i32* %0
	%8 = getelementptr [9 x i8], [9 x i8]* @string.2dcc97a590ca083991ffe9b43c08dd02, i32 0, i32 0
	%9 = load i32, i32* %0
	%10 = call i32 (i8*, ...) @printf(i8* %8, i32 %9)
	br label %exit


exit:
	ret void

}

define void @test.binary() {
entry:
	%0 = alloca i32
	br label %body


body:
	store i32 0, i32* %0
	store i32 5, i32* %0
	%1 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%2 = load i32, i32* %0
	%3 = call i32 (i8*, ...) @printf(i8* %1, i32 %2)
	%4 = load i32, i32* %0
	%5 = add i32 %4, 5
	store i32 %5, i32* %0
	%6 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%7 = load i32, i32* %0
	%8 = call i32 (i8*, ...) @printf(i8* %6, i32 %7)
	%9 = load i32, i32* %0
	%10 = sub i32 %9, 1
	store i32 %10, i32* %0
	%11 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%12 = load i32, i32* %0
	%13 = call i32 (i8*, ...) @printf(i8* %11, i32 %12)
	%14 = load i32, i32* %0
	%15 = mul i32 %14, 2
	store i32 %15, i32* %0
	%16 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%17 = load i32, i32* %0
	%18 = call i32 (i8*, ...) @printf(i8* %16, i32 %17)
	%19 = load i32, i32* %0
	%20 = sdiv i32 %19, 9
	store i32 %20, i32* %0
	%21 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%22 = load i32, i32* %0
	%23 = call i32 (i8*, ...) @printf(i8* %21, i32 %22)
	store i32 11, i32* %0
	%24 = load i32, i32* %0
	%25 = srem i32 %24, 4
	store i32 %25, i32* %0
	%26 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%27 = load i32, i32* %0
	%28 = call i32 (i8*, ...) @printf(i8* %26, i32 %27)
	%29 = load i32, i32* %0
	%30 = shl i32 %29, 2
	store i32 %30, i32* %0
	%31 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%32 = load i32, i32* %0
	%33 = call i32 (i8*, ...) @printf(i8* %31, i32 %32)
	%34 = load i32, i32* %0
	%35 = ashr i32 %34, 1
	store i32 %35, i32* %0
	%36 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%37 = load i32, i32* %0
	%38 = call i32 (i8*, ...) @printf(i8* %36, i32 %37)
	%39 = load i32, i32* %0
	%40 = or i32 %39, 15
	store i32 %40, i32* %0
	%41 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%42 = load i32, i32* %0
	%43 = call i32 (i8*, ...) @printf(i8* %41, i32 %42)
	%44 = load i32, i32* %0
	%45 = xor i32 %44, 8
	store i32 %45, i32* %0
	%46 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%47 = load i32, i32* %0
	%48 = call i32 (i8*, ...) @printf(i8* %46, i32 %47)
	%49 = load i32, i32* %0
	%50 = and i32 %49, 6
	store i32 %50, i32* %0
	%51 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%52 = load i32, i32* %0
	%53 = call i32 (i8*, ...) @printf(i8* %51, i32 %52)
	%54 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%55 = or i32 15, 8
	%56 = call i32 (i8*, ...) @printf(i8* %54, i32 %55)
	%57 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%58 = xor i32 15, 8
	%59 = call i32 (i8*, ...) @printf(i8* %57, i32 %58)
	%60 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%61 = and i32 15, 8
	%62 = call i32 (i8*, ...) @printf(i8* %60, i32 %61)
	%63 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%64 = icmp eq i32 10, 10
	%65 = sext i1 %64 to i32
	%66 = call i32 (i8*, ...) @printf(i8* %63, i32 %65)
	%67 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%68 = icmp ne i32 10, 10
	%69 = sext i1 %68 to i32
	%70 = call i32 (i8*, ...) @printf(i8* %67, i32 %69)
	%71 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%72 = icmp slt i32 10, 10
	%73 = sext i1 %72 to i32
	%74 = call i32 (i8*, ...) @printf(i8* %71, i32 %73)
	%75 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%76 = icmp sle i32 10, 10
	%77 = sext i1 %76 to i32
	%78 = call i32 (i8*, ...) @printf(i8* %75, i32 %77)
	%79 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%80 = icmp sgt i32 10, 10
	%81 = sext i1 %80 to i32
	%82 = call i32 (i8*, ...) @printf(i8* %79, i32 %81)
	%83 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%84 = icmp sge i32 10, 10
	%85 = sext i1 %84 to i32
	%86 = call i32 (i8*, ...) @printf(i8* %83, i32 %85)
	%87 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%88 = shl i32 10, 3
	%89 = call i32 (i8*, ...) @printf(i8* %87, i32 %88)
	%90 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%91 = ashr i32 10, 1
	%92 = call i32 (i8*, ...) @printf(i8* %90, i32 %91)
	%93 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%94 = add i32 5, 3
	%95 = call i32 (i8*, ...) @printf(i8* %93, i32 %94)
	%96 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%97 = sub i32 5, 3
	%98 = call i32 (i8*, ...) @printf(i8* %96, i32 %97)
	%99 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%100 = mul i32 5, 3
	%101 = call i32 (i8*, ...) @printf(i8* %99, i32 %100)
	%102 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%103 = sdiv i32 5, 3
	%104 = call i32 (i8*, ...) @printf(i8* %102, i32 %103)
	%105 = getelementptr [8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0
	%106 = srem i32 5, 3
	%107 = call i32 (i8*, ...) @printf(i8* %105, i32 %106)
	%108 = getelementptr [11 x i8], [11 x i8]* @string.a753fba743a9e6b08cb7a2627f69b75d, i32 0, i32 0
	%109 = or i1 true, false
	%110 = sext i1 %109 to i32
	%111 = call i32 (i8*, ...) @printf(i8* %108, i32 %110)
	%112 = getelementptr [11 x i8], [11 x i8]* @string.a753fba743a9e6b08cb7a2627f69b75d, i32 0, i32 0
	%113 = and i1 true, false
	%114 = sext i1 %113 to i32
	%115 = call i32 (i8*, ...) @printf(i8* %112, i32 %114)
	br label %exit


exit:
	ret void

}

define void @test.parentheses() {
entry:
	br label %body


body:
	%0 = getelementptr [18 x i8], [18 x i8]* @string.9155af3e03234ca6017e6a626fa48d60, i32 0, i32 0
	%1 = mul i32 add (i32 5, i32 6), 3
	%2 = call i32 (i8*, ...) @printf(i8* %0, i32 %1)
	br label %exit


exit:
	ret void

}

define void @test.literal() {
entry:
	%0 = alloca i8
	%1 = alloca i1
	%2 = alloca i8
	%3 = alloca half
	%4 = alloca float
	%5 = alloca double
	%6 = alloca i8*
	br label %body


body:
	store i8 97, i8* %0
	%7 = getelementptr [8 x i8], [8 x i8]* @string.6815af516458351e77683ead5f501317, i32 0, i32 0
	%8 = load i8, i8* %0
	%9 = sext i8 %8 to i32
	%10 = call i32 (i8*, ...) @printf(i8* %7, i32 %9)
	store i1 true, i1* %1
	%11 = getelementptr [8 x i8], [8 x i8]* @string.d58ddb72e75f1acfc4203e33bddc08a1, i32 0, i32 0
	%12 = load i1, i1* %1
	%13 = sext i1 %12 to i32
	%14 = call i32 (i8*, ...) @printf(i8* %11, i32 %13)
	store i8 123, i8* %2
	%15 = getelementptr [8 x i8], [8 x i8]* @string.5b8b2fafadbddfa000cd0e716725d4a4, i32 0, i32 0
	%16 = load i8, i8* %2
	%17 = sext i8 %16 to i32
	%18 = call i32 (i8*, ...) @printf(i8* %15, i32 %17)
	store half 0x4009200000000000, half* %3
	%19 = getelementptr [10 x i8], [10 x i8]* @string.dd42ef93dc06a72b063baa72848d660c, i32 0, i32 0
	%20 = load half, half* %3
	%21 = fpext half %20 to double
	%22 = call i32 (i8*, ...) @printf(i8* %19, double %21)
	store float 0x40091EB860000000, float* %4
	%23 = getelementptr [10 x i8], [10 x i8]* @string.dc24ff6a55a1c588a346f9dff66c25a0, i32 0, i32 0
	%24 = load float, float* %4
	%25 = fpext float %24 to double
	%26 = call i32 (i8*, ...) @printf(i8* %23, double %25)
	store double 0x40091EB851EB851F, double* %5
	%27 = getelementptr [10 x i8], [10 x i8]* @string.7a828f7c003ac662930a932d14c84f48, i32 0, i32 0
	%28 = load double, double* %5
	%29 = call i32 (i8*, ...) @printf(i8* %27, double %28)
	%30 = getelementptr [12 x i8], [12 x i8]* @string.3aff445dea2b63e4d3b135c5219ba7dc, i32 0, i32 0
	store i8* %30, i8** %6
	%31 = load i8*, i8** %6
	%32 = call i32 @puts(i8* %31)
	br label %exit


exit:
	ret void

}

define void @test.subscripting() {
entry:
	%0 = alloca [5 x i8]
	%1 = alloca i8*
	%2 = alloca i8*
	%3 = alloca i8*
	%4 = alloca [5 x i8]
	br label %body


body:
	store [5 x i8] zeroinitializer, [5 x i8]* %0
	%5 = getelementptr [15 x i8], [15 x i8]* @string.ccbd06f65fb69a974bb7bbe132352fd5, i32 0, i32 0
	%6 = getelementptr [5 x i8], [5 x i8]* %0, i32 0, i32 0
	%7 = load i8, i8* %6
	%8 = sext i8 %7 to i32
	%9 = call i32 (i8*, ...) @printf(i8* %5, i32 %8)
	%10 = getelementptr [5 x i8], [5 x i8]* %0, i32 0, i32 3
	store i8 3, i8* %10
	%11 = getelementptr [15 x i8], [15 x i8]* @string.502edb90c5d63a7982b92c4846005a12, i32 0, i32 0
	%12 = getelementptr [5 x i8], [5 x i8]* %0, i32 0, i32 3
	%13 = load i8, i8* %12
	%14 = sext i8 %13 to i32
	%15 = call i32 (i8*, ...) @printf(i8* %11, i32 %14)
	%16 = getelementptr [5 x i8], [5 x i8]* %0, i32 0, i32 0
	store i8* %16, i8** %1
	%17 = getelementptr [16 x i8], [16 x i8]* @string.f078cc2571d60d58e6a551d92df567c4, i32 0, i32 0
	%18 = load i8*, i8** %1
	%19 = getelementptr i8, i8* %18, i32 0
	%20 = load i8, i8* %19
	%21 = sext i8 %20 to i32
	%22 = call i32 (i8*, ...) @printf(i8* %17, i32 %21)
	%23 = getelementptr [16 x i8], [16 x i8]* @string.966d3dfa4a8527741a06bd9fbaa21f93, i32 0, i32 0
	%24 = load i8*, i8** %1
	%25 = getelementptr i8, i8* %24, i32 3
	%26 = load i8, i8* %25
	%27 = sext i8 %26 to i32
	%28 = call i32 (i8*, ...) @printf(i8* %23, i32 %27)
	%29 = load i8*, i8** %1
	store i8* %29, i8** %2
	%30 = getelementptr [16 x i8], [16 x i8]* @string.f9fe529bc21937b1ee14af38e842590b, i32 0, i32 0
	%31 = load i8*, i8** %2
	%32 = getelementptr i8, i8* %31, i32 0
	%33 = load i8, i8* %32
	%34 = sext i8 %33 to i32
	%35 = call i32 (i8*, ...) @printf(i8* %30, i32 %34)
	%36 = getelementptr [16 x i8], [16 x i8]* @string.85ac8a2eb899f708bbb12d753db07868, i32 0, i32 0
	%37 = load i8*, i8** %2
	%38 = getelementptr i8, i8* %37, i32 3
	%39 = load i8, i8* %38
	%40 = sext i8 %39 to i32
	%41 = call i32 (i8*, ...) @printf(i8* %36, i32 %40)
	store i8* zeroinitializer, i8** %3
	%42 = load i8*, i8** %2
	store i8* %42, i8** %3
	%43 = getelementptr [16 x i8], [16 x i8]* @string.8824f4dbad52bbb3684995480092775e, i32 0, i32 0
	%44 = load i8*, i8** %3
	%45 = getelementptr i8, i8* %44, i32 0
	%46 = load i8, i8* %45
	%47 = sext i8 %46 to i32
	%48 = call i32 (i8*, ...) @printf(i8* %43, i32 %47)
	%49 = getelementptr [16 x i8], [16 x i8]* @string.5f0dc37317410eea89a43776ec4ac6e1, i32 0, i32 0
	%50 = load i8*, i8** %3
	%51 = getelementptr i8, i8* %50, i32 3
	%52 = load i8, i8* %51
	%53 = sext i8 %52 to i32
	%54 = call i32 (i8*, ...) @printf(i8* %49, i32 %53)
	store [5 x i8] [i8 1, i8 2, i8 3, i8 4, i8 5], [5 x i8]* %4
	%55 = getelementptr [19 x i8], [19 x i8]* @string.db08573c403e33d25bb325d1df98c844, i32 0, i32 0
	%56 = getelementptr [5 x i8], [5 x i8]* %4, i32 0, i32 0
	%57 = load i8, i8* %56
	%58 = sext i8 %57 to i32
	%59 = call i32 (i8*, ...) @printf(i8* %55, i32 %58)
	%60 = getelementptr [19 x i8], [19 x i8]* @string.b1faf43818ca0e7e4b4b1e24b441f795, i32 0, i32 0
	%61 = getelementptr [5 x i8], [5 x i8]* %4, i32 0, i32 3
	%62 = load i8, i8* %61
	%63 = sext i8 %62 to i32
	%64 = call i32 (i8*, ...) @printf(i8* %60, i32 %63)
	br label %exit


exit:
	ret void

}

define void @test.member_access() {
entry:
	%0 = alloca %test.SubData
	%1 = alloca %test.Data
	br label %body


body:
	%2 = getelementptr [14 x i8], [14 x i8]* @string.d18cb31ff3a37014a9ed64a2687344d4, i32 0, i32 0
	%3 = load i8, i8* @test.Color.g
	%4 = sext i8 %3 to i32
	%5 = call i32 (i8*, ...) @printf(i8* %2, i32 %4)
	%6 = getelementptr [22 x i8], [22 x i8]* @string.9fcfb18ceb0d348e69c2e13fa41b241d, i32 0, i32 0
	%7 = getelementptr [5 x i8], [5 x i8]* @test.global_array, i32 0, i32 2
	%8 = load i8, i8* %7
	%9 = sext i8 %8 to i32
	%10 = call i32 (i8*, ...) @printf(i8* %6, i32 %9)
	store %test.SubData zeroinitializer, %test.SubData* %0
	%11 = getelementptr [18 x i8], [18 x i8]* @string.1daf4144552c4db57e99d55450ed346e, i32 0, i32 0
	%12 = getelementptr %test.SubData, %test.SubData* %0, i32 0, i32 0
	%13 = load i8, i8* %12
	%14 = sext i8 %13 to i32
	%15 = call i32 (i8*, ...) @printf(i8* %11, i32 %14)
	%16 = getelementptr %test.SubData, %test.SubData* %0, i32 0, i32 0
	store i8 5, i8* %16
	%17 = getelementptr [18 x i8], [18 x i8]* @string.1daf4144552c4db57e99d55450ed346e, i32 0, i32 0
	%18 = getelementptr %test.SubData, %test.SubData* %0, i32 0, i32 0
	%19 = load i8, i8* %18
	%20 = sext i8 %19 to i32
	%21 = call i32 (i8*, ...) @printf(i8* %17, i32 %20)
	%22 = getelementptr %test.SubData, %test.SubData* %0, i32 0, i32 1
	store float 0x40091EB860000000, float* %22
	%23 = getelementptr [16 x i8], [16 x i8]* @string.560e3347d8fe3fd15f15ce5db418664f, i32 0, i32 0
	%24 = getelementptr %test.SubData, %test.SubData* %0, i32 0, i32 1
	%25 = load float, float* %24
	%26 = fpext float %25 to double
	%27 = call i32 (i8*, ...) @printf(i8* %23, double %26)
	%28 = getelementptr %test.SubData, %test.SubData* %0, i32 0, i32 2
	%29 = getelementptr [5 x i8], [5 x i8]* %28, i32 0, i32 3
	store i8 3, i8* %29
	%30 = getelementptr [19 x i8], [19 x i8]* @string.b585a7adc3e8d68bbf60cb859044df1e, i32 0, i32 0
	%31 = getelementptr %test.SubData, %test.SubData* %0, i32 0, i32 2
	%32 = getelementptr [5 x i8], [5 x i8]* %31, i32 0, i32 3
	%33 = load i8, i8* %32
	%34 = sext i8 %33 to i32
	%35 = call i32 (i8*, ...) @printf(i8* %30, i32 %34)
	store %test.Data zeroinitializer, %test.Data* %1
	%36 = getelementptr %test.Data, %test.Data* %1, i32 0, i32 1
	store i8 5, i8* %36
	%37 = getelementptr [17 x i8], [17 x i8]* @string.84ad90c9c520f1a4e80779cfa15248b6, i32 0, i32 0
	%38 = getelementptr %test.Data, %test.Data* %1, i32 0, i32 1
	%39 = load i8, i8* %38
	%40 = sext i8 %39 to i32
	%41 = call i32 (i8*, ...) @printf(i8* %37, i32 %40)
	%42 = getelementptr %test.Data, %test.Data* %1, i32 0, i32 0
	%43 = getelementptr %test.SubData, %test.SubData* %42, i32 0, i32 0
	store i8 8, i8* %43
	%44 = getelementptr [23 x i8], [23 x i8]* @string.07ce14d972194d598243322dc9f50250, i32 0, i32 0
	%45 = getelementptr %test.Data, %test.Data* %1, i32 0, i32 0
	%46 = getelementptr %test.SubData, %test.SubData* %45, i32 0, i32 0
	%47 = load i8, i8* %46
	%48 = sext i8 %47 to i32
	%49 = call i32 (i8*, ...) @printf(i8* %44, i32 %48)
	%50 = getelementptr %test.Data, %test.Data* %1, i32 0, i32 0
	%51 = getelementptr %test.SubData, %test.SubData* %50, i32 0, i32 2
	%52 = getelementptr [5 x i8], [5 x i8]* %51, i32 0, i32 3
	store i8 9, i8* %52
	%53 = getelementptr [24 x i8], [24 x i8]* @string.6db0fbcde59d77fa7fc3126dc45321f0, i32 0, i32 0
	%54 = getelementptr %test.Data, %test.Data* %1, i32 0, i32 0
	%55 = getelementptr %test.SubData, %test.SubData* %54, i32 0, i32 2
	%56 = getelementptr [5 x i8], [5 x i8]* %55, i32 0, i32 3
	%57 = load i8, i8* %56
	%58 = sext i8 %57 to i32
	%59 = call i32 (i8*, ...) @printf(i8* %53, i32 %58)
	br label %exit


exit:
	ret void

}

define void @test.conversion() {
entry:
	%0 = alloca i32
	%1 = alloca i16
	%2 = alloca i8
	%3 = alloca float
	%4 = alloca i8
	%5 = alloca i8
	%6 = alloca i8
	%7 = alloca half
	%8 = alloca double
	%9 = alloca float
	%10 = alloca float*
	%11 = alloca i32*
	%12 = alloca i32*
	%13 = alloca i32*
	br label %body


body:
	store i32 65636, i32* %0
	%14 = load i32, i32* %0
	%15 = trunc i32 %14 to i16
	store i16 %15, i16* %1
	%16 = getelementptr [25 x i8], [25 x i8]* @string.3ef4b443add330c8a95ca5f96c0ff649, i32 0, i32 0
	%17 = load i16, i16* %1
	%18 = sext i16 %17 to i32
	%19 = call i32 (i8*, ...) @printf(i8* %16, i32 %18)
	%20 = load i16, i16* %1
	%21 = trunc i16 %20 to i8
	store i8 %21, i8* %2
	%22 = getelementptr [24 x i8], [24 x i8]* @string.4ed512d2145b6a7ca1372858fd18ec55, i32 0, i32 0
	%23 = load i8, i8* %2
	%24 = sext i8 %23 to i32
	%25 = call i32 (i8*, ...) @printf(i8* %22, i32 %24)
	%26 = fsub float 0x0, 0x40091EB860000000
	store float %26, float* %3
	%27 = load float, float* %3
	%28 = fptosi float %27 to i8
	store i8 %28, i8* %4
	%29 = getelementptr [26 x i8], [26 x i8]* @string.22f4c765769b2c3e259cdec04f7ab9fc, i32 0, i32 0
	%30 = load i8, i8* %4
	%31 = sext i8 %30 to i32
	%32 = call i32 (i8*, ...) @printf(i8* %29, i32 %31)
	%33 = load float, float* %3
	%34 = fptoui float %33 to i8
	store i8 %34, i8* %5
	%35 = getelementptr [26 x i8], [26 x i8]* @string.b8daa6164c5a20e98d469e71a9da3125, i32 0, i32 0
	%36 = load i8, i8* %5
	%37 = sext i8 %36 to i32
	%38 = call i32 (i8*, ...) @printf(i8* %35, i32 %37)
	%39 = load float, float* %3
	%40 = fsub float 0x0, %39
	%41 = fptoui float %40 to i8
	store i8 %41, i8* %6
	%42 = getelementptr [26 x i8], [26 x i8]* @string.b8daa6164c5a20e98d469e71a9da3125, i32 0, i32 0
	%43 = load i8, i8* %6
	%44 = sext i8 %43 to i32
	%45 = call i32 (i8*, ...) @printf(i8* %42, i32 %44)
	%46 = load float, float* %3
	%47 = fptrunc float %46 to half
	store half %47, half* %7
	%48 = getelementptr [25 x i8], [25 x i8]* @string.c3c73595ffd209f6a04a209eb631fdd3, i32 0, i32 0
	%49 = load half, half* %7
	%50 = fpext half %49 to double
	%51 = call i32 (i8*, ...) @printf(i8* %48, double %50)
	%52 = load float, float* %3
	%53 = fpext float %52 to double
	store double %53, double* %8
	%54 = getelementptr [25 x i8], [25 x i8]* @string.68c37f72a2c1fdcc1cdcc2f848d60eae, i32 0, i32 0
	%55 = load double, double* %8
	%56 = call i32 (i8*, ...) @printf(i8* %54, double %55)
	%57 = load i32, i32* %0
	%58 = sitofp i32 %57 to float
	store float %58, float* %9
	%59 = getelementptr [25 x i8], [25 x i8]* @string.ef448b1e1b7d3a83d28a1ce23787d883, i32 0, i32 0
	%60 = load float, float* %9
	%61 = fpext float %60 to double
	%62 = call i32 (i8*, ...) @printf(i8* %59, double %61)
	%63 = alloca float*
	store float* %3, float** %63
	%64 = load float*, float** %63
	store float* %64, float** %10
	%65 = getelementptr [21 x i8], [21 x i8]* @string.b0cbe247e2f4af55fd9dcce28109e925, i32 0, i32 0
	%66 = load float*, float** %10
	%67 = call i32 (i8*, ...) @printf(i8* %65, float* %66)
	%68 = load float*, float** %10
	%69 = bitcast float* %68 to i32*
	store i32* %69, i32** %11
	%70 = getelementptr [21 x i8], [21 x i8]* @string.b25abb667c59353b4d5e4cda2e7cc58e, i32 0, i32 0
	%71 = load i32*, i32** %11
	%72 = call i32 (i8*, ...) @printf(i8* %70, i32* %71)
	%73 = load i32*, i32** %11
	store i32* %73, i32** %12
	%74 = getelementptr [18 x i8], [18 x i8]* @string.825793f8ecd43c3dc72996595a4131e4, i32 0, i32 0
	%75 = load i32*, i32** %12
	%76 = getelementptr i32, i32* %75, i32 0
	%77 = load i32, i32* %76
	%78 = call i32 (i8*, ...) @printf(i8* %74, i32 %77)
	%79 = load float*, float** %10
	%80 = bitcast float* %79 to i32*
	store i32* %80, i32** %13
	%81 = getelementptr [18 x i8], [18 x i8]* @string.825793f8ecd43c3dc72996595a4131e4, i32 0, i32 0
	%82 = load i32*, i32** %13
	%83 = getelementptr i32, i32* %82, i32 0
	%84 = load i32, i32* %83
	%85 = call i32 (i8*, ...) @printf(i8* %81, i32 %84)
	br label %exit


exit:
	ret void

}

define void @test.functions() {
entry:
	%0 = alloca void (i8*)*
	br label %body


body:
	%1 = getelementptr [40 x i8], [40 x i8]* @string.80c523c134f2b89c9ec7f6652a2dbdd7, i32 0, i32 0
	%2 = call i32 @puts(i8* %1)
	store void (i8*)* @test.do_something, void (i8*)** %0
	%3 = load void (i8*)*, void (i8*)** %0
	%4 = getelementptr [15 x i8], [15 x i8]* @string.44083ed8ce984d51a6ecfdba2a6c2105, i32 0, i32 0
	call void %3(i8* %4)
	%5 = load void (i8*)*, void (i8*)** @test.ff
	%6 = getelementptr [15 x i8], [15 x i8]* @string.b5b7eec21a3c4ab41dc70340c8ae1d93, i32 0, i32 0
	call void %5(i8* %6)
	br label %exit


exit:
	ret void

}

define void @test.do_something(i8* %msg) {
entry:
	%0 = alloca i8*
	store i8* %msg, i8** %0
	br label %body


body:
	%1 = load i8*, i8** %0
	%2 = call i32 @puts(i8* %1)
	br label %exit


exit:
	ret void

}

define void @test.statement() {
entry:
	%0 = alloca i8
	%1 = alloca i32
	%2 = alloca i32
	br label %body


body:
	%3 = getelementptr [41 x i8], [41 x i8]* @string.5f0f1578abd44713c746ded55bf898ea, i32 0, i32 0
	%4 = call i32 @puts(i8* %3)
	store i8 10, i8* %0
	%5 = load i8, i8* %0
	%6 = icmp uge i8 %5, 10
	br i1 %6, label %10, label %7


exit:
	ret void


7:
	%8 = load i8, i8* %0
	%9 = icmp ugt i8 %8, 100
	br i1 %9, label %16, label %19


10:
	%11 = getelementptr [14 x i8], [14 x i8]* @string.07af74d61c4bcfd65e300c22c36df6a3, i32 0, i32 0
	%12 = load i8, i8* %0
	%13 = sext i8 %12 to i32
	%14 = call i32 (i8*, ...) @printf(i8* %11, i32 %13)
	br label %7


15:
	store i8 0, i8* %0
	br label %24


16:
	%17 = getelementptr [17 x i8], [17 x i8]* @string.12625b519c0ef75b350a9963cafc3f42, i32 0, i32 0
	%18 = call i32 @puts(i8* %17)
	br label %15


19:
	%20 = getelementptr [9 x i8], [9 x i8]* @string.7c13f0ed550e89d5fe0dab15a8790a6b, i32 0, i32 0
	%21 = call i32 @puts(i8* %20)
	br label %15


22:
	store i8 3, i8* %0
	%23 = load i8, i8* %0
	switch i8 %23, label %40 [
		i8 0, label %41
		i8 3, label %44
	]


24:
	%25 = load i8, i8* %0
	%26 = icmp ult i8 %25, 10
	br i1 %26, label %30, label %22


27:
	%28 = load i8, i8* %0
	%29 = add i8 %28, 1
	store i8 %29, i8* %0
	br label %24


30:
	%31 = getelementptr [11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i32 0, i32 0
	%32 = load i8, i8* %0
	%33 = sext i8 %32 to i32
	%34 = call i32 (i8*, ...) @printf(i8* %31, i32 %33)
	br label %27


35:
	store i32 123, i32* %1
	%36 = load i32, i32* %1
	store i32 %36, i32* %2
	%37 = getelementptr [9 x i8], [9 x i8]* @string.ec374cb30dabe78ccd41f1bcfddac7db, i32 0, i32 0
	%38 = load i32, i32* %2
	%39 = call i32 (i8*, ...) @printf(i8* %37, i32 %38)
	br label %exit


40:
	br label %35


41:
	%42 = getelementptr [14 x i8], [14 x i8]* @string.ba86886fe05268c3936c4741a0d07a6e, i32 0, i32 0
	%43 = call i32 @puts(i8* %42)
	br label %35


44:
	%45 = getelementptr [14 x i8], [14 x i8]* @string.162d9796d41e74535694f9688ea21a49, i32 0, i32 0
	%46 = call i32 @puts(i8* %45)
	br label %35

}

define void @test.structs() {
entry:
	%0 = alloca %test.Printer
	%1 = alloca %test.Printer*
	br label %body


body:
	%2 = getelementptr [38 x i8], [38 x i8]* @string.91a35f7e30ee87849a8fb990c35dabf1, i32 0, i32 0
	%3 = call i32 @puts(i8* %2)
	store %test.Printer { i32 100, [8 x i8] [i8 1, i8 2, i8 3, i8 4, i8 5, i8 6, i8 7, i8 8], %test.Driver { i8 99 } }, %test.Printer* %0
	%4 = getelementptr [19 x i8], [19 x i8]* @string.8c16759f16bae00294081efad1d55ec3, i32 0, i32 0
	%5 = getelementptr %test.Printer, %test.Printer* %0, i32 0, i32 0
	%6 = load i32, i32* %5
	%7 = call i32 (i8*, ...) @printf(i8* %4, i32 %6)
	%8 = getelementptr [24 x i8], [24 x i8]* @string.c316f30584ee0ac304e8eed7e3af175f, i32 0, i32 0
	%9 = getelementptr %test.Printer, %test.Printer* %0, i32 0, i32 1
	%10 = getelementptr [8 x i8], [8 x i8]* %9, i32 0, i32 7
	%11 = load i8, i8* %10
	%12 = sext i8 %11 to i32
	%13 = call i32 (i8*, ...) @printf(i8* %8, i32 %12)
	%14 = getelementptr [26 x i8], [26 x i8]* @string.09e58fc876babc8908c9040bd77d8624, i32 0, i32 0
	%15 = getelementptr %test.Printer, %test.Printer* %0, i32 0, i32 2
	%16 = getelementptr %test.Driver, %test.Driver* %15, i32 0, i32 0
	%17 = load i8, i8* %16
	%18 = sext i8 %17 to i32
	%19 = call i32 (i8*, ...) @printf(i8* %14, i32 %18)
	%20 = getelementptr %test.Printer, %test.Printer* %0, i32 0, i32 2
	%21 = getelementptr %test.Printer, %test.Printer* %0, i32 0, i32 2
	%22 = alloca %test.Driver*
	store %test.Driver* %21, %test.Driver** %22
	%23 = load %test.Driver*, %test.Driver** %22
	%24 = getelementptr [15 x i8], [15 x i8]* @string.263c2d145bd0257bade41874fd5a73ec, i32 0, i32 0
	call void @test.Driver.print(%test.Driver* %23, i8* %24)
	%25 = alloca %test.Printer*
	store %test.Printer* %0, %test.Printer** %25
	%26 = load %test.Printer*, %test.Printer** %25
	%27 = getelementptr [15 x i8], [15 x i8]* @string.263c2d145bd0257bade41874fd5a73ec, i32 0, i32 0
	call void @test.Printer.print(%test.Printer* %26, i8* %27)
	%28 = alloca %test.Printer*
	store %test.Printer* %0, %test.Printer** %28
	%29 = load %test.Printer*, %test.Printer** %28
	store %test.Printer* %29, %test.Printer** %1
	%30 = load %test.Printer*, %test.Printer** %1
	%31 = load %test.Printer*, %test.Printer** %1
	%32 = getelementptr [15 x i8], [15 x i8]* @string.8c85cb3ae23186673c0ee88126a99c83, i32 0, i32 0
	call void @test.Printer.print(%test.Printer* %31, i8* %32)
	%33 = alloca %test.Printer*
	store %test.Printer* %0, %test.Printer** %33
	%34 = load %test.Printer*, %test.Printer** %33
	call void @test.pass_struct_pointer(%test.Printer* %34)
	%35 = getelementptr %test.Printer, %test.Printer* %0, i32 0, i32 1
	%36 = getelementptr [8 x i8], [8 x i8]* %35, i32 0, i32 0
	call void @test.pass_array(i8* %36)
	%37 = getelementptr [26 x i8], [26 x i8]* @string.6e04f1d448592af0a363c48cd79347e3, i32 0, i32 0
	%38 = getelementptr %test.Printer, %test.Printer* @test.global_printer, i32 0, i32 0
	%39 = load i32, i32* %38
	%40 = call i32 (i8*, ...) @printf(i8* %37, i32 %39)
	%41 = getelementptr [31 x i8], [31 x i8]* @string.569e8d7da8dcd242b4520ca536accffb, i32 0, i32 0
	%42 = getelementptr %test.Printer, %test.Printer* @test.global_printer, i32 0, i32 1
	%43 = getelementptr [8 x i8], [8 x i8]* %42, i32 0, i32 7
	%44 = load i8, i8* %43
	%45 = sext i8 %44 to i32
	%46 = call i32 (i8*, ...) @printf(i8* %41, i32 %45)
	%47 = getelementptr [33 x i8], [33 x i8]* @string.e1297fae8db86112c4fd38cff8aca961, i32 0, i32 0
	%48 = getelementptr %test.Printer, %test.Printer* @test.global_printer, i32 0, i32 2
	%49 = getelementptr %test.Driver, %test.Driver* %48, i32 0, i32 0
	%50 = load i8, i8* %49
	%51 = sext i8 %50 to i32
	%52 = call i32 (i8*, ...) @printf(i8* %47, i32 %51)
	%53 = getelementptr [17 x i8], [17 x i8]* @string.7ffa0c893047b73021f29c1b48c9892b, i32 0, i32 0
	%54 = call i32 (i8*, ...) @printf(i8* %53, i32 4)
	%55 = getelementptr [17 x i8], [17 x i8]* @string.451c6e0c15d560a4cfc5a46a02d53bfa, i32 0, i32 0
	%56 = call i32 (i8*, ...) @printf(i8* %55, i32 8)
	%57 = getelementptr [21 x i8], [21 x i8]* @string.adf1b889a9023b93577417ca23d21793, i32 0, i32 0
	%58 = call i32 (i8*, ...) @printf(i8* %57, i32 ptrtoint (i8** getelementptr (i8*, i8** null, i32 1) to i32))
	%59 = getelementptr [21 x i8], [21 x i8]* @string.846c6e4a2aac8de8fa1cce667d7cd481, i32 0, i32 0
	%60 = call i32 (i8*, ...) @printf(i8* %59, i32 ptrtoint (%test.Printer* getelementptr (%test.Printer, %test.Printer* null, i32 1) to i32))
	%61 = getelementptr [22 x i8], [22 x i8]* @string.fbd4ad59a9864656f8875863ac3b7dab, i32 0, i32 0
	%62 = call i32 (i8*, ...) @printf(i8* %61, i32 ptrtoint ([5 x i8]* getelementptr ([5 x i8], [5 x i8]* null, i32 1) to i32))
	br label %exit


exit:
	ret void

}

define void @test.pass_struct_pointer(%test.Printer* %printer) {
entry:
	%0 = alloca %test.Printer*
	store %test.Printer* %printer, %test.Printer** %0
	br label %body


body:
	%1 = load %test.Printer*, %test.Printer** %0
	%2 = getelementptr %test.Printer, %test.Printer* %1, i32 0, i32 2
	%3 = getelementptr %test.Driver, %test.Driver* %2, i32 0, i32 0
	store i8 3, i8* %3
	%4 = load %test.Printer*, %test.Printer** %0
	%5 = load %test.Printer*, %test.Printer** %0
	%6 = getelementptr [15 x i8], [15 x i8]* @string.263c2d145bd0257bade41874fd5a73ec, i32 0, i32 0
	call void @test.Printer.print(%test.Printer* %5, i8* %6)
	br label %exit


exit:
	ret void

}

define void @test.pass_array(i8* %buffer) {
entry:
	%0 = alloca i8*
	store i8* %buffer, i8** %0
	br label %body


body:
	%1 = load i8*, i8** %0
	%2 = getelementptr i8, i8* %1, i32 2
	store i8 2, i8* %2
	%3 = getelementptr [16 x i8], [16 x i8]* @string.b5abd14716ff1d42a2c76d0bae14c3cf, i32 0, i32 0
	%4 = load i8*, i8** %0
	%5 = getelementptr i8, i8* %4, i32 2
	%6 = load i8, i8* %5
	%7 = sext i8 %6 to i32
	%8 = call i32 (i8*, ...) @printf(i8* %3, i32 %7)
	br label %exit


exit:
	ret void

}

define void @test.Driver.print(%test.Driver* %this, i8* %message) {
entry:
	%0 = alloca %test.Driver*
	store %test.Driver* %this, %test.Driver** %0
	%1 = alloca i8*
	store i8* %message, i8** %1
	br label %body


body:
	%2 = load i8*, i8** %1
	%3 = call i32 @puts(i8* %2)
	%4 = getelementptr [10 x i8], [10 x i8]* @string.f229d6156f4a2e6f6e5c4ee96406192b, i32 0, i32 0
	%5 = load %test.Driver*, %test.Driver** %0
	%6 = getelementptr %test.Driver, %test.Driver* %5, i32 0, i32 0
	%7 = load i8, i8* %6
	%8 = sext i8 %7 to i32
	%9 = call i32 (i8*, ...) @printf(i8* %4, i32 %8)
	br label %exit


exit:
	ret void

}

define void @test.Printer.print(%test.Printer* %this, i8* %message) {
entry:
	%0 = alloca %test.Printer*
	store %test.Printer* %this, %test.Printer** %0
	%1 = alloca i8*
	store i8* %message, i8** %1
	br label %body


body:
	%2 = load %test.Printer*, %test.Printer** %0
	%3 = getelementptr %test.Printer, %test.Printer* %2, i32 0, i32 2
	%4 = load %test.Printer*, %test.Printer** %0
	%5 = getelementptr %test.Printer, %test.Printer* %4, i32 0, i32 2
	%6 = alloca %test.Driver*
	store %test.Driver* %5, %test.Driver** %6
	%7 = load %test.Driver*, %test.Driver** %6
	%8 = load i8*, i8** %1
	call void @test.Driver.print(%test.Driver* %7, i8* %8)
	br label %exit


exit:
	ret void

}

define void @test.test() {
entry:
	br label %body


body:
	call void @test.expression()
	call void @test.statement()
	call void @test.structs()
	call void @test.functions()
	br label %exit


exit:
	ret void

}
