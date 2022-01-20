#include <stdio.h>
#include <stdlib.h>

#define BUFFERSIZE 10000

// C original rand
// [0, 2147483647(=2^31-1)]
// https://www.cplusplus.com/reference/cstdlib/rand/
int main(int argc, char **argv) {
    unsigned int seed = (unsigned int)strtol(argv[1], NULL, 10);
    srand(seed);
    size_t size = sizeof(unsigned int);
    unsigned int* outBuffer = malloc(size * BUFFERSIZE);

    while(1) {
        // Writing
        for(int i = 0; i < BUFFERSIZE; i++) {
            outBuffer[i] = (unsigned int)rand();
        }
        fwrite(outBuffer, 1, BUFFERSIZE, stdout);
        fflush(stdout);
    }

}