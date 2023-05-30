```
#include <iostream>
#include <unordered_map>
#include <any>
#include <string>
#include "frameEvent.h"
using namespace std;
```
#CompartmentParams $[state_param:int] >[enter_param:int]

  -machine-

  $S0[state_param:int]
    var state_var:int = 100

    |>| [enter_param:int]
      print(to_string(state_param)
          + " " + to_string(state_var) 
                    + " " + to_string(enter_param)
          )
      -> => $S1(state_param+20)
       ^

  $S1[state_param:int]
    var state_var:int = 200

    |>|[enter_param:int]
       print(to_string(state_param)
          + " " + to_string(state_var)
          + " " + to_string(enter_param)  
          )
       ^     

  

  -actions-

  print[s:string] {` std::cout << s << std::endl;`}

##