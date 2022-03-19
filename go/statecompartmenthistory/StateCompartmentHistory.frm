```
package main

import (

	"github.com/frame-lang/frame-demos/statecompartmenthistory/framelang"
)
```
#HistoryStateContext

    -interface-

    start @(|>>|)
    switchState
    gotoDeadEnd
    back

    -machine-

    $Start
        |>>| -> $S0 ^
    
    $S0
        var enterMsg:string = "Enter S0"
        
        |>| print(enterMsg) ^
        |switchState| -> "Switch\nState" $S1 ^
        |gotoDeadEnd| $$[+] -> "Goto\nDead End" $DeadEnd ^

    $S1
        |>| print("Enter S1") ^
        |switchState| -> "Switch\nState" $S0 ^
        |gotoDeadEnd| $$[+] -> "Goto\nDead End" $DeadEnd ^

    $DeadEnd
        |>| print("Enter $DeadEnd") ^
        |back| -> "Go Back" $$[-] ^

    -actions-

    print[msg:string] 

##