package main

import (
	"fmt"
)

type Node struct {
	data int
	Next *Node
}
type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) enqueue(value int) {
	newNode := &Node{data: value}
	if q.rear == nil {
		q.front, q.rear = newNode, newNode
		return
	}
	q.rear.Next = newNode
	q.rear = newNode
}
func (q *Queue) dequeue() {
	if q.front == nil {
		fmt.Println("Queue is Empty")
	} else {
		q.front = q.front.Next
	}
}
func (q *Queue) display() {
	//temp := &Node{}

	if q.front == nil {
		fmt.Println("Queue is Empty")
	} else {
		temp := q.front
	}
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.Next

	}
}
func main() {
	qu := Queue{}
	qu.enqueue(5)
	qu.enqueue(6)
	qu.enqueue(8)
	qu.enqueue(10)
	qu.enqueue(52)
	qu.enqueue(58)
	qu.enqueue(9)
	qu.dequeue()
	qu.display()
}
