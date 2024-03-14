package main

import "fmt"


const arraySize=5


type HashTable struct{
	array [arraySize]*LinkedList
}

type Node struct{
	val string
	Next *Node
}

type LinkedList struct{
	head *Node
}

func hash(key string)int{
	value:=0
	for _,val :=range key{
		value+=int(val)
	}
	return value%arraySize
}

func (l *LinkedList) search(key string)bool{
	curNode:=l.head

	for curNode!=nil{
		if curNode.val == key{
			return true
		}
		curNode=curNode.Next

	}
	return false
}

func (l *LinkedList) insert(key string){

	if !l.search(key){
		newNode:=&Node{val:key}
		newNode.Next=l.head
		l.head=newNode
	}else{
		fmt.Println("The key already exist")
	}
}

func (h *HashTable) Insert(key string){
	index:=hash(key)
	h.array[index].insert(key)
}

func (l *LinkedList) delete(key string){
	curNode:=l.head

	for curNode!=nil{
		if curNode.Next.val==key{
			curNode.Next=curNode.Next.Next
		}
		curNode=curNode.Next
	}
}

func (h *HashTable) Delete(key string){
	index:=hash(key)
	h.array[index].delete(key)
}
func (h *HashTable) Search(key string){
	index:=hash(key)
	h.array[index].search(key)
}


