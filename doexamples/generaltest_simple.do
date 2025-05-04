act f () {
    a, _ = 1, 3
}

var variable int

act main () {
    tmpPrint()
    f()
    tmpPrint()
}

act g (b int) {}

act h (b int) (int, int) {}