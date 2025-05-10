@tmpOut = global i32 zeroinitializer
@a = global i32 zeroinitializer
@b = global i32 zeroinitializer
@c = global i32 zeroinitializer
@d = global i32 zeroinitializer
@e = global i32 zeroinitializer

declare {} @tmpPrint()

define {} @f() {
0:
	store i32 1, i32* @a
	%1 = alloca {}
	%_ret = load {}, {}* %1
	ret {} %_ret
}

define {} @main() {
0:
	%1 = call {} @tmpPrint()
	%2 = call {} @f()
	%3 = call {} @tmpPrint()
	%4 = call { i32, i32 } @h()
	%5 = extractvalue { i32, i32 } %4, 0
	%6 = extractvalue { i32, i32 } %4, 1
	store i32 %5, i32* @b
	store i32 %6, i32* @c
	%7 = alloca {}
	%_ret = load {}, {}* %7
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
	%2 = call {} @g()
	br label %4

3:
	br label %4

4:
	br i1 true, label %5, label %10

5:
	%6 = call {} @g()
	br label %11

7:
	br i1 false, label %8, label %9

8:
	store i32 2, i32* @e
	br label %10

9:
	store i32 4, i32* @c
	br label %10

10:
	br label %11

11:
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
