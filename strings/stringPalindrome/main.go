package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {
	res := []rune(strings.ToLower(s))

	l := len(res)
	for i := 0; i < l/2; i++ {
		if res[i] != res[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var s string

	s = "MALAYALAM"
	res := []rune(strings.ToLower(s))
	if IsPalindrome(s) == false {
		fmt.Println("Not Palindrome")
	} else {
		fmt.Println("Is Palindrome")
	}
	fmt.Println(IsPalindrome(s))
	fmt.Println(res)

}
