#TransitionEventForwarding >[cycles:int]

    -machine-

    $Start
        |>| [cycles:int] 
            cycles == 0 ?
               ("stopping") -> => $Stop ^
            :
               ("keep going") -> => $ForwardEventAgain
            :: ^
        |<| [msg:string] print(msg)  ^
            
    $ForwardEventAgain
        |>| [cycles:int] -> => $Decrement ^

    $Decrement
        |>| [cycles:int] 
            print(cycles.toString())
            -> ((cycles - 1)) $Start ^
        

    $Stop 
        |>| [cycles:int]
            print(cycles.toString())
            print("done") ^

    -actions-

    print[msg:string]
##