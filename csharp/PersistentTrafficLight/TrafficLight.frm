```
using PersistentTrafficLight;
#nullable disable
namespace csharp.persistenttrafficlight
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

    enterRed {`this._manager.enterRed_do();`}
    enterGreen {`this._manager.enterGreen_do();`}
    enterYellow {`this._manager.enterYellow_do();`}
    enterFlashingRed
    exitFlashingRed
    startWorkingTimer {`return;`}
    stopWorkingTimer
    startFlashingTimer
    stopFlashingTimer
    changeColor [color: string]
    startFlashing
    stopFlashing
    changeFlashingAnimation
    log[msg:string]

    -domain-
    var flashColor:string = ""
    

##