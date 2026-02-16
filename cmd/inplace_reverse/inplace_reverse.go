package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Add(newNode *Node) {
	if list.Head == nil {
		list.Head = newNode
	}
}

func main() {
	list := &LinkedList{}
	node := Node{Value: 1}
	list.Add(&node)
	
	current := list.Head
	fmt.Printf("1-й элемент : %v",current.Value)

	for current != nil {

	}

}