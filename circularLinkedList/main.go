package main

import "fmt"

type Node struct {
	Data  int
	Next  *Node
	index int
}

type CircularLinkedList struct {
	head   *Node
	tail   *Node

}

// complexity
// O(1)TS
// time always constant because adding a new node to tail its not depend on input
// space always take 1 because alwatys allocate 1 memory for a new node is constant
func (c *CircularLinkedList) Append(value int) {
	newNode := &Node{Data: value}
	if c.head == nil {
		c.head = newNode
	} else {
		c.tail.Next = newNode
	}
	c.tail = newNode
	c.tail.Next = c.head

}

func (c *CircularLinkedList) Display() {
	if c.head == nil {
		fmt.Println("The List is Empty")
		return
	}
	fmt.Println(c.head.Data, " ")

	currNode := c.head.Next

	for currNode != c.head && currNode.Next != currNode {
		fmt.Println(currNode.Data, " ")

		currNode = currNode.Next
	}
	fmt.Println()
}

func (c *CircularLinkedList) AppendValOnLimit() {
	var limit int
	fmt.Println("Enter the limt of the variables")
	fmt.Scan(&limit)

	for i := 1; i <= limit; i++ {
		c.Append(i)
	}
	println()
}

func main() {
	clist := &CircularLinkedList{}

	clist.AppendValOnLimit()

	clist.Display()

}
