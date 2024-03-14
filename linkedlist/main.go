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
	l.tail.Next = nil

}
func (l *SinglyLinkedList) DeleteNode(target int) {
	if l.head.data == target && l.head != nil {
		l.head = l.head.Next
		return
	}
	temp := &Node{Next: l.head}
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
func (l *SinglyLinkedList) InsertNodeAfter(after int, value int) {
	newNode := &Node{data: value}
	temp := &Node{Next: l.head}
	if after == temp.data {
		newNode.Next = temp.Next
		temp.Next = newNode
		return
	}
	for temp.Next != nil {
		if temp.Next.data == after {
			if temp.Next == l.tail {
				l.tail.Next = newNode
				l.tail = newNode
			} else {
				newNode.Next = temp.Next.Next
				temp.Next.Next = newNode
			}
		}
		temp = temp.Next
	}

}
func (l *SinglyLinkedList) InsertNodeBefore(before int, value int) {
	newNode := &Node{data: value}
	temp := l.head

	if before == l.head.data {

		newNode.Next = temp
		l.head = newNode
		return
	}

	for temp.Next != nil {

		if temp.Next.data == before {
			newNode.Next = temp.Next
			temp.Next = newNode
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

func (l *SinglyLinkedList) Reverse() {
	ReverseHelpr(l.head)
}
func ReverseHelpr(n *Node) {
	if n == nil {
		return
	}

	ReverseHelpr(n.Next)
	fmt.Println(n.data, " ")
}

func main() {
	list := SinglyLinkedList{}
	list.Display()
	list.AddNode(1)
	list.AddNode(5)
	list.AddNode(7)
	list.AddNode(8)
	list.AddNode(2)
	list.DeleteNode(8)
	list.InsertNodeAfter(2, 20)
	list.InsertNodeBefore(1, 9)
	list.InsertNodeBefore(5, 19)
	list.Display()
	fmt.Println()
	list.Reverse()
}
