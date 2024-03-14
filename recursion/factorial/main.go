package main

import "fmt"

func Factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else {
		return num * Factorial(num-1)
	}

}

func main() {
	num := 5
	fmt.Println(Factorial(num))
}
