@tmpOut = global i32 zeroinitializer

declare {} @tmpPrint()

define {} @f(i32 %arg_newVal) {
0:
	%newVal = alloca i32
	store i32 %arg_newVal, i32* %newVal
	%1 = load i32, i32* @tmpOut
	%2 = load i32, i32* %newVal
	store i32 %2, i32* @tmpOut
	%3 = alloca {}
	%_ret = load {}, {}* %3
	ret {} %_ret
}

define {} @main() {
0:
	%1 = call {} @tmpPrint()
	%2 = call {} @f(i32 2)
	%3 = call {} @tmpPrint()
	%4 = alloca {}
	%_ret = load {}, {}* %4
	ret {} %_ret
}
