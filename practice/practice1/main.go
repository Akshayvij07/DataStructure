package main

import "fmt"

type Node struct {
	data int
	Next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(value int) {
	newNode := &Node{data: value}

	if s.top == nil {
		s.top = newNode
	} else {
		newNode.Next = s.top
		s.top = newNode
	}
}

func (s *Stack) pop() int {
	if s.top == nil {
		return -1
	}
	data := s.top.data
	s.top = s.top.Next
	return data

}

type Queue struct {
	inbox  Stack
	outbox Stack
}

func (q *Queue) Enqueue(data int) {
	q.inbox.Push(data)
}
func (q *Queue) dequeue() int {
	if q.outbox.top == nil {
		for q.inbox.top != nil {
			q.outbox.Push(q.inbox.pop())
		}
	}
	return q.outbox.pop()

}

func (q *Queue) Display() {
	if q.outbox.top == nil {
		fmt.Println("empty queue")
	}
	temp := q.outbox.top
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.Next
	}
}

func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	q.Enqueue(4)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
}
