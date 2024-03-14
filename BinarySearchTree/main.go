package main

import (
	"fmt"
)

type Node struct {
	Data   int
	rChild *Node
	lChild *Node
}

type BinarySearchtree struct {
	root *Node
}

func (t *BinarySearchtree) addNode() {
	var data int
	newNode := &Node{}
	fmt.Println("Enter the data :")
	fmt.Scan(&data)
	newNode.Data = data

	if t.root == nil {
		t.root = newNode
		return
	}

	temp := t.root

	for temp != nil {
		if newNode.Data < temp.Data {
			if temp.lChild == nil {
				temp.lChild = newNode
				return
			}
			temp = temp.lChild

		} else if newNode.Data > temp.Data {
			if temp.rChild == nil {
				temp.rChild = newNode
				return
			}
			temp = temp.rChild

		}
	}
}

func (t *BinarySearchtree) contain() {

	var data int
	temp := t.root

	fmt.Print("\nEnter the data to check whether the tree contains or not : ")
	fmt.Scan(&data)

	if temp == nil {

		fmt.Println("Tree is empty!...")
		return
	}

	for temp != nil {

		if data < temp.Data {

			temp = temp.lChild

		} else if data > temp.Data {

			temp = temp.rChild

		} else {

			fmt.Println("Match found!")
			return

		}
	}

	fmt.Println("No match found!")

}

func (t *BinarySearchtree) remove() {
	var data int
	fmt.Println("Enter the Node to remove: ")
	fmt.Scan(&data)
	curNode := t.root
	if curNode == nil {
		fmt.Println("The tree is empty")
	}

	for curNode != nil {
		if data < curNode.Data {
			curNode = curNode.lChild
		} else if data > curNode.Data {
			curNode = curNode.rChild
		} else {
			if curNode.rChild != nil {
				curNode.Data = getMin(curNode)
			} else if curNode.lChild != nil {
				curNode.Data = getMax(curNode)
			} else {
				curNode.Data = getMin(curNode)
				fmt.Println("The node removed succesfully")
				return
			}
		}
	}
}

func getMax(curNode *Node) int {
	prevNode := curNode
	temp := curNode.lChild
	for temp.rChild != nil {
		prevNode = temp
		temp = temp.rChild
	}
	if prevNode == curNode {
		prevNode.lChild = temp.lChild
	} else {
		prevNode.rChild = temp.lChild
	}
	return temp.Data
}

func getMin(curNode *Node) int {
	prevNode := curNode
	temp := curNode.rChild
	for temp.lChild != nil {
		prevNode = temp
		temp = temp.lChild
	}
	if prevNode == curNode {
		prevNode.rChild = temp.rChild
	} else {
		prevNode.lChild = temp.rChild
	}
	return temp.Data
}

func (t *BinarySearchtree) inOrder() {

	fmt.Print("Inorder : ")
	t.inOrderHelper(t.root)

}

func (t *BinarySearchtree) inOrderHelper(temp *Node) {

	if temp != nil {

		t.inOrderHelper(temp.lChild)
		fmt.Print(temp.Data, " ")
		t.inOrderHelper(temp.rChild)
	}

}

func (t *BinarySearchtree) preOrder() {
	fmt.Println("PreOrder  :")
	t.preOrderHelper(t.root)
}

func (t *BinarySearchtree) postOrder() {
	fmt.Println("PostOrder  :")
	t.postOrderHelper(t.root)
}
func (t *BinarySearchtree) preOrderHelper(temp *Node) {
	if temp != nil {
		fmt.Print(temp.Data, " ")
		t.preOrderHelper(temp.lChild)
		t.preOrderHelper(temp.rChild)
	}
}
func (t *BinarySearchtree) postOrderHelper(temp *Node) {
	if temp != nil {

		t.postOrderHelper(temp.lChild)
		t.postOrderHelper(temp.rChild)
		fmt.Print(temp.Data, " ")
	}
}

func main() {
	tree := BinarySearchtree{}
	tree.addNode()
	tree.addNode()
	tree.addNode()
	tree.addNode()
	tree.addNode()
	tree.inOrder()
	fmt.Println()
	tree.preOrder()
	fmt.Println()
	tree.postOrder()
	tree.contain()
	tree.remove()

}
