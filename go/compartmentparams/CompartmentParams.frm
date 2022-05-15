```
package main

import (

	"github.com/frame-lang/frame-demos/compartmentparams/framelang"
)
```
#CompartmentParams $[state_param:int] >[enter_param:int]

  -machine-

  $S0[state_param:int]
    var state_var:int = 100

    |>| [enter_param:int]
      print(strconv.Itoa(state_param) 
          + " " + strconv.Itoa(state_var) 
          + " " + strconv.Itoa(enter_param)   
          )
      -> => $S1(state_param+20)
       ^

  $S1[state_param:int]
    var state_var:int = 200

    |>|[enter_param:int]
       print(strconv.Itoa(state_param) 
          + " " + strconv.Itoa(state_var) 
          + " " + strconv.Itoa(enter_param)   
          )
       ^     

  -actions-

  print[s:string] {`fmt.Println(s)`}

##