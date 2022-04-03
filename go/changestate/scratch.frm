      i = i - 1
      i = i*5 - 1
      i = (i*5) - 1
      i = ((i*5) - 1)
      i = ((i*5) - i*1) + f(i)

#DebugBug $[i:int] >[name:"hi"]

        $S0[state_param:int]
    var state_var:int = 0

    |>| [enter_param:int]
      $[state_param] = 1
      state_param = 2
      $.state_var = 3
      state_var = 4
      ||[enter_param] = 5
      enter_param = 6
      (enter_param) -> (state_param) $S0 
      ^

--------

      Bugs

      $.missing_var = 1

--------

```
package main

import (

	"github.com/frame-lang/frame-demos/changestate/framelang"
)
```
#DebugBug $[state_param:int] >[enter_param:int]

  -machine-

   $S0[state_param:int]
    var state_var:int = 100

    |>| [enter_param:int]
      $[state_param] = 1
      state_param = 2
      $.state_var = 3
      state_var = 4
      ||[enter_param] = 5
      enter_param = 6
 
      ^

##