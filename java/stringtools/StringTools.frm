
#StringTools

-interface-

reverse [str:String] : String
makePalindrome [str:String] : String

-machine-

$Router
    |makePalindrome| [str:String] : String
        -> "make\npalindrome" => $MakePalindrome ^
    |reverse| [str:String] : String
        -> "reverse" => $Reverse ^

$Reverse
    |reverse| [str:String] : String
        @^ = reverse_str(str)
        -> "ready" $Router ^

$MakePalindrome
    |makePalindrome| [str:String] : String
        @^ = str + reverse_str(str)
        -> "ready" $Router ^

-actions-

reverse_str[str:String]

##