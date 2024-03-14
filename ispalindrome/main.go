package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type SinglyList struct {
	Len  int
	Head *Node
}

func (s *SinglyList) AddNode(num int) {
	element := &Node{
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

func isPalindrome(head *Node) bool {
	if head == nil {
		return true
	}

	var data []int
	for currentNode := head; currentNode != nil; {
		data = append(data, currentNode.Val)
		currentNode = currentNode.Next
	}

	for i, j := 0, len(data)-1; i < len(data); {
		if data[i] != data[j] {
			return false
		}
		i++
		j--
	}
	return true
}
func main() {
	f := &SinglyList{}
	f.AddNode(1)
	f.AddNode(2)
	f.AddNode(1)

	fmt.Println(isPalindrome(f.Head))
}
