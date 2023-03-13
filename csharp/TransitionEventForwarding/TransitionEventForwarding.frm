```
using TransitionEventForwarding;
#nullable disable
namespace csharp.transitioneventforwarding
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
        |<| [msg:string] print(msg)  ^
            
    $ForwardEventAgain
        |>| [cycles:int] -> => $Decrement ^

    $Decrement
        |>| [cycles:int] 
            print(cycles.ToString())
            -> ((cycles - 1)) $Start ^
        

    $Stop 
        |>| [cycles:int]
            print(cycles.ToString())
            print("done") ^

    -actions-

    print[msg:string] {`Console.WriteLine(msg);`}
##