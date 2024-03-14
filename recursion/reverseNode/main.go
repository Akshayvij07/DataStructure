package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
}

func (l *SinglyLinkedList) AddNode(data int) {
	newNode := &Node{Val: data}
	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.Next = newNode
	}
	l.tail = newNode
}
func (l *SinglyLinkedList) DeleteNode(target int) {
	/*if l.head == nil {
		fmt.Println("The list is empty")
	}*/

	if l.head == nil && l.head.Val == target {
		l.head = l.head.Next
		return
	}

	temp := l.head

	for temp != nil && temp.Val == target {
		if temp.Next.Val == target {
			if temp.Next == l.tail {
				l.tail = temp
				l.tail.Next = nil
			}
		}
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
	fmt.Println(n.Val, " ")
}

/*func(l *SinglyLinkedList) AddnodeAfter(after int,data int){
	newNode:=&Node{Val: data}

	temp:=l.head

	for temp.Next!=nil{
		if temp.Val==after{
if temp.Next=nil
		}
	}

}*/
