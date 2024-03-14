package main

import (
	"fmt"
)

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (l *DoublyLinkedList) addNode(value int) {

	newNode := &Node{Data: value}

	if l.head == nil {
		l.head = newNode
	} else {

		l.tail.Next = newNode
		newNode.Prev = l.tail
	}

	l.tail = newNode
}

func (l *DoublyLinkedList) displayForward() {
	temp := &Node{}

	temp = l.head

	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.Next
	}

}

func (l *DoublyLinkedList) displayBackward() {
	temp := &Node{}

	temp = l.tail

	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.Prev
	}
}

func (l *DoublyLinkedList) insertBefore(Before int, value int) {
	//temp := &Node{}
	newNode := &Node{}
	newNode.Data = value

	temp := l.head
	if l.head.Data == Before {
		l.head.Prev = newNode
		newNode.Next = l.head
		l.head = newNode
	} else {
		for temp.Data != Before {

			if temp.Next.Data == Before {
				newNode.Next = temp.Next
				newNode.Prev = temp
				temp.Next = newNode
				break
			}
			temp = temp.Next
		}
	}

}
func (l *DoublyLinkedList) insertNodeAfter(after int, value int) {
	temp := &Node{}
	newNode := &Node{}
	newNode.Data = value

	temp = l.head
	if l.tail.Data == after {
		l.tail.Next = newNode
		newNode.Prev = l.tail
	} else {
		for temp != nil {

			if temp.Data == after {
				temp.Next = newNode
				newNode.Prev = temp
				break
			}
			temp = temp.Next
		}
	}
}

func (l *DoublyLinkedList) delete(data int) {

	currentNode := &Node{}

	currentNode = l.head

	for currentNode != nil {

		if currentNode.Data == data {

			if currentNode == l.head {
				l.head = l.head.Next
				l.head.Prev = nil
			} else if currentNode == l.tail {
				l.tail = l.tail.Prev
				l.tail.Next = nil
			} else {
				currentNode.Prev.Next = currentNode.Next
				currentNode.Next.Prev = currentNode.Prev

				break
			}
		}
		currentNode = currentNode.Next
	}

}

func main() {

	list := DoublyLinkedList{}
	list.addNode(4)
	list.addNode(3)
	list.addNode(9)
	list.addNode(0)
	list.displayBackward()
	fmt.Println()
	list.insertNodeAfter(9, 89)
	list.insertBefore(9, 3)
	list.delete(4)
	list.delete(0)
	list.displayForward()
}
