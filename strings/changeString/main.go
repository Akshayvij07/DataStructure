package main

import "fmt"

func stringChange(word string, n int) string {
	newVal := n % 26
	chars := []rune(word)
	for i, Val := range chars {
		letterpos := int(Val) + newVal
		if letterpos > 122 {
			chars[i] = rune(96 + letterpos%122)
		} else {
			chars[i] = rune(letterpos)
		}
	}
	return string(chars)
}

func main() {
	str := "dfsd"
	fmt.Println(stringChange(str, 54))
}
