#include "machine.h"
#include <iostream>

int main() {
    StringToolsController stc;
    std::string name = stc.reverse("Deepak");

    std::string palindrome = stc.makePalindrome("Deepak");

    std::cout << name << std::endl;
    std::cout << palindrome << std::endl;

    return 0;
}