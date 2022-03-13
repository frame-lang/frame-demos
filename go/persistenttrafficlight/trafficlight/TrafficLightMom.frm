```
package trafficlight

import (
	"encoding/json"

	"github.com/frame-lang/frame-demos/persistenttrafficlight/framelang"
)
```

#MOM

    -interface-

    start @(|>>|)
    stop @(|<<|)
    tick 
    
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

    $Working
        |>| 
            trafficLight = LoadTrafficLight(# data) 
            trafficLight.Tick()
            -> "Done" $Saving ^

    $End

    -domain-

    var trafficLight:TrafficLight = null
    var data:`[]byte` = null

##