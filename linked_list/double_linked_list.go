package main

import (
	"fmt"
	"log"
	"strings"
)

// DLLNode : double linkedin list node
type DLLNode struct {
	data int
	next *DLLNode
	prev *DLLNode
}

// ListLength : count length of list
//
// params: head node
func ListLength(head *DLLNode) (count int) {
	if head == nil {
		return 0
	}
	var current *DLLNode
	count = 1
	current = head
	for current.next != nil {
		count++
		current = current.next
	}
	return
}

// DDLInsertNodeAtPos :
//
// insert a node into [head] at [pos] with [data]
func DDLInsertNodeAtPos(head **DLLNode, pos int, data int) {
	var (
		newNode *DLLNode
		p       *DLLNode
		q       *DLLNode
		k       int
	)
	newNode = &DLLNode{
		data: data,
		next: nil,
		prev: nil,
	}

	p = *head
	if pos == 0 {
		newNode.next = *head
		newNode.prev = nil

		if *head != nil {
			(*head).prev = newNode
		}

		*head = newNode
	} else {
		for k < pos && p != nil {
			k++
			q = p
			p = p.next
		}

		newNode.next = q.next
		newNode.prev = q

		if p != nil {
			p.prev = newNode
		}

		q.next = newNode
	}
}

func printNode(head **DLLNode) {
	var (
		p       *DLLNode
		strNode string
	)
	p = *head
	for p != nil {
		strNode = fmt.Sprintf("%s %d", strNode, p.data)
		p = p.next
	}

	log.Println(strings.TrimSpace(strNode))
}

func main() {
	var head *DLLNode
	head = &DLLNode{
		data: 0,
		prev: nil,
		next: nil,
	}
	buildNodeList(&head)

	// Count the list
	count := ListLength(head)
	log.Println(count)

	printNode(&head)
}

func buildNodeList(head **DLLNode) {
	DDLInsertNodeAtPos(head, 1, 0)
	DDLInsertNodeAtPos(head, 2, 1)
	DDLInsertNodeAtPos(head, 3, 2)
}
