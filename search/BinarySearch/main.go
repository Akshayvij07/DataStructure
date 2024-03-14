package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
func binRec(arr []int, target int, start int, end int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2

	if arr[mid] == target {
		return mid
	} else if target < arr[mid] {
		return binRec(arr, target, start, mid-1)
	} else {
		return binRec(arr, target, mid+1, end)
	}

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 11
	index := binRec(arr, target, 0, len(arr)-1)

	if index != -1 {
		fmt.Printf("Element found at index %d\n", index+1)
	} else {
		fmt.Println("Element not found in the array")
	}
}
