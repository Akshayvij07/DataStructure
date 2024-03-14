package main

import "fmt"

type Node struct {
	data int
	Next *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
}

func (l *SinglyLinkedList) AddNode(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.Next = newNode
	}
	l.tail = newNode

}
func (l *SinglyLinkedList) DeleteNode(target int) {
	if l.head.data == target && l.head != nil {
		l.head = l.head.Next
		return
	}
	//temp := &Node{Next: l.head}
	temp := l.head.Next
	for temp != nil {
		if temp.Next.data == target {
			if temp.Next == l.tail {
				l.tail = temp
				l.tail.Next = nil
			} else {
				temp.Next = temp.Next.Next
			}
			return
		}
		temp = temp.Next
	}
}

func (l *SinglyLinkedList) Display() {
	temp := &Node{}
	if l.head == nil {
		fmt.Println("List Empty")
	}
	temp = l.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.Next
	}
}
func (l *SinglyLinkedList) deleteAllDuplicate() {
	temp := l.head

	for temp != nil {

		next := temp.Next

		for next != nil && next.data == temp.data {
			next = next.Next
		}
		temp.Next = next
		if next == l.tail && temp.data == next.data {
			l.tail = temp
			l.tail.Next = nil
		}
		temp = next

	}

}

func main() {
	list := SinglyLinkedList{}
	list.AddNode(2)
	list.AddNode(5)
	list.AddNode(7)
	list.AddNode(9)
	list.AddNode(9)
	list.AddNode(9)
	list.AddNode(10)
	list.DeleteNode(5)
	list.deleteAllDuplicate()
	list.Display()

}
