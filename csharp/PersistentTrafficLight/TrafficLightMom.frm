```
using PersistentTrafficLight;
#nullable disable
namespace csharp.persistenttrafficlight

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
            _trafficLight = new TrafficLightController(#)
            -> "Traffic Light\nStarted" $Saving ^
 
    $Saving 
        |>|
            _data = _trafficLight.marshal() 
            _trafficLight = nil 
            -> "Saved" $Persisted ^

    $Persisted 
        |tick| -> "Tick"  =>  $Working ^
        |systemError| -> "System Error" =>  $Working ^
        |stop| -> "Stop" $End ^

    $Working => $TrafficLightApi
        |>|    
            _trafficLight = TrafficLightController.loadTrafficLight(`(TrafficLightMomController) this, (TrafficLightController) this._data`)  ^
        |tick|  
            _trafficLight.tick() -> "Done" $Saving ^
        |systemError| 
            systemError() -> "Done" $Saving ^

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
            _trafficLight = TrafficLightController.loadTrafficLight(`(TrafficLightMomController) this, (TrafficLightController) this._data`) 
            trafficLight.stop() 
            trafficLight = nil ^

    -actions-

    enterRed {`Console.WriteLine("Red");`}
    enterGreen {`Console.WriteLine("Green");`}
    enterYellow {`Console.WriteLine("Yellow");`}
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
    var trafficLight:TrafficLightMomController = null
    var _data:TrafficLight = null
    var _trafficLight:TrafficLight = null

##