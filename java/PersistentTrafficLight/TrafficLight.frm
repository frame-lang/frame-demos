```
package java.PersistentTrafficLight;
import java.util.*;
```
#[derive(Marshal)] 
#[managed(TrafficLightMom)]

#TrafficLight

    -interface-

    stop 
    tick
    systemError
    systemRestart

    -machine-

    $Begin
        |>|
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

    enterRed {`try {this._manager.enterRed_do();}catch(Exception e){}`}
    enterGreen {`try {this._manager.enterGreen_do();}catch(Exception e){}`}
    enterYellow {`try {this._manager.enterYellow_do();}catch(Exception e){}`}
    enterFlashingRed
    exitFlashingRed
    startWorkingTimer
    stopWorkingTimer
    startFlashingTimer
    stopFlashingTimer
    changeColor [color: String]
    startFlashing
    stopFlashing
    changeFlashingAnimation
    log[msg:String]

    -domain-
    var flashColor:String = ""
    

##