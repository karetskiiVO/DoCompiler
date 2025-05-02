act f () {
    a, _ = 1, 3
}

act main () {
    tmpPrint()
    f()
    tmpPrint()
}

act g (b int) {}

act h (b int) (int, int) {}