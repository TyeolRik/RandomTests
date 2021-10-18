#include <iostream>
#include <random>

int main(int argc, char **argv) {
    unsigned int seed = (unsigned int)strtol(argv[1], NULL, 10);
    std::mt19937 mt((std::uint_fast32_t)seed);
    printf("%lu\n", (std::uint_fast32_t)mt());
}