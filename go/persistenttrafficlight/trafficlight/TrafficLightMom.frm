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
            trafficLight = new NewTrafficLight(#)
            -> "Traffic Light\nStarted" $Saving ^
 
    $Saving 
        |>|
            data = trafficLight.Marshal() 
            trafficLight = nil 
            -> "Saved" $Persisted ^

    $Persisted 
        |tick| -> "Tick"  =>  $Working ^
        |systemError| -> "System Error" =>  $Working ^
        |stop| -> "Stop" $End ^

    $Working => $TrafficLightApi
        |>|    
            trafficLight = LoadTrafficLight(# data)  ^
        |tick|  
            trafficLight.Tick() -> "Done" $Saving ^
        |systemError| 
            trafficLight.SystemError() -> "Done" $Saving ^

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
            trafficLight = LoadTrafficLight(# data) 
            trafficLight.Stop() 
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