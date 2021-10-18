#include <stdio.h>
#include <stdlib.h>

// C original rand
// https://www.cplusplus.com/reference/cstdlib/rand/
int main(int argc, char **argv) {
    unsigned int seed = (unsigned int)strtol(argv[1], NULL, 10);
    srand(seed);
    printf("%d\n", rand());
}