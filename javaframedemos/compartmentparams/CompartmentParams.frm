```
import java.util.*;
```
#CompartmentParams $[state_param:int] >[enter_param:int]

  -machine-

  $S0[state_param:int]
    var state_var:int = 100

    |>| [enter_param:int]
      print(state_param
          + " " + state_var 
                    + " " + enter_param
          )
      -> => $S1(state_param+20)
       ^

  $S1[state_param:int]
    var state_var:int = 200

    |>|[enter_param:int]
       print(state_param
          + " " + state_var
          + " " + enter_param  
          )
       ^     

  

  -actions-

  print[s:String] {`console.log(s)`}

##