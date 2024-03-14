package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func arrayToList(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}
	head := &Node{Data: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &Node{Data: nums[i]}
		current = current.Next
	}
	return head
}

func Display(N *Node) {
	for N != nil {
		fmt.Println(N.Data)
		N = N.Next
	}
}

func main() {
	var res = []int{5, 9, 10, 25, 20, 14}

	Display(arrayToList(res))
}
