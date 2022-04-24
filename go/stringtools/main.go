package main

import "fmt"

func main() {
	tools := NewStringTools()
	name := tools.Reverse("mark")
	fmt.Println(name)
	palindrome := tools.MakePalindrome(name)
	fmt.Println(palindrome)
}
