```
package trafficlight

import (
	"github.com/frame-lang/frame-demos/persistenttrafficlight/framelang"
	"github.com/gorilla/websocket"
)
```

#TrafficLightMom [clntId: string conn: `*websocket.Conn`]

    -interface-

    start @(|>>|)
    stop
    initTrafficLight
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
    startFlashing
    stopFlashing
    changeFlashingAnimation
    systemError
    systemRestart
    log [msg:string]
    destroyTrafficLight
    getConnection:`*websocket.Conn`

    -machine-

    $New => $TrafficLightApi
        |>>|

            trafficLight = NewTrafficLight(#)
            trafficLight.Start()
            -> "Traffic Light\nStarted" $Saving ^
        |getConnection|:`*websocket.Conn` @^ = connection ^
 
    $Saving 
        |>|
            data = trafficLight.Marshal() 
            trafficLight = nil 
            -> "Saved" $Persisted ^

    $Persisted 
        |tick| -> "Tick"  =>  $Working ^
        |systemError| -> "System Error" =>  $Working ^
        |systemRestart| -> "System Error" =>  $Working ^
        |getConnection|:`*websocket.Conn` @^ = connection ^
        |stop| -> "Stop" $End ^

    $Working => $TrafficLightApi
        |>| trafficLight = LoadTrafficLight(# data)  ^
        |tick| trafficLight.Tick() -> "Done" $Saving ^
        |systemError| trafficLight.SystemError() -> "Done" $Saving ^
        |systemRestart| trafficLight.SystemRestart() -> "Done" $Saving ^

    $TrafficLightApi
        |initTrafficLight| initTrafficLight() ^
        |enterRed| enterRed() ^
        |enterGreen| enterGreen()  ^
        |enterYellow| enterYellow() ^
        |enterFlashingRed| enterFlashingRed() ^
        |exitFlashingRed| exitFlashingRed() ^
        |startWorkingTimer| startWorkingTimer() ^
        |stopWorkingTimer| stopWorkingTimer() ^
        |startFlashingTimer| startFlashingTimer() ^
        |stopFlashingTimer| stopFlashingTimer() ^
        |startFlashing| startFlashing() ^
        |stopFlashing| stopFlashing() ^
        |changeFlashingAnimation| changeFlashingAnimation() ^
        |log| [msg:string] log(msg) ^
        |destroyTrafficLight| destroyTrafficLight() ^

    $End => $TrafficLightApi
        |>|
            trafficLight = LoadTrafficLight(# data) 
            trafficLight.Stop()
            -> $New ^

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
    startFlashing
    stopFlashing
    initTrafficLight
    changeFlashingAnimation
    log [msg:string]
    destroyTrafficLight

    -domain-

    var trafficLight:TrafficLight = null
    var data:`[]byte` = null
    var connection:`*websocket.Conn` = conn
    var clientId:string = clntId
    var stopper:`chan<- bool` = null
    var clntId:string = null
    var conn:string = null
##