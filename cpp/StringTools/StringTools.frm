```
#include <iostream>
#include <unordered_map>
#include <string>
#include <any>
#include "frameEvent.h"
using namespace std;
```
#StringTools

-interface-

reverse [str:string] : string
makePalindrome [str:string] : string

-machine-

$Router
    |makePalindrome| [str:string] : string
        -> "make\npalindrome" => $MakePalindrome ^
    |reverse| [str:string] : string
        -> "reverse" => $Reverse ^

$Reverse
    |reverse| [str:string] : string
        @^ = reverse_str(str)
        -> "ready" $Router ^

$MakePalindrome
    |makePalindrome| [str:string] : string
        @^ = str + reverse_str(str)
        -> "ready" $Router ^

-actions-

reverse_str[str:string] : string{` 
    string result = "";
    for(char c : str) {
        result = c + result;
    }
    return result;
`}


##