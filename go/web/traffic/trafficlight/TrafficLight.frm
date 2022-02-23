#TrafficLight

    -interface-

    Start @(|>>|)
    Timer
    SystemError
    SystemRestart

    -machine-

    $Begin
        |>>|
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

    $Working
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