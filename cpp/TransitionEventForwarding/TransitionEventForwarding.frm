```
#include <iostream>
#include <unordered_map>
#include <string>
#include <any>
#include "frameEvent.h"
using namespace std;
```
#TransitionEventForwarding >[cycles:int]

    -machine-

    $Start
        |>| [cycles:int] 
            cycles == 0 ?
               (string("stopping")) -> => $Stop ^
            :
               (string("keep going")) -> => $ForwardEventAgain
            :: ^
        |<| [msg:string] print(msg)  ^
            
    $ForwardEventAgain
        |>| [cycles:int] -> => $Decrement ^

    $Decrement
        |>| [cycles:int] 
            print(to_string(cycles))
            -> ((cycles - 1)) $Start ^
        

    $Stop 
        |>| [cycles:int]
            print(to_string(cycles))
            print("done") ^

    -actions-

    print[msg:string] {`cout << msg << endl;`}
##