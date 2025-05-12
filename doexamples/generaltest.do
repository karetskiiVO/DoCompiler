with A struct {}

act f (newVal int) {
    tmpOut = newVal
}

act main () {
    tmpPrint()
    f(2)
    tmpPrint()
}