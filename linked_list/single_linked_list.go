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
	if head == nil {
		return 0
	}
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

// DeleteNode : head, pos
//
// delete a note at pos
func DeleteNode(head **Node, pos int) {
	var (
		q *Node
		p *Node
		k int
	)

	if head == nil {
		return
	}

	p = *head
	if pos == 0 {
		*head = p.next
	} else {
		for k < pos && p != nil {
			k++
			q = p
			p = p.next
		}

		if p == nil {
			log.Println("this pos is not found")
		} else {
			q.next = p.next
		}
	}
}

// DeleteListOfNode : head
//
// delete list of node
func DeleteListOfNode(head **Node) {
	var p, q *Node
	p = *head
	for p.next != nil {
		q = p.next
		p = nil
		p = q
	}
	*head = nil
}

func main() {
	var head *Node
	head = &Node{
		data: 0,
		next: nil,
	}

	buildNodeList(&head)

	// Count the list
	count := ListLength(head)
	log.Println(count)

	// Delete a node at beginning
	DeleteNode(&head, 0)

	// Count the list
	count = ListLength(head)
	log.Println(count)

	// Delete a node at custom pos
	DeleteNode(&head, 1)

	// Count the list
	count = ListLength(head)
	log.Println(count)

	// Delete a node at custom pos
	DeleteNode(&head, 100)

	// Count the list
	count = ListLength(head)
	log.Println(count)

	// Delete all nodes
	DeleteListOfNode(&head)
	count = ListLength(head)
	log.Println(count)
}

func buildNodeList(head **Node) {
	InsertNodeAtPos(head, 1, 0)
	InsertNodeAtPos(head, 2, 1)
	InsertNodeAtPos(head, 3, 2)
}
