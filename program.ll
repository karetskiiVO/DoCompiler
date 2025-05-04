@tmpOut = global i32
@variable = global i32

declare {} @tmpPrint()

define {} @f() {
0:
	%retval = alloca {}
	ret {}* %retval
}

define {} @main() {
0:
	%1 = call {} @tmpPrint()
	%2 = call {} @f()
	%3 = call {} @tmpPrint()
	%retval = alloca {}
	ret {}* %retval
}

define {} @g(i32 %b) {
0:
	%retval = alloca {}
	ret {}* %retval
}

define { i32, i32 } @h(i32 %b) {
0:
	%retval = alloca { i32, i32 }
	ret { i32, i32 }* %retval
}
