package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)

}

func merge(left, right []int) (result []int) {
	final := []int{}
	i := 0
	j := 0
	for i < len(right) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		final = append(final, left[i])
	}
	for ; j < len(right); j++ {
		final = append(final, right[j])
	}
	return final
}
func main() {
	arr := []int{54, 5, 8, 98, 42, 1, 5, 6}
	fmt.Println("Original Array", arr)
	sortedArray := mergeSort(arr)
	fmt.Println("Sorted Array", sortedArray)
}
