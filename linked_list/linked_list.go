package main

import "fmt"

// Node : struct
type Node struct {
	data int
	next *Node
}

// NewNode : data
//
// create a node instance
func NewNode(data int) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

// ListLength : head node
//
// count length of list
func ListLength(head *Node) (count int) {
	var current *Node
	count = 1
	current = head
	for current.next != nil {
		count++
		current = current.next
	}
	return
}

// InsertNodeAtPos : head, data, position
//
// insert new node to head at posistion
func InsertNodeAtPos(head **Node, data int, pos int) {
	var newNode *Node
	newNode = &Node{
		data: data,
		next: nil,
	}
	if pos == 0 {
		newNode.next = *head
		*head = newNode
	}
}

func main() {
	var head *Node
	head = &Node{
		data: 0,
		next: nil,
	}

	// Insert a node at beginning
	InsertNodeAtPos(&head, 1, 0)

	// Count the list
	count := ListLength(head)
	fmt.Println(count)

}
