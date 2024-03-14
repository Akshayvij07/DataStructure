package main

import (
	"fmt"
)

func Sorting(input []int) []int {

	for i := 0; i < len(input); i++ {
		var min int = input[i]
		loc := i

		for j := i + 1; j < len(input); j++ {

			if input[j] < min {

				min = input[j]
				loc = j
				input[loc] = input[i]
				input[i] = min
			}

		}

	}

	return input

}

func display(input []int) {

	for _, val := range input {
		fmt.Print(val, " ")
	}

}

func main() {

	var input []int
	var limit int
	fmt.Print("Enter the limit : ")
	fmt.Scan(&limit)
	fmt.Println("Enter the inputs : ")

	for i := 0; i < limit; i++ {
		var element int
		fmt.Scan(&element)
		input = append(input, element)

	}

	display(Sorting(input))
}
