```
#include <iostream>
#include <unordered_map>
#include <any>
#include <string>
#include "frameEvent.h"
using namespace std;
```     
#Countdown $[i:int]

  -machine-

  $S0[i:int]
    var dec:int = 1

    |>| 
      i = i - dec
      print(to_string(i))
      i == 0 ? -> $STOP ^ :: 
      -> (i) $S1 ^

  $S1
    |>| [i:int]
      -> $S0(i) ^

  $Stop
    |>| print("done") ^

  -actions-

  print[s:string] {`std::cout << s << std::endl;`}

##