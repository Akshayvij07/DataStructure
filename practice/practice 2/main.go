package main

import "fmt"

const arraySize = 5

type HashTable struct {
	array [arraySize]*linkedList
}

type Node struct {
	data string
	Next *Node
}

type linkedList struct {
	head *Node
}

func Init() *HashTable {
	newNode := &HashTable{}

	for i := range newNode.array {
		newNode.array[i] = &linkedList{}
	}
	return newNode
}

func Hash(key string) int {
	sum := 0

	for _, val := range key {
		sum += int(val)
	}
	return sum
}

func (l *linkedList) search(key string) bool {
	temp := l.head

	for temp != nil {
		if temp.data == key {
			return true
		}
	}
	return false

}

func (l *linkedList) insert(key string) {
	if !l.search(key) {
		newNode := &Node{}
		newNode.data = key
		newNode.Next = l.head
		l.head = newNode
	} else {
		fmt.Println("The key already exist")
	}

}

func (h *HashTable) Insert(key string) {
	index := Hash(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) {
	index := Hash(key)
	h.array[index].search(key)
}

func (h *HashTable) Delete(key string) {
	index := Hash(key)
	h.array[index].delete(key)
}


func (l *linkedList) delete(key string) {
	if l.head.data == key {
		l.head = l.head.Next
		return
	}
	temp := l.head
	for temp.Next != nil {
		if temp.Next.data == key {
			temp.Next = temp.Next.Next
		}
		temp=temp.Next
	}
}
