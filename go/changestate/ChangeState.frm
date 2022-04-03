```
package main

import (

	"github.com/frame-lang/frame-demos/changestate/framelang"
)
```
#ChangeState[i:int]

  -machine-

  $Start
    |>| [i:int]
      ->> $S0(i) ^

  $S0[i:int]
    |>| 
      i = i - 1
      i == 0 ? -> $Stop ^ :: 
      -> $S1(i)
      ^

  $S1[i:int]
    |>|
      itoa(i)
      ->> $S0(i) ^

  $Stop
    |>| print("done") ^


  -actions-

  itoa[i:int] {`strconv.Itoa(i)`}
  print[s:string] {`fmt.Println("done")`}
##