```
using Countdown;
#nullable disable
namespace csharp.countdown
```
#Countdown $[i:int]

  -machine-

  $S0[i:int]
    var dec:int = 1

    |>| 
      i = i - dec
      print(i.ToString())
      i == 0 ? -> $STOP ^ :: 
      -> (i) $S1 ^

  $S1
    |>| [i:int]
      -> $S0(i) ^

  $Stop
    |>| print("done") ^

  -actions-

  print[s:String] {`Console.WriteLine(s);`}

##