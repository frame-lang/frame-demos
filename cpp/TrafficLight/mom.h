#include "machine.h"

class Mom {
public:
    Mom() {
        TrafficLightController* mom = new TrafficLightController();
    }

    void Start() {
        mom->start(mom);
    }

    void Stop() {
        mom->stop();
    }

private:
    TrafficLightController* mom;
};
