package main

import "fmt"

func numSumTarget(arr []int, t int) []int {
	var res []int
	var empty []int

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == t {
				res = append(res, arr[i], arr[j])
			}
		}
	}
	if len(res) == 0 {
		return empty
	} else {
		return res
	}
}

func main() {
	var arr = []int{5, 9, 6, 7, 2, 3}
	var res []int
	var target = 14
	res = numSumTarget(arr, target)
	fmt.Println(res)
}
