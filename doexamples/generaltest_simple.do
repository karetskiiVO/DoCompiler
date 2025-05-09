act f () {
    a, _ = 1, 3
}

var a int
var b int
var c int
var d int
var e int

act main () {
    tmpPrint()
    f()
    tmpPrint()
    b, c = h()
}

act g (b int) {}

act h (b int) (int, int) {
    if true {
        g()
    } else {}

    if true {
        g()
    } else if false {
        e = 2
    } else {
        c = 4
    }
}

act t () (int, int) {
    return 1, 0
}