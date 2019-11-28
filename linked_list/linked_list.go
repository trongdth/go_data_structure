package main

import "log"

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
	count = 0
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
	var (
		newNode *Node
		p       *Node
		q       *Node
		k       int
	)
	newNode = &Node{
		data: data,
		next: nil,
	}
	p = *head
	if pos == 0 {
		newNode.next = p
		*head = newNode

	} else {
		for k < pos && p != nil {
			k++
			q = p
			p = p.next
		}

		q.next = newNode
		newNode.next = p
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

	// Insert a node at custom pos
	InsertNodeAtPos(&head, 3, 1)

	// Count the list
	count := ListLength(head)
	log.Println(count)

}
