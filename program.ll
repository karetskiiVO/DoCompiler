@tmpOut = global i32 zeroinitializer

declare {} @tmpPrint()

define {} @f(i32 %arg_newVal) {
0:
	%newVal = alloca i32
	store i32 %arg_newVal, i32* %newVal
	%1 = load i32, i32* %newVal
	store i32 %1, i32* @tmpOut
	%2 = alloca {}
	%_ret = load {}, {}* %2
	ret {} %_ret
}

define {} @main() {
0:
	%struc = alloca { i32, i32, i32, i1 }
	store { i32, i32, i32, i1 } zeroinitializer, { i32, i32, i32, i1 }* %struc
	%1 = getelementptr i32, { i32, i32, i32, i1 }* %struc, i32 0
	store i32 10, i32* %1
	%2 = getelementptr i32, { i32, i32, i32, i1 }* %struc, i32 1
	store i32 11, i32* %2
	%3 = getelementptr i32, { i32, i32, i32, i1 }* %struc, i32 2
	store i32 12, i32* %3
	%4 = getelementptr i1, { i32, i32, i32, i1 }* %struc, i32 3
	store i1 true, i1* %4
	%5 = call {} @tmpPrint()
	%6 = call {} @f(i32 2)
	%7 = call {} @tmpPrint()
	%a = alloca i32
	store i32 zeroinitializer, i32* %a
	%8 = load i32, i32* @tmpOut
	store i32 %8, i32* %a
	%9 = getelementptr i32, { i32, i32, i32, i1 }* %struc, i32 0
	%10 = load i32, i32* %9
	%11 = call {} @println(i32 %10)
	br i1 true, label %12, label %16

12:
	%13 = getelementptr i32, { i32, i32, i32, i1 }* %struc, i32 1
	%14 = load i32, i32* %13
	%15 = call {} @println(i32 %14)
	br label %20

16:
	%17 = getelementptr i32, { i32, i32, i32, i1 }* %struc, i32 1
	%18 = load i32, i32* %17
	%19 = call {} @println(i32 %18)
	br label %20

20:
	%21 = alloca {}
	%_ret = load {}, {}* %21
	ret {} %_ret
}

define {} @println(i32 %arg_val) {
0:
	%val = alloca i32
	store i32 %arg_val, i32* %val
	%1 = load i32, i32* %val
	store i32 %1, i32* @tmpOut
	%2 = call {} @tmpPrint()
	%3 = alloca {}
	%_ret = load {}, {}* %3
	ret {} %_ret
}
