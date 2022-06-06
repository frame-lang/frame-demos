```
package trafficlight
import (
    "github.com/frame-lang/frame-demos/persistenttrafficlight/framelang"
)
```
#TrafficLightMom

    -interface-
    
    stop
    tick    
    enterRed
    enterGreen
    enterYellow
    enterFlashingRed
    exitFlashingRed
    startWorkingTimer
    stopWorkingTimer
    startFlashingTimer
    stopFlashingTimer
    changeColor [color:string]
    startFlashing
    stopFlashing
    changeFlashingAnimation
    systemError
    systemRestart
    log [msg:string]

    -machine-

    $New => $TrafficLightApi
        |>| 
            trafficLight = new TrafficLightController(#)
            -> "Traffic Light\nStarted" $Saving ^
 
    $Saving 
        |>|
            data = trafficLight.marshal() 
            trafficLight = nil 
            -> "Saved" $Persisted ^

    $Persisted 
        |tick| -> "Tick"  =>  $Working ^
        |systemError| -> "System Error" =>  $Working ^
        |stop| -> "Stop" $End ^

    $Working => $TrafficLightApi
        |>|    
            trafficLight = TrafficLightController.loadTrafficLight(# data)  ^
        |tick|  
            trafficLight.tick() -> "Done" $Saving ^
        |systemError| 
            trafficLight.systemError_do() -> "Done" $Saving ^

    $TrafficLightApi
        |enterRed| enterRed() ^
        |enterGreen| enterGreen()  ^
        |enterYellow| enterYellow() ^
        |enterFlashingRed| enterFlashingRed() ^
        |exitFlashingRed| exitFlashingRed() ^
        |startWorkingTimer| startWorkingTimer() ^
        |stopWorkingTimer| stopWorkingTimer() ^
        |startFlashingTimer| startFlashingTimer() ^
        |stopFlashingTimer| stopFlashingTimer() ^
        |changeColor| [color:string] changeColor(color) ^
        |startFlashing| startFlashing() ^
        |stopFlashing| stopFlashing() ^
        |changeFlashingAnimation| changeFlashingAnimation() ^
        |systemError| systemError() ^
        |systemRestart| systemRestart() ^
        |log| [msg:string] log(msg) ^

    $End => $TrafficLightApi
        |>|
            trafficLight = TrafficLightController.loadTrafficLight(# data) 
            trafficLight.stop() 
            trafficLight = nil ^

    -actions-

    enterRed
    enterGreen
    enterYellow
    enterFlashingRed
    exitFlashingRed
    startWorkingTimer
    stopWorkingTimer
    startFlashingTimer
    stopFlashingTimer   
    changeColor [color:string]
    startFlashing
    stopFlashing
    changeFlashingAnimation
    systemError
    systemRestart
    log [msg:string]
    -domain-
    var trafficLight:TrafficLight = null
    var data:`[]byte` = null

##