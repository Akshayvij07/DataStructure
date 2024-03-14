package main

import "fmt"

func arrangeToLast(arr []int, t int) []int {

	for j := len(arr) - 1; j > 0; j-- {
		if arr[j] == t {
			continue
		} else {
			for i := 0; i < j; i++ {
				if arr[i] == t {
					temp := arr[i]
					arr[i] = arr[j]
					arr[j] = temp
				}
			}
		}
	}
	return arr
}

func main() {
	var arr = []int{6, 5, 9, 6, 6, 7, 6, 2, 3}
	var res []int
	var target = 6
	res = arrangeToLast(arr, target)
	fmt.Println(res)
}
