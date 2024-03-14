package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Display() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

type SingleList struct {
	Len  int
	Head *ListNode
}

func (s *SingleList) AddFront(num int) {
	element := &ListNode{
		Val: num,
	}
	if s.Head == nil {
		s.Head = element
	} else {
		element.Next = s.Head
		s.Head = element
	}
	s.Len++
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	i := &ListNode{}
	if l1 == nil && l2 == nil {
		return &ListNode{}
	}
	sum := l1.Val + l2.Val
	res := &ListNode{Val: sum % 10}
	carry := sum / 10
	i = res
	for l1.Next != nil && l2.Next != nil {
		l1, l2 = l1.Next, l2.Next
		sum = l1.Val + l2.Val + carry
		i.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		i = i.Next
	}
	for l1.Next != nil {
		l1 = l1.Next
		sum = l1.Val + carry
		i.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		i = i.Next
	}
	for l2.Next != nil {
		l2 = l2.Next
		sum = l2.Val + carry
		i.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		i = i.Next
	}
	if carry != 0 {
		i.Next = &ListNode{Val: carry}
		i = i.Next
	}
	return res
}

func main() {
	first := &SingleList{}
	first.AddFront(4)
	first.AddFront(7)
	first.AddFront(8)

	second := &SingleList{}
	second.AddFront(3)
	second.AddFront(3)
	second.AddFront(9)

	result := addTwoNumbers(first.Head, second.Head)
	result.Display()
}
