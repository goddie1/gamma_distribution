
#include <iostream>
#include <random>

extern "C" {
#include "library.h"
}

double gamma_random(double a,double b) {
    std::cout << "Hello, World!" << std::endl;
    std::random_device rd;
    std::mt19937 gen(rd());
    // A gamma distribution with alpha=1, and beta=2
    // approximates an exponential distribution.
    std::gamma_distribution<> d(a,b);
    auto index = d(gen);
    return index ;
}
