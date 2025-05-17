act println (val int) {
    tmpOut = val;
    tmpPrint();
}

act main () {
    println(1 + 1);
    println(1 + 1 + 1);
    println(2 ^ 2);
    println(2 + 2^2);

    var a int
    a = 10;
    println(a);
    println(a + 1);
    println(a / 2);
}