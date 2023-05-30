```
#include <iostream>
#include <unordered_map>
#include <string>
#include <stack>
#include <any>
#include "frameEvent.h"
using namespace std;
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
        var enterMsg:string = string("Enter S0")
        
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

    print[msg:string] {`cout << msg << endl;`}

##