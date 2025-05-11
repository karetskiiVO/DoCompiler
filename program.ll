@tmpOut = global i32 zeroinitializer
@a = global i32 zeroinitializer
@b = global i32 zeroinitializer
@c = global i32 zeroinitializer
@d = global i32 zeroinitializer
@e = global i32 zeroinitializer

declare {} @tmpPrint()

define {} @f() {
0:
	%1 = load i32*, i32* @a
	store i32 1, i32* @a
	%2 = alloca {}
	%_ret = load {}, {}* %2
	ret {} %_ret
}

define {} @main() {
0:
	%a = alloca i32
	store i32 zeroinitializer, i32* %a
	%aboba = alloca i32
	store i32 zeroinitializer, i32* %aboba
	%c = alloca i32
	store i32 zeroinitializer, i32* %c
	%1 = call {} @tmpPrint()
	%2 = call {} @f()
	%3 = call {} @tmpPrint()
	%4 = load i32*, i32* @b
	%5 = load i32*, i32* %c
	%6 = load i32*, i32* %a
	%7 = call { i32, i32 } @h(i32* %6)
	%8 = extractvalue { i32, i32 } %7, 0
	%9 = extractvalue { i32, i32 } %7, 1
	store i32 %8, i32* @b
	store i32 %9, i32* %c
	%10 = alloca {}
	%_ret = load {}, {}* %10
	ret {} %_ret
}

define {} @g(i32 %b) {
0:
	%1 = alloca {}
	%_ret = load {}, {}* %1
	ret {} %_ret
}

define { i32, i32 } @h(i32 %b) {
0:
	br i1 true, label %1, label %3

1:
	%2 = call {} @g(i32 1)
	br label %4

3:
	br label %4

4:
	br i1 true, label %5, label %12

5:
	%6 = call {} @g(i32 2)
	br label %13

7:
	br i1 false, label %8, label %10

8:
	%9 = load i32*, i32* @e
	store i32 2, i32* @e
	br label %12

10:
	%11 = load i32*, i32* @c
	store i32 4, i32* @c
	br label %12

12:
	br label %13

13:
	unreachable
}

define { i32, i32 } @t() {
0:
	%1 = alloca { i32, i32 }
	%2 = getelementptr i32, { i32, i32 }* %1, i64 0
	store i32 1, i32* %2
	%3 = getelementptr i32, { i32, i32 }* %1, i64 1
	store i32 0, i32* %3
	%_ret = load { i32, i32 }, { i32, i32 }* %1
	ret { i32, i32 } %_ret

4:
	unreachable
}
