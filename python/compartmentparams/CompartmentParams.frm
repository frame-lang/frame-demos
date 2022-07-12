
#CompartmentParams $[state_param:int] >[enter_param:int]

  -machine-

  $S0[state_param:int]
    var state_var:int = 100

    |>| [enter_param:int]
      print(str(state_param)
          + " " + str(state_var) 
                    + " " + str(enter_param)
          )
      -> => $S1(state_param+20)
       ^

  $S1[state_param:int]
    var state_var:int = 200

    |>|[enter_param:int]
       print(str(state_param)
          + " " + str(state_var) 
                    + " " + str(enter_param))
       ^     

  

  -actions-

  print[s:string] {`console.log(s)`}

##