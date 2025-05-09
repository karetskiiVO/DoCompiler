@tmpOut = global i32
@a = global i32
@b = global i32
@c = global i32
@d = global i32
@e = global i32

declare {} @tmpPrint()

define {} @f() {
0:
	store i32 1, i32* @a
	%"#ret" = alloca {}
	ret {}* %"#ret"
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
	%"#ret" = alloca {}
	ret {}* %"#ret"
}

define {} @g(i32 %b) {
0:
	%"#ret" = alloca {}
	ret {}* %"#ret"
}

define { i32, i32 } @h(i32 %b) {
0:
	%"#ret" = alloca { i32, i32 }
	ret { i32, i32 }* %"#ret"
}

define { i32, i32 } @t() {
0:
	%"#ret" = alloca { i32, i32 }
	ret { i32, i32 }* %"#ret"
}
