package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func reverseList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	newHead := reverseList(head.next)
	head.next.next = head
	head.next = nil
	return newHead
}

func printList(head *Node) {
	for head != nil {
		fmt.Printf("%d -> ", head.data)
		head = head.next
	}
	fmt.Printf("nil\n")
}

func main() {
	// create the linked list: 1 -> 2 -> 3 -> 4 -> 5
	head := &Node{data: 1, next: &Node{data: 2, next: &Node{data: 3, next: &Node{data: 4, next: &Node{data: 5, next: nil}}}}}
	fmt.Printf("Original list: ")
	printList(head)

	// reverse the linked list using recursion
	head = reverseList(head)
	fmt.Printf("Reversed list: ")
	printList(head)
}
