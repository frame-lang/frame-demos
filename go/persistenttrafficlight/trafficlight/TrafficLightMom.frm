```
package trafficlight
import (
    "github.com/frame-lang/frame-demos/persistenttrafficlight/framelang"
)
```
#TrafficLightMom
    -interface-
    start @(|>>|)
    stop @(|<<|)
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
        |>>| 
            trafficLight = NewTrafficLight(#)
            trafficLight.Start()
            -> "Traffic Light\nStarted" $Saving ^
 
    $Saving 
        |>|
            data = trafficLight.Marshal() 
            trafficLight = nil 
            -> "Saved" $Persisted ^

    $Persisted 
        |tick| -> "Tick"  =>  $Working ^
        |systemError| -> "System Error" =>  $Working ^
        |<<| -> "Stop" $End ^

    $Working => $TrafficLightApi
        |>| [msg:string]    trafficLight = LoadTrafficLight(# data)  ^
        |tick|  trafficLight.Tick() -> "Done" $Saving ^
        |systemError| trafficLight.SystemError() -> "Done" $Saving ^

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
            trafficLight.Stop() ^
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