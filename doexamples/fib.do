act main () {
    println(first(fibonacci(0)));
    println(first(fibonacci(1)));
    println(first(fibonacci(2)));
    println(first(fibonacci(3)));
    println(first(fibonacci(4)));
    println(first(fibonacci(5)));
    println(first(fibonacci(6)));
    println(first(fibonacci(7)));
    println(first(fibonacci(8)));
}

act println (val int) {
    tmpOut = val;
    tmpPrint();
}

act fibonacci (n int) (int, int){
    if n == 0 {
        return 0, 1;
    } else {
        a, b := fibonacci(n - 1);
        return b, a + b;
    }
}

act first (a int, b int) int {
    return a;
}