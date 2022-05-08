```
package main

import (
	"fmt"
    "github.com/frame-lang/frame-demos/systemparams/trafficlight"
)
```

#SystemParams $[stateMsg:string] >[enterMsg:string]

    -machine-

    $Begin [stateMsg:string]
        |>|[enterMsg:string]
            print(stateMsg + " " + enterMsg) ^

    -actions-

    print[msg:string] {`
        fmt.Println(msg)
    `}

##