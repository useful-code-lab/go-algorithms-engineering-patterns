package main

type Node struct {
	Value int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (*LinkedList) Add(newNode *Node) {

}

func main() {
	list := &LinkedList{}
	node := Node{Value: 1}
	list.Add(&node)
}