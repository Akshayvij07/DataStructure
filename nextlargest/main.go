package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type SinglyList struct {
	Len  int
	Head *ListNode
}

func (s *SinglyList) AddNode(num int) {
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

func (l *ListNode) Display() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func nextLargerNodes(head *ListNode) []int {
	var data []int
	j := &ListNode{}
	k := 0
	for head != nil {
		if head.Next == nil {
			data = append(data, 0)
			break
		}
		j = head.Next

		if head.Val < j.Val {
			data = append(data, j.Val)
			k++
		} else if head.Val >= j.Val {
			for j != nil && head.Val >= j.Val {
				j = j.Next
			}
			if j != nil {
				data = append(data, j.Val)
				k++
			} else {
				data = append(data, 0)
				k++
			}

		} else {
			data = append(data, 0)
			k++
		}
		head = head.Next
	}
	return data
}
func main() {
	//var res []int
	f := &SinglyList{}
	f.AddNode(2)
	f.AddNode(7)
	f.AddNode(4)
	f.AddNode(3)
	f.AddNode(5)
	f.Head.Display()
	res := nextLargerNodes(f.Head)
	fmt.Println(res)

}
