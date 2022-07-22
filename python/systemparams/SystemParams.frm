```
from framelang.framelang import FrameEvent
```
#SystemParams $[stateMsg:str] >[enterMsg:str]

    -machine-

    $Begin [stateMsg:str]
        |>|[enterMsg:str]
            print(stateMsg + " " + enterMsg) ^

    -actions-

    print[msg:str] {`
        print(msg)
    `}

##