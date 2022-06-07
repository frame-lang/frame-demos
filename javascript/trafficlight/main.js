const Mom = require("./trafficlight/mom");


const mom = new Mom()
mom.start()
setTimeout(()=>{
mom.stop()
}, 10000)