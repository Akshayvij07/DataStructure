package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{54, 5, 8, 98, 42, 1, 5, 6}
	fmt.Println("Original Array", arr)
	sortedArray := quicksort(arr)
	fmt.Println("Sorted Array", sortedArray)
}

// Generates a slice of size, size filled with random numbers

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
