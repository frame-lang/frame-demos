
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

reverse_str[str:string]

##