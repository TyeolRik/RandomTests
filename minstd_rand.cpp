#include <random>
#include <iostream>

int main(int argc, char **argv) {
    unsigned int seed = (unsigned int)strtol(argv[1], NULL, 10);
    std::minstd_rand minstd_rand;
    minstd_rand.seed((std::uint_fast32_t)seed);

    printf("%lu\n", (std::uint_fast32_t)minstd_rand());
}