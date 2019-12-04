package main

// DLLNode : double linkedin list node
type DLLNode struct {
	data int
	next *DLLNode
	prev *DLLNode
}

// DDLInsertNodeAtPos :
//
// insert a node into [head] at [pos] with [data]
func DDLInsertNodeAtPos(head **DLLNode, pos int, data int) {
	var (
		newNode *DLLNode
		p       *DLLNode
	)
	newNode = &DLLNode{
		data: data,
		next: nil,
		prev: nil,
	}
}

func main() {

}
