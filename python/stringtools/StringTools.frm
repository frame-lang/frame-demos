```
from framelang.framelang import FrameEvent
```
#StringTools

-interface-

reverse [str:str] : str
makePalindrome [str:str] : str

-machine-

$Router
    |makePalindrome| [str:str] : str
        -> "make\npalindrome" => $MakePalindrome ^
    |reverse| [str:str] : str
        -> "reverse" => $Reverse ^

$Reverse
    |reverse| [str:str] : str
        @^ = reverse_str(str)
        -> "ready" $Router ^

$MakePalindrome
    |makePalindrome| [str:str] : str
        @^ = str + reverse_str(str)
        -> "ready" $Router ^

-actions-

reverse_str[str:str]

##