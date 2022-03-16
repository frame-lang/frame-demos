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
    log [msg:string]
    -machine-

    $New 
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
        |tick| -> "Tick" $Working ^
        |<<| -> "Stop" $End ^

    $Working => $TrafficLightApi
        |>| 
            trafficLight = LoadTrafficLight(# data) 
            trafficLight.Tick()
            -> "Done" $Saving ^

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
        |log| [msg:string] log(msg) ^

    $End

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
    log [msg:string]

    -domain-

    var trafficLight:TrafficLight = null
    var data:`[]byte` = null

##