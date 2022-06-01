const TrafficLightMomController = require("./trafficlight/mom")

let a = new TrafficLightMomController

const tick = setInterval(()=>{
    console.log("Tick")
    a.tick()
},1000)

setTimeout(()=>{
    clearInterval(tick)
    a.stop()
    console.log("Finished")
}, 5000)