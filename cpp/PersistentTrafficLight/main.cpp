#include <iostream>
#include <chrono>
#include <thread>
#include "frameEvent.h"
#include "TrafficLightMom.h"
class TrafficLightController;
class TrafficLight;
class TrafficLightMom;
class TrafficLightMomController;

int main() {
    TrafficLightMomController* mom;
    int count = 0;
    while (count <= 3) {
        std::cout << "Tick" << std::endl;
        std::this_thread::sleep_for(std::chrono::seconds(1));
        mom->tick();
        count++;
    }
    std::cout << "Finished" << std::endl;
    delete mom;

    return 0;
}