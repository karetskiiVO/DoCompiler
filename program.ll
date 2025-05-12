@tmpOut = global i32 zeroinitializer

declare {} @tmpPrint()

define {} @f() {
0:
	%1 = load i32, i32* @tmpOut
	store i32 1, i32* @tmpOut
	%2 = alloca {}
	%_ret = load {}, {}* %2
	ret {} %_ret
}

define {} @main() {
0:
	%1 = call {} @tmpPrint()
	%2 = call {} @f()
	%3 = call {} @tmpPrint()
	%4 = alloca {}
	%_ret = load {}, {}* %4
	ret {} %_ret
}
