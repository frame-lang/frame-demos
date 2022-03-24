```
package main

import (
	"fmt"
    "github.com/frame-lang/frame-demos/systemparams/trafficlight"
)
)
```

#SystemParams[msg:string]

    -machine-

    $Begin
        |>|[msg:string]
            print(msg) ^

    -actions-

    print[msg:string] {`
        fmt.Println(msg)
    `}

##