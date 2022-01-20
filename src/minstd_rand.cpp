#include <random>
#include <stdio.h>

#define BUFFERSIZE 10000

int main(int argc, char **argv) {
    unsigned int seed = (unsigned int)strtol(argv[1], NULL, 10);
    std::minstd_rand minstd_rand;
    minstd_rand.seed((std::uint_fast32_t)seed);
    size_t size = sizeof(unsigned int);
    unsigned int* outBuffer = (unsigned int*)malloc(size * BUFFERSIZE);

    while(1) {
        // Writing
        for(int i = 0; i < BUFFERSIZE; i++) {
            outBuffer[i] = (std::uint_fast32_t)minstd_rand();
        }
        fwrite(outBuffer, 1, BUFFERSIZE, stdout);
        fflush(stdout);
    }
}