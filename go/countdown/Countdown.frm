```
package main

import (

	"github.com/frame-lang/frame-demos/countdown/framelang"
)
```
#Countdown $[i:int]

  -machine-

  $S0[i:int]
    var dec:int = 1

    |>| 
      i = i - dec
      print(strconv.Itoa(i))
      i == 0 ? -> $Stop ^ :: 
      -> (i) $S1 ^

  $S1
    |>| [i:int]
      -> $S0(i) ^

  $Stop
    |>| print("done") ^

  -actions-

  print[s:string] {`fmt.Println(s)`}

##