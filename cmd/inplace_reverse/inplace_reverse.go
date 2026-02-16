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
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func main() {
	list := &LinkedList{}
	node := Node{Value: 1}
	list.Add(&node)
	

	for i := range(5) {
		list.Add(&Node{Value: i+6})
	}

	current := list.Head
	i := 1

	for current.Next != nil {
		fmt.Printf("%d-й элемент : %d\n", i, current.Value)
		current = current.Next
		i++
	}
	fmt.Printf("%d-й элемент : %d\n", i, current.Value)

}