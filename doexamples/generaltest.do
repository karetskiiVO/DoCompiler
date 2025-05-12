with A struct {}

act f (newVal int) {
    tmpOut = newVal
}

act main () {
    tmpPrint()
    f(2)
    tmpPrint()

    if true {
        println(3)
    } else {
        println(4)
    }
}

act println (val int) {
    tmpOut = val
    tmpPrint()
}