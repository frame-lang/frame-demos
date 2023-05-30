#include <chrono>
#include <thread>
#include "Mom.h"

int main() {
    Mom mom;
    mom.Start();
    std::this_thread::sleep_for(std::chrono::seconds(10));
    mom.Stop();
    return 0;
}
