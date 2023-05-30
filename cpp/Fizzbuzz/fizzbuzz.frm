```
#include <iostream>
#include <unordered_map>
#include <any>
#include <string>
#include "frameEvent.h"
using namespace std;
```
#FizzBuzz

  -interface-

  start @(|>>|)

  -machine-

  $Begin
    |>>|
        -> (1) "start" $Fizz ^

  $Fizz
    |>| [i:int]
      gt_100(i) ?
        -> "i > 100" $End ^
      ::
      mod3_eq0(i) ?
        print("Fizz")
        -> (i true) "i % 3 == 0" $Buzz
      :
        -> (i false) "i % 3 != 0" $Buzz
      :: ^

  $Buzz
    |>| [i:int fizzed:bool]
      mod5_eq0(i) ?
        print("Buzz")
        (string(" ")) -> (plus_1(i)) "i % 5 == 0" $Fizz ^
      ::
      fizzed ?
        (string(" ")) -> (plus_1(i)) "fizzed" $Fizz ^
      ::
      (string("")) -> (i) "! mod3 | mod5" $Digit ^
    |<| [output:string]
      print(output) ^

    $Digit
      |>| [i:int]
        print(to_string(i))
        print(" ")
        ->  (plus_1(i))  "loop" $Fizz ^

   $End

   -actions-

   print    [msg:string] {`cout << msg << " ";`}
   gt_100   [i:int]:bool {`return i > 100;`}
   mod3_eq0 [i:int]:bool {`return i % 3 == 0;`}
   mod5_eq0 [i:int]:bool {`return i % 5 == 0;`}
   plus_1   [i:int]:int {`return i + 1;`}

##