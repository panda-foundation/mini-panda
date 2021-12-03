%global.Data = type { i8, float, [8 x i8] }
%global.NewData = type { %global.Data, i8 }

@global.Color.r = global i8 0
@global.Color.g = global i8 1
@global.Color.b = global i8 2
@global.values = global [5 x i8] [i8 1, i8 2, i8 3, i8 4, i8 5]
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
@string.12625b519c0ef75b350a9963cafc3f42 = constant [17 x i8] c"shouldn't happen\00"
@string.7c13f0ed550e89d5fe0dab15a8790a6b = constant [9 x i8] c"I'm else\00"
@string.e509c213bf338f03d246b720ec617c01 = constant [11 x i8] c"loop: %d \0A\00"
@string.ba86886fe05268c3936c4741a0d07a6e = constant [14 x i8] c"switch case 0\00"
@string.162d9796d41e74535694f9688ea21a49 = constant [14 x i8] c"switch case 3\00"
@string.ba4ed99596c7e9aa2595a8f23577c2a9 = constant [16 x i8] c"values[0]: %d \0A\00"
@string.291bd270faa3b66dd92c4af584f01044 = constant [16 x i8] c"values[4]: %d \0A\00"
@string.bcfa829c5c86235c99443fb88b9d9699 = constant [15 x i8] c"array[2]: %d \0A\00"
@string.47b89087c0546b3ff5a4ec613cfd034c = constant [19 x i8] c"this.integer: %d \0A\00"

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
	call void @global.extern()
	call void @global.expression()
	call void @global.statement()
	call void @global.structs()
	call void @global.pointers()
	call void @global.conversions()
	call void @global.functions()
	br label %exit


exit:
	ret void

}

define void @global.extern() {
entry:
	%0 = alloca i32
	br label %body


body:
	%1 = call i32 @puts(i8* getelementptr ([12 x i8], [12 x i8]* @string.cb091131e20d7842e7627e8736856b45, i32 0, i32 0))
	store i32 1, i32* %0
	%2 = load i32, i32* %0
	%3 = call i32 (i8*, ...) @printf(i8* getelementptr ([12 x i8], [12 x i8]* @string.17dc8feeedfe47c12c0d109e5e0da235, i32 0, i32 0), i32 %2)
	store i32 2, i32* %0
	%4 = load i32, i32* %0
	%5 = call i32 (i8*, ...) @printf(i8* getelementptr ([12 x i8], [12 x i8]* @string.17dc8feeedfe47c12c0d109e5e0da235, i32 0, i32 0), i32 %4)
	%6 = call i32 @global.add(i32 1, i32 2)
	%7 = call i32 (i8*, ...) @printf(i8* getelementptr ([13 x i8], [13 x i8]* @string.328de3303ca25a967f81f9c8c805e8a1, i32 0, i32 0), i32 %6)
	br label %exit


exit:
	ret void

}

define i32 @global.add(i32 %a, i32 %b) {
entry:
	%0 = alloca i32
	store i32 %a, i32* %0
	%1 = alloca i32
	store i32 %b, i32* %1
	%2 = alloca i32
	br label %body


body:
	%3 = load i32, i32* %0
	%4 = load i32, i32* %1
	%5 = add i32 %3, %4
	store i32 %5, i32* %2
	br label %exit


exit:
	%6 = load i32, i32* %2
	ret i32 %6

}

define void @global.expression() {
entry:
	br label %body


body:
	call void @global.unary()
	call void @global.increment_decrement()
	call void @global.binary()
	call void @global.parentheses()
	call void @global.literal()
	call void @global.subscripting()
	call void @global.member_access()
	br label %exit


exit:
	ret void

}

define void @global.unary() {
entry:
	%0 = alloca i32
	%1 = alloca i32
	%2 = alloca i32
	%3 = alloca i1
	%4 = alloca i1
	br label %body


body:
	store i32 1, i32* %0
	%5 = load i32, i32* %0
	%6 = call i32 (i8*, ...) @printf(i8* getelementptr ([9 x i8], [9 x i8]* @string.f9b6d891c5ca674309c459ad55eb01c8, i32 0, i32 0), i32 %5)
	%7 = load i32, i32* %0
	%8 = sub i32 0, %7
	store i32 %8, i32* %1
	%9 = load i32, i32* %1
	%10 = call i32 (i8*, ...) @printf(i8* getelementptr ([9 x i8], [9 x i8]* @string.f52c63a936a31e2b2d03c5c746b8d5b9, i32 0, i32 0), i32 %9)
	%11 = xor i32 1, -1
	store i32 %11, i32* %2
	%12 = load i32, i32* %2
	%13 = call i32 (i8*, ...) @printf(i8* getelementptr ([9 x i8], [9 x i8]* @string.da88a2e8a843d3d238dc43a4378c6887, i32 0, i32 0), i32 %12)
	store i1 true, i1* %3
	%14 = load i1, i1* %3
	%15 = sext i1 %14 to i32
	%16 = call i32 (i8*, ...) @printf(i8* getelementptr ([9 x i8], [9 x i8]* @string.c3c4ff0f83dad5387535d315826c22f8, i32 0, i32 0), i32 %15)
	%17 = load i1, i1* %3
	%18 = xor i1 %17, true
	store i1 %18, i1* %4
	%19 = load i1, i1* %4
	%20 = sext i1 %19 to i32
	%21 = call i32 (i8*, ...) @printf(i8* getelementptr ([9 x i8], [9 x i8]* @string.e40a403ffbdf9c2c70921d6bb7739cd8, i32 0, i32 0), i32 %20)
	br label %exit


exit:
	ret void

}

define void @global.increment_decrement() {
entry:
	%0 = alloca i32
	br label %body


body:
	store i32 10, i32* %0
	%1 = load i32, i32* %0
	%2 = add i32 %1, 1
	store i32 %2, i32* %0
	%3 = load i32, i32* %0
	%4 = call i32 (i8*, ...) @printf(i8* getelementptr ([9 x i8], [9 x i8]* @string.2dcc97a590ca083991ffe9b43c08dd02, i32 0, i32 0), i32 %3)
	store i32 20, i32* %0
	%5 = load i32, i32* %0
	%6 = sub i32 %5, 1
	store i32 %6, i32* %0
	%7 = load i32, i32* %0
	%8 = call i32 (i8*, ...) @printf(i8* getelementptr ([9 x i8], [9 x i8]* @string.2dcc97a590ca083991ffe9b43c08dd02, i32 0, i32 0), i32 %7)
	br label %exit


exit:
	ret void

}

define void @global.binary() {
entry:
	%0 = alloca i32
	br label %body


body:
	store i32 0, i32* %0
	store i32 5, i32* %0
	%1 = load i32, i32* %0
	%2 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %1)
	%3 = load i32, i32* %0
	%4 = add i32 %3, 5
	store i32 %4, i32* %0
	%5 = load i32, i32* %0
	%6 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %5)
	%7 = load i32, i32* %0
	%8 = sub i32 %7, 1
	store i32 %8, i32* %0
	%9 = load i32, i32* %0
	%10 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %9)
	%11 = load i32, i32* %0
	%12 = mul i32 %11, 2
	store i32 %12, i32* %0
	%13 = load i32, i32* %0
	%14 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %13)
	%15 = load i32, i32* %0
	%16 = sdiv i32 %15, 9
	store i32 %16, i32* %0
	%17 = load i32, i32* %0
	%18 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %17)
	store i32 11, i32* %0
	%19 = load i32, i32* %0
	%20 = srem i32 %19, 4
	store i32 %20, i32* %0
	%21 = load i32, i32* %0
	%22 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %21)
	%23 = load i32, i32* %0
	%24 = shl i32 %23, 2
	store i32 %24, i32* %0
	%25 = load i32, i32* %0
	%26 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %25)
	%27 = load i32, i32* %0
	%28 = ashr i32 %27, 1
	store i32 %28, i32* %0
	%29 = load i32, i32* %0
	%30 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %29)
	%31 = load i32, i32* %0
	%32 = or i32 %31, 15
	store i32 %32, i32* %0
	%33 = load i32, i32* %0
	%34 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %33)
	%35 = load i32, i32* %0
	%36 = xor i32 %35, 8
	store i32 %36, i32* %0
	%37 = load i32, i32* %0
	%38 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %37)
	%39 = load i32, i32* %0
	%40 = and i32 %39, 6
	store i32 %40, i32* %0
	%41 = load i32, i32* %0
	%42 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %41)
	%43 = or i32 15, 8
	%44 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %43)
	%45 = xor i32 15, 8
	%46 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %45)
	%47 = and i32 15, 8
	%48 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %47)
	%49 = icmp eq i32 10, 10
	%50 = sext i1 %49 to i32
	%51 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %50)
	%52 = icmp ne i32 10, 10
	%53 = sext i1 %52 to i32
	%54 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %53)
	%55 = icmp slt i32 10, 10
	%56 = sext i1 %55 to i32
	%57 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %56)
	%58 = icmp sle i32 10, 10
	%59 = sext i1 %58 to i32
	%60 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %59)
	%61 = icmp sgt i32 10, 10
	%62 = sext i1 %61 to i32
	%63 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %62)
	%64 = icmp sge i32 10, 10
	%65 = sext i1 %64 to i32
	%66 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %65)
	%67 = shl i32 10, 3
	%68 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %67)
	%69 = ashr i32 10, 1
	%70 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %69)
	%71 = add i32 5, 3
	%72 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %71)
	%73 = sub i32 5, 3
	%74 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %73)
	%75 = mul i32 5, 3
	%76 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %75)
	%77 = sdiv i32 5, 3
	%78 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %77)
	%79 = srem i32 5, 3
	%80 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.fc631314303f7db146188786b60902e8, i32 0, i32 0), i32 %79)
	%81 = or i1 true, false
	%82 = sext i1 %81 to i32
	%83 = call i32 (i8*, ...) @printf(i8* getelementptr ([11 x i8], [11 x i8]* @string.a753fba743a9e6b08cb7a2627f69b75d, i32 0, i32 0), i32 %82)
	%84 = and i1 true, false
	%85 = sext i1 %84 to i32
	%86 = call i32 (i8*, ...) @printf(i8* getelementptr ([11 x i8], [11 x i8]* @string.a753fba743a9e6b08cb7a2627f69b75d, i32 0, i32 0), i32 %85)
	br label %exit


exit:
	ret void

}

define void @global.parentheses() {
entry:
	br label %body


body:
	%0 = mul i32 add (i32 5, i32 6), 3
	%1 = call i32 (i8*, ...) @printf(i8* getelementptr ([18 x i8], [18 x i8]* @string.9155af3e03234ca6017e6a626fa48d60, i32 0, i32 0), i32 %0)
	br label %exit


exit:
	ret void

}

define void @global.literal() {
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
	%7 = load i8, i8* %0
	%8 = sext i8 %7 to i32
	%9 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.6815af516458351e77683ead5f501317, i32 0, i32 0), i32 %8)
	store i1 true, i1* %1
	%10 = load i1, i1* %1
	%11 = sext i1 %10 to i32
	%12 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.d58ddb72e75f1acfc4203e33bddc08a1, i32 0, i32 0), i32 %11)
	store i8 123, i8* %2
	%13 = load i8, i8* %2
	%14 = sext i8 %13 to i32
	%15 = call i32 (i8*, ...) @printf(i8* getelementptr ([8 x i8], [8 x i8]* @string.5b8b2fafadbddfa000cd0e716725d4a4, i32 0, i32 0), i32 %14)
	store half 0x4009200000000000, half* %3
	%16 = load half, half* %3
	%17 = fpext half %16 to double
	%18 = call i32 (i8*, ...) @printf(i8* getelementptr ([10 x i8], [10 x i8]* @string.dd42ef93dc06a72b063baa72848d660c, i32 0, i32 0), double %17)
	store float 0x40091EB860000000, float* %4
	%19 = load float, float* %4
	%20 = fpext float %19 to double
	%21 = call i32 (i8*, ...) @printf(i8* getelementptr ([10 x i8], [10 x i8]* @string.dc24ff6a55a1c588a346f9dff66c25a0, i32 0, i32 0), double %20)
	store double 0x40091EB851EB851F, double* %5
	%22 = load double, double* %5
	%23 = call i32 (i8*, ...) @printf(i8* getelementptr ([10 x i8], [10 x i8]* @string.7a828f7c003ac662930a932d14c84f48, i32 0, i32 0), double %22)
	store i8* getelementptr ([12 x i8], [12 x i8]* @string.3aff445dea2b63e4d3b135c5219ba7dc, i32 0, i32 0), i8** %6
	%24 = load i8*, i8** %6
	%25 = call i32 @puts(i8* %24)
	br label %exit


exit:
	ret void

}

define void @global.subscripting() {
entry:
	%0 = alloca [8 x i8]
	br label %body


body:
	store [8 x i8] zeroinitializer, [8 x i8]* %0
	%1 = getelementptr [8 x i8], [8 x i8]* %0, i32 0, i32 0
	%2 = load i8, i8* %1
	%3 = sext i8 %2 to i32
	%4 = call i32 (i8*, ...) @printf(i8* getelementptr ([15 x i8], [15 x i8]* @string.ccbd06f65fb69a974bb7bbe132352fd5, i32 0, i32 0), i32 %3)
	%5 = getelementptr [8 x i8], [8 x i8]* %0, i32 0, i32 3
	store i8 3, i8* %5
	%6 = getelementptr [8 x i8], [8 x i8]* %0, i32 0, i32 3
	%7 = load i8, i8* %6
	%8 = sext i8 %7 to i32
	%9 = call i32 (i8*, ...) @printf(i8* getelementptr ([15 x i8], [15 x i8]* @string.502edb90c5d63a7982b92c4846005a12, i32 0, i32 0), i32 %8)
	br label %exit


exit:
	ret void

}

define void @global.member_access() {
entry:
	%0 = alloca %global.Data
	%1 = alloca %global.NewData
	br label %body


body:
	store %global.Data zeroinitializer, %global.Data* %0
	%2 = getelementptr %global.Data, %global.Data* %0, i32 0, i32 0
	%3 = load i8, i8* %2
	%4 = sext i8 %3 to i32
	%5 = call i32 (i8*, ...) @printf(i8* getelementptr ([19 x i8], [19 x i8]* @string.5b8a1afb98c4b2718e7e1f29b27539e6, i32 0, i32 0), i32 %4)
	%6 = getelementptr %global.Data, %global.Data* %0, i32 0, i32 0
	store i8 5, i8* %6
	%7 = getelementptr %global.Data, %global.Data* %0, i32 0, i32 0
	%8 = load i8, i8* %7
	%9 = sext i8 %8 to i32
	%10 = call i32 (i8*, ...) @printf(i8* getelementptr ([19 x i8], [19 x i8]* @string.5b8a1afb98c4b2718e7e1f29b27539e6, i32 0, i32 0), i32 %9)
	%11 = getelementptr %global.Data, %global.Data* %0, i32 0, i32 1
	store float 0x40091EB860000000, float* %11
	%12 = getelementptr %global.Data, %global.Data* %0, i32 0, i32 1
	%13 = load float, float* %12
	%14 = fpext float %13 to double
	%15 = call i32 (i8*, ...) @printf(i8* getelementptr ([17 x i8], [17 x i8]* @string.4703b4d82797dc9d0990618793a935a5, i32 0, i32 0), double %14)
	%16 = getelementptr %global.Data, %global.Data* %0, i32 0, i32 2
	%17 = getelementptr [8 x i8], [8 x i8]* %16, i32 0, i32 3
	store i8 3, i8* %17
	%18 = getelementptr %global.Data, %global.Data* %0, i32 0, i32 2
	%19 = getelementptr [8 x i8], [8 x i8]* %18, i32 0, i32 3
	%20 = load i8, i8* %19
	%21 = sext i8 %20 to i32
	%22 = call i32 (i8*, ...) @printf(i8* getelementptr ([20 x i8], [20 x i8]* @string.e839e54fd3fe1d952dd8a33030d97634, i32 0, i32 0), i32 %21)
	store %global.NewData zeroinitializer, %global.NewData* %1
	%23 = getelementptr %global.NewData, %global.NewData* %1, i32 0, i32 1
	store i8 5, i8* %23
	%24 = getelementptr %global.NewData, %global.NewData* %1, i32 0, i32 1
	%25 = load i8, i8* %24
	%26 = sext i8 %25 to i32
	%27 = call i32 (i8*, ...) @printf(i8* getelementptr ([21 x i8], [21 x i8]* @string.7499c41e0f87337e5f3f93200f97701e, i32 0, i32 0), i32 %26)
	%28 = getelementptr %global.NewData, %global.NewData* %1, i32 0, i32 0
	%29 = getelementptr %global.Data, %global.Data* %28, i32 0, i32 0
	store i8 8, i8* %29
	%30 = getelementptr %global.NewData, %global.NewData* %1, i32 0, i32 0
	%31 = getelementptr %global.Data, %global.Data* %30, i32 0, i32 0
	%32 = load i8, i8* %31
	%33 = sext i8 %32 to i32
	%34 = call i32 (i8*, ...) @printf(i8* getelementptr ([28 x i8], [28 x i8]* @string.5c65bb89388b87cc845b7ed6cc4e0933, i32 0, i32 0), i32 %33)
	%35 = getelementptr %global.NewData, %global.NewData* %1, i32 0, i32 0
	%36 = getelementptr %global.Data, %global.Data* %35, i32 0, i32 2
	%37 = getelementptr [8 x i8], [8 x i8]* %36, i32 0, i32 3
	store i8 9, i8* %37
	%38 = getelementptr %global.NewData, %global.NewData* %1, i32 0, i32 0
	%39 = getelementptr %global.Data, %global.Data* %38, i32 0, i32 2
	%40 = getelementptr [8 x i8], [8 x i8]* %39, i32 0, i32 3
	%41 = load i8, i8* %40
	%42 = sext i8 %41 to i32
	%43 = call i32 (i8*, ...) @printf(i8* getelementptr ([29 x i8], [29 x i8]* @string.18fc68733fbf6df1ade57d0706714eec, i32 0, i32 0), i32 %42)
	br label %exit


exit:
	ret void

}

define void @global.statement() {
entry:
	%0 = alloca i8
	br label %body


body:
	store i8 10, i8* %0
	%1 = load i8, i8* %0
	%2 = icmp uge i8 %1, 10
	br i1 %2, label %6, label %3


exit:
	ret void


3:
	%4 = load i8, i8* %0
	%5 = icmp ugt i8 %4, 100
	br i1 %5, label %11, label %13


6:
	%7 = load i8, i8* %0
	%8 = sext i8 %7 to i32
	%9 = call i32 (i8*, ...) @printf(i8* getelementptr ([14 x i8], [14 x i8]* @string.07af74d61c4bcfd65e300c22c36df6a3, i32 0, i32 0), i32 %8)
	br label %3


10:
	store i8 0, i8* %0
	br label %17


11:
	%12 = call i32 @puts(i8* getelementptr ([17 x i8], [17 x i8]* @string.12625b519c0ef75b350a9963cafc3f42, i32 0, i32 0))
	br label %10


13:
	%14 = call i32 @puts(i8* getelementptr ([9 x i8], [9 x i8]* @string.7c13f0ed550e89d5fe0dab15a8790a6b, i32 0, i32 0))
	br label %10


15:
	store i8 3, i8* %0
	%16 = load i8, i8* %0
	switch i8 %16, label %28 [
		i8 0, label %29
		i8 3, label %31
	]


17:
	%18 = load i8, i8* %0
	%19 = icmp ult i8 %18, 10
	br i1 %19, label %23, label %15


20:
	%21 = load i8, i8* %0
	%22 = add i8 %21, 1
	store i8 %22, i8* %0
	br label %17


23:
	%24 = load i8, i8* %0
	%25 = sext i8 %24 to i32
	%26 = call i32 (i8*, ...) @printf(i8* getelementptr ([11 x i8], [11 x i8]* @string.e509c213bf338f03d246b720ec617c01, i32 0, i32 0), i32 %25)
	br label %20


27:
	br label %exit


28:
	br label %27


29:
	%30 = call i32 @puts(i8* getelementptr ([14 x i8], [14 x i8]* @string.ba86886fe05268c3936c4741a0d07a6e, i32 0, i32 0))
	br label %27


31:
	%32 = call i32 @puts(i8* getelementptr ([14 x i8], [14 x i8]* @string.162d9796d41e74535694f9688ea21a49, i32 0, i32 0))
	br label %27

}

define void @global.pointers() {
entry:
	br label %body


body:
	br label %exit


exit:
	ret void

}

define void @global.conversions() {
entry:
	br label %body


body:
	br label %exit


exit:
	ret void

}

define void @global.structs() {
entry:
	%0 = alloca %global.Data
	br label %body


body:
	%1 = getelementptr [5 x i8], [5 x i8]* @global.values, i32 0, i32 0
	%2 = load i8, i8* %1
	%3 = sext i8 %2 to i32
	%4 = call i32 (i8*, ...) @printf(i8* getelementptr ([16 x i8], [16 x i8]* @string.ba4ed99596c7e9aa2595a8f23577c2a9, i32 0, i32 0), i32 %3)
	%5 = getelementptr [5 x i8], [5 x i8]* @global.values, i32 0, i32 4
	%6 = load i8, i8* %5
	%7 = sext i8 %6 to i32
	%8 = call i32 (i8*, ...) @printf(i8* getelementptr ([16 x i8], [16 x i8]* @string.291bd270faa3b66dd92c4af584f01044, i32 0, i32 0), i32 %7)
	store %global.Data { i8 1, float 0x40091EB860000000, [8 x i8] [i8 1, i8 2, i8 3, i8 4, i8 5, i8 6, i8 7, i8 8] }, %global.Data* %0
	%9 = getelementptr %global.Data, %global.Data* %0, i32 0, i32 0
	%10 = load i8, i8* %9
	%11 = sext i8 %10 to i32
	%12 = call i32 (i8*, ...) @printf(i8* getelementptr ([19 x i8], [19 x i8]* @string.5b8a1afb98c4b2718e7e1f29b27539e6, i32 0, i32 0), i32 %11)
	%13 = getelementptr %global.Data, %global.Data* %0, i32 0, i32 1
	%14 = load float, float* %13
	%15 = fpext float %14 to double
	%16 = call i32 (i8*, ...) @printf(i8* getelementptr ([17 x i8], [17 x i8]* @string.4703b4d82797dc9d0990618793a935a5, i32 0, i32 0), double %15)
	%17 = getelementptr %global.Data, %global.Data* %0, i32 0, i32 2
	%18 = getelementptr [8 x i8], [8 x i8]* %17, i32 0, i32 3
	%19 = load i8, i8* %18
	%20 = sext i8 %19 to i32
	%21 = call i32 (i8*, ...) @printf(i8* getelementptr ([20 x i8], [20 x i8]* @string.e839e54fd3fe1d952dd8a33030d97634, i32 0, i32 0), i32 %20)
	call void @global.Data.print_integer(%global.Data* %0)
	call void @global.call_print(%global.Data* %0)
	%22 = getelementptr %global.Data, %global.Data* %0, i32 0, i32 2
	%23 = getelementptr [8 x i8], [8 x i8]* %22, i32 0, i32 0
	call void @global.call_array(i8* %23)
	br label %exit


exit:
	ret void

}

define void @global.call_print(%global.Data* %data) {
entry:
	br label %body


body:
	%0 = getelementptr %global.Data, %global.Data* %data, i32 0, i32 0
	store i8 3, i8* %0
	call void @global.Data.print_integer(%global.Data* %data)
	br label %exit


exit:
	ret void

}

define void @global.call_array(i8* %data) {
entry:
	br label %body


body:
	%0 = getelementptr i8, i8* %data, i32 2
	store i8 2, i8* %0
	%1 = getelementptr i8, i8* %data, i32 2
	%2 = load i8, i8* %1
	%3 = sext i8 %2 to i32
	%4 = call i32 (i8*, ...) @printf(i8* getelementptr ([15 x i8], [15 x i8]* @string.bcfa829c5c86235c99443fb88b9d9699, i32 0, i32 0), i32 %3)
	br label %exit


exit:
	ret void

}

define void @global.functions() {
entry:
	br label %body


body:
	br label %exit


exit:
	ret void

}

define void @global.Data.print_integer(%global.Data* %this) {
entry:
	br label %body


body:
	%0 = getelementptr %global.Data, %global.Data* %this, i32 0, i32 0
	%1 = load i8, i8* %0
	%2 = sext i8 %1 to i32
	%3 = call i32 (i8*, ...) @printf(i8* getelementptr ([19 x i8], [19 x i8]* @string.47b89087c0546b3ff5a4ec613cfd034c, i32 0, i32 0), i32 %2)
	br label %exit


exit:
	ret void

}
