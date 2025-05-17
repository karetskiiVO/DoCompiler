#include <stdio.h>
#include <stdint.h>

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

tupple printbool (uint8_t b) {
    printf(b ? "true\n" : "false\n");
    tupple res;
    return res;
}