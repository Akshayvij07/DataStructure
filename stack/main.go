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
func (s *Stack) Pop() {
	if s.top == nil {
		fmt.Println("Stack is empty")
		return
	}
	s.top = s.top.Next
}
func (s *Stack) Display() {
	temp := s.top
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.Next
	}
}

func (s *Stack) Peek() int {
	if s.top == nil {
		return -1
	}
	return s.top.data
}

func (s *Stack) middleElement() int {
	curNode := s.top.Next
	midNode := s.top

	for curNode != nil {
		midNode = midNode.Next
		curNode = curNode.Next.Next
	}
	return midNode.data
}
func main() {
	s := Stack{}
	s.Push(1)
	s.Push(16)
	s.Push(3)
	s.Push(7)
	s.Push(8)
	s.Push(5)
	s.Push(9)
	s.Push(10)
	s.Pop()
	s.Display()
	fmt.Println(s.Peek())
	fmt.Println("The middle element is : ", s.middleElement())
}
