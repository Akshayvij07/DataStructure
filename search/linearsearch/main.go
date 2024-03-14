package main

import "fmt"

func LinearSearch(arr []int) int {
	var value int

	fmt.Println("Enter the value to be searched")
	fmt.Scan(&value)

	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return i
		}
	}
	return -1
}
func InsertValue(arr []int, limit int) {

	fmt.Println("Enter a Sorted Array")
	for i := 0; i < limit; i++ {
		fmt.Scan(&arr[i])
	}
}

func main() {
	var limit int
	fmt.Println("Enter the limit of the array")
	fmt.Scan(&limit)
	arr := make([]int, limit)
	InsertValue(arr, limit)
	fmt.Println("Array", arr)
	index := LinearSearch(arr)
	if index != -1 {
		fmt.Printf("Element found at index %d\n", index+1)
	} else {
		fmt.Println("Element not found in the array")
	}
}
