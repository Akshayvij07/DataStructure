package main

import "fmt"

const arraySize = 5

type HashTable struct {
	array [arraySize]*LinkedList
}

type Node struct {
	Key  string
	Next *Node
}

type LinkedList struct {
	head *Node
}

func Init() *HashTable {
	newNode := &HashTable{}
	for i := range newNode.array {
		newNode.array[i] = &LinkedList{}
	}
	return newNode
}

func (h *HashTable) Insert(key string) {
	/*var key string
	fmt.Println("Enter  the String   :")
	fmt.Scanf("%s", &key)*/

	index := hash(key)
	h.array[index].insert(key)

}

func (l *LinkedList) insert(key string) {

	if !l.search(key) {
		newNode := &Node{}
		newNode.Key = key
		newNode.Next = l.head
		l.head = newNode

	} else {
		fmt.Println("The Key Already Exist")
	}
}

func (h *HashTable) Search(key string) bool {

	index := hash(key)
	return h.array[index].search(key)

}

func (l *LinkedList) search(key string) bool {
	curNode := l.head

	for curNode != nil {
		if curNode.Key == key {
			return true
		}
		curNode = curNode.Next

	}
	return false
}

func (h *HashTable) Delete() {
	var key string
	fmt.Println("Enter  the String  to delete  :")
	fmt.Scanf("%s", &key)
	index := hash(key)
	h.array[index].delete(key)

}

func (l *LinkedList) delete(key string) {

	/*if l.head.Key == key {
		l.head = l.head.Next
		return
	}*/

	prevNode := l.head
	for prevNode != nil {
		if prevNode.Next.Key == key {
			prevNode.Next = prevNode.Next.Next
		}
		prevNode = prevNode.Next
	}
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % arraySize
}

/*func insertString(a []string, index int, value string) []string {
    if len(a) == index { // nil or empty slice or after last element
        return append(a, value)
    }
    a = append(a[:index+1], a[index:]...) // index < len(a)
    a[index] = value
    return a
}*/

func GenerateSting(arr []string) {
	for i := 0; i < arraySize; i++ {
		var element string
		fmt.Scanf("%s", &element)
		arr = append(arr, element)

	}
}

func main() {
	//var v string
	Hashtable := Init()
	/*list := []string{
		"RANDY",
		"SANDY",
		"KENNY",
		"KYLE",
		"ERIC",
		"STAN",
		"MATE",
	}
	/*fmt.Println("Enter the key values to the list")
	for i := range list {
		fmt.Scanf("%s", &list[i])
	}*/
	var list []string
	fmt.Println("Enter the inputs : ")

	/*for i := 0; i < arraySize; i++ {
		var element string
		fmt.Scanf("%s", &element)
		list = append(list, element)

	}*/
	GenerateSting(list)
	for _, v := range list {

		(Hashtable).Insert(v)
	}

	(Hashtable).Delete()
	var key string
	fmt.Println("Enter  the String  to search to search :")
	fmt.Scanf("%s", &key)
	fmt.Println("The Searched key is ", key, (Hashtable).Search(key))
}
