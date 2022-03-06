#MOM

    -interface-

    start @(|>>|)
    stop @(|<<|)
    tick 
    
    -machine-

    $New 
        |>>| 
            trafficLight = New()
            trafficLight.Start()
            -> "Traffic Light\nStarted" $Saving ^
 
    $Saving 
        |>|
            data = trafficLight.Save() 
            trafficLight = nil 
            -> "Saved" $Persisted ^

    $Persisted 
        |tick| -> "Tick" $Working ^
        |<<| -> "Stop" $End ^

    $Working
        |>| 
            trafficLight = New() 
            trafficLight.Tick()
            -> "Done" $Saving ^

    $End

    -domain-

    var trafficLight:TrafficLight = null
    var data:`[]byte` = null
##