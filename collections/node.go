package collections

import "fmt"

// Node
type Node struct {
	ID    string
	Edges []*Node
}

// Create a new Node
func NewNode(id string) *Node {
	return &Node{
		ID:        id,
		Edges:     []*Node{},
	}
}


// Print a node
func (n *Node) ToString() string {
	return fmt.Sprintf("%v", n.ID)
}

