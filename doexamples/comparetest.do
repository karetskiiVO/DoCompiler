act main () {
    
    if true {
        println(1);
    } else {
        println(0);
    }

    if 10 == 10 {
        println(2);
    }
 
    if 10 == 20 {
        println(0);
    } else {
        println(3);
    }

    if false {
        println(0);
    } else {
        println(4);
    }

    var a int
    a = 10;
    if a == 10 {
        println(5);
    }

    if a == 20 {
        println(0);
    } else {
        println(6);
    }

    
    if true && true {
        println(7);
    } else {
        println(0);
    }

    if true && false {
        println(0);
    } else {
        println(8);
    }

    if false && false {
        println(0);
    } else {
        println(9);
    }

    if true || true {
        println(10);
    } else {
        println(0);
    }

    if true || false {
        println(11);
    } else {
        println(0);
    }

    if false || false {
        println(0);
    } else {
        println(12);
    }
}

act println (val int) {
    tmpOut = val;
    tmpPrint();
}

act h () int {
    return 2;
}