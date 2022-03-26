
```
package main

import (
	"github.com/frame-lang/frame-demos/transitioneventforwarding/framelang"
)
```

#TransitionEventForwarding[cycles:int]

    -machine-

    $One
        |>| [cycles:int] 
            cycles == 0 ?
                -> => $Stop ^
            :
                -> => $Two
            :: ^
            
    $Two
        |>| [cycles:int] -> => $Three ^

    $Three
        |>| [cycles:int] 
            print(strconv.Itoa(cycles))
            -> (cycles - 1) $One ^
        

    $Stop 
        |>| [cycles:int]
            print(strconv.Itoa(cycles))
            print("done") ^

    -actions-

    print[msg:string]
##