from machine import StringTools

class StringToolsController(StringTools):

    def __init__(self):
        super().__init__()
    
    def reverse_str_do(self, str):
        
        temp_str = ""

        for i in str:

            temp_str = i + temp_str

        return temp_str
    

sm = StringToolsController()
name = sm.reverse_str_do("MARK")
print(name)
palindrome = sm.makePalindrome(name)
print(palindrome)
