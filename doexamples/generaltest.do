with Struct struct {
    fst     int
    snd     int
    thrd    int
    c       bool
}

act f (newVal int) {
    tmpOut = newVal
}

act main () {
    var struc Struct

    struc.fst  = 10
    struc.snd  = 11
    struc.thrd = 12

    tmpPrint()
    f(2)
    tmpPrint()

    var a int
    a = tmpOut

    println(struc.fst)
    if true {
        println(struc.snd)
    } else {
        println(struc.snd)
    }
}

act println (val int) {
    tmpOut = val
    tmpPrint()
}