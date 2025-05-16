#include <stdio.h>

extern int tmpOut;

typedef struct {} tupple;

tupple tmpPrint () {
    printf("%d\n", tmpOut);
    tupple res;
    return res;
}

int one () {
    return 1;
}