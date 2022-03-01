#TrafficLight

    -interface-

    start[mom:`*MOM`] @(|>>|)
    stop 
    timer
    systemError
    systemRestart

    -machine-

    $Begin
        |>>|[mom:`*MOM`] 
            startWorkingTimer()
            -> $Red ^

    $Red => $Working
        |>|
            enterRed() ^
        |timer|
            -> $Green ^

    $Green => $Working
        |>|
            enterGreen() ^
        |timer|
            -> $Yellow ^

    $Yellow => $Working
        |>|
            enterYellow() ^
        |timer|
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
        |timer|
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
    var mom:`*MOM` = nil
    var ticker:`*time.Ticker` = nil 
##