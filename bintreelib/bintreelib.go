// Package bintreelib: Basic Binary Tree Stuff
package bintreelib

type Node struct {
	left  *Node
	data  string
	right *Node
}

// Node's implementation of the fmt.Stringer interface
func (node Node) String() string {
	return node.data
}

type BinaryTree struct {
	root *Node
}

