```
package trafficlight 

import (
	"encoding/json"

	"github.com/frame-lang/frame-demos/persistenttrafficlight/framelang"
)
```
#[derive(MOM,Marshal)]
#[stateType="TrafficLightFrameState"]
#TrafficLight

    -interface-

    start @(|>>|)
    stop 
    tick
    systemError
    systemRestart

    -machine-

    $Begin
        |>>|
            startWorkingTimer()
            -> $Red ^

    $Red => $Working
        |>|
            enterRed() ^
        |tick|
            -> $Green ^

    $Green => $Working
        |>|
            enterGreen() ^
        |tick|
            -> $Yellow ^

    $Yellow => $Working
        |>|
            enterYellow() ^
        |tick|
            -> $Red ^

    $FlashingRed
        |>|
            enterFlashingRed()
            stopWorkingTimer()
            startFlashingTimer() ^
        |<|
            exitFlashingRed()
            stopFlashingTimer()
            startWorkingTimer() ^
        |tick|
            changeFlashingAnimation() ^
        |systemRestart|
            -> $Red  ^
        |stop| 
            -> $End ^

    $End 
        |>| stopWorkingTimer() ^

    $Working
        |stop| 
            -> $End ^
        |systemError|
            -> $FlashingRed ^

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

    var flashColor:string = ""

##