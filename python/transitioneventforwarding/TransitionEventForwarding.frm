```
from framelang.framelang import FrameEvent
```
#TransitionEventForwarding >[cycles:int]

    -machine-

    $Start
        |>| [cycles:int] 
            cycles == 0 ?
               ("stopping") -> => $Stop ^
            :
               ("keep going") -> => $ForwardEventAgain
            :: ^
        |<| [msg:str] print(msg)  ^
            
    $ForwardEventAgain
        |>| [cycles:int] -> => $Decrement ^

    $Decrement
        |>| [cycles:int] 
            print(str(cycles))
            -> ((cycles - 1)) $Start ^
        

    $Stop 
        |>| [cycles:int]
            print(str(cycles))
            print("done") ^

    -actions-

    print[msg:str]
##