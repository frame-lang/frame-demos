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
            data = trafficLight.Save()
            -> $Waiting ^
       |<| 
            data = trafficLight.Save() 
            trafficLight = nil ^
 
    $Waiting 
        |tick| -> $Working ^
        |<<| -> $End ^

    $Working
        |>| 
            trafficLight = New() 
            trafficLight.Tick()
            -> $Waiting ^
       |<| 
            data = trafficLight.Save() 
            trafficLight = nil ^

    $End

    -domain-

    var trafficLight:TrafficLight = null
    var data:`[]byte` = null
##