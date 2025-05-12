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
	%struc = alloca { i32, i32, i32, i1 }
	store { i32, i32, i32, i1 } zeroinitializer, { i32, i32, i32, i1 }* %struc
	%1 = getelementptr i32, { i32, i32, i32, i1 }* %struc, i64 0
	%2 = load i32, i32* %1
	%3 = getelementptr i32, { i32, i32, i32, i1 }* %struc, i64 0
	store i32 10, i32* %3
	%4 = getelementptr i32, { i32, i32, i32, i1 }* %struc, i64 1
	%5 = load i32, i32* %4
	%6 = getelementptr i32, { i32, i32, i32, i1 }* %struc, i64 1
	store i32 11, i32* %6
	%7 = getelementptr i32, { i32, i32, i32, i1 }* %struc, i64 2
	%8 = load i32, i32* %7
	%9 = getelementptr i32, { i32, i32, i32, i1 }* %struc, i64 2
	store i32 12, i32* %9
	%10 = call {} @tmpPrint()
	%11 = call {} @f(i32 2)
	%12 = call {} @tmpPrint()
	%a = alloca i32
	store i32 zeroinitializer, i32* %a
	%13 = load i32, i32* %a
	%14 = load i32, i32* @tmpOut
	store i32 %14, i32* %a
	%15 = getelementptr i32, { i32, i32, i32, i1 }* %struc, i64 0
	%16 = load i32, i32* %15
	%17 = call {} @println(i32 %16)
	br i1 true, label %18, label %22

18:
	%19 = getelementptr i32, { i32, i32, i32, i1 }* %struc, i64 1
	%20 = load i32, i32* %19
	%21 = call {} @println(i32 %20)
	br label %26

22:
	%23 = getelementptr i32, { i32, i32, i32, i1 }* %struc, i64 1
	%24 = load i32, i32* %23
	%25 = call {} @println(i32 %24)
	br label %26

26:
	%27 = alloca {}
	%_ret = load {}, {}* %27
	ret {} %_ret
}

define {} @println(i32 %arg_val) {
0:
	%val = alloca i32
	store i32 %arg_val, i32* %val
	%1 = load i32, i32* @tmpOut
	%2 = load i32, i32* %val
	store i32 %2, i32* @tmpOut
	%3 = call {} @tmpPrint()
	%4 = alloca {}
	%_ret = load {}, {}* %4
	ret {} %_ret
}
