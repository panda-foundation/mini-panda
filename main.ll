%global.Data = type { i8, float, [8 x i8] }

@global.Color.r = global i8 0
@global.Color.g = global i8 1
@global.Color.b = global i8 2
@string.cb091131e20d7842e7627e8736856b45 = constant [12 x i8] c"hello world\00"

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
	%0 = call i32 @puts(i8* bitcast ([12 x i8]* @string.cb091131e20d7842e7627e8736856b45 to i8*))
	br label %exit


exit:
	ret void

}
