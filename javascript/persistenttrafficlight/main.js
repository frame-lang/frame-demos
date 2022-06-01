const TrafficLightMomController = require("./trafficlight/mom")

let mom = new TrafficLightMomController()

const tick = setInterval(()=>{
    console.log("Tick")
    mom.tick()
},1000)

setTimeout(()=>{
    clearInterval(tick)
    mom.stop()
    console.log("Finished")
}, 5000)