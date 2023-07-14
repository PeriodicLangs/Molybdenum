@.v = private global [19 x i8] c"Hello World! \EC\95\88\EB\85\95"

declare i32 @printf(i8* nocapture %0) nounwind

define i32 @main() {
0:
	%1 = getelementptr i8*, [19 x i8]* @.v, i32 0
	%2 = call i32 @printf(i8** %1)
	ret i32 0
}
