package main

import (
	"fmt"
)

func Sorting(Uinput []int) []int {

	for i := 0; i < len(Uinput)-1; i++ {

		for j := 0; j < len(Uinput)-1; j++ {

			if Uinput[j] > Uinput[j+1] {

				temp := Uinput[j]
				Uinput[j] = Uinput[j+1]
				Uinput[j+1] = temp
			}

		}
	}

	return Uinput

}

func display(Uinput []int) {

	for _, val := range Uinput {
		fmt.Print(val, " ")
	}

}

func main() {

	var input []int
	var limit int
	fmt.Print("Enter the limit of array: ")
	fmt.Scan(&limit)
	fmt.Println("Enter the inputs : ")

	for i := 0; i < limit; i++ {
		var element int
		fmt.Scan(&element)
		input = append(input, element)

	}

	display(Sorting(input))
}
