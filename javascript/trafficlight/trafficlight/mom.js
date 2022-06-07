const TrafficLightController = require("./machine");

class Mom {
    constructor() {
        this.mom = new TrafficLightController()
    }

    start() {
        this.mom.start(this.mom)
    }

    stop() {
        this.mom.stop()
    }
}

module.exports = Mom