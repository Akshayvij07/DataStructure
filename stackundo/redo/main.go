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

type Stacks struct {
	stack1 Stack
	stack2 Stack
}

func (s *Stacks) undo() {
	//stack1 := NewStack()
	val := s.stack1.pop()
	//newNode := &Node{data: val}
	s.stack2.Push(val)

}

func (s *Stacks) redo() {
	value := s.stack2.pop()
	//newnode := &Node{data: value}
	s.stack1.Push(value)
}

func (s *Stack) Display() {
	temp := s.top
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.Next
	}
}
func main() {
	s := Stacks{}
	s.stack1.Push(1)
	s.stack1.Push(16)
	s.stack1.Push(3)
	s.stack1.Push(7)
	s.stack1.Push(8)
	s.stack1.Push(5)
	s.stack1.pop()
	//s.stack1.Display()
	s.undo()
	s.redo()
	s.stack1.Display()
}

