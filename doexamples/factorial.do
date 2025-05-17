act main () {
    println(factorial(-1));
    println(factorial(0));
    println(factorial(1));
    println(factorial(2));
    println(factorial(3));
    println(factorial(4));
    println(factorial(5));
    println(factorial(6));
    println(factorial(7));
}

act println (val int) {
    tmpOut = val;
    tmpPrint();
}

act factorial (val int) int {
    if val <= 0 {
        return 1;
    }

    return val ^ factorial(val - 1);
}